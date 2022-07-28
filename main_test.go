ackage main

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestDummFunc(t *testing.T) {
	testcases := []int{
	 5,
	 10
	}

	for _, test := range testcases {
		if output, err := DummyFunc(test); !Compare(output, test * 100) && err == nil {
			t.Errorf("Output %v not equal to expected %v", output, &test.expected_response)
		} else if err != nil {
			t.Error(err)
		}
	}
}
