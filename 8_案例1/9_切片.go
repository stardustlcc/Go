package main

import "fmt"

func main()  {
	sli1 := []int{1,2,3,4,5,6,7}
	fmt.Println("sli1=", sli1)

	sli2 := make([]int,5,10)
	sli2[0]=1
	sli2[1]=2
	sli2[2]=1
	sli2[3]=2
	sli2 = append(sli2,555)
	sli2 = append(sli2,555)
	sli2 = append(sli2,555)
	fmt.Println("sli2=", sli2)

	test(sli2)
	fmt.Println("sli2=", sli2)



	arr := [9]int{1,2,3,4,5,6,7,8}
	sli3 := arr[2:6]//3,4,5,6
	sli3[0] = 100
	//sli3 = append(sli3,200)
	//sli3 = append(sli3,200)
	//sli3 = append(sli3,200)
	//sli3 = append(sli3,200)
	sli3 = append(sli3,200)

	//sli3[6] = 4
	sli3[0]=2000
	arr[3]=100000
	fmt.Println("sli3=", sli3, "cap=", cap(sli3))
	fmt.Println("arr=", arr)


	var sli4 []int
	if sli4 == nil {
		fmt.Println("sli4==nil")
	}
	sli4 = make([]int, 2)
	copy(sli4,sli3)
	fmt.Println(sli3)
	fmt.Println(sli4)
	sli4[0]=1
	sli4[1]=1
	sli3[0]=222
	fmt.Println(sli3)
	fmt.Println(sli4, "cap=", cap(sli4))

	for i:=0; i<2; i++{
		fmt.Println(sli4[i])
	}

	for _, v := range sli4 {
		fmt.Println(v)
	}

	str := "hello,world"
	sli5 := str[1:5]
	fmt.Println(sli5)
	fmt.Printf("%T", sli5)

	s := []byte(sli5)
	s[0] = 'H'
	fmt.Println(string(s))

	str1 := "你好,hello"
	s1 := []rune(str1)
	fmt.Println(string(s1))
	s1[0] = '我'
	fmt.Println(string(s1))

	sli6 := make([]string,2,9)
	sli6[0] = "你好"
	sli6[1] ="我"
	sli6 = append(sli6,"我")
	fmt.Println("sli6=", sli6)

	sli7 := []int{1,2,3,4,5,}

}

func test(sli []int)  {
	fmt.Println(sli)
	sli[1] = 10000
}
