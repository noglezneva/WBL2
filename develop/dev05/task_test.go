package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFilterFile(t *testing.T) {
	content := []byte("hello\nworld\nhello world\n")
	err := ioutil.WriteFile("test.txt", content, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.txt")

	testCases := []struct {
		name           string
		file           string
		pattern        string
		after          int
		before         int
		context        int
		count          bool
		ignoreCase     bool
		invert         bool
		fixed          bool
		lineNum        bool
		expectedOutput string
	}{
		{
			name:           "No options",
			file:           "test.txt",
			pattern:        "hello",
			expectedOutput: "hello\nhello world\n",
		},
		{
			name:           "After option",
			file:           "test.txt",
			pattern:        "hello",
			after:          1,
			expectedOutput: "hello\nworld\nhello world\n",
		},
		{
			name:           "Before option",
			file:           "test.txt",
			pattern:        "world",
			before:         1,
			expectedOutput: "hello\nworld\n",
		},
		{
			name:           "Context option",
			file:           "test.txt",
			pattern:        "world",
			context:        1,
			expectedOutput: "hello\nworld\nhello world\n",
		},
		{
			name:           "Count option",
			file:           "test.txt",
			pattern:        "hello",
			count:          true,
			expectedOutput: "2\n",
		},
		{
			name:           "Ignore case option",
			file:           "test.txt",
			pattern:        "HeLlO",
			ignoreCase:     true,
			expectedOutput: "hello\nhello world\n",
		},
		{
			name:           "Invert option",
			file:           "test.txt",
			pattern:        "hello",
			invert:         true,
			expectedOutput: "world\n",
		},
		{
			name:           "Fixed option",
			file:           "test.txt",
			pattern:        "hello",
			fixed:          true,
			expectedOutput: "hello\n",
		},
		{
			name:           "Line number option",
			file:           "test.txt",
			pattern:        "hello",
			lineNum:        true,
			expectedOutput: "1:hello\n3:hello world\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := captureOutput(func() {
				filterFile(testCase.file, testCase.pattern, testCase.after, testCase.before, testCase.context, testCase.count, testCase.ignoreCase, testCase.invert, testCase.fixed, testCase.lineNum)
			})

			if output != testCase.expectedOutput {
				t.Errorf("Expected output: %s, but got: %s", testCase.expectedOutput, output)
			}
		})
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
