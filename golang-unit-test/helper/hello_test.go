// TO RUN TEST FROM ROOT FOLDER USE COMMAND --> go test -v ./...
// TO MAKE UNIT TEST EASIER USE LIBRARY TESTIFY FROM github.com/stretchr/testify
package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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
