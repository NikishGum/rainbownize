package main

import (
	"fmt"
	"math"
	"bufio"
	"io"
	"os"
)

// get_rgb returns the next hue of the rainbow using sin function 
func get_rgb (j int) (int, int, int) {
	var speed = 0.1 // Lenght of a single rainbow gradient
	return int (math.Sin(speed * float64(j))* 127 + 128),
	int (math.Sin(speed * float64(j) +  math.Pi / 2) * 127 + 128), 
	int (math.Sin(speed * float64(j) +  math.Pi)* 127 + 128) 
}

// output rainbow rune slice to stdout
func print_rainbow(output []rune) {
	for j := 0; j < len(output); j++  {
		r,g,b := get_rgb(j)

		// \033 - escape sequence
		// 38;2 - subsequences tells, that subsequence after contains color in rgb format
		// [0m - resets color for next rune
        fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}

func main() {
	// Feth information about Stdin file descriptor
	info,_ := os.Stdin.Stat()

	// Added for education puposes
	// println(info.Mode().String())
	// println(os.ModeCharDevice.String())

	var output []rune

	// Using bit mask identify if given file mode is character device  (c bit in file mode)
	if info.Mode() & os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes only")
	}

	// Bufferize stream to iterate throught it
	reader := bufio.NewReader(os.Stdin)

	for {
		// Read one rune (unicode code point)
		input, _, err := reader.ReadRune()

		if err != nil && err == io.EOF {
			break
		}

		// Form string
		output = append(output, input)
	}

	print_rainbow(output)

}