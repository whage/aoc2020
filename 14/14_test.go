package main

import (
	"testing"
)

func TestGetPowerSet(t *testing.T) {
	base := []int{1,2,3}
	expected := [][]int{
		[]int{1,2,3},
		[]int{1,2},
		[]int{1,3},
		[]int{1},
		[]int{2,3},
		[]int{2},
		[]int{3},
		[]int{},
	}
	result := getPowerSet(base, []int{})

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if expected[i][j] != result[i][j] {
				t.Errorf("Expected %v, got %v at [%d][%d]", expected[i][j], result[i][j], i, j)
			}
		}
	}
}
