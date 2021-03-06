package jianshu

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	rootUrl        = "https://www.jianshu.com/"
	rootUrlOnce    = "https://www.jianshu.com"
	trendingSearch = "https://www.jianshu.com/trending_search"
	homePageAuthor = "https://www.jianshu.com/recommendations/users?utm_source=desktop&utm_medium=index-users"
)

func MakeCompleteUrl(element string) string {
	if element == "" {
		panic("element should not be nil.")
	}
	if strings.HasPrefix(element, "/") {
		return rootUrlOnce + element
	}
	return rootUrl + element
}

func SplitString(element string) string {
	newElement := strings.SplitAfter(element, "-")
	return newElement[1]
}

func StringToInt(element string) int {
	newInt, _ := strconv.Atoi(element)
	return newInt
}

func Pages(allPassage int) int {
	return allPassage/9 + 1
}

func StringSpace(element string) string {
	return strings.TrimSpace(element)
}

func GetUuidFromLink(user *User) string {
	if user.Link == "" {
		return "None"
	}
	urlList := strings.Split(user.Link, "/")
	return urlList[4]
}

// 各种URL集合
func LikedNotesUrl(user *User) string {
	id := GetUuidFromLink(user)
	if id == "None" {
		panic("id should not be nil.")
	}
	return rootUrl + "users/" + id + "/liked_notes?_pjax=%23list-container"
}

func SubscriptionUrl(user *User) string {
	id := GetUuidFromLink(user)
	if id == "None" {
		panic("id should not be nil.")
	}
	return rootUrl + "users/" + id + "/subscriptions?_pjax=%23list-container"
}

func TimeLineUrl(user *User) string {
	id := GetUuidFromLink(user)
	if id == "None" {
		panic("id should not be nil.")
	}
	return rootUrl + "users/" + id + "/timeline?_pjax=%23list-container"
}

func CommentedUrl(user *User) string {
	id := GetUuidFromLink(user)
	if id == "None" {
		panic("id should not be nil.")
	}
	return user.Link + "?order_by=commented_at&_pjax=%23list-container"
}

func HotPassageUrl(user *User) string {
	id := GetUuidFromLink(user)
	if id == "None" {
		panic("id should not be nil.")
	}
	return user.Link + "?order_by=top&_pjax=%23list-container"
}

func SearchQuery(query string, page int) string {
	return rootUrl + fmt.Sprintf("/search?q=%s&page=%d&type=note", query, page)
}

func GetTrendSearch() string {
	return trendingSearch
}

func GetSpecialSubjectURL(id int) string {
	return fmt.Sprintf("https://www.jianshu.com/collections/%s/side_list", strconv.Itoa(id))
}

// 首页推荐作者
func HomePageAuthorUrl() string {
	return homePageAuthor
}

func StringGetInt(value string) int {
	tempInt := strings.SplitAfter(value, " ")
	return StringToInt(tempInt[1])
}

func StringHandle(value string) (string, string) {
	newStringFirst := StringSpace(strings.Replace(value, " ", "", -1))
	newList := strings.Split(newStringFirst, "\n")
	//fmt.Println(newList[0], newList[len(newList)-1])
	return newList[0], newList[len(newList)-1]
}

func StringCommon(value string) string {
	newStringFirst := strings.Replace(value, "\n", "", -1)
	newString := strings.Replace(newStringFirst, " ", "", -1)
	return StringSpace(newString)
}
func StringSplitWith(value string) (string, string) {
	newString := strings.Split(value, "，")
	return newString[0], newString[len(newString)-1]
}

func RegexpFindSting(value string) (int, int) {
	reg, _ := regexp.Compile(`\d+`)
	newList := reg.FindAllString(value, -1)
	numberOne, _ := strconv.Atoi(newList[0])
	numberTwo, _ := strconv.Atoi(newList[1])
	return numberOne, numberTwo
}

func JsonPretty(value interface{}) {
	data, _ := json.MarshalIndent(value, "", " ")
	fmt.Println(string(data))
}
