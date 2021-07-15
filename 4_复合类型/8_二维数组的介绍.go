package main

import "fmt"

func main()  {

	/**
		有多少[]就是多少维
	 */
	var a [3][4]int
	k := 0
	for i:=0;i<3; i++ {
		for j:=0;j<4 ;j++  {
			k++
			a[i][j] = k
		}
	}

	fmt.Println("a=", a)

	//部分初始化，没有初始化的元素取默认值
	b :=[3][4]int{{1,1,1},{2,2,2,2},{3,3,3,3}}
	fmt.Println("b=", b)

}
