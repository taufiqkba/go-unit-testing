package helper

import (
	"fmt"
	"testing"
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

}
