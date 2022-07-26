package main


import "fmt"

func main() {
	a := []int{5,7, 8, 11,15,16,20}
	// fmt.Println(FindElement(a,20,0,len(a)-1))
	fmt.Println(RecursiveFind(a,8,0,len(a)-1))
}