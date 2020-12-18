// Advent Calendar day 2 - https://adventofcode.com/2020/day/2
package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type policy struct {
  letter string
  min int
  max int
}

func newPolicy(policySection string) policy {
  sections := strings.Split(policySection, " ")
  letter := sections[1]
  minMax := strings.Split(sections[0], "-")

  min, err := strconv.Atoi(minMax[0])
  utils.Check(err)
  max, err := strconv.Atoi(minMax[1])
  utils.Check(err)

  return policy{letter: letter, min: min, max: max}
}

type password struct {
  policy policy
  password string
}

func newPassword(policySection string, passwordSection string) password {
  return password{policy: newPolicy(policySection), password: strings.Trim(passwordSection, " ")}
}

func processInput(data string) [] password{
  passwords := [] password{}

  for _, line := range strings.Split(data, "\n") {
    sections := strings.Split(line, ": ")
    passwords = append(passwords, newPassword(sections[0], sections[1]))
  }

  return passwords
}

// Part 1
func isPasswordCorrect1(password password) bool{
  pwd := password.password
  policy := password.policy

  repetition := strings.Count(pwd, policy.letter)

  return repetition >= policy.min && repetition <= policy.max
}

// Part 2
func isPasswordCorrect2(password password) bool{
  pwd := password.password
  policy := password.policy
  letter := policy.letter

  minChar := string(pwd[policy.min - 1])
  maxChar := string(pwd[policy.max - 1])

  valid := (minChar == letter || maxChar == letter) && minChar != maxChar
  return valid
}

// RunDay2 Main runner for day 2
func main(){
  passwords := processInput(utils.LoadFile("input.txt"))

  sum := 0
  for _, password := range passwords {
    ok := isPasswordCorrect2(password)
    if ok {
      sum++
    }
  }

  fmt.Printf("Solution: %d", sum)
}