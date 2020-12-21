package main

import (
	"testing"
)

func TestGetPowerSet(t *testing.T) {
	base := []int{1,2,3}
	expected := [][]int{
		[]int{},
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{1,2},
		[]int{1,3},
		[]int{2,3},
		[]int{1,2,3},
	}
	result := getPowerSet(base)

	if (len(result) != len(expected)) {
		t.Errorf("Expected length %v. Results: %v", len(expected), result)
	}
}