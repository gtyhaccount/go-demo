package main

import (
	"fmt"
	"sort"
)

type Hero struct {
	Sort     int64  `json:"sort"`
	ImageUrl string `json:"image_url"`
}

type heroList []Hero

func (h heroList) Len() int {
	return len(h)
}

func (h heroList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heroList) Less(i, j int) bool {
	return h[i].Sort < h[j].Sort
}

func TestSlice() {
	heros := []Hero{{Sort: 3, ImageUrl: "aaa"},
		{Sort: 2, ImageUrl: "bbb"},
		{Sort: 1, ImageUrl: "ccc"}}
	heros2 := []Hero{{Sort: 1, ImageUrl: "aaa"},
		{Sort: 2, ImageUrl: "bbb"},
		{Sort: 1, ImageUrl: "ccc"}}
	sort.Sort(heroList(heros))
	sort.Sort(heroList(heros2))

	fmt.Printf("heros1 %+v", heros)
	fmt.Printf("heros2 %+v", heros2)
}

func (Hero) SortArray(hero []Hero) []Hero {
	sort.Sort(heroList(hero))
	fmt.Printf("heros1 %+v", hero)
	return hero
}
func main() {

	//heros:=[]Hero{{Sort: 3,ImageUrl: "aaa"},
	//	{Sort: 4,ImageUrl: "bbb"},
	//	{Sort: 1,ImageUrl: "ccc"}}
	//
	var heros2 []Hero

	Hero{}.SortArray(heros2)
	fmt.Printf("h2 %+v", heros2)

	//TestSlice()
	//s8 := []int{1, 2, 3, 4, 5, 6}
	//s9 := []int{7, 8, 9}
	//copy(s8, s9)    // 将s9的值拷贝到s8中
	//fmt.Println(s8) // [7,8,9,1,2,3]
	//copy(s9,s8) // 将s9的值拷贝到s8中
	//fmt.Println(s9)// [1,2,3]
}
