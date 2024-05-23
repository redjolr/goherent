package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// os.Exit(cmd.Main())
	/**
	- Position the Cursor:
	\033[<L>;<C>H
		Or
	\033[<L>;<C>f
	puts the cursor at line L and column C.
	- Move the cursor up N lines:
	\033[<N>A
	- Move the cursor down N lines:
	\033[<N>B
	- Move the cursor forward N columns:
	\033[<N>C
	- Move the cursor backward N columns:
	\033[<N>D

	- Clear the screen, move to (0,0):
	\033[2J
	- Erase to end of line:
	\033[K
	*/
	fmt.Println("This is a sentence")
	fmt.Println(os.Getenv("GOMAXPROCS"), runtime.NumCPU())
	fmt.Print("This is another line")
	fmt.Print("\033[1D\033[1D✅")

}
