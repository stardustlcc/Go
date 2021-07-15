package main

import "fmt"

func main()  {

	/**
		数组
		指针
		切片
		map
		结构体struct

		参考资料：http://www.topgoer.com/
	*/

	//数组
	arr := [4]int{1,2,3,4}
	fmt.Println("arr=", arr)

	arr2 := [3][2]string{{"aa","bb"},{"cc","dd"},{"ee","ff"}}
	fmt.Println("arr2=", arr2)

	arr3 := [...]string{"aaaa", "bbbb", "cccc"}
	fmt.Println("arr3=", arr3)


	//指针
	//操作符 "&" 取变量地址，"*" 通过指针访问目标对象
	var p int
	p = 100
	a := &p
	*a = 200
	fmt.Println("p=", p)
	fmt.Println("a=", *a)


	//切片
	slip1 := []string{"aaaa", "bbbb"}
	fmt.Println("slip1=", slip1)
	slip1 = append(slip1,"cccc")
	fmt.Println("slip1=", slip1)

	// 通过make 申请内存空间
	slip2 := make([]int, 4, 6)
	fmt.Println("slip2=", slip2,"len(slip2)=",len(slip2),  "cap(slip2)=", cap(slip2))
	slip2[1] = 111
	slip2[3] = 333
	// 同append 可以动态扩容切片空间
	slip2 = append(slip2,4444)
	slip2 = append(slip2,4444)
	slip2 = append(slip2,4444)
	slip2 = append(slip2,4444)
	slip2 = append(slip2,4444)
	fmt.Println("slip2=", slip2, "len(slip2)=",len(slip2),"cap(slip2)=", cap(slip2))
	//从切片中截取，前闭后开
	slip3 := slip2[1:5]
	fmt.Println("slip3=", slip3)


	//map
	var map1 map[int]string
	map1 = make(map[int]string)
	map1[1] = "aaa"
	map1[2] = "bbb"
	fmt.Println("map1=", map1)

	map2 := map[string]int{"a":100,"b":200}
	fmt.Println("map2=", map2)
	map2["c"] = 100
	fmt.Println("map2=", map2)
	//map的遍历
	for key, val := range map2 {
		fmt.Println("key=", key, "val=", val)
	}
	//map的删除
	delete(map2, "c")
	fmt.Println("map2=", map2)


	//结构体
	var p1 Person
	p1 = Person{1, "cici", 18, 1}
	fmt.Println("p1=", p1)
	p1.name = "make"
	fmt.Println("p1=", p1)

	p2 := Person{2,"nite", 11, 2}
	fmt.Println("p2=", p2)

	var p3 Person
	p3.id = 4
	fmt.Println("p3=", p3)

	p4 := Student{Person{1, "CICI", 18, 1}, 10001, "sh"}
	fmt.Println("p4=", p4)
	fmt.Printf("%+v\n", p4)
	// 通过new 申请内存空间
	p5 := new(Student)
	p5.Person.id = 1
	p5.Person.name = "cici"
	p5.adder = "sh"
	p5.classId = 10003
	fmt.Println("p5=", *p5)


	mapInfo := make(map[string]Student)
	stu1 := Student{Person{1, "CICI", 18, 1}, 20, "SH"}
	mapInfo["cici"] = stu1
	fmt.Println("mapInfo=", mapInfo)
	fmt.Println("mapInfo[cici]=", mapInfo["cici"])

}

type Person struct {
	id int
	name string
	age int
	sex byte
}

type Student struct{
	Person
	classId int
	adder string
}


