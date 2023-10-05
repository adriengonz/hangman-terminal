package main

import (
	"os"
)

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func Exit() {
	os.Exit(0)
}