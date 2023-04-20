package main

import (
	"todo_ditector"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(todo_ditector.Analyzer) }
