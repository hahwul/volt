package util

import (
	"reflect"
	"testing"
)

func TestUniqueStringSlice(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "not changed case",
			args: []string{"1", "2", "3", "4"},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "changed case1",
			args: []string{"1", "2", "3", "4", "4"},
			want: []string{"1", "2", "3", "4"},
		},
		{
			name: "changed case2",
			args: []string{"1", "2", "4", "3", "2"},
			want: []string{"1", "2", "4", "3"},
		},
		{
			name: "changed case3",
			args: []string{"1", "1", "2", "1", "2"},
			want: []string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueStringSlice(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueIntSlice(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "not changed case",
			args: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "changed case1",
			args: []int{1, 2, 3, 4, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "changed case2",
			args: []int{1, 2, 4, 3, 2},
			want: []int{1, 2, 4, 3},
		},
		{
			name: "changed case3",
			args: []int{1, 2, 1, 1, 1, 1, 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueIntSlice(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
