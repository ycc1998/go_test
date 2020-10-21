package mysql

import (
	"fmt"
	"test/global"
	"time"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
)

type postData struct{
	Name string `json:"name"`
	Content string `json:"content"`
	Tag1 string `json:"tag1"`
	Tag2 string `json:"tag2"`
	Tag3 string `json:"tag3"`
	Tag4 string `json:"tag4"`
	Host string `json:"host"`
	Logo string `json:"logo"`
	Cate1 string `json:"cate1"`
}

func Insert(){
	for{
		select {
		case data := <- global.DataInfoUrl:
			fmt.Println(data)
			//Send_post(data)
		default:
			time.Sleep(1000 * time.Microsecond)
		}
	}
	
}

func Send_post(d *global.DataUrl){
	url := "https://5913.com"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"
	// json
	
	post_data := postData{
		Name:d.Title,
		Content: d.Desc,
		Host:d.Url,
		Logo:d.ImgPath,
	}
	tag := strings.Split(d.Keywords,",");

	for i := 0; i<len(tag);i++{
		switch i{
			case 0:
				post_data.Tag1 = tag[0];
				post_data.Cate1 = tag[0];
			break;
			case 1:
				post_data.Tag2 = tag[1];
			break;
			case 2:
				post_data.Tag3 = tag[2];
			break;
			case 3:
				post_data.Tag4 = tag[3];
			break;
			default:
				break 
		}
	}
	

	contentType := "application/json"

	post_data_json, err := json.Marshal(post_data)

	if err != nil{
		fmt.Println("request API err ：",err)
		return 
	}
	//fmt.Println(string(post_data_json))

	resp, err := http.Post(url, contentType, strings.NewReader(string(post_data_json)))

	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
