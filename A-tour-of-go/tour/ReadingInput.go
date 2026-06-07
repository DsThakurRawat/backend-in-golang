package tour

import (
	"bufio"
	"fmt"
	"os"
)

// ReadInput reads multiple whitespace-separated values in one go.
// fmt.Scan splits on whitespace, so it CANNOT read a line containing spaces
// into a single string variable.
func ReadInput() {
	var name string
	var age int
	fmt.Print("Enter name and age (space-separated): ")
	fmt.Scan(&name, &age)  // taking multiple inputs
	fmt.Println(name, age) // Println so they're spaced apart
}

// ReadAgain reads a single full line, including spaces.
// This is the preferred way to read a whole line of input.
func ReadAgain() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a line: ")
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)
}

// ReadingMultipleLine keeps reading lines until EOF (Ctrl+D on Unix/macOS,
// Ctrl+Z then Enter on Windows).
func ReadingMultipleLine() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Got:", line)
	}
}

