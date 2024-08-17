package main

import (
	"bufio"
	"fmt"
	"math"
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

// 星座を構成する星の上限m: 200
// 星座の数n: 1000

var star_sign [][]int = make([][]int, 200)
var star_sign_layout [][]int = make([][]int, 200)
var star_list [][]int = make([][]int, 1000)
var star_presence map[[2]int]bool = make(map[[2]int]bool)

func main() {
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < m; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		star_sign[i] = []int{x, y}

		if i != 0 {
			star_sign_layout[i] = []int{
				x - star_sign[i-1][0],
				y - star_sign[i-1][1],
			}
		}
	}

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	// for i := 0; i <= 1000000; i++ {
	// 	star_presence[i] = make([]bool, 1000001)
	// }
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())
		star_list[i] = []int{x, y}
		star_presence[[2]int{x, y}] = true
	}

	ansx := 0
	ansy := 0

	for i := 0; i < n; i++ {
		x := star_list[i][0]
		y := star_list[i][1]

		for j := 1; j < m; j++ {
			x += star_sign_layout[j][0]
			y += star_sign_layout[j][1]

			if outOfArea(x) || outOfArea(y) {
				break
			}
			if !star_presence[[2]int{x, y}] {
				break
			}

			if j == m-1 {
				ansx = star_list[i][0] - star_sign[0][0]
				ansy = star_list[i][1] - star_sign[0][1]
			}
		}
	}

	fmt.Printf("%d %d\n", ansx, ansy)
}

func outOfArea(x int) bool {
	if x < 0 || x > 1000000 {
		return true
	}
	return false
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
