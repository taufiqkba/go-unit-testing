package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Taufiq")
	if result != "Hello Taufiq" {
		// unit test failed
		// t.Fail() // Still continue to running another code
		t.Error("Result must be 'Hello Taufiq'")
	}

	fmt.Println("Test DONE!")
}

func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("Kurniawan")
	if result != "Hello Kurniawan" {
		// unit test failed
		// t.FailNow() // Code stop running with FailNow()
		t.Fatal("Result must be 'Hello Kurniawan'")
	}

	fmt.Println("Test DONE 2")
}

// Test with Assertion
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Taufiq")
	assert.Equal(t, "Hello Taufiq", result, "Result must be 'Hello Taufiq'")

	fmt.Println("TestHelloWorld with assertion DONE!")
}

// Test with Require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Taufiq")
	require.Equal(t, "Hello Taufiq", result, "Result must be 'Hello Taufiq'")

	fmt.Println("TestHelloWorld with Require DONE!")
}

// Test with Test Skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on LinuxOS")
	}

	result := HelloWorld("Taufiq")
	require.Equal(t, "Hello Taufiq", result, "Result must be 'Hello Taufiq'")
}
