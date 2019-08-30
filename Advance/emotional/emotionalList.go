package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/djimenez/iconv-go"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ArticleDetail struct {
	Title string	`json:"title"`
	Images []string	`json:"images"`
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

func HttpGetContent(url string)(result string, err error)  {

	var response *http.Response

	if response,err = http.Get(url); err!=nil {
		return
	}

	defer response.Body.Close()


	input, err := ioutil.ReadAll(response.Body)
	out := make([]byte, len(input))
	out = out[:]

	iconv.Convert(input, out, "gb2312", "utf-8")
	result += string(out[:])

	return
}

func workConstellationsItem(url string,bbsID string,contextTitle string )(err error) {

	fmt.Println(url)
	result,err := HttpGetContent(url)
	if err != nil {
		println("workItemBBS ==> Http Get Error!",err)
		return
	}

	//fmt.Println(result)

	regexp1 := regexp.MustCompile(`<!--文章内容-->(?s:(.*?))<!--文章内容end-->`)
	if len(regexp1.FindAllStringSubmatch(result, 1)) ==0 {
		err = errors.New("Get Content Error!==>"+url)
		return
	}
	content := regexp1.FindAllStringSubmatch(result, 1)[0][1]

	//fmt.Println(content)

	content = strings.Replace(content, "</a>", "", -1)
	regexpHTMLA := regexp.MustCompile(`<a(.*?)target="_blank">`)
	filterHTMLA := regexpHTMLA.FindAllStringSubmatch(content,-1)

	for len(filterHTMLA)!=0 {
		for i:=0; i<len(filterHTMLA) ;i++  {
			aStr := filterHTMLA[0][0]
			content = strings.Replace(content, aStr, "", 1)
		}
		filterHTMLA = regexpHTMLA.FindAllStringSubmatch(content,-1)
	}

	regexpHTMLA2 := regexp.MustCompile(`<a(.*?)target="_blank" class="cmsLink">`)
	filterHTMLA2 := regexpHTMLA2.FindAllStringSubmatch(content,-1)
	for len(filterHTMLA2)!=0 {
		for i:=0; i<len(filterHTMLA2) ;i++  {
			aStr := filterHTMLA2[0][0]
			content = strings.Replace(content, aStr, "", 1)
		}
		filterHTMLA2 = regexpHTMLA2.FindAllStringSubmatch(content,-1)
	}

	//println(content)

	imagesList := make([]string, 0)
	regexpImages := regexp.MustCompile(`<p style="text-align: center;">(.*?)</p>`)
	filterImages := regexpImages.FindAllStringSubmatch(content,1)
	regexpImage := regexp.MustCompile(`src="(.*?)" style="`)
	if len(filterImages)!=0 {
		for len(filterImages)!=0 {
			strImage :=  filterImages[0][0]
			filterImage := regexpImage.FindAllStringSubmatch(strImage,1)
			if len(filterImage) ==1 {
				imagesList = append(imagesList,"https:"+filterImage[0][1])
			}
			//fmt.Println(strImage)
			content = strings.Replace(content, strImage, "", 1)
			filterImages = regexpImages.FindAllStringSubmatch(content,1)
		}
	}else  {
		regexpImages := regexp.MustCompile(`<p style="text-align: center">(.*?)</p>`)
		filterImages := regexpImages.FindAllStringSubmatch(content,1)
		for len(filterImages)!=0 {
			strImage :=  filterImages[0][0]
			filterImage := regexpImage.FindAllStringSubmatch(strImage,1)
			if len(filterImage) ==1 {
				imagesList = append(imagesList,"https:"+filterImage[0][1])
			}
			//fmt.Println(strImage)
			content = strings.Replace(content, strImage, "", 1)
			filterImages = regexpImages.FindAllStringSubmatch(content,1)
		}
	}


	//println(content)

	regexpContent := regexp.MustCompile(`<p><strong>(.*?)</strong></p>`)
	regexpContent2 := regexp.MustCompile(`<p>(.*?)</p>`)


	articleList := make([]string, 0)

	goNext := true

	for goNext  {
		//println("------",content)
		var (
				strContent, str string
			)
		indexBR := strings.Index(content,"<p><strong>")
		indexPP := strings.Index(content,"<p>")

		if indexBR!=-1 && indexPP!=-1 && indexBR == indexPP {
			filterContent := regexpContent.FindAllStringSubmatch(content,1)
			strContent =  filterContent[0][0]
			str = filterContent[0][1]
		}else if indexBR!=-1 && indexPP!=-1 && indexBR > indexPP {
			filterContent2 := regexpContent2.FindAllStringSubmatch(content,1)
			strContent =  filterContent2[0][0]
			str = filterContent2[0][1]
		}else if indexBR==-1 && indexPP!=-1 {
			filterContent2 := regexpContent2.FindAllStringSubmatch(content,1)
			if len(filterContent2) ==0 {
				content = strings.Trim(content,"\r\n")
				if content=="<p>" {
					break
				}
				println(content)
				//panic(errors.New(" Error !!!!"+url))
				err = errors.New(" Error !!!!"+url)
				return
			}
			strContent =  filterContent2[0][0]
			str = filterContent2[0][1]
		}else {
			println(indexBR,indexPP)
			content = strings.Trim(content,"\r\n")
			println("====",content,"===")
			goNext = false
		}
		//println(str)
		articleList = append(articleList,str)
		content = strings.Replace(content, strContent, "", 1)

		if len(content) == 0 {
			goNext = false
		}
		//println(content)
	}


	articleDetail := ArticleDetail{}
	articleDetail.Title = contextTitle
	articleDetail.Images = imagesList
	articleDetail.Articles = articleList

	fmt.Println(articleDetail)
	fmt.Println("--------------------------------------")

	jsonCommutity, err := json.Marshal(articleDetail)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/emotional/Article/ArticleDetail/art_"+bbsID+".json"
	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(string(jsonCommutity))

	return nil
}

func workConstellations(pageNumber int,category int) {

//https://emotion.pclady.com.cn/raiders1/
	url := "https://emotion.pclady.com.cn/raiders/"  //情感攻略
	if pageNumber > 1 {
		url = "https://emotion.pclady.com.cn/raiders/index_"+strconv.Itoa(pageNumber-1)+".html"
	}

	if category==2 {//https://emotion.pclady.com.cn/000028934/  婚姻家庭
		url = "https://emotion.pclady.com.cn/000028934/"
		if pageNumber > 1 {
			url = "https://emotion.pclady.com.cn/000028934/index_"+strconv.Itoa(pageNumber-1)+".html"
		}
	}

	if category==3 {//职场解读
		url = "https://emotion.pclady.com.cn/raiders1/"
		if pageNumber > 1 {
			url = "https://emotion.pclady.com.cn/raiders1/index_"+strconv.Itoa(pageNumber-1)+".html"
		}
	}

	result, err := HttpGetContent(url)
	if err != nil {
		println("Http Get Error!", err)
		return
	}
	//println(result)
	fmt.Printf("========第 %d 页 抓取成功，开始分析页面============\n", 1)



	regexp1 := regexp.MustCompile(`<ul class="clearfix">(?s:(.*?))</ul>`)
	filterName := regexp1.FindAllStringSubmatch(result, 1)

	//println(len(filterName))
	resultContent := filterName[0][0]

	//println(resultContent)
	regexpLi := regexp.MustCompile(`<li>(?s:(.*?))</li>`)
	filterLi := regexpLi.FindAllStringSubmatch(resultContent, -1)

	regexpURL := regexp.MustCompile(`<a href="(.*?)">`)
	regexpImage := regexp.MustCompile(`<img src="(.*?)" alt="`)
	regexpTitle := regexp.MustCompile(`target="_blank">(.*?)</a></em><br/>`)
	regexpDesc := regexp.MustCompile(`<span class="sDes">(?s:(.*?))<a href="`)
	regexpID := regexp.MustCompile(`<a href="//emotion.pclady.com.cn/(.*?).html">`)


	length := len(filterLi)
	articleList := make([]Article, 0)
	for i:=0; i<length ;i++  {
		strLiContent := filterLi[i][0]

		filterURL := regexpURL.FindAllStringSubmatch(strLiContent,1)
		filterImage := regexpImage.FindAllStringSubmatch(strLiContent,1)
		filterTitle := regexpTitle.FindAllStringSubmatch(strLiContent,1)
		filterDesc := regexpDesc.FindAllStringSubmatch(strLiContent,1)
		filterID := regexpID.FindAllStringSubmatch(strLiContent,1)

		article := Article{}
		article.Image = "https:"+filterImage[0][1]
		article.Url = "https:"+filterURL[0][1]
		article.Title = filterTitle[0][1]
		article.Desc = strings.Trim(filterDesc[0][1],"\r\n")
		article.Id = strings.Replace(filterID[0][1], "/", "_", 1)
		//fmt.Println(filterURL[0][1],filterImage[0][1],filterTitle[0][1],filterDesc[0][1],filterID[0][1],)

		//fmt.Println(article)

		err := workConstellationsItem(article.Url,article.Id,article.Title)
		if err == nil {
			articleList = append(articleList,article)
		}
		//
	}

	//workConstellationsItem("https://emotion.pclady.com.cn/182/1828113.html","184/1842662","遭遇职场潜规则")

	commutity := Commutity{}
	commutity.PageNum = pageNumber
	commutity.Articles = articleList

	//Marshal失败时err!=nil
	jsonCommutity, err := json.Marshal(commutity)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonCommutity))

	var path string
	if category==1 {
		path = "/Users/zyh/GolandProjects/GoLangLecture/Advance/emotional/Article/QJGL_"+strconv.Itoa(pageNumber)+".json"
	}else if category==2 {
		path = "/Users/zyh/GolandProjects/GoLangLecture/Advance/emotional/Article/HYJT_"+strconv.Itoa(pageNumber)+".json"
	}else if category==3 {
		path = "/Users/zyh/GolandProjects/GoLangLecture/Advance/emotional/Article/ZCJD_"+strconv.Itoa(pageNumber)+".json"
	}

	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(string(jsonCommutity))

}



func main()  {
	var pageNumber, category int
	fmt.Print("请输入PageNumber（>=1): ")
	fmt.Scan(&pageNumber)
	fmt.Println("情感攻略 请输入:[1] ","     婚姻家庭 请输入:[2] ","      职场解读 请输入:[3]","     ")
	fmt.Scan(&category)
	workConstellations(pageNumber,category)
}