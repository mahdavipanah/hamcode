// Encode, decode and correct (single bit) a binary code using Hamming code.
// Usage: hamcode [command] [binary code]
//
// Available Commands:
//  correct            Print the corrected binary code
//  encode             Print the encoded data binary using Hamming code
//  decode             Print the data binary code inside the input Hamming code
//  help, -h, --help   Print the help
//
// Author: Hamidreza Mahdavipanah
// Repository: github.com/mahdavipanah/hamcode
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Checks if program has any argument
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// Checks if help output is asked
	if os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "help" {
		printHelp()
		return
	}

	// Checks if no valid command entered
	if os.Args[1] != "correct" && os.Args[1] != "decode" && os.Args[1] != "encode" {
		fmt.Fprintf(os.Stderr, "%v: Unknown command\nSee 'hamcode help'.\n", os.Args[1])
		os.Exit(1)
	}

	// Checks if program has at least 3 arguments
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "%v: Command needs a binary code.", os.Args[1])
		os.Exit(1)
	}

	bin := os.Args[2]

	if !isValidBinary(bin) {
		fmt.Fprintln(os.Stderr, "Invalid binary code.")
		os.Exit(1)
	}

	// Determines if the code should interpreted right to left
	rtl := false

	// Checks if there is any option for the command
	if len(os.Args) > 3 {
		if os.Args[3] == "--rtl" {
			rtl = true
		} else {
			fmt.Fprintf(os.Stderr, "Unknown option '%v'.\nSee 'hamcode help'.\n", os.Args[3])
			os.Exit(1)
		}
	}

	switch os.Args[1] {
	case "correct":
		corrected, _ := correct(bin)
		fmt.Println(corrected)
	case "encode":
		if rtl {
			bin = Reverse(bin)
		}
		fmt.Println(encode(bin))
	case "decode":
		decoded := decode(bin)
		if rtl {
			decoded = Reverse(decoded)
		}
		fmt.Println(decoded)
	}
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Corrects the bin and returns the corrected binary and error bit position.
// Returns bin and 0 if there are no errors.
func correct(bin string) (string, int) {
	// Parities
	p := make([]int, 0)

	// Finds parities and calculates their values
	for pos, _ := range bin {
		pos++

		if isPerfectSquare(pos) {
			c := 0
			for i, _ := range bin {
				i++
				// Checks if the bit should be calculated
				if i != pos && (i&pos != 0) {
					c ^= cToB(bin[i-1])
				}
			}

			// Appends the calculated parity bit
			p = append(p, c^cToB(bin[pos-1]))
		}
	}

	errPos := errorPosition(p) - 1

	// Checks if there was any error in the binary code
	if errPos != -1 {
		c := "0"

		if bin[errPos] == '0' {
			c = "1"
		}

		// Corrects the binary code
		bin = bin[:errPos] + c + bin[errPos+1:]
	}

	return bin, errPos + 1
}

// Returns the data code inside an Hamming code.
// This function first corrects the code using correct function.
func decode(bin string) (code string) {
	bin, _ = correct(bin)

	for pos, _ := range bin {
		if !isPerfectSquare(pos + 1) {
			code += string(bin[pos])
		}
	}

	return
}

// Returns encoded data code as an Hamming code
func encode(data string) (hcode string) {
	bits := make([]int, 0)

	for pos, i := 0, 0; i < len(data); pos++ {
		if isPerfectSquare(pos + 1) {
			bits = append(bits, 0)
		} else {
			bits = append(bits, cToB(data[i]))
			i++
		}
	}

	for pos, _ := range bits {
		if isPerfectSquare(pos + 1) {
			p := 0
			for i, _ := range bits {
				// Checks if the bit should be calculated
				if i+1 != pos+1 && ((i+1)&(pos+1) != 0) {
					p ^= bits[i]
				}
			}

			bits[pos] = p
		}
	}

	for _, bit := range bits {
		hcode += string(bit + 48)
	}

	return

}

// Prints help to standard output
func printHelp() {
	fmt.Println(`Encode, decode and correct (single bit) a binary code using Hamming code.

Usage: hamcode [command] [binary code]

Available Commands:
  correct            Print the corrected binary code
  encode             Print the encoded data binary using Hamming code
  decode             Print the data binary code inside the input Hamming code
  help, -h, --help   Print the help

Author: Hamidreza Mahdavipanah
Repository: http://github.com/mahdavipanah/hamcode`)
}

// Checks if all characters in a strings are 0 or 1
func isValidBinary(bin string) bool {
	// Checks if all chars are 0 or 1
	for _, char := range bin {
		// Checks if character is 0 or 1
		if char != '0' && char != '1' {
			return false
		}
	}

	return true
}

// Checks if a number is perfect square
func isPerfectSquare(n int) bool {
	return n == (n & -n)
}

// Converts char{0,1} to int{0,1}
func cToB(b byte) int {
	if b == '1' {
		return 1
	}
	return 0
}

// Returns the error bit position based of p slice
func errorPosition(p []int) int {
	str := ""
	for _, val := range p {
		str = string(val+48) + str
	}

	number, _ := strconv.ParseInt(str, 2, 0)

	return int(number)
}
