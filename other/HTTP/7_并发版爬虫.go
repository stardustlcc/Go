package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string)(result string, err error)  {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 1024 * 4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.read err=", err)
			break
		}

		result += string(buf[:n])
	}

	return
}

func SpiderPage(i int, page chan int)  {
	url := "https://m.api.haoshiqi.net/product/coupleskulist_v1?udid=FL5vXc2zQaVE_hPHq71oIms1WQiqQrkz&platform=h5&channel=h5&spm=h5&v=4.5.8&terminal=wap&swidth=3440&sheight=1440&zoneId=857&cityId=857&pageLimit=10&needPagination=1&pageNum="+strconv.Itoa(i)
	fmt.Sprintf("正在爬虫第%d页面\n", i )

	result, err:= HttpGet(url)
	if err != nil {
		fmt.Println("http.Get err=", err)
		return
	}

	//把内容写入到文件
	fileName := strconv.Itoa(i)+".json"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err1=", err1)
		return
	}

	f.WriteString(result)
	f.Close()

	page <- i
}

func DoWork(start, end int)  {
	fmt.Printf("正在爬取 %d -- %d 的页面\n", start, end)
	page := make(chan int)
	for i:=start; i<=end; i++ {
		go SpiderPage(i, page)
	}

	for i:= start; i<=end;i++ {
		fmt.Printf("第%d爬虫完成", <-page )
	}
}

func main()  {

	var start, end int
	fmt.Printf("请输入起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=起始页）:")
	fmt.Scan(&end)

	DoWork(start, end)
}
