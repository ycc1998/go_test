package main

import (
	"fmt"
	"sync"
	"test/chromedp"
	"test/global"
	"test/mysql"
	"test/tail"
	"time"
)

func main() {
	global.Init()

	err := tail.Init("./test.log")
	if err != nil {
		fmt.Println("tail init:", err)
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	for i := 1; i <= 4; i++ {
		go mysql.Insert()
		go run()
	}

	wg.Wait()

	// go func() {
	// 	for {
	// 		select {
	// 		case l := <-tail.AllUrl:
	// 			fmt.Println(l)
	// 		}
	// 	}
	// }()

	// time.Sleep(50 * time.Second)

}

func run() {
	for {
		select {
		case url := <-global.AllUrl:
			if len(url) <= 0 {
				continue
			}

			d := &global.DataUrl{
				Url: url,
			}
			err := chromedp.Run(d)
			if err != nil {
				fmt.Println("err run:", err)
			}
		default:
			time.Sleep(100 * time.Microsecond)
		}

	}
}
