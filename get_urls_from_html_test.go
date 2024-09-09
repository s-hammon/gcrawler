package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name       string
		htmlBody   string
		rawBaseURL string
		want       []string
	}{
		{
			name: "parse abs & rel paths",
			htmlBody: `
<html>
	<body>
		<a href="path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			rawBaseURL: "https://blog.boot.dev",
			want: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/one",
			},
		},
		{
			name: "nil on bad href",
			htmlBody: `
<html>
	<body>
		<a href=":\\badhref">
			<span>Boot.dev</span>
		</a>
	</body>
</html>`,
			rawBaseURL: "https://blog.boot.dev",
			want:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLsFromHTML(tt.htmlBody, tt.rawBaseURL)
			if err != nil {
				t.Errorf("Test %s failed, unexpected error: %v", tt.name, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Test failed: expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestGetURLsFromHTMLError(t *testing.T) {
	tests := []struct {
		name       string
		htmlBody   string
		rawBaseURL string
	}{
		{
			name: "invalid base URL",
			htmlBody: `
<html>
	<body>
		<a href="/path">Boot.dev</a>
	</body>
</html>
			`,
			rawBaseURL: ":\\invalidURL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getURLsFromHTML(tt.htmlBody, tt.rawBaseURL)
			if err == nil {
				t.Errorf("Test %s failed, expected error", tt.name)
			}
		})
	}
}
