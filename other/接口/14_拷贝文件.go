package 接口

import (
	"fmt"
	"io"
	"os"
)

func main()  {

	list := os.Args

	if len(list) != 3 {
		fmt.Println(1111)
		return
	}

	srcFile := list[1]
	dstFile := list[2]
	if srcFile == dstFile {
		fmt.Println("名字不能相同")
		return
	}

	//打开源文件
	sf, err1 := os.Open(srcFile)
	if err1 != nil {
		fmt.Println("err1=", err1)
		return
	}

	//新建目标文件
	df, err2 := os.Create(dstFile)
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}

	defer sf.Close()
	defer df.Close()

	//从源文件往目的文件中写，读多少，写多少
	buf := make([]byte, 1024*4)

	for {
		n, err := sf.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err=", err)
		}

		df.Write(buf[:n])
	}





}
