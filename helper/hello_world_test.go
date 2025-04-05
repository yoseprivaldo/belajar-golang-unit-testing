package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// * Table BanchMark
func BenchmarkTable(b *testing.B) {
	// TODO: Initialize
	bencmarks := []struct{
		name string
		request string
	}{
		{
			name: "Yosep", 
			request: "Yosep",
		},
		{
			name: "Rivaldo",
			request : "Rivaldo",
		},
	}
	for _, bencmark := range bencmarks {
		b.Run(bencmark.name, func(b *testing.B) {
			for b.Loop() {
				HelloWorld(bencmark.request)
			}
		})
	}
}

// * BanchMark Test
func Benchmark(b *testing.B) {
	b.Run("Yosep", func(b *testing.B) {
		for b.Loop() {
			HelloWorld("Yosep")
		}
	})
	b.Run("Rivaldo", func(b *testing.B) {
		for b.Loop() {
			HelloWorld("Rivaldo")
		}
	})
}

// * Table test
func TestTableHelloWorld(t *testing.T) {
	 testCases  := []struct {
		name	string
		request string
		expected string
	}{
		{
			name: "HelloWorld(Yosep)",
			request: "Yosep",
			expected: "Hello Yosep",
		},
		{
			name: "HelloWorld(Rivaldo)",
			request: "Rivaldo",
			expected: "Hello Rivaldo",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
			
		})
	}
}

// * Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Yosep", func (t *testing.T)  {
		result := HelloWorld("Yosep")
		require.Equal(t, "Hello Yosep", result, "Result must be 'Hello Yosep'")
	})

	t.Run("Rivaldo", func (t *testing.T)  {
		result := HelloWorld("Rivaldo")
		require.Equal(t, "Hello Rivaldo", result, "Result must be 'Hello Rivaldo'")
	})
}

// * Before and After Test
func TestMain(m *testing.M){
	// * before
	fmt.Println("Sebelum Unit Test")

	m.Run()

	// * after
	fmt.Println("Setelah Unit Test")
}

// * Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("WORLD")
	require.Equal(t, "Hello WORLD", result, "Result must be 'Hello WORLD'")
}


func TestHelloWorldYosep(t *testing.T) {
	result := HelloWorld("Yosep")

	if result != "Hello Yosep" {
		t.Error("Result must be Hello Yosep")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldRivaldo(t *testing.T) {
	result := HelloWorld("Rivaldo")

	if result != "Hello Rivaldo" {
		t.Fatal("Result must be Hello Rivaldo")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Silaban")
	assert.Equal(t, "Hello Silaban", result)
	fmt.Println("Diekseskusi")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Silaban")
	require.Equal(t, "Hello Silaban", result)
	fmt.Println("Diekseskusi")
}