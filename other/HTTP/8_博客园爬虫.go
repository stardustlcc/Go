package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func HttpGet( url string) (result string, err error)  {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 4*1024)
	for {
		n, _:= resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}

	return
}

func SpiderOneBook(url string) (title, content string, err error)  {
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}

	//取关键信息
	re := regexp.MustCompile(`<span>(?s:(.*?))</span>`)
	if re == nil {
		//fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("%s","regexp.MustCompile err")
		return
	}
	//取内容
	tmpTtile := re.FindAllStringSubmatch(result, 1)//只过滤第一个
	for _, data := range tmpTtile {
		title = data[1]
		break
	}


	//取关键信息
	re1 := regexp.MustCompile(`<p>(?s:(.*?))</p>`)
	if re1 == nil {
		//fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("%s","regexp.MustCompile err")
		return
	}
	//取内容
	tmpContent := re1.FindAllStringSubmatch(result, 1)//只过滤第一个
	for _, data := range tmpContent {
		content = data[1]
		break
	}
	return
}

//把内容写到文件中
func StoreToFile(i int,fileTitle, fileContent []string)  {
	//新建文件
	fileName := strconv.Itoa(i)+".txt"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}

	defer f.Close()

	//写内容
	n := len(fileTitle)
	for i:=0; i<n; i++ {
		f.WriteString(fileTitle[i]+"\n")
		f.WriteString(fileContent[i]+"\n")
		f.WriteString("\n-----------------------------------------\n")
	}
}

func SpiderPage(i int, page chan int)  {
	//明确爬取的网址 url
	url := "https://www.cnblogs.com/#p" + strconv.Itoa(i)
	fmt.Printf("正在爬取第%d个网页：%s\n", i, url)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet  err =", err)
		return
	}

	//fmt.Println("r=", result)
	//取博客园url链接
	//解释表达式
	re := regexp.MustCompile(` <a class="post-item-title" href="(?s:(.*?))" target="_blank">`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	//取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1)
	//fmt.Println("joyUrls=", joyUrls)

	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	//取网址
	//第一个返回下标，第二个返回内容
	for _, data := range joyUrls {
		//fmt.Println("url=", data[1])
		//爬取每一个博客内容
		title, content, err := SpiderOneBook(data[1])
		if err != nil {
			fmt.Println("SpiderOneBook err=", err)
			continue
		}
		//fmt.Println("title=", title)
		//fmt.Println("content=", content)

		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)

	}

	//把内容写到文件
	StoreToFile(i,fileTitle, fileContent)

	page<-i
}

func DoWork(start, end int)  {
	fmt.Printf("准备爬取第%d页--%d页的网址\n", start, end)

	page := make(chan int)
	for i:= start; i<=end; i++ {
		//定义一个函数，爬主页面
		go SpiderPage(i, page)
	}

	for i:= start; i<=end; i++ {
		fmt.Printf("第%d个页面爬虫完成\n", <-page)
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
