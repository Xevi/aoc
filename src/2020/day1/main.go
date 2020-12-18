// Advent Calendar day 1 - https://adventofcode.com/2020/day/1
package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func processInput(input string) [] int{
  res := [] int{}

  for _, val := range strings.Split(input, "\n") {
    iVal, err := strconv.Atoi(val)
    utils.Check(err)
    res = append(res, iVal)
  }

  return res
}

// Part 1
func twoSum(input []int, numbers map[int]bool, target int) (int, error) {
  for _, val := range input {
    rest := target - val

    _, ok := numbers[rest]
    if ok {
      return rest * val, nil
    }
  }

  return -1, fmt.Errorf("Not found!")
}

// Part 2
func threeSum(input []int, numbers map[int]bool, target int) (int, error){
  for _, val := range input {
    solution, err := twoSum(input, numbers, target - val)
    if err == nil {
      return solution * val, nil
    }
  }

  return -1, nil
}

func main(){
  input := processInput(utils.LoadFile("input.txt"))

  numbers := map[int]bool{}

  for _, val := range input {
    numbers[val] = true
  }

  solution, err := threeSum(input, numbers, 2020)
  utils.Check(err)

  fmt.Println("Solution:", solution)
}
