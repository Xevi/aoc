// Advent Calendar day 3 -https://adventofcode.com/2020/day/3
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type slope struct {
  right int
  down int
}

func newSlope(right int, down int) slope{
  return slope{right: right, down: down}
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func processInput(input string) [] [] string {
  plan := [][]string{}

  for _, line := range strings.Split(input, "\n"){
    plan = append(plan, strings.Split(line, ""))
  }

  return plan
}

func getSolution(plan [][]string, slopes []slope) []int{
  results := []int{}

  for _, s := range slopes {
    trees := 0
    row := 0
    col := 0

    for row < len(plan) - 1 {
      line := plan[row]

      col = (col + s.right) % len(line)
      row = row + s.down

      if plan[row][col] == "#" {
        trees++
      }
    }
    results = append(results, trees)
  }

  return results
}

func main() {
  data, err := ioutil.ReadFile("./inputs/day_3_input.txt")
  check(err)

  plan := processInput(string(data))

  // slopesPart1 := []slope{newSlope(3, 1)}

  slopesPart2 := []slope{newSlope(1, 1), newSlope(3, 1), newSlope(5, 1), newSlope(7, 1), newSlope(1,g p 2)}

  solutions := getSolution(plan, slopesPart2)

  solution := 1
  for _, val := range solutions {
    solution = val * solution
  }

  fmt.Println("Solution:", solution)
}