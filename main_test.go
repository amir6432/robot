package main

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/amir6432/robot/robot"
)

var runTests = []struct {
	commands string
	want     string
}{
	{"PLACE 0,0,NORTH \n MOVE \n REPORT", "0, 1, NORTH"},
	{"PLACE 0,0,NORTH \n LEFT \n REPORT", "0, 0, WEST"},
	{"PLACE 1,2,EAST \n MOVE \n MOVE \n LEFT \n MOVE \n REPORT", "3, 3, NORTH"},
}

func TestRun(t *testing.T) {
	for i, rt := range runTests {
		table := robot.NewTable(5, 5)
		r := robot.New(table)
		c := robot.NewController(r)

		input := new(bytes.Buffer)
		output := new(bytes.Buffer)

		io.WriteString(input, rt.commands)
		run(c, input, output, false)
		if got := strings.TrimSpace(output.String()); got != rt.want {
			t.Errorf("test case %d resulted in %s; want: %s", i, got, rt.want)
		}
	}
}
