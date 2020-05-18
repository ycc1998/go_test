package chromedp

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"io/ioutil"
	"bytes"
	"math/rand"
	//"sync"
	"time"
	"test/global"
	//"fmt"
	
	
)
var r *rand.Rand

func init() {
    r = rand.New(rand.NewSource(time.Now().Unix()))
}

func Run(d *global.DataUrl)(err error){
	var buffer bytes.Buffer

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	
	ctx, cancel = context.WithTimeout(ctx, 45*time.Second)
	defer cancel()

	var  img []byte

	err = chromedp.Run(ctx,
        chromedp.Emulate(device.Reset),
        chromedp.EmulateViewport(1500, 1024),
		chromedp.Navigate(d.Url),
		chromedp.InnerHTML(`title`, &d.Title,  chromedp.ByQuery),

		chromedp.Evaluate(`
			var key = document.getElementsByName("keywords");
			if(key.length > 0){
				key[0].content;
			}else{
				var key2 = document.getElementsByName("Keywords");
				if(key2.length > 0){
					 key2[0].content;
				}else{
					 '';
				}
			}`,
		&d.Keywords),
		chromedp.Evaluate(`
			var key = document.getElementsByName("description");
			if(key.length > 0){
				key[0].content;
			}else{
				var key2 = document.getElementsByName("Description");
				if(key2.length > 0){
					 key2[0].content;
				}else{
					 '';
				}
			}`,
		&d.Desc),
		//chromedp.OuterHTML(`meta[name="keywords"]`, &key,chromedp.ByQuery),
		//chromedp.OuterHTML(`meta[name="description"]`, &desc,chromedp.ByQuery),
		
		chromedp.CaptureScreenshot(&img),
		// chromedp.JavascriptAttribute(`meta[name="keywords"]`,`content`, &d.Keywords,chromedp.ByQuery),
		// chromedp.JavascriptAttribute(`meta[name="description"]`, `content`,&d.Desc,chromedp.ByQuery),
		
    );

	if err != nil {
		return
	}


	if bytes.Count(img,nil) > 10{
		buffer.WriteString("./img/")
		buffer.WriteString(RandString(10))
		buffer.WriteString(".png")

		fileName := buffer.String()		
			
		err = ioutil.WriteFile(fileName, img, 0777);

		if err != nil {
			return
		}

		d.ImgPath = fileName
		global.DataInfoUrl <- d
		return
	}
	return
}


// RandString 生成随机字符串
func RandString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}