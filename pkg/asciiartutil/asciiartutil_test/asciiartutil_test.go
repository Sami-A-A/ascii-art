package asciiartutil_test

import (
	"asciiart/pkg/asciiartutil"
	"fmt"
	"testing"
)

func TestCheckFormat(t *testing.T) {

	sampleInput := []string{
		"--output=output.txt",
		"--alignment=right",
		"--color=blue",
		"abc",
		"--reverse=true",
		"This is a test",
		"standard",
	}

	resultA, resultB, resultC, err := asciiartutil.CheckFormat(sampleInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	expectedA, expectedB, expectedC, err := "abc", "This is a test", "standard", nil

	if resultA != expectedA {
		t.Errorf("Test A Failed")
		
	} else {
		fmt.Println("Test A Passed")
	}
	if resultB != expectedB {
		t.Errorf("Test B Failed")
	} else {
		fmt.Println("Test B Passed")
	}
	if resultC != expectedC {
		t.Errorf("Test C Failed")
	} else {
		fmt.Println("Test C Passed")
	}
}