package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/amir6432/robot/robot"
)

const (
	tableWidth  = 5
	tableHeight = 5
)

func main() {
	table := robot.NewTable(tableWidth, tableHeight)
	r := robot.New(table)
	c := robot.NewController(r)

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		run(c, file, os.Stdout, false)
	} else {
		run(c, os.Stdin, os.Stdout, true)
	}
}

func run(controller robot.Controller, input io.Reader, output io.Writer, verbose bool) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		cmd := scanner.Text()

		res, err := controller.Command(cmd)
		if err != nil {
			if verbose {
				io.WriteString(output, "Error: "+err.Error()+"\n")
			}
		} else if res != "" {
			io.WriteString(output, res+"\n")
		} else if verbose {
			io.WriteString(output, "Accepted\n")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("reading input:", err)
	}
}
