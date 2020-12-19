package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

type passport struct {
  fields map[string]string
  valid bool
  size int
}

var validFields = map[string]func(string)bool {
  // byr (Birth Year) - four digits; at least 1920 and at most 2002.
  "byr": func (val string) bool {
    value, err := strconv.Atoi(val)
    utils.Check(err)
    return value >= 1920 && value <= 2002
  },
  // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
  "iyr": func (val string) bool {
    value, err := strconv.Atoi(val)
    utils.Check(err)
    return value >= 2010 && value <= 2020
  },
  // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
  "eyr": func (val string) bool {
    value, err := strconv.Atoi(val)
    utils.Check(err)

    return value >= 2020 && value <= 2030
  },
  // hgt (Height) - a number followed by either cm or in:
  //     If cm, the number must be at least 150 and at most 193.
  //     If in, the number must be at least 59 and at most 76.
  "hgt": func (value string) bool {
    inCm := regexp.MustCompile(`^(\d{3})cm$`)
    inIn := regexp.MustCompile(`^(\d{2})in$`)

    if inCm.MatchString(value) {
      matches := inCm.FindStringSubmatch(value)
      num, err := strconv.Atoi(matches[1])
      utils.Check(err)
      return num >= 150 && num <= 193
    }

    if inIn.MatchString(value) {
      matches := inIn.FindStringSubmatch(value)
      num, err := strconv.Atoi(matches[1])
      utils.Check(err)
      return num >= 59 && num <= 76
    }

    return false
  },
  // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
  "hcl": func (value string) bool {
    validRegex := regexp.MustCompile(`#[0-9|a-f]{6}`)
    return validRegex.MatchString(value)
  },
  // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
  "ecl": func (value string) bool {
    validRegex := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
    return validRegex.MatchString(value)
  },
  // pid (Passport ID) - a nine-digit number, including leading zeroes.
  "pid": func (value string) bool {
    validRegex := regexp.MustCompile(`^[0-9]{9}$`)
    return validRegex.MatchString(value)
  },
  // cid (Country ID) - ignored, missing or not.
  "cid": func (value string) bool {
    return true
  },
}

func newPassport(entry string, validate bool) passport{
  p := passport{
    fields: map[string]string{},
    valid: true,
    size: 0,
  }

  for _, field := range strings.Split(strings.ReplaceAll(entry, "\n", " "), " ") {
    props := strings.Split(field, ":")
    prop := props[0]
    value := props[1]

    validator, ok := validFields[prop]
    if(!ok) {
      p.valid = false
      return p
    }else{
      p.fields[prop] = value
      p.size++
    }

    if validate {
      p.valid = p.valid && validator(value)
    }
  }

  if p.size < 7 || p.size > 8 {
    p.valid = false
  }

  if p.size == 7 {
    _, hasCid := p.fields["cid"]
    p.valid = p.valid && !hasCid
  }

  return p
}

func processInput(data string, validate bool) []passport{
  passports := []passport{}
  for _, entry := range strings.Split(data, "\n\n"){
    passports = append(passports, newPassport(entry, validate))
  }

  return passports
}

func getSolution(passports []passport) int {
  valid := 0
  for _, passport := range passports {
    if passport.valid {
      valid++
    }
  }

  return valid
}

func main(){
  // Solution 1 does not require validate input (second param false)
  // Solution 2 requires to validate input (second param true)
  passports := processInput(utils.LoadFile("input.txt"), true)
  solution := getSolution(passports)
  fmt.Println("Solution:", solution)
}