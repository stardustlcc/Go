package main

import (
	"fmt"
	"regexp"
)

func main()  {

	// ``原生字符串，原来怎么样就怎么样
	str := `
		<html>
			<h1>你好</h1>
			<div>
				<div>nihao 111111</div>
			</div>
			<div>
				<div>nihao 2222222</div>
			</div>
			<div>
				<div>
					nihao 333333
					44444444
					5555555
				</div>
			</div>
		</html>
	`

	//reg := regexp.MustCompile(`<div>(.*)</div>`)
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	if reg == nil {
		fmt.Println("err=", reg)
		return
	}

	result := reg.FindAllStringSubmatch(str, -1)
	//fmt.Println("result=", result)

	//过滤 <> </>

	for _, data := range result {
		//fmt.Println("data[0]", data[0])
		fmt.Println("data[1]", data[1])

	}


}
