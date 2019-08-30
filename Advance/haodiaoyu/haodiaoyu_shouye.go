package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)



type ImageContent struct {
	Image string	`json:"image"`
	Content string `json:"content"`
}


type ArticleDetail struct {
	Title string	`json:"title"`
	Sorting []string `json:"sorting"`
	Articles []map[string]string `json:"articles"`
	Images []map[string]ImageContent `json:"images"`


}


type Publisher struct {
	Name string	`json:"name"`
	Time int64 `json:"time"`
	Portrait string `json:"portrait"`
}



type Article struct {
	Id int	`json:"id"`
	Title string	`json:"title"`
	Images []string	`json:"images"`
	GiveLike int `json:"giveLike"`
	Url string `json:"url"`
	PublisherUser *Publisher `json:"publisher"`
}


type Commutity struct {
	BaseURL string	`json:"baseURL"`
	PageNum int	`json:"pageNum"`
	Articles []Article  `json:"content"`
}


func HttpGetDouPan(url string)(result string, err error)  {

	var response *http.Response

	if response,err = http.Get(url); err!=nil {
		return
	}

	defer response.Body.Close()

	buffer := make([]byte,4096)

	for {
		var n int
		n,err = response.Body.Read(buffer)

		if n==0 {
			err = nil
			break
		}

		if err!=nil && err != io.EOF {
			return
		}
		err = nil
		result += string(buffer[:n])
	}
	return
}

func saveInfoLocalFile(index int,jsonString string)  {

	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/community_"+strconv.Itoa(index)+".json"

	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(jsonString)
}


func saveArticleDetailLocalFile(index int,jsonString string)  {

	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/ArticleDetail/at_"+strconv.Itoa(index)+".json"

	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(jsonString)
}




func workItemBBS(url string,bbsID int,contextTitle string,pageChannel chan int ) {

	println(url)

	result,err := HttpGetDouPan(url)
	if err != nil {
		println("workItemBBS ==> Http Get Error!",err)
		return
	}

	//fmt.Println("workItemBBS result=>", result)
	articleDetail:= ArticleDetail{
		Title:contextTitle,
	}
	var sorting []string
	var articles []map[string]string
	var imagesSlice []map[string]ImageContent


	regexp0 := regexp.MustCompile(`<table cellspacing="0" cellpadding="0"><tr><td class="t_f" id="(.*?)<br />

<ignore_js_op>`)
	filterTitle :=  regexp0.FindAllStringSubmatch(result,1)
	var artCount = 0
	if len(filterTitle)==1 {
		title := filterTitle[0][1]
		titleSP :=strings.Split(title, ">")
		//fmt.Println(titleSP[1])
		map1 := make(map[string]string)
		map1["art1"] = titleSP[1]
		articles = append(articles, map1)
		sorting = append(sorting,"art1")
		artCount = 1
	}else {
		artCount = 0
	}


	regexp1 := regexp.MustCompile(`</ignore_js_op>
<br />(?s:(.*?))<ignore_js_op>`)
	filterName :=  regexp1.FindAllStringSubmatch(result,-1)
	fmt.Println(len(filterName))

	regexp2 := regexp.MustCompile(`<p style="font-size: 15px; text-align: center; color: #666;">(.*?)</p>`)
	filterImageContent :=  regexp2.FindAllStringSubmatch(result,-1)

	regexp3 := regexp.MustCompile(`<ignore_js_op>
<img(?s:(.*?))<p style="font-size: 15px; text-align: center; color: #666;">`)
	filterImages :=  regexp3.FindAllStringSubmatch(result,-1)
	regexp4 := regexp.MustCompile(`" file="(.*?)" class="zoom"`)

	length2:= len(filterImageContent)
	length := len(filterName)
	println(length2)

	lens := length
	if length>length2{
		lens = length2
	}
	for i:=1;i<=lens ;i++ {
		//fmt.Println("===========>  :",filterName[i-1][1])
		articleStr := filterName[i-1][1]
		articleStr = strings.Trim(articleStr,"\r\n")
		articleStr = strings.Trim(articleStr,"<br />")
		articleStr = strings.Replace(articleStr, "<br />", "", -1)
		fmt.Println("===========>  :",articleStr)

		var imageTag = "imageShow_"+strconv.Itoa(i)
		sorting = append(sorting,imageTag)
		if len(articleStr) == 0 {

		} else {
			artCount++
			tag := "art"+strconv.Itoa(artCount)
			sorting = append(sorting,tag)
			map2 := make(map[string]string)
			map2[tag] = articleStr
			articles = append(articles, map2)
		}

		str := filterImages[i-1][0]
		filterImageURL := regexp4.FindAllStringSubmatch(str,1)
		fmt.Println(filterImageURL[0][1])

		imageContent := filterImageContent[i-1][1]
		ic :=ImageContent{
			Image:filterImageURL[0][1],
			Content:imageContent,
		}
		icMap := make(map[string]ImageContent)
		imageShowTag := "imageShow_"+strconv.Itoa(i)
		icMap[imageShowTag] = ic
		imagesSlice = append(imagesSlice, icMap)
	}

	articleDetail.Articles = articles
	articleDetail.Sorting = sorting
	articleDetail.Images = imagesSlice
	str, err := json.Marshal(articleDetail)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json", string(str))
	saveArticleDetailLocalFile(bbsID,string(str))
	pageChannel <- bbsID
}



func workDouban(pageNumber int)  {

	url := "https://www.haodiaoyu.com/"
	result,err := HttpGetDouPan(url)
	if err != nil {
		println("Http Get Error!",err)
		return
	}

	fmt.Printf("========第 %d 页 抓取成功，开始分析页面============\n",1)

	pageChannel := make(chan  int)
	articleList := make([]Article,0)

	regexp1 := regexp.MustCompile(`<div class="article">(?s:(.*?))</span></span>`)
	filterName :=  regexp1.FindAllStringSubmatch(result,20)
	for _,nameList := range filterName {
		fmt.Println("-----------------------********---------------------------------------")
		//fmt.Println(nameList[0])
		article := nameList[0]

		regexp1 := regexp.MustCompile(`.html" target="_blank">(?s:(.*?))</a></h3>`)
		filterTitle :=  regexp1.FindAllStringSubmatch(article,1)
		fmt.Println(filterTitle[0][1])

		regexp2 := regexp.MustCompile(`">(.*?)</a></span>`)
		filterUserName :=  regexp2.FindAllStringSubmatch(article,1)
		fmt.Println(filterUserName[0][1])

		regexp3 := regexp.MustCompile(`data-tid="(.*?)">`)
		filterID:=  regexp3.FindAllStringSubmatch(article,1)
		fmt.Println(filterID[0][1])

		regexp4 := regexp.MustCompile(`class="praise-number">(?s:(.*?))</span></span>`)
		filterGLike:=  regexp4.FindAllStringSubmatch(article,1)
		fmt.Println(filterGLike[0][1])

		regexp5 := regexp.MustCompile(`<a href="(.*?)" target="_blank">`)
		filterURL:=  regexp5.FindAllStringSubmatch(article,1)
		fmt.Println(filterURL[0][1])

		regexp6 := regexp.MustCompile(`><img src="(.*?)" alt=""></a></i>`)
		filterAvrtar:=  regexp6.FindAllStringSubmatch(article,1)
		fmt.Println(filterAvrtar[0][1])

		regexp7 := regexp.MustCompile(`<div class="article-thumb">(?s:(.*?))</div>`)
		filterImages:=  regexp7.FindAllStringSubmatch(article,1)
		imagesString := filterImages[0][1]

		regexp8 := regexp.MustCompile(`<img src="(.*?)" alt="" />`)
		filterImagesURL:=  regexp8.FindAllStringSubmatch(imagesString,3)
		//fmt.Println(filterImagesURL[0][1])
		//fmt.Println(filterImagesURL[1][1])
		//fmt.Println(filterImagesURL[2][1])

		randNumber := rand.Intn(100000)+10000
		now := time.Now().Unix()-int64(randNumber)
		publisheruser := Publisher{}
		publisheruser.Name = filterUserName[0][1]
		publisheruser.Portrait = filterAvrtar[0][1]
		publisheruser.Time= now


		filterIDValue, _ := strconv.Atoi(filterID[0][1])
		giveLikeValue, _ := strconv.Atoi(filterGLike[0][1])
		imagesList := make([]string,0)
		imagesList = append(imagesList,filterImagesURL[0][1],filterImagesURL[1][1],filterImagesURL[2][1])



		articleObj := Article{
		}
		articleObj.Id = filterIDValue
		articleObj.Title = filterTitle[0][1]
		articleObj.Images =imagesList
		articleObj.GiveLike = giveLikeValue
		articleObj.Url = filterURL[0][1]
		articleObj.PublisherUser = &publisheruser
		articleList = append(articleList,articleObj)

		go workItemBBS(filterURL[0][1],filterIDValue,filterTitle[0][1],pageChannel)

		fmt.Println("-----------------------#######--------------------------------------")
	}

	commutity := Commutity{}
	commutity.BaseURL = "http://localhost:5000"
	commutity.PageNum = pageNumber
	commutity.Articles = articleList

	//Marshal失败时err!=nil
	jsonCommutity, err := json.Marshal(commutity)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	//fmt.Println(string(jsonCommutity))

	saveInfoLocalFile(pageNumber,string(jsonCommutity))

	//workItemBBS("https://bbs.haodiaoyu.com/thread-264798-1-1.html",264798,"昨日再觅新钓点，今日平台装备上阵")

	n := len(filterName)
	for i:=0;i<n ;i++  {
		bbsID := <-pageChannel
		fmt.Printf("workItemBBS %d Finish \n",bbsID)
	}
}

func main()  {
	var pageNumber int
	fmt.Print("请输入PageNumber（>=1): ")
	fmt.Scan(&pageNumber)
	workDouban(pageNumber)
}