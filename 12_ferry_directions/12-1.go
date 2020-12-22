package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
)

type Vec2 struct {
	X, Y int
}

func (v Vec2) add(other Vec2) Vec2 {
	return Vec2{v.X+other.X, v.Y+other.Y}
}

func (v Vec2) multiplyByMatrix(m [2]Vec2) Vec2 {
	x := m[0].X*v.X+m[1].X*v.Y
	y := m[0].Y*v.X+m[1].Y*v.Y
	return Vec2{x, y}
}

func (v Vec2) multiplyByConstant(c int) Vec2 {
	return Vec2{v.X*c, v.Y*c}
}

func main() {
	data, _ := ioutil.ReadFile("data-large.txt")
	split := strings.Split(string(data), "\n")
	allLines := split[:len(split)-1]
	
	position := Vec2{0,0}
	direction := Vec2{1, 0}

	for _, navInst := range allLines {
		action := navInst[0]
		rawValue := navInst[1:]
		value, _ := strconv.Atoi(rawValue)
		switch action {
		case 'N':
			position = position.add(Vec2{0, value})
		case 'S':
			position = position.add(Vec2{0, value*-1})
		case 'E':
			position = position.add(Vec2{value, 0})
		case 'W':
			position = position.add(Vec2{value*-1, 0})
		case 'L', 'R':
			if action == 'R' { value *= -1 }
			direction = direction.multiplyByMatrix([2]Vec2{
				Vec2{int(math.Cos(float64(value)/180*math.Pi)), int(math.Sin(float64(value)/180*math.Pi))},
				Vec2{int(-1*math.Sin(float64(value)/180*math.Pi)), int(math.Cos(float64(value)/180*math.Pi))},
			})
		case 'F':
			position = position.add(direction.multiplyByConstant(value))
		}
	}

	fmt.Println("Manhattan distance:", math.Abs(float64(position.X)) + math.Abs(float64(position.Y)))
}
