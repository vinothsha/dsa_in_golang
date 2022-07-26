package main

func FindElement(list []int,f int,low int,high int)(int,int){
	for low<=high{
		mid:=low+(high-low)/2
		if f==list[mid]{
			return list[mid],mid
		}else if f>list[mid]{
			low=mid+1
		}else{
			high=mid-1
		}
	}
	return -1,-1
}