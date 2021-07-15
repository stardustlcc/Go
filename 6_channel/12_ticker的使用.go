package main

import (
	"fmt"
	"time"
)

func main()  {

	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for {
		<-ticker.C
		i++
		if i == 5 {
			ticker.Stop()
			break
		}
		fmt.Println(i)
	}
}