package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []Page
	}{
		{
			name: "order count desc",
			input: map[string]int{
				"url1": 5,
				"url2": 3,
				"url3": 1,
				"url4": 10,
				"url5": 2,
			},
			expected: []Page{
				{"url4", 10},
				{"url1", 5},
				{"url2", 3},
				{"url5", 2},
				{"url3", 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortPages(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Test %s failed, expected %v, got %v", tt.name, tt.expected, got)
			}
		})
	}
}
