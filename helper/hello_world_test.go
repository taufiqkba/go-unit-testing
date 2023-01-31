package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Testing with Test Main

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run() // Running all testing

	fmt.Println("After Unit Test")
}

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

// Test with SubTest

func TestSubTest(t *testing.T) {
	t.Run("Sub Test 1", func(t *testing.T) {
		result := HelloWorld("Taufiq")
		require.Equal(t, "Hello Taufiq", result, "Result must be 'Hello Taufiq'")
	})
	t.Run("Sub Test 2", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'")
	})
}

// Testing with Table Test

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Test_1",
			request:  "Taufiq",
			expected: "Hello Taufiq",
		},
		{
			name:     "Test_2",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
		{
			name:     "Test_3",
			request:  "Bayu",
			expected: "Hello Bayu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Benchmark test

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Taufiq")
	}
}

// Sub Benchmark test

func BenchmarkSub(b *testing.B) {
	b.Run("Sub Benchmark 1", func(b *testing.B) {
		for x := 0; x < b.N; x++ {
			HelloWorld("SubBench_1")
		}
	})
	b.Run("Sub Benchmark 2", func(b *testing.B) {
		for y := 0; y < b.N; y++ {
			HelloWorld("SubBench_2")
		}
	})
}
