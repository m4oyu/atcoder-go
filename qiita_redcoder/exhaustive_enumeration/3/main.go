package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	const buf = 200100
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, buf), buf)
}

func debug() {
	inputFile, err := os.Open("./.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no inputFile.")
		os.Exit(1)
	}

	scanner = bufio.NewScanner(inputFile)
}

func main() {
	scanner.Scan()
	s := scanner.Text()

	// fmt.Println(len(s))

	re, _ := regexp.Compile("^[ACGT]+$")
	ans := 0
	for i := 0; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			// fmt.Println(s[i:j])
			if re.MatchString(s[i:j]) && ans < j-i {
				ans = j - i
			}
		}
	}

	fmt.Println(ans)
}
