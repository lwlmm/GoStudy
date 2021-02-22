package main

var a = "G"		// global variable

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a := "O"
	print(a)	// this a is newly initialized
}				//and only lives in the function