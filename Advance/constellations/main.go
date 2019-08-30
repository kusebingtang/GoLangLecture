package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type ArticleDetail struct {
	Title string	`json:"title"`
	Image string	`json:"image"`
	Articles []string `json:"articles"`
}


type Article struct {
	Id string	`json:"id"`
	Title string	`json:"title"`
	Image string	`json:"image"`
	Url string `json:"url"`
	Desc string `json:"desc"`
	Time string `json:"time"`
}

type Commutity struct {
	PageNum int	`json:"pageNum"`
	Articles [] Article  `json:"content"`
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

func workConstellationsItem(url string,bbsID string,contextTitle string )(err error) {
	fmt.Println(url)
	result,err := HttpGetDouPan(url)
	if err != nil {
		println("workItemBBS ==> Http Get Error!",err)
		return
	}

	regexp1 := regexp.MustCompile(`<div class="common_det_con"><p>(?s:(.*?))</p></div>`)

	content := regexp1.FindAllStringSubmatch(result, 1)[0][0]

	//println(content)

	regexpImage := regexp.MustCompile(`<img src="(.*?)" title="`)

	imageURL := regexpImage.FindAllStringSubmatch(content,1)
	println(imageURL[0][1])

	regexpStart := regexp.MustCompile(`<div class="common_det_con"><p>(?s:(.*?))</p><p>`)
	filterStart := regexpStart.FindAllStringSubmatch(content,1)
	if len(filterStart)==0 {
		err = errors.New("Get Error!==>"+url)
		return
	}


	str1 := regexpStart.FindAllStringSubmatch(content,1)[0][0]
	str2 := `<p style="text-align: center;"><strong><span style="text-align: center; color: rgb(255, 0, 0);">第一星座网原创文章，转载请联系网站管理人员，否则视为侵权。</span></strong></p>`
	str3 :=`</p><p style="text-align: center;"><span style="color: rgb(255, 0, 0);"><strong>第一星座网原创文章，转载请联系网站管理人员，否则视为侵权。</strong></span>`
	content = strings.Replace(content, str1, "", 1)
	content = strings.Replace(content, str2, "", 1)
	content = strings.Replace(content, str3, "", 1)
	content = strings.Replace(content, "</p></div>", "</p><p>", 1)

	regexpFilter2 := regexp.MustCompile(`<a href="https://www.d1xz.net/sx/zonghe/art267678.aspx"(?s:(.*?))</a></p><p>`)
	filterStart2 := regexpFilter2.FindAllStringSubmatch(content,1)
	if len(filterStart2) ==1 {
		start2Str := filterStart2[0][0]
		content = strings.Replace(content, start2Str, "", 1)
	}
	//fmt.Println(content)

	regexpBR := regexp.MustCompile(`</p><p>(.*?)<br/>`)
	regexPP := regexp.MustCompile(`</p><p>(.*?)</p><p>`)

	regexHtmlA := regexp.MustCompile(`<a(.*?)>`)
	regexHtmltextAlign := regexp.MustCompile(`</p><p style="text-align(.*?)/>`)

	articleList := make([]string, 0)

	for len(content)!=0 {
		indexBR := strings.Index(content,"<br/>")
		indexPP := strings.Index(content,"</p><p>")

		//println(indexBR,"      ",indexPP)

		content = "</p><p>"+content

		var filterStr [][]string

		if indexBR!=-1 && indexBR < indexPP {
			filterStr = regexpBR.FindAllStringSubmatch(content,1)
		}else {
			filterStr = regexPP.FindAllStringSubmatch(content,1)
		}

		jsonStr := filterStr[0][1]
		filterHtmlA := regexHtmlA.FindAllStringSubmatch(jsonStr,-1)
		for j:=0;j<len(filterHtmlA) ;j++  {
			htmlAValue := filterHtmlA[0][0]
			jsonStr = strings.Replace(jsonStr, htmlAValue, "", 1)
		}
		filterHtmlTextAlign := regexHtmltextAlign.FindAllStringSubmatch(jsonStr,-1)
		for j:=0;j<len(filterHtmlTextAlign) ;j++  {
			htmlAValue := filterHtmlTextAlign[0][0]
			jsonStr = strings.Replace(jsonStr, htmlAValue, "", 1)
		}
		jsonStr = strings.Replace(jsonStr, "</a>", "", -1)
		println(jsonStr)
		articleList = append(articleList,jsonStr)

		content = strings.Replace(content, filterStr[0][0], "", 1)

		if strings.HasPrefix(content, "</p><p>") {
			content = strings.Replace(content, "</p><p>", "", 1)
			//println(content)
		}
	}
	articleDetail :=  ArticleDetail{

	}
	articleDetail.Title = contextTitle
	articleDetail.Image = imageURL[0][1]
	articleDetail.Articles = articleList

	//Marshal失败时err!=nil
	jsonCommutity, err := json.Marshal(articleDetail)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/constellations/Article/ArticleDetail/art_"+bbsID+".json"
	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(string(jsonCommutity))
	return nil
}


func workConstellations(pageNumber int) {

	url := "https://www.d1xz.net/astro/2019nianyunchengcesuan/"+strconv.Itoa(pageNumber)+"/"
	result, err := HttpGetDouPan(url)
	if err != nil {
		println("Http Get Error!", err)
		return
	}

	fmt.Printf("========第 %d 页 抓取成功，开始分析页面============\n", 1)

	//pageChannel := make(chan int)
	articleList := make([]Article, 0)

	regexp1 := regexp.MustCompile(`<ul class="words_list_ui">(?s:(.*?))</ul>`)
	filterName := regexp1.FindAllStringSubmatch(result, 1)

	//println(len(filterName))
	resultContent := filterName[0][0]

	//println(resultContent)
	regexpLi := regexp.MustCompile(`<li>(?s:(.*?))</li>`)
	filterLi := regexpLi.FindAllStringSubmatch(resultContent, 15)

	regexpTitle := regexp.MustCompile(`alt="(.*?)" /></a>`)
	regexpURL := regexp.MustCompile(`href="(.*?)" class="pic fl"`)
	regexpID := regexp.MustCompile(`/art(.*?).aspx" class="pic fl">`)
	regexpImage := regexp.MustCompile(`<img src="(.*?)" alt="`)
	regexpDesc := regexp.MustCompile(`<div class="txt"><p>(?s:(.*?))</p></div>`)
	regexpInfoTime := regexp.MustCompile(`<p class="infor_time">(.*?)</p>`)

	//println(filterLi[2][0])
	for i:=0;i<15 ;i++  {
		liResult := filterLi[i][0]
		//println(regexpTitle.FindAllStringSubmatch(liResult, 1)[0][1])
		//println(regexpID.FindAllStringSubmatch(liResult, 1)[0][1])
		//println(regexpURL.FindAllStringSubmatch(liResult, 1)[0][1])
		//println(regexpImage.FindAllStringSubmatch(liResult, 1)[0][1])
		//println(regexpDesc.FindAllStringSubmatch(liResult, 1)[0][1])
		//println(regexpInfoTime.FindAllStringSubmatch(liResult, 1)[0][1])

		objUrl := regexpURL.FindAllStringSubmatch(liResult, 1)[0][1]
		objTitle := regexpTitle.FindAllStringSubmatch(liResult, 1)[0][1]
		filterID := regexpID.FindAllStringSubmatch(liResult, 1)
		if len(filterID)==0 {
			continue
		}
		objId :=regexpID.FindAllStringSubmatch(liResult, 1)[0][1]

		err := workConstellationsItem(objUrl,objId,objTitle)

		if err == nil {
			articleObj := Article{
			}
			articleObj.Id = objId
			articleObj.Title = objTitle
			articleObj.Image = regexpImage.FindAllStringSubmatch(liResult, 1)[0][1]
			articleObj.Url = objUrl
			articleObj.Desc = strings.Trim(regexpDesc.FindAllStringSubmatch(liResult, 1)[0][1],"　　")
			articleObj.Time = regexpInfoTime.FindAllStringSubmatch(liResult, 1)[0][1]

			articleList = append(articleList,articleObj)
			fmt.Println("-----------------------#######--------------------------------------")
		} else {
			println(err.Error())
		}






	}

	commutity := Commutity{}
	commutity.PageNum = pageNumber
	commutity.Articles = articleList

	//Marshal失败时err!=nil
	jsonCommutity, err := json.Marshal(commutity)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	//fmt.Println(string(jsonCommutity))

	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/constellations/Article/constellations_"+strconv.Itoa(pageNumber)+".json"
	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(string(jsonCommutity))

	//workConstellationsItem("https://www.d1xz.net/astro/pisces/art273149.aspx","273149","https://www.d1xz.net/astro/pisces/art273149.aspx",nil )
}



func main()  {
	var pageNumber int
	fmt.Print("请输入PageNumber（>=1): ")
	fmt.Scan(&pageNumber)
	workConstellations(pageNumber)
}