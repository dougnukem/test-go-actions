package main_test

import "testing"

func TestExample(t *testing.T) {
	expected := true
	testValue := true
	if expected != testValue {
		t.Errorf("Test fails expected: %v but was: %v", expected, testValue)
	}
}
