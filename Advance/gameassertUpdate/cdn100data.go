package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main()  {

	var(
		content []byte
		err error
	)
	//1读取文件内容
	if content,err = ioutil.ReadFile("/Users/zyh/GolandProjects/GoLangLecture/Advance/gameassertUpdate/update.txt"); err != nil {
		fmt.Println(err.Error())
		return
	}

	contentStr := string(content)
	//fmt.Println(string(content))

	regexp0 := regexp.MustCompile(`http://cdn.game100.cn/unity_games/release/98BY/(.*?)\?time`)
	filterContent :=  regexp0.FindAllStringSubmatch(contentStr,-1)
	length := len(filterContent)

	fmt.Println(length)

	for i:=0;i<length ;i++  {
		fmt.Println(filterContent[i][0],"     ",filterContent[i][1])
	}
}