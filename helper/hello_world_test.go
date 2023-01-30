package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Taufiq" {
		// unit test failed
		panic("Result is no Hello Taufiq")
	}
}

func TestHelloWorldKurniawan(t *testing.T) {
	result := HelloWorld("Kurniawan")
	if result != "Hello Kurniawan" {
		// unit test failed
		panic("Result is no Hello Kurniawan")
	}
}
