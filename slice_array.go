package main

import(
	"fmt"
)

func main(){
	a := [...]int{1,2,3,4,5,6,7}
	fmt.Printf("len and cap of array %v is:%d and %d\n",a,len(a),cap(a))
	fmt.Printf("item in array:%v is:",a)
	for _,value := range a{
		fmt.Printf("%d ",value)
	}

	fmt.Println()
		
	s1 := a[3:6]
	fmt.Printf("a[3:6] len and cap of slice: %v is: %d and %d\n",s1,len(s1),cap(s1))
	fmt.Printf("item in slice:%v is:",s1)
	for _,value := range s1{
		fmt.Printf("%d ",value)
	}

	fmt.Println()

	s2 := a[:]
	fmt.Printf("a[:] len and cap of slice: %v is: %d and %d\n",s2,len(s2),cap(s2))
	fmt.Printf("item in slice:%v is:",s2)
	for _,value := range s2{
		fmt.Printf("%d ",value)
	}

	fmt.Println()

	s3 := a[:0]
	fmt.Printf("a[:0] len and cap of slice: %v is: %d and %d\n",s3,len(s3),cap(s3))
	fmt.Printf("item in slice:%v is:",s3)
	for _,value := range s3{
		fmt.Printf("%d",value)
	}

	fmt.Println()

	s4 := a[:4]
	fmt.Printf("a[:4] len and cap of slice: %v is: %d and %d\n",s4,len(s4),cap(s4))
	fmt.Printf("item in slice:%v is:",s4)
	for _,value := range s4{
		fmt.Printf("%d ",value)
	}

	fmt.Println()

	a2 := a[2:]
	fmt.Printf("a[2:] len and cap of slice: %v is: %d and %d\n",a2,len(a2),cap(a2))
	fmt.Printf("item in slice:%v is:",a2)
	for _,value := range a2{
		fmt.Printf("%d ",value)
	}

	fmt.Println()

	s1[0] = 456
	fmt.Printf("a item in array change after changing slice:%v is:",s1)
	for _,value := range a{
		fmt.Printf("%d ",value)
	}

	fmt.Println()

	s5 := make([]int,10,20)
	s5[4] = 5
	fmt.Printf("len and cap of slice:%v is:%d and %d\n",s5,len(s5),cap(s5))
	fmt.Printf("item in slice %v is:",s5)
	for _,value := range s5{
		fmt.Printf("%d ",value)
	}

	 fmt.Println()
}






