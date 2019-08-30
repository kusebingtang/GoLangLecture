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



type ImageContentDY struct {
	Image string	`json:"image"`
	Content string `json:"content"`
}


type ArticleDetailDY struct {
	Title string	`json:"title"`
	Sorting []string `json:"sorting"`
	Articles []map[string]string `json:"articles"`
	Images []map[string]ImageContentDY `json:"images"`
}


type JiqiaoItem struct {
	Id int	`json:"id"`
	Title string	`json:"title"`
	Introduction string	`json:"introduction"`
	Url string `json:"url"`
	ImageUrl string `json:"imageUrl"`
}


type JiqiaoList struct {
	PageNum int	`json:"pageNum"`
	Content []JiqiaoItem	`json:"Content"`
}


func HttpGetDYRM(url string)(result string, err error)  {

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

func saveJiQiaoDiaoyuLocalFile(index,category int,jsonString string)  {

	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/jiqiao/diaoyu/diaoyu_"+strconv.Itoa(index)+".json"
	switch category {
	case 1:
		path ="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/jiqiao/diaoyu/diaoyu_"+strconv.Itoa(index)+".json"
		break
	case 2:
		path ="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/jiqiao/diaoyu/diaoji_"+strconv.Itoa(index)+".json"
		break
	case 3:
		path ="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/jiqiao/diaoyu/yuju_"+strconv.Itoa(index)+".json"
		break
	case 4:
		path ="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/jiqiao/diaoyu/shuiyu_"+strconv.Itoa(index)+".json"
		break
	default:
		panic(errors.New("category  Error !!!!"))
	}

	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(jsonString)
}

func saveJiQiaoDiaoyuItemLocalFile(index int,jsonString string)  {

	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/haodiaoyu/jiqiao/diaoyu/item/DY_Item_"+strconv.Itoa(index)+".json"

	file,err := os.Create(path)
	if err != nil {
		return
	}
	defer  file.Close()
	file.WriteString(jsonString)
}



func workDiaoYuRuMenItem(url string,itemID int,contextTitle string,pageChannel chan int )  {

	result,err := HttpGetDYRM(url)
	if err != nil {
		println("workDiaoYuRuMenItem ==> Http Get Error!",err)
		return
	}
	//fmt.Println("workDiaoYuRuMenItem result=>", result)

	regexp0 := regexp.MustCompile(`<div class="content">(?s:(.*?))<div class="view-operate">`)
	filterContent :=  regexp0.FindAllStringSubmatch(result,1)

	//println(filterContent[0][0])
	//fmt.Println()

	regexp1 := regexp.MustCompile(`<p>(.*?)</p>`)
	regexp2 := regexp.MustCompile(`<p align="center"><img alt="(.*?)</p>`)
	regexp3 := regexp.MustCompile(`src="(.*?)" width="700" /></p>`)
	regexp4 := regexp.MustCompile(`src="(.*?)" />`)
	regexp5 := regexp.MustCompile(`<h2>(.*?)</h2>`)
	regexp6 := regexp.MustCompile(`src="(.*?)" alt="" />`)

	contentStr := filterContent[0][1]
	contentStr = strings.Trim(contentStr,"\r\n")
	contentStr =strings.TrimSpace(contentStr)
	//fmt.Println(contentStr)

	var goNext = true

	var (
		artCount = 0
		imageCount = 0
	)

	articleDetail:= ArticleDetailDY{
		Title:contextTitle,
	}
	var sorting []string
	var articles []map[string]string
	var imagesSlice []map[string]ImageContentDY


	for goNext  {

		if strings.HasPrefix(contentStr, "<p><img") {
			filterStr :=  regexp1.FindAllStringSubmatch(contentStr,1)
			contentStr = strings.Replace(contentStr, filterStr[0][0], "", -1)
			filterURL :=  regexp6.FindAllStringSubmatch(filterStr[0][0],1)

			if len(filterURL)==0 {
				fmt.Println(contentStr)
				panic(errors.New("url get Error!==>"+url))
			}

			ic :=ImageContentDY{
				Image:filterURL[0][1],
				Content:"",
			}
			imageCount++
			icMap := make(map[string]ImageContentDY)
			imageShowTag := "imageShow_"+strconv.Itoa(imageCount)
			sorting = append(sorting,imageShowTag)
			icMap[imageShowTag] = ic
			imagesSlice = append(imagesSlice, icMap)
		}else if strings.HasPrefix(contentStr, "<p>") {
			filterStr :=  regexp1.FindAllStringSubmatch(contentStr,1)
			//println(filterStr[0][1])
			contentStr = strings.Replace(contentStr, filterStr[0][0], "", -1)


			contentJson := strings.Replace(filterStr[0][1], "&ldquo;", "“", -1)
			contentJson = strings.Replace(contentJson, "&rdquo;", "”", -1)
			artCount++
			tag := "art"+strconv.Itoa(artCount)
			sorting = append(sorting,tag)
			map2 := make(map[string]string)
			map2[tag] = contentJson
			articles = append(articles, map2)

		}else if strings.HasPrefix(contentStr, `<p align="center">`) {
			filterStr2 :=  regexp2.FindAllStringSubmatch(contentStr,1)
			filterURL :=  regexp3.FindAllStringSubmatch(contentStr,1)
			filterURL2 :=  regexp4.FindAllStringSubmatch(contentStr,1)
			//println(filterURL[0][1])
			if len(filterURL)+len(filterURL2) == 0 {
				fmt.Println(contentStr)
				panic(errors.New("url get Error!==>"+url))
			}
			var imageUrl string
			if len(filterURL)>0 {
				imageUrl = filterURL[0][1]
			}else if len(filterURL2)>0 {
				imageUrl = filterURL2[0][1]
			}

			if len(filterStr2)==0 {
				fmt.Println(contentStr)
				panic(errors.New("url get Error!==>"+url))
			}
			contentStr = strings.Replace(contentStr, filterStr2[0][0], "", -1)

			ic :=ImageContentDY{
				Image:imageUrl,
				Content:"",
			}
			imageCount++
			icMap := make(map[string]ImageContentDY)
			imageShowTag := "imageShow_"+strconv.Itoa(imageCount)
			sorting = append(sorting,imageShowTag)
			icMap[imageShowTag] = ic
			imagesSlice = append(imagesSlice, icMap)

		}else if strings.HasPrefix(contentStr, `<h2>`) {
			filterStr :=  regexp5.FindAllStringSubmatch(contentStr,1)
			contentStr = strings.Replace(contentStr, filterStr[0][0], "", -1)

			contentJson := strings.Replace(filterStr[0][1], "&ldquo;", "“", -1)
			contentJson = strings.Replace(contentJson, "&rdquo;", "”", -1)
			artCount++
			tag := "art"+strconv.Itoa(artCount)
			sorting = append(sorting,tag)
			map2 := make(map[string]string)
			map2[tag] = contentJson
			articles = append(articles, map2)

		} else {
			fmt.Println(contentStr)
			index_p := strings.Index(contentStr,"</p>")
			indexP := strings.Index(contentStr,"<p>")
			if index_p>0 && indexP>0 && index_p<indexP{
				contentStr="<p>"+contentStr  //特殊处理
			} else
			{
				panic(errors.New("Not <p>or<p align=\"center\">  Error !!!!"+url))
			}
		}
		contentStr = strings.TrimSpace(contentStr)
		contentStr = strings.Trim(contentStr,"\r\n")
		contentStr = strings.TrimSpace(contentStr)
		if contentStr=="</div>" {
			goNext = false
		}
	}

	articleDetail.Articles = articles
	articleDetail.Sorting = sorting
	articleDetail.Images = imagesSlice
	str, err := json.Marshal(articleDetail)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json", string(str))
	saveJiQiaoDiaoyuItemLocalFile(itemID,string(str))
	pageChannel<-itemID
}


func workDiaoYuRuMen(pageNumber,category int)  {
	url := "https://www.haodiaoyu.com/jiqiao/diaoyu/"
	switch category {
	case 1:
		url = "https://www.haodiaoyu.com/jiqiao/diaoyu/"
		break
	case 2:
		url = "https://www.haodiaoyu.com/jiqiao/diaoji/"
		break
	case 3:
		url = "https://www.haodiaoyu.com/jiqiao/yuju/"
		break
	case 4:
		url = "https://www.haodiaoyu.com/jiqiao/shuiyu/"
		break
	default:
		panic(errors.New("category  Error !!!!"))
	}

	result,err := HttpGetDYRM(url)
	if err != nil {
		println("Http Get Error!",err)
		return
	}

	pageChannel := make(chan  int)

	jiqiaoList := JiqiaoList{
		PageNum:pageNumber,
	}

	var itemList []JiqiaoItem


	//fmt.Println("==========>result:",result)

	regexp0 := regexp.MustCompile(`<h3><a target="_blank"(.*?)</a></h3>`)
	filterContent :=  regexp0.FindAllStringSubmatch(result,-1)
	length := len(filterContent)


	regexp1 := regexp.MustCompile(`<p class="cnt">(?s:(.*?))</p>`)
	filterCnt :=  regexp1.FindAllStringSubmatch(result,length)

	regexpTitle := regexp.MustCompile(`target="_blank">(.*?)</a></h3>`)
	regexpID := regexp.MustCompile(`<a target="_blank" href="https://www.haodiaoyu.com/view/(.*?).html" target="_blank">`)
	regexpURL := regexp.MustCompile(`<a target="_blank" href="(.*?)" target="_blank">`)
	regexpImageURL := regexp.MustCompile(`data-src="(.*?)" alt="`)

	filterImageURL := regexpImageURL.FindAllStringSubmatch(result,length)

	lengthCount := 0

	for i:=0;i<length ;i++  {
		strFind := filterContent[i][0]
		strCnt := filterCnt[i][1]

		filterTitle := regexpTitle.FindAllStringSubmatch(strFind,1)
		filterID := regexpID.FindAllStringSubmatch(strFind,1)
		filterURL := regexpURL.FindAllStringSubmatch(strFind,1)

		//fmt.Println(filterTitle[0][1],"   |   ",strCnt,filterID[0][1],filterURL[0][1])
		//fmt.Println("------------------------------------------------")
		//fmt.Println(filterImageURL[i][1])

		filterIDValue, _ := strconv.Atoi(filterID[0][1])

		jiqiaoItem := JiqiaoItem{
			Title : filterTitle[0][1],
			Introduction : strCnt,
			Url:filterURL[0][1],
			Id:filterIDValue,
			ImageUrl:filterImageURL[i][1],
		}
		itemList = append(itemList,jiqiaoItem)

		urlStr := filterURL[0][1]

		if urlStr =="https://www.haodiaoyu.com/view/44967.html" {

		}else {
			lengthCount = lengthCount+1
			go workDiaoYuRuMenItem(urlStr,filterIDValue,filterTitle[0][1],pageChannel)
		}


	}
	jiqiaoList.Content = itemList

	jsonCommutity, err := json.Marshal(jiqiaoList)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonCommutity))
	saveJiQiaoDiaoyuLocalFile(pageNumber,category,string(jsonCommutity))

	for i:=0;i<lengthCount ;i++  {
		bbsID := <-pageChannel
		fmt.Printf("workDiaoYuRuMenItem %d Finish \n",bbsID)
	}
}


func main()  {
	var pageNumber ,category  int
	fmt.Print("请输入PageNumber（>=1): ")
	fmt.Scan(&pageNumber)
	fmt.Println("钓鱼秘籍 请输入:[1] ","     钓鱼大全 请输入:[2] ","      渔具技巧 请输入:[3]","      水域绝招 请输入:[4]")
	fmt.Scan(&category)
	workDiaoYuRuMen(pageNumber,category)
	//pageChannel := make(chan  int)
	//workDiaoYuRuMenItem("https://www.haodiaoyu.com/view/44310.html",44310,"夏天水库野钓的这三个黄金位，爆护渔获",pageChannel)


}