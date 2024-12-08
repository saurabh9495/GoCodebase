package main

import "fmt"

type Solution struct{}

func (s Solution) candy(ratings []int) int{
	n := len(ratings)
	if n == 0{
		return 0
	}
	candies := make([]int, n)
	for i := range candies{
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1] + 1
        }
    }

    for i := n - 2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] {
            candies[i] = max(candies[i], candies[i+1]+1)
        }
    }

	totalCandies := 0
	for _,c := range candies{
		totalCandies += c
	}

	return totalCandies

}

func main(){
	s := Solution{}
	ratings := []int{1,2,2}
	fmt.Println(s.candy(ratings))
}