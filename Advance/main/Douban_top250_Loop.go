package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func HttpGetDouPan2(url string)(result string, err error)  {

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

func saveInfoLocalFile2(index int,filterName,filterScore,filterCount [][]string)  {

	path :="/Users/zyh/GolandProjects/GoLangLecture/Advance/main/douban_index_"+strconv.Itoa(index)+".txt"

	file,err := os.Create(path)
	if err != nil {
		return
	}

	defer  file.Close()

	n := len(filterName)

	file.WriteString("电影名称"+"\t\t\t\t\t\t"+"评分"+"\t\t\t\t\t\t"+"评分人数"+"\n")

	for i:=0;i<n ;i++  {
		file.WriteString(filterName[i][1]+"\t\t\t\t\t\t"+filterScore[i][1]+"\t\t\t\t\t\t"+filterCount[i][1]+"\n")
	}

}

func spidePageDB(index int,pageChannel chan int)  {
	url := "https://movie.douban.com/top250?start="+strconv.Itoa((index-1)*25)+"&filter="

	fmt.Println(url)

	result,err := HttpGetDouPan2(url)
	if err != nil {
		println("Http Get Error!",err)
		return
	}

	fmt.Printf("第 %d 页 抓取成功，开始分析页面\n",index)
	//fmt.Println(result)

	regexp1 := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)

	filterName :=  regexp1.FindAllStringSubmatch(result,-1)
	//for _,nameList := range filterName {
	//	println(nameList[1])
	//}

	regexp2 := regexp.MustCompile(`<span class="rating_num" property="v:average">(.*?)</span>`)
	filterScore :=  regexp2.FindAllStringSubmatch(result,-1)
	//for _,filterScoreList := range filterScore {
	//	println(filterScoreList[1])
	//}


	regexp3 := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	filterCount :=  regexp3.FindAllStringSubmatch(result,-1)
	//for _,pCount := range filterCount {
	//	println(pCount[1])
	//}

	saveInfoLocalFile2(index,filterName,filterScore,filterCount)

	pageChannel <- index
}



func workDouban2(start, end int)  {

	fmt.Printf("开始抓取 %d 到 %d 页面的数据  \n",start,end)
	pageChannel := make(chan  int)
	for i := start; i<=end; i++ {
		go spidePageDB(i,pageChannel)

	}

	for i := start; i<=end; i++ {
		page := <-pageChannel
		fmt.Printf("第 %d 输入分析保存success",page)
	}
}




func main()  {

	var (
		start, end int
	)

	fmt.Print("请输入抓取的起始页（>=1): ")
	fmt.Scan(&start)

	fmt.Print("请输入抓取的结束页（>=start): ")
	fmt.Scan(&end)

	workDouban2(start,end)
}