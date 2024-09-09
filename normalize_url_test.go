package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "return url",
			url:  "blog.boot.dev/path",
			want: "blog.boot.dev/path",
		},
		{
			name: "remove scheme",
			url:  "https://blog.boot.dev/path",
			want: "blog.boot.dev/path",
		},
		{
			name: "remove params",
			url:  "https://blog.boot.dev/path?views=1000",
			want: "blog.boot.dev/path",
		},
		{
			name: "remove multiple params",
			url:  "https://blog.boot.dev/path?views=1000&updated=True",
			want: "blog.boot.dev/path",
		},
		{
			name: "remove fragment",
			url:  "https://blog.boot.dev/path#header-1",
			want: "blog.boot.dev/path",
		},
		{
			name: "to lowercase",
			url:  "https://blog.boot.dev/PATH",
			want: "blog.boot.dev/path",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeURL(tt.url)
			if err != nil {
				t.Errorf(
					"Test %s failed, unexpected error: %v",
					tt.name,
					err,
				)
			}

			if tt.want != got {
				t.Errorf(
					"Test %s failed: expected '%s', got '%s'",
					tt.name,
					tt.want,
					got,
				)
			}
		})
	}
}

func TestNormalizeURLError(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "invalid URL",
			url:  `:/invalid`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := normalizeURL(tt.url)
			if err == nil {
				t.Errorf(
					"Test %s failed: expected error: %v",
					tt.name,
					err,
				)
			}
		})
	}
}
