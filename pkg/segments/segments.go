package main

type segment interface {
	output() string
	len() int
	coloredOutput() string
}
