package main

import "fmt"

func main()  {

	//声明字符类型
	var ch byte
	ch = 97
	fmt.Printf("%c, %d\n", ch,ch)

	//字符，单引号
	ch = 'a'
	fmt.Printf("%c, %d\n", ch, ch)

	//大写转小写，小写转大写
	fmt.Printf("A大写：%d ; a小写 %d\n", 'A', 'a')
	fmt.Printf("A->a %c\n", 'A'+32)
	fmt.Printf("a->A %c\n", 'a'-32)

	// '\' 以反斜杠开头的字符是转移字符，不会输出
	fmt.Printf("hello go%c",'\n')
	fmt.Printf("hello go")
}
