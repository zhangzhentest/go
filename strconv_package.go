package main

import(
	"fmt"
	"strconv"
)

func main(){
	var ori string = "123456"
	var i int
	var s string
	
	fmt.Println("The Size of ints is:%d\n",strconv.IntSize)

	i,_ = strconv.Atoi(ori)
	fmt.Println("The integer is %d\n",i)
	i = i + 5
	s =strconv.Itoa(i)
	fmt.Println("The new string is %s\n",s)
}
