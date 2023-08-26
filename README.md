# strangeslice
A Go based parser for string based number range -> number slice.

Examples:
```golang
package strangeslice_test

import (
	"github.com/danesparza/strangeslice"
	"reflect"
	"testing"
)

func TestStrRangeToInts(t *testing.T) {

	tests := []struct {
		name   string
		source string
		want   []int
	}{
		{
			name:   "Simple sequence to sorted slice",
			source: "3, 2, 4, 9, 5, 13",
			want:   []int{2, 3, 4, 5, 9, 13},
		},
		{
			name:   "Sequence and range to sorted slice",
			source: "3, 2, 4, 5-10, 25-26",
			want:   []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 25, 26},
		},
		{
			name:   "Overlapping ranges to sorted slice",
			source: "5-10, 8-12",
			want:   []int{5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name:   "Reversed overlapped ranges to sorted slice",
			source: "10-5, 12-8",
			want:   []int{5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name:   "Ranges with single items",
			source: "10-10, 12-12",
			want:   []int{10, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strangeslice.StrRangeToInts(tt.source)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrRangeToInts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
```
