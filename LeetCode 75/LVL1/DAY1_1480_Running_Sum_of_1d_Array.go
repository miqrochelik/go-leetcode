package main

import(
	"fmt"
)

func main(){
	nums:=[]int{3,1,2,10,1}
	fmt.Print(runningSum(nums))
}

func runningSum(nums []int) []int {
    items:=make([]int, len(nums))
    sum:=0
    for i,_ := range(nums){
    	sum+=nums[i]
    	items[i]=sum
    }
    return items
}