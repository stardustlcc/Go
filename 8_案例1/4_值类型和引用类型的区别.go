package main

func main()  {
	/**
		值类型，引用类型的区别

		值类型：int,bool,string,float，数组 和 struct
		特点：变量直接指向地址空间的值，申请变量时，自动分配内存，内存被分配在栈中。栈在函数调用完会被自动释放

		引用类型：指针, slice, map, chan
		特点：变量存储的是一个地址，地址存储的是具体的值。需要借助 new 和 make 申请内存，内存通常分配在堆上，通过 GC回收

		new 和 make 的区别
		new可以给所有类型分配内存空间，而make只能给特定的几个类型分配
		new的返回值是该类型的指针，而make的返回值是类型的本身
	 */
}
