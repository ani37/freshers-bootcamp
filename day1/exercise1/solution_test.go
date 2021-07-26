package main
import (
	"testing"
)

var tests = []struct {
	matrices  Matrix
	expected int
}{
	{Matrix{3, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		3},
	{Matrix{4, 4, [][]int{{1, 2, 3,4}, {4, 5, 6,7}, {7, 8, 9,10},{11,12,13,14}}},
		4},
}

func TestGetRows(t *testing.T) {

		for _, test :=range tests{
			if output := test.matrices.getRows(); output != test.expected {
				t.Error("GetRows has error!", test.expected, output)
			}
		}
}
