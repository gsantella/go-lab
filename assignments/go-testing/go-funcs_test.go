package main

import "testing"

func TestFunc1(t *testing.T) {
	expected := 3
	result := func1(2, 1)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFunc2(t *testing.T) {
	expected := 3
	result := func2(2, 1)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSmallestNumber(t *testing.T) {
	expected := 1
	result := smallestNumber([]int{1, 2, 3, 4, 5})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestLargestNumber(t *testing.T) {
	expected := 5
	result := largestNumber([]int{1, 2, 3, 4, 5})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestFirstNumber(t *testing.T) {
	expected := 1
	result := firstNumber([]int{1, 2, 3, 4, 5})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestLastNumber(t *testing.T) {
	expected := 5
	result := lastNumber([]int{1, 2, 3, 4, 5})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestAverageNumber(t *testing.T) {
	expected := 3
	result := averageNumber([]int{1, 2, 3, 4, 5})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCountNumbers(t *testing.T) {
	expected := 5
	result := countNumbers([]int{1, 2, 3, 4, 5})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSumNumbers(t *testing.T) {
	expected := 15
	result := sumNumbers([]int{1, 2, 3, 4, 5})
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
