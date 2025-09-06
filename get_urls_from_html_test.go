package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		inputBody   string
		expected    []string
		expectedErr bool
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://starcraft.gg",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>starcraft.gg</span>
		</a>
		<a href="https://other.com/path/one">
			<span>starcraft.gg</span>
		</a>
	</body>
</html>
`,
			expected:    []string{"https://starcraft.gg/path/one", "https://other.com/path/one"},
			expectedErr: false,
		},
		{
			name:     "relative urls only",
			inputURL: "https://starcraft.gg",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>starcraft.gg</span>
		</a>
		<a href="/path/two">
			<span>starcraft.gg</span>
		</a>
	</body>
</html>
`,
			expected:    []string{"https://starcraft.gg/path/one", "https://starcraft.gg/path/two"},
			expectedErr: false,
		},
		{
			name:     "absolute urls only",
			inputURL: "https://starcraft.gg",
			inputBody: `
<html>
	<body>
		<a href="https://starcraft.gg/path/one">
			<span>starcraft.gg</span>
		</a>
		<a href="https://other.com/path/one">
			<span>starcraft.gg</span>
		</a>
	</body>
</html>
`,
			expected:    []string{"https://starcraft.gg/path/one", "https://other.com/path/one"},
			expectedErr: false,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			base, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}
			actual, err := getURLsFromHTML(tc.inputBody, base)
			if err != nil && !tc.expectedErr {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
