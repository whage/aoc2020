package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Vec2 struct {
	X, Y int
}

func (v *Vec2) add(other Vec2) Vec2 {
	return Vec2{v.X+other.X, v.Y+other.Y}
}

func (v *Vec2) multiply(m [2]Vec2) Vec2 {
	x := m[0].X*v.X+m[1].X*v[1].X
	y := m[0].Y*v.Y+m[1].Y*v[1].Y
	return Vec2{x, y}
}

func main() {
	data, _ := ioutil.ReadFile("data-small.txt")
	split := strings.Split(string(data), "\n")
	allLines := split[:len(split)-1]
	
	direction := Vec2{1, 0} // starts facing "East"
	position := Vec2{0,0}

	for _, navInst := range allLines {
		action := navInst[0]
		rawValue := navInst[1:]
		value, _ := strconv.Atoi(rawValue)

		switch action {
		case 'N':
			pos = pos.add(Vec2{value, 0})
		}
		case 'S':
			pos = pos.add(Vec2{value*-1, 0})
		}
		case 'E':
			pos = pos.add(Vec2{0, value})
		}
		case 'W':
			pos = pos.add(Vec2{0, value*-1})
		}
		case 'L':
			pos = pos.multiply([2]Vec2{Vec2{0, 1}, Vec2{-1 0})
		}
		case 'R':
			pos = pos.multiply([2]Vec2{Vec2{0, -1}, Vec2{1 0})
		}
		case 'F':
			pos = pos.add(Vec2{0, value*-1})
		}
	}

	fmt.Println("Final direction:", direction)
	fmt.Println("Final position:", position)
}
