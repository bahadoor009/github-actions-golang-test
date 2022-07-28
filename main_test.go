package main

import (
	"testing"
)

func TestDummFunc(t *testing.T) {
	testcases := []int{
	 5,
	 10,
	}

	for _, test := range testcases {
		if output := DummyFunc(test); output != test * 100 {
			t.Errorf("Output %v not equal to expected", output)
		}
	}
}
