package main

import (
    "testing"
	"reflect"
    //"github.com/stretchr/testify/assert"
    //"github.com/stretchr/testify/require"
	//"fmt"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name		string
		inputURL	string
		inputBody	string
		expected	[]string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "Multiple URLS",
			inputURL: "https://test.test.com",
			inputBody: `
			<html>
				<body>
					<a href="https://something.com">
						<span>Boot.dev</span>
					</a>
					<a href="https://spannerworks.net">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://something.com", "https://spannerworks.net", "https://other.com/path/one"},
		},
		// Add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s Failed: Unexpected error: %v", i, tc.name, err)
				return
			}
			testCheck := reflect.DeepEqual(actual, tc.expected)
			if testCheck != true {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		
		})
	}

}