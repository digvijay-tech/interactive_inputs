package utilities_test

import (
	"testing"

	"github.com/digvijay-tech/interactive_inputs/internal/utilities"
)

func TestFindArrayType(t *testing.T) {
	names := []string{"x", "y", "z"}
	integers := []int{1, 2, 3}
	points := []float32{1.1}
	points2 := []float64{}
	smallNums := []int8{}

	typeName, _ := utilities.FindArrayType(names, true)
	if typeName != "string" {
		t.Errorf("expected %s, got %s\n", "string", typeName)
	}

	typeName, _ = utilities.FindArrayType(names, false)
	if typeName != "[]string" {
		t.Errorf("expected %s, got %s\n", "string", typeName)
	}

	typeName, _ = utilities.FindArrayType(integers, true)
	if typeName != "int" {
		t.Errorf("expected %s, got %s\n", "int", typeName)
	}

	typeName, _ = utilities.FindArrayType(integers, false)
	if typeName != "[]int" {
		t.Errorf("expected %s, got %s\n", "[]int", typeName)
	}

	typeName, _ = utilities.FindArrayType(points, true)
	if typeName != "float32" {
		t.Errorf("expected %s, got %s\n", "float32", typeName)
	}

	typeName, _ = utilities.FindArrayType(points, false)
	if typeName != "[]float32" {
		t.Errorf("expected %s, got %s\n", "[]float32", typeName)
	}

	typeName, _ = utilities.FindArrayType(points2, true)
	if typeName != "float64" {
		t.Errorf("expected %s, got %s\n", "float64", typeName)
	}

	typeName, _ = utilities.FindArrayType(points2, false)
	if typeName != "[]float64" {
		t.Errorf("expected %s, got %s\n", "[]float64", typeName)
	}

	typeName, _ = utilities.FindArrayType(smallNums, true)
	if typeName != "int8" {
		t.Errorf("expected %s, got %s\n", "int8", typeName)
	}

	typeName, _ = utilities.FindArrayType(smallNums, false)
	if typeName != "[]int8" {
		t.Errorf("expected %s, got %s\n", "[]int8", typeName)
	}
}
