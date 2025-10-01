package utilities_test

import (
	"testing"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

func TestFindType(t *testing.T) {
	names := []string{"x", "y", "z"}
	integers := []int{1, 2, 3}
	points := []float32{1.1}
	points2 := []float64{}
	smallNums := []int8{}

	typeName := utilities.FindType(names, true)
	if typeName != "string" {
		t.Errorf("expected %s, got %s\n", "string", typeName)
	}

	typeName = utilities.FindType(names, false)
	if typeName != "[]string" {
		t.Errorf("expected %s, got %s\n", "string", typeName)
	}

	typeName = utilities.FindType(integers, true)
	if typeName != "int" {
		t.Errorf("expected %s, got %s\n", "int", typeName)
	}

	typeName = utilities.FindType(integers, false)
	if typeName != "[]int" {
		t.Errorf("expected %s, got %s\n", "[]int", typeName)
	}

	typeName = utilities.FindType(points, true)
	if typeName != "float32" {
		t.Errorf("expected %s, got %s\n", "float32", typeName)
	}

	typeName = utilities.FindType(points, false)
	if typeName != "[]float32" {
		t.Errorf("expected %s, got %s\n", "[]float32", typeName)
	}

	typeName = utilities.FindType(points2, true)
	if typeName != "float64" {
		t.Errorf("expected %s, got %s\n", "float64", typeName)
	}

	typeName = utilities.FindType(points2, false)
	if typeName != "[]float64" {
		t.Errorf("expected %s, got %s\n", "[]float64", typeName)
	}

	typeName = utilities.FindType(smallNums, true)
	if typeName != "int8" {
		t.Errorf("expected %s, got %s\n", "int8", typeName)
	}

	typeName = utilities.FindType(smallNums, false)
	if typeName != "[]int8" {
		t.Errorf("expected %s, got %s\n", "[]int8", typeName)
	}
}
