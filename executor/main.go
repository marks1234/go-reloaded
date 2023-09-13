package main

import "os"

// simple error catcher
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	files := os.Args[1:]

	if len(files) > 2 || len(files) == 0 {
		println("Not the right amount of arguments!")
		return
	}

	dat, err := os.ReadFile(files[0])
	check(err)

	println(string(dat))
}
