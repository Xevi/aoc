package main

import (
	"fmt"
	"strings"
	"utils"
)

func processInput(entries string) [][]string{
	results := [][]string{}
	for _, group := range strings.Split(entries, "\n\n"){
		results = append(results, strings.Split(group, "\n"))
	}

	return results
}

func countAnyoneYes(group []string) int{
	yesMap := map[string]bool{}
	sum := 0

	for _, answers := range group{
		for _, answer := range strings.Split(answers, ""){
			_, alreadyIn := yesMap[answer]

			if !alreadyIn {
				yesMap[answer] = true
				sum++
			}
		}
	}

	return sum
}

func countEvereyoneYes(group []string) int{
	yesMap := map[string]int{}

	for _, answers := range group{
		for _, answer := range strings.Split(answers, ""){
			_, alreadyIn := yesMap[answer]

			if !alreadyIn {
				yesMap[answer] = 1
			}else{
				yesMap[answer]++
			}
		}
	}

	sum := 0
	groupLen := len(group)
	for _, numYes := range yesMap {
		if(numYes == groupLen){
			sum++
		}
	}

	return sum
}

func getSolution(groups [][]string, count func(group []string) int) int{
	sum := 0
	for _, group := range groups {
		sum = sum + count(group)
	}

	return sum
}

func main(){
	groups := processInput(utils.LoadFile("input.txt"))

	solution1 := getSolution(groups, countAnyoneYes)

	fmt.Println("Solution1:", solution1)

	solution2 := getSolution(groups, countEvereyoneYes)

	fmt.Println("Solution2:", solution2)
}