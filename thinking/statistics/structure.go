package main

import "os"

type config struct {
	path string
	lang []string
}

type fb struct {
	file *os.File
	path string
}

type pcb struct {
	chanFile      chan fb
	chanFileClose chan bool
	conf          *config
}

type scanner interface {
	scan()
}

func newPCB(conf *config) *pcb {
	bufSize := 2048
	return &pcb{
		chanFile:      make(chan fb, bufSize),
		chanFileClose: make(chan bool),
		conf:          conf,
	}
}

func newFB(path string, f *os.File) fb {
	return fb{
		file: f,
		path: path,
	}
}

func matchScanner() {

}
