package robot_test

import (
	"testing"

	"github.com/amir6432/robot/robot"
)

var inBoundsTests = []struct {
	width, height, x, y int
	want                bool
}{
	{4, 8, 0, 0, true},
	{4, 8, 0, 7, true},
	{4, 8, 3, 0, true},
	{4, 8, 3, 7, true},
	{4, 8, 1, 2, true},

	{4, 8, -1, 1, false},
	{4, 8, 1, -1, false},
	{4, 8, -2, -2, false},
	{4, 8, 4, 0, false},
	{4, 8, 5, 3, false},
	{4, 8, 0, 8, false},
	{4, 8, 3, 9, false},
}

func TestTable_InBounds(t *testing.T) {
	for _, ibt := range inBoundsTests {
		tbl := robot.NewTable(ibt.width, ibt.height)
		if got := tbl.InBounds(ibt.x, ibt.y); got != ibt.want {
			t.Errorf("Table(%d, %d).InBounds(%d, %d) = %v; want: %v", ibt.width, ibt.height, ibt.x, ibt.y, got, ibt.want)
		}
	}
}
