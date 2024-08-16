package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

// 買い物客の目的となる品物で全探索？
//
// -> a1 ~ an の中央値が入口、出口も同様 -> わかるわけない
// |𝑥 − 𝑎1| + |𝑥 − 𝑎2| + ⋯ + |𝑥 − 𝑎𝑁| の最小値が a1 ~ an の中央値であることを知ってればいけそう
// 解説: https://img.atcoder.jp/s8pc-6/editorial.pdf?_gl=1*yska7r*_ga*MzQ3ODE0NjI1LjE3MjMyNjcxMDE.*_ga_RC512FD18N*MTcyMzc4NzYyMS4xOC4xLjE3MjM3ODgxNTkuMC4wLjA.
func main() {
	// scanner.Scan()
	// n, _ := strconv.Atoi(scanner.Text())

	ans := 0

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateArea(x1, y1, x2, y2, x3, y3, x4, y4 int) int {
	// Shoelace formula calculation
	area := math.Abs(float64(
		(x1*y2 + x2*y3 + x3*y4 + x4*y1) -
			(y1*x2 + y2*x3 + y3*x4 + y4*x1),
	))

	return int(area / 2)
}
