package main

// Пишите тесты в этом файле

import "testing"

func TestMaximumFunctions(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{42},
			want: 42,
		},
		{
			name: "equal elements",
			data: []int{7, 7, 7},
			want: 7,
		},
		{
			name: "maximum at beginning",
			data: []int{100, 10, 50},
			want: 100,
		},
		{
			name: "maximum in middle",
			data: []int{10, 100, 50},
			want: 100,
		},
		{
			name: "maximum at end",
			data: []int{10, 50, 100},
			want: 100,
		},
		{
			name: "fewer elements than chunks",
			data: []int{3, 1, 2},
			want: 3,
		},
		{
			name: "length not divisible by chunks",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
	}

	functions := []struct {
		name string
		fn   func([]int) int
	}{
		{name: "maximum", fn: maximum},
		{name: "maxChunks", fn: maxChunks},
	}

	for _, function := range functions {
		t.Run(function.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := function.fn(tt.data)

					if got != tt.want {
						t.Errorf(
							"%s(%v) = %d; want %d",
							function.name,
							tt.data,
							got,
							tt.want,
						)
					}
				})
			}
		})
	}
}

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size 0",
			size: 0,
		},
		{
			name: "size 1",
			size: 1,
		},
		{
			name: "size 10",
			size: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if len(got) != tt.size {
				t.Errorf(
					"generateRandomElements(%d): got slice length %d, want %d",
					tt.size,
					len(got),
					tt.size,
				)
			}
		})
	}
}
