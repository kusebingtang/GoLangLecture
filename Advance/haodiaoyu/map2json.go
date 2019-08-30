package main
// 导入其它的包
import (
	"encoding/json"
	"fmt"
)





func main() {
	map2json2map()
}

func map2json2map() {

	var s0 []map[string]string


	map1 := make(map[string]string)
	map1["art1"] = "hello"

	map2 := make(map[string]string)
	map2["art2"] = "world"

	s0 = append(s0, map1)
	s0 = append(s0, map2)
	//return []byte
	str, err := json.Marshal(s0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json", string(str))



	var icMapSlice []map[string]ImageContent
	icMap := make(map[string]ImageContent)
	ic :=ImageContent{
		Image:"https://p4.diaoyur.cn/group4/M00/08/F0/cjd0iVzBvoT1npPKOlbmHSqN6LI-2.jpg",
		Content:"奔波渔获多多的钓点，一路向鱼路进军",
	}
	icMap["imageShow_1"] = ic
	icMapSlice = append(icMapSlice,icMap)

	str, err = json.Marshal(icMapSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json", string(str))



	//json([]byte) to map
	//map2 := make(map[string]interface{})
	//err = json.Unmarshal(str, &map2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("json to map ", map2)
	//fmt.Println("The value of key1 is", map2["1"])
}