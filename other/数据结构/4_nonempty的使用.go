package main

import "fmt"

func nonempty(strings []string) []string  {
	i:=0
	for _, s := range strings  {
		if  s != ""  {
			strings[i] = s
			i++
		}
	}
	return  strings
}

func main()  {
	data := []string{"cici", "", "hello"}
	c := nonempty(data)
	fmt.Println( c)
	fmt.Println( data)
}
