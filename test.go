package main

import(
//	"strings"
	"fmt"
	"unicode"
)

func main(){
/*	s := "12 34 5#6"
	a := func(char rune) bool { return unicode.IsLetter(char)}
	s1 := strings.FieldsFunc(s,a)
	for _,c := range s1{
		fmt.Println(c)
	}*/
	a := unicode.IsLetter('a')
	fmt.Println(a)
}
