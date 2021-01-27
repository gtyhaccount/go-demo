package main

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"strconv"
	"time"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Results *Items      `json:"results"`
	Result  interface{} `json:"result"`
}

type Items struct {
	List     interface{} `json:"list"`
	PageSize int         `json:"page_size"`
	PageNum  int         `json:"page_num"`
	Total    int         `json:"total"`
}

type BatchInputSkuFromPIMRequest struct {
	Env  string            `json:"env"`
	Skus []*SkuInfoFromPIM `json:"skus"`
}

type BatchGetSkuExtensionRequest struct {
	Skus string `json:"skus"`
}

type BatchGetSkuExtensionResponse struct {
	List []*SkuExtensionInfo `json:"list"`
}

type SkuInfoFromPIM struct {
	ProductTitle       string          `json:"product_title" validate:"required"`
	IncludeInSearch    bool            `json:"include_in_search" validate:"required"` // 是否可搜索
	ProductType        int             `json:"product_type" validate:"required"`
	SectionCode        int             `json:"section_code" validate:"required"`
	ProductLine        int             `json:"product_line" validate:"required"`
	Sku                string          `json:"sku" validate:"required"`
	SkuTitle           string          `json:"sku_title" validate:"required"`
	RetailPrice        float64         `json:"retail_price" validate:"required"`
	SuggestRetailPrice float64         `json:"suggest_retail_price" validate:"required"`
	RedeemPoint        int             `json:"redeem_point"`
	ChildProducts      []*ChildProduct `json:"child_products"`
}

type ChildProduct struct {
	Sku string `json:"sku" validate:"required"`
	Num int    `json:"num" validate:"required"`
}

type SkuExtensionInfo struct {
	BasicInfo              BasicInfoBean              `json:"basic_info"`
	MarketingConfiguration MarketingConfigurationBean `json:"marketing_configuration"`
	SkuImages              SkuImagesBean              `json:"sku_images"`
	SheSays                SheSaysBean                `json:"she_says"`
	ShareAndRecommend      ShareAndRecommendBean      `json:"share_and_recommend"`
	CustomParameters       []CustomParameter          `json:"custom_parameters"`
	XmlContent             XmlContentBean             `json:"xml_content"`
}

type BasicInfoBean struct {
	Sku                 string              `json:"sku"`
	SubstituteSku       string              `json:"substitute_sku"`
	EnglishProductTitle string              `json:"english_product_title"`
	Volume              string              `json:"volume"`
	ProductGroupName    string              `json:"product_group_name"`
	RgbColor            string              `json:"rgb_color"`
	RgbColorName        string              `json:"rgb_color_name"`
	DeliveryMethod      string              `json:"delivery_method"`
	SkuDescription      string              `json:"sku_description"`
	AdditionalInfo      string              `json:"additional_info"`
	ImportTaxUrl        string              `json:"import_tax_url"`
	ImportCountryUrl    string              `json:"import_country_url"`
	Specification       []SpecificationBean `json:"specification"`
	ServiceDescription  []SpecificationBean `json:"service_description"`
}

type SpecificationBean struct {
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
	ImageUrl    string `json:"image_url"`
}

type MarketingConfigurationBean struct {
	SkuSellingPoints string  `json:"sku_selling_points"`
	MarketingTags    []Label `json:"marketing_tags"`
	ProductTags      []Label `json:"product_tags"`
	RecommendTag     Label   `json:"recommend_tag"`
	AwardTags        []Label `json:"award_tags"`
	RecommendWords   string  `json:"recommend_words"`
}

type SkuImagesBean struct {
	PngImageUrl      string    `json:"png_image_url"`
	ScanCodeImageUrl string    `json:"scan_code_image_url"`
	ListImage        ImageBean `json:"list_image"`
	HeadImage        ImageBean `json:"head_image"`
	DetailImage      ImageBean `json:"detail_image"`
}

type ImageBean struct {
	OfficialWebsitePc     []Image `json:"official_website_pc"`
	OfficialWebsiteMobile []Image `json:"official_website_mobile"`
	App                   []Image `json:"app"`
	Applets               []Image `json:"applets"`
}

type Image struct {
	ImageUrl      string `json:"image_url"`
	VideoCoverUrl string `json:"video_cover_url"`
	VideoUrl      string `json:"video_url"`
	Sort          int    `json:"sort"`
}

type SheSaysBean struct {
	SheSaysTag         string `json:"she_says_tag"`
	SheSaysDescription string `json:"she_says_description"`
	SheSaysImageUrl    string `json:"she_says_image_url"`
	SheSaysVideoUrl    string `json:"she_says_video_url"`
}

type ShareAndRecommendBean struct {
	ShareTitle     string   `json:"share_title"`
	ShareCoverUrl  string   `json:"share_cover_url"`
	ShareWords     []string `json:"share_words"`
	SharePosterUrl []string `json:"share_poster_url"`
}

type CustomParameter struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type         int    `json:"type"`
	Sort         int    `json:"sort"`
	DefaultValue string `json:"default_value"`
	CreatedBy    string `json:"created_by"`
}

type XmlContentBean struct {
	CustomPrice      string             `json:"custom_price"`
	RelatedProducts  []RecommendProduct `json:"related_products"`
	ExtendProperties string             `json:"extend_properties"`
	ImageVersion     string             `json:"image_version"`
}

type RecommendProduct struct {
	Sku      string `json:"sku"`
	SkuTitle string `json:"sku_title"`
	Sort     int    `json:"sort"`
}

type SkuListRequest struct {
	Skus            string `json:"skus" query:"skus"`
	SkuTitle        string `json:"sku_title" query:"sku_title"`
	UpdateTimeStart int64  `json:"update_time_start" query:"update_time_start"`
	UpdateTimeEnd   int64  `json:"update_time_end" query:"update_time_end"`
	ProductType     int    `json:"product_type" query:"product_type"`
	PublishType     int    `json:"publish_type" query:"publish_type"`
	PageSize        int    `json:"page_size" query:"page_size"`
	PageNum         int    `json:"page_num" query:"page_num"`
}

type PublishSkuRequest struct {
	Skus      string `json:"skus"`
	TargetEnv string `json:"target_env"`
}

type GetListRequest struct {
	PageSize int `json:"page_size" query:"page_size"`
	PageNum  int `json:"page_num" query:"page_num"`
}

type Award struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	CreatedBy string `json:"created_by"`
}

type Label struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	CreatedBy string `json:"created_by"`
}

//func main() {
//	body := `{
//    "sku_extension_info":{
//        "basic_info":{
//            "substitute_sku":"2020001606",
//            "english_product_title":"奶油唇膏",
//            "volume":"160g",
//            "product_group_name":"产品分组名称",
//            "rgb_color":"168,133,66",
//            "rgb_color_name":"色号名称",
//            "delivery_method":"配送方式",
//            "sku_description":"产品描述",
//            "additional_info":"补充说明",
//            "import_tax_url":"进口税说明图片url",
//            "import_country_url":"进口国显示图片url",
//            "specification":{
//                "title":"标题",
//                "explanation":"描述",
//                "image_url":"图片url"
//            },
//            "service_description":{
//                "title":"标题",
//                "explanation":"描述",
//                "image_url":"图片url"
//            }
//        },
//        "marketing_configuration":{
//            "sku_selling_points":"产品卖点",
//            "marketing_tags":[
//                {
//                    "id":1,
//                    "name":"畅销单品",
//                    "type":2,
//                    "sort":1
//                }],
//            "product_tags":[
//                {
//                    "id":2,
//                    "name":"畅销组合",
//                    "type":1,
//                    "sort":1
//                }],
//            "recommend_tags":[
//                {
//                    "id":3,
//                    "name":"畅销组合",
//                    "type":1,
//                    "sort":1
//                }],
//            "recommend_words":"榜单推荐话术",
//            "award_tags":[
//                {
//                    "id":6,
//                    "name":"《时尚COSMO》年度精华产品大奖",
//                    "image_url":"http://1.jpg",
//                    "sort":1
//                }]
//        },
//        "sku_images":{
//            "png_image_url":"产品png图片url",
//            "scan_code_image_url":"扫码识别产品图片url",
//            "list_image":{
//                "official_website_pc":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "official_website_mobile":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "app":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "applet":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }]
//            },
//            "head_image":{
//                "official_website_pc":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "official_website_mobile":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "app":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "applet":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }]
//            },
//            "detail_image":{
//                "official_website_pc":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "official_website_mobile":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "app":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }],
//                "applet":[
//                    {
//                        "image_url":"http://image.jpg",
//                        "video_cover_url":"http://cover.jpg",
//                        "video_url":"http://video.jpg",
//                        "sort":1
//                    }]
//            }
//        },
//        "she_says":{
//            "she_says_tag":[
//                {
//                    "id":8,
//                    "name":"买大赠小",
//                    "type":1
//                }],
//            "she_says_description":"文案",
//            "she_says_image_url":"图片url",
//            "she_says_video_url":"视频url"
//        },
//        "share_and_recommend":{
//            "share_title":"分享标题",
//            "share_words":[
//                "分享话术"],
//            "share_cover_url":"分享封面url",
//            "share_poster_urls":[
//                "分享海报url"]
//        },
//        "custom_parameters":[
//            {
//                "name":"参数名",
//                "type":1,
//                "default_value":"默认值"
//            }
//        ]
//    }
//}
//`
//
//	//skuExtensionInfoBytes, _, _, _ := jsonparser.Get([]byte(body), "sku_extension_info")
//	//basicInfoByte, _, _, _ := jsonparser.Get(skuExtensionInfoBytes, "basic_info")
//	//marketingConfigurationMarketingTags, _, _, _ := jsonparser.Get(skuExtensionInfoBytes, "marketing_configuration", "marketing_tags")
//	//
//	//var marketingTags []Label
//	//json.Unmarshal(marketingConfigurationMarketingTags, &marketingTags)
//	//fmt.Println(marketingTags[0].Name)
//	//
//	//var basicInfo BasicInfoBean
//	//json.Unmarshal(basicInfoByte, &basicInfo)
//	//fmt.Println(basicInfo)
//
//	// 遍历
//	jsonparser.ObjectEach([]byte(body), func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
//		var tags []struct {
//			Id   int
//			Type int
//			Name string
//		}
//		switch dataType {
//		case jsonparser.String:
//			fmt.Printf("jsonparser object each string:%s", string(value))
//			fmt.Println()
//			break
//		case jsonparser.Array:
//			json.Unmarshal(value, &tags)
//
//			for _, tag := range tags {
//				fmt.Println(tag.Id)
//			}
//			break
//		}
//
//		fmt.Println("-------------------")
//		return nil
//	}, "sku_extension_info", "marketing_configuration")
//
//}

func main() {
	bytes, err := ioutil.ReadFile("D:\\go\\project\\github.com\\study\\basic\\json\\pim_data.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	type childrenProduct struct {
		Sku      string `json:"sku"`
		Quantity int    `json:"quantity"`
	}
	type product struct {
		ProductTitle        *string            `json:"product_title,omitempty"`
		IncludeInSearch     *int8              `json:"-,omitempty"`
		IncludeInSearchBool *bool              `json:"include_in_search,omitempty" gorm:"-"`
		ProductType         *int               `json:"-,omitempty"`
		ProductTypeStr      *string            `json:"product_type,omitempty" gorm:"-"`
		SectionCode         *int               `json:"-,omitempty"`
		SectionCodeStr      *string            `json:"section_code,omitempty" gorm:"-"`
		ProductLine         *int               `json:"-,omitempty"`
		ProductLineStr      *string            `json:"product_line,omitempty" gorm:"-"`
		Sku                 *string            `json:"sku,omitempty"`
		Name                *string            `json:"name,omitempty"`
		RetailPrice         *float64           `json:"retail_price,omitempty"`
		SuggestRetailPrice  *float64           `json:"suggest_retail_price,omitempty"`
		RedeemPoint         *string            `json:"redeem_point,omitempty"`
		ChildProducts       *string            `json:"-,omitempty"`
		ChildProductsArray  *[]childrenProduct `json:"child_products,omitempty" gorm:"-"`
		// basic info
		EnglishName     *string   `json:"english_name,omitempty"`
		Volume          *string   `json:"volume,omitempty"`
		ShadeGroup      *string   `json:"shade_group,omitempty"`
		RgbColor        *string   `json:"rgb_color,omitempty"`
		RgbColorName    *string   `json:"rgb_color_name,omitempty"`
		ExtendProperies *string   `json:"extend_properies,omitempty"`
		CustomPrice     *float64  `json:"custom_price,omitempty"`
		ImageVersion    *string   `json:"image_version,omitempty"`
		CreatedBy       string    `json:"created_by,omitempty"`
		CreatedTime     time.Time `json:"created_time,omitempty"`
		UpdatedBy       string    `json:"updated_by,omitempty"`
		UpdatedTime     time.Time `json:"updated_time,omitempty"`
	}

	db, _ := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/product?charset=utf8mb4&parseTime=True&loc=Local")
	tx := db.Begin()
	jsonparser.ArrayEach(bytes, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		productSku := new(product)

		json.Unmarshal(value, productSku)
		productSku.CreatedBy = "oe"
		productSku.CreatedTime = time.Now()
		productSku.UpdatedBy = "oe"
		productSku.UpdatedTime = time.Now()

		a := switchBoolToInt(*productSku.IncludeInSearchBool)
		productSku.IncludeInSearch = &a

		pt, _ := strconv.Atoi(*productSku.ProductTypeStr)
		productSku.ProductType = &pt

		sc, _ := strconv.Atoi(*productSku.SectionCodeStr)
		productSku.SectionCode = &sc

		pl, _ := strconv.Atoi(*productSku.ProductLineStr)
		productSku.ProductLine = &pl

		s, _ := json.Marshal(productSku.ChildProductsArray)
		s1 := string(s)
		productSku.ChildProducts = &s1
		fmt.Println("current sku:", *productSku.Sku)

		if err := tx.Table("product_sku").Create(productSku).Error; err != nil {
			tx.Rollback()
			panic(err)
		}
	})

	tx.Commit()
}

func switchBoolToInt(b bool) int8 {
	if b {
		return 1
	}

	return 0

}
