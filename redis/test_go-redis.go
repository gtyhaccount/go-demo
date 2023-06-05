package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {

	rc := redis.NewClient(&redis.Options{
		Addr: "192.168.152.1:6379",
	})

	defer rc.Close()
	rc.Set("user-test1", "5555", time.Minute)

	a := &ArticleInDDB{}
	a.ID = "abcdef"
	str, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	obj := string(str)

	rc.HSet("recommend-article", a.ID, obj)

	r2, err := rc.HMGet("recommend-article", a.ID).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r2)
	rc.Del("recommend-article")
	r1, err := rc.HGetAll("recommend-article").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r1)

}

type ArticleInDDB struct {
	ID                   string `json:"id" gorm:"primary_key"`
	ResourceID           int64  `json:"resource_id"`
	Content              string `json:"content"`
	UserID               string `json:"user_id"`
	UserName             string `json:"user_name"`
	AuthorDirectorUserID string `json:"author_director_user_id"` // 作者的督导
	AuthorNsdUserID      string `json:"author_nsd_user_id"`      // 作者的首席

}
