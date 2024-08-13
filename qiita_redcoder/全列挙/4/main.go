package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		ai := make([]int, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			ai[j], _ = strconv.Atoi(scanner.Text())
		}
		a[i] = ai
	}

	// 選曲が100*99/2通り
	// 100人の合計をとる
	ans := 0

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += max(a[k][i], a[k][j])
			}
			ans = max(ans, sum)
		}
	}

	fmt.Println(ans)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
