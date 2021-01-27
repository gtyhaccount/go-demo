package s3

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"
)

const (
	tableName               = "product_recommend_user_product"
	ExpireTimeOneMonth      = 2592000
	userRecommendProductKey = "product:recommend:user:%s:algorithms:%s"
)

var ctx = context.Background()

type UserAlgorithms struct {
	ContactId   string
	Algorithms  string
	CreatedTime time.Time
}

type UserProductRecommends struct {
	ContactId   string
	Sku         string
	Scores      string
	Algorithms  string
	CreatedTime time.Time
}

type ProductProductRecommends struct {
	Sku          string
	RecommendSku string
	Similarity   float64
	Price1       float64
	Price2       float64
	DiffPrice    float64
	Count        int32
	CreatedTime  time.Time
}

func TestHeadBucket(t *testing.T) {
	sess := getS3Session()

	// Create S3 service client
	svc := s3.New(sess)

	//if err != nil {
	//	fmt.Errorf("create s3 service error: %s", err)
	//}

	result, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String("bjs-s3-mdm-prod-mkc-mobile-resources"),
	})
	if err != nil {
		t.Errorf("list buckets error: %s", err)
	}

	fmt.Printf("head bucket input response:%s \n", result)
}

func TestGetObject(t *testing.T) {
	sess := getS3Session()

	downloader := s3manager.NewDownloader(sess)

	bucket := "bjs-s3-mdm-prod-mkc-mobile-resources"
	s3FilePath := "/MyBiz-ML/recommend_result_china/"
	item := "item_popularity_china.csv"

	file, err := os.Create(item)
	if err != nil {
		exitErrorf("Unable to open file %q,%v", item, err)
	}

	defer file.Close()

	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(s3FilePath + item),
	})

	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func TestReadCsv(t *testing.T) {
	fileName := "recommend_result_china.csv"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("can not open the file,err is %+v \n", err)
	}
	defer file.Close()

	sql := fmt.Sprintf("insert into %s (`contact_id`,`sku`,`scores`,`algorithms`,`created_time`) values ", tableName)
	batchInsertSql := sql

	db := getDb()

	var insertNum int
	r := csv.NewReader(file)
	for {
		insertNum++

		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}

		if err == io.EOF {
			log.Infof("read to the end of the file")
			break
		}

		batchInsertSql += fmt.Sprintf("('%s','%s','%s','%s','%s'),", row[0], row[1], row[2], row[3], time.Now().Format("2006-01-02 15:04:05"))

		if insertNum >= 1000 {
			batchInsertSql = batchInsertSql[:len(batchInsertSql)-1]

			// insert table
			if err := db.Table(tableName).Exec(batchInsertSql).Error; err != nil {
				log.Fatalf("insert error:%s", err)
			}

			batchInsertSql = sql
			insertNum = 0
		}
	}

	if insertNum > 0 {
		batchInsertSql = batchInsertSql[:len(batchInsertSql)-1]
		if err := db.Table(tableName).Exec(batchInsertSql).Error; err != nil {
			log.Fatalf("last insert error:%s", err)
		}
	}

	defer db.Close()
}

/**
分页查询数据库数据然后同步到redis，设置过期时间为一个月
*/
func TestSyncDataToRedis(t *testing.T) {
	db := getDb()

	// 分组用户和算法
	var allProductRecommends []UserProductRecommends
	if err := db.Table(tableName).Select("contact_id,algorithms").Group("contact_id,algorithms").Find(&allProductRecommends).Error; err != nil {
		log.Errorf("group user algorithms err:%s", err)
		return
	}

	conn := getRedisConn()
	defer conn.Close()

	var userRecommends []UserProductRecommends
	if allProductRecommends != nil && len(allProductRecommends) > 0 {
		for _, recommend := range allProductRecommends {
			if err := db.Table(tableName).Where("contact_id=? and algorithms=?", recommend.ContactId, recommend.Algorithms).
				Order("scores desc").Find(&userRecommends).Error; err != nil {
				log.Errorf("find recommendation by userID and algorithms err:%s", err)
				return
			}

			// 每个用户、一种算法设置一个key
			var recommendSku []string
			for _, userRecommend := range userRecommends {
				recommendSku = append(recommendSku, userRecommend.Sku)
			}

			go setRedisKey(conn, fmt.Sprintf(userRecommendProductKey, recommend.ContactId, recommend.Algorithms), recommendSku)
		}
	}
}

func TestRandRecommend(t *testing.T) {
	skus := []string{"10020713", "10023682", "10023683", "10023689", "10035569", "10035580", "10036259", "10048969", "10041961", "10049628"}

	rand.Seed(time.Now().Unix())
	fmt.Printf("random recommend index,%d \n", rand.Intn(len(skus)))
	fmt.Printf("random recommend,%v \n", skus[rand.Intn(len(skus))])
	indexArr := rand.Perm(len(skus))
	fmt.Printf("random recommend,%v \n", indexArr)
	arr := indexArr[:len(indexArr)-3]
	fmt.Printf("random recommend,%v \n", arr)
	sort.Ints(arr)
	fmt.Printf("random recommend,%v \n", arr)
}

func TestTime(t *testing.T) {
	temp := 0
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		for j := 0; j < 10000; j++ {
			temp++
		}
	}
	endTime := time.Now()

	fmt.Println(endTime.Sub(startTime))
	fmt.Println(endTime.Sub(startTime).Milliseconds())

	fmt.Println(time.Now().Format("200601021504"))
}

func TestCron(t *testing.T) {
	c := cron.New()

	c.AddFunc("23 20/1 * * *", func() {
		fmt.Println("……")
		fmt.Println(c.Entries()[0].Next)
	})

	c.Start()

	for {
		time.Sleep(60 * time.Second)
	}

}

func TestString(t *testing.T) {
	str := "product_recommend_popularity_product_20201120"
	fmt.Println(str[:len(str)-9])
}

func changeIntValue(i *int) {
	*i++
}

func setRedisKey(rdb *redis.Client, key string, value []string) {
	if err := rdb.Del(ctx, key).Err(); err != nil {
		log.Fatalf("del key error:%s", err)
	}

	if err := rdb.RPush(ctx, key, value).Err(); err != nil {
		log.Fatalf("rpush error:%s", err)
	}

	if err := rdb.Expire(ctx, key, ExpireTimeOneMonth*time.Second).Err(); err != nil {
		log.Fatalf("set expire error:%s", err)
	}
}

func getS3Session() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("AKIAQEWZO6CISBH2AGGM", "jFwhZzvkZ8/gNrJ4t4Rln2I/UlvU+5LGON+t8nu9", ""),
		Region:      aws.String(endpoints.CnNorth1RegionID),
	})

	if err != nil {
		fmt.Errorf("get s3 session error: %s", err)
	}

	return sess
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func getDb() *gorm.DB {
	// init mysql
	mysqlAddr := "root:123456@tcp(localhost:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", mysqlAddr)
	if err != nil {
		log.Fatal(err)
	}

	// show sql
	db = db.Debug()

	return db
}

func getRedisConn() *redis.Client {

	// init redis pool
	rdc := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		DB:          0,
		IdleTimeout: 240 * time.Second,
	})

	return rdc
}
