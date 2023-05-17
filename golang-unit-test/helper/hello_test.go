// TO RUN TEST FROM ROOT FOLDER USE COMMAND --> go test -v ./...
// TO MAKE UNIT TEST EASIER USE LIBRARY TESTIFY FROM github.com/stretchr/testify
// ITS PROPER TO USE TABLE TO AVOID ITERATE WRITING CODE
package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// BENCHMARK TABLE
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Juan",
			request: "Juan",
		},
		{
			name:    "Ananda",
			request: "Ananda",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// SUB BENCHMARK
func BenchmarkSub(b *testing.B) {
	b.Run("Juan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Juan")
		}
	})
	b.Run("Ananda", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ananda")
		}
	})
}

// TEST BENCHMARK
func BenchmarkHelloWorldJuan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Juan")
	}
}
func BenchmarkHelloWorldAnanda(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ananda")
	}
}

// TABLE TEST
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Juan",
			request:  "Juan",
			expected: "Hello Juan",
		},
		{
			name:     "Ananda",
			request:  "Ananda",
			expected: "Hello Ananda",
		},
		{
			name:     "Jamal",
			request:  "Jamal",
			expected: "Hello Jamal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// SUB TEST
func TestSubTest(t *testing.T) {
	t.Run("Juan", func(t *testing.T) {
		result := HelloWorld("Juan")
		assert.Equal(t, "Hello Juan", result, "Result is not Hello Juan")
	})
	t.Run("2", func(t *testing.T) {
		result := HelloWorld("2")
		assert.Equal(t, "Hello 2", result, "Result is not Hello Juan")
	})
}

// MAIN TEST
func TestMain(m *testing.M) {
	//before unit test
	fmt.Println("Before Unit Test")

	m.Run()

	//after unit test
	fmt.Println("After Unit Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Juan")

	if result != "Hello Juan" {
		//unit test failed
		t.Error("Result is not Hello Juan")
	}
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("2")

	if result != "Hello 2" {
		//unit test failed
		t.Fatal("Result is not Hello 2")
	}
}

func TestHelloWorldJuan(t *testing.T) {
	result := HelloWorld("Juan")

	if result != "Hi Juan" {
		t.Error("Result must be Hi Juan")
	}

	fmt.Println("TestHelloWorld DONE")
}

func TestHelloWorldAnanda(t *testing.T) {
	result := HelloWorld("Ananda")

	if result != "Hi Ananda" {
		t.Fatal("Result must be Hi Ananda")
	}

	fmt.Println("TestHelloWorldAnanda DONE")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Juan")

	assert.Equal(t, "Hi Juan", result, "Result is not Hello Juan")

	fmt.Println("Test Hello World with Assert DONE")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Juan")

	require.Equal(t, "Hi Juan", result, "Result is not Hello Juan")

	fmt.Println("Test Hello World with Require DONE")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on MAC") //to skip test in specific situations
	}
	result := HelloWorld("Juan")
	require.Equal(t, "Hi Juan", result, "Result is not Hello Juan")
}
