// Linear search is a sequential searching algorithm
// where we start from one end and check every element of the list until the desired element is found.
// It is the simplest searching algorithm.

package main

import (
	"fmt"
)


func main(){
	var a=make([]int,5,7)
	var n int
	for i:=0;i<5;i++{
		fmt.Scanf("%v",&a[i])
	}
	// fmt.Println(len(a))
	// fmt.Println(cap(a))
	// fmt.Println(a)
	fmt.Println("enter the finding number n")
	fmt.Scanln(&n)
	LinearSearch(a,n)
}
func LinearSearch(list []int ,n int){
	var found bool
	var SaveIndex int
	for index,val:=range(list){
		if val==n{
			found=true
			SaveIndex=index
			break
		}
	}
	fmt.Println(FoundOrNot(found,SaveIndex))
}
func FoundOrNot(f bool,i int)string{

	if f{

		return fmt.Sprintf("Given number is found at index of %v", i)
	}
	return "not found"
}