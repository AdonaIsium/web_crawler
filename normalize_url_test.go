package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		expected    string
		expectedErr bool
	}{
		{
			name:        "remove scheme",
			inputURL:    "https://starcraft.gg/iscool",
			expected:    "starcraft.gg/iscool",
			expectedErr: false,
		},
		{
			name:        "already no scheme",
			inputURL:    "starcraft.gg/iscool",
			expected:    "starcraft.gg/iscool",
			expectedErr: false,
		},
		{
			name:        "malformed scheme",
			inputURL:    "htps://starcraft.gg/iscool",
			expected:    "",
			expectedErr: true,
		},
		{
			name:        "malformed scheme",
			inputURL:    "ftp://starcraft.gg/iscool",
			expected:    "",
			expectedErr: true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !tc.expectedErr {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
