package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
	"utils"
)

func createRange(min int, max int) (func() (int, int), func() (int, int)) {
  currentMin := min
  currentMax := max

    //upper
  return func() (int, int) {
      diff := currentMax - currentMin
      currentMax = currentMax - int(math.Floor(float64(diff)/2.0)) - 1
      return currentMin, currentMax
    },
    //lower
    func() (int, int) {
      diff := currentMax - currentMin
      currentMin = currentMin + int(math.Ceil(float64(diff)/2.0))
      return currentMin, currentMax
    }
}

// Seat type
type Seat struct {
  row int
  col int
}

func (s Seat) getID() int {
  return s.row * 8 + s.col
}

func newSeat(id string) (Seat, error){
  seatRegExp := regexp.MustCompile(`([F|B]{7})([R|L]{3})`)

  if !seatRegExp.MatchString(id) {
    return Seat{}, fmt.Errorf("invalid seat")
  }

  matches := seatRegExp.FindStringSubmatch(id)

  rows := matches[1]
  frontRow, backRow := createRange(0, 127)
  var minRow int
  var maxRow int

  for _, rowEntry := range rows {
    if rowEntry == 'F' {
      minRow, maxRow = frontRow()
    }

    if rowEntry == 'B' {
      minRow, maxRow = backRow()
    }
  }

  if minRow != maxRow {
    return Seat{}, fmt.Errorf("wrong row %d != %d", minRow, maxRow)
  }

  cols := matches[2]
  leftRow, rightRow := createRange(0, 7)
  var minCol int
  var maxCol int

  for _, colEntry := range cols {
    if colEntry == 'L' {
      minCol, maxCol = leftRow()
    }

    if colEntry == 'R' {
      minCol, maxCol = rightRow()
    }
  }

  if minCol != maxCol {
    return Seat{}, fmt.Errorf("wrong col %d != %d", minCol, maxCol)
  }

  return Seat{
    row: minRow,
    col: minCol,
  }, nil
}

// SeatList of seats definition
type SeatList []Seat

func (s SeatList) Len() int {
  return len(s)
}

func (s SeatList) Less(i, j int) bool {
  return s[i].getID() < s[j].getID()
}

func (s SeatList) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func getSeats(passes [] string) []Seat{
  seats := SeatList{}
  for _, pass := range passes {
    s, error := newSeat(pass)
    utils.Check(error)
    seats = append(seats, s)
  }

  return seats
}

func getSolution1(seats SeatList) int{
  max := 0
  for _, s := range seats{
    current := s.getID()
    if current > max {
      max = current
    }
  }

  return max
}

func getSolution2(seats SeatList) int{
  sort.Sort(seats)

  for i, s := range seats{
    current := s.getID()
    if i > 0 {
      expectedCurrent := seats[i-1].getID() + 1
      if expectedCurrent != current {
        return expectedCurrent
      }
    }
  }

  return -1
}

func main(){
  seats := getSeats(strings.Split(utils.LoadFile("input.txt"), "\n"))
  solution1 := getSolution1(seats)
  fmt.Println("Solution1:", solution1)
  solution2 := getSolution2(seats)
  fmt.Println("Solution2:", solution2)
}