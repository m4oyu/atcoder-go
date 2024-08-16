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

// 柱(n) : 1~3000
// xy : 0~5000
// 3000 x 3000 x 2 = 18,000,000

var pillar_list = make([][]int, 3001)
var pillar_location = make([][]bool, 5001)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < 5001; i++ {
		pillar_location[i] = make([]bool, 5001)
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		y, _ := strconv.Atoi(scanner.Text())

		pillar_list[i] = []int{x, y}
		pillar_location[x][y] = true
	}

	ans := 0

	// 最初に２つの柱を選ぶ
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, left_turn(i, j))
			ans = max(ans, right_turn(i, j))
		}
	}

	fmt.Println(ans)
}

func left_turn(a, b int) int {
	// left turn
	// c: (cx, cy) = bx-(by-ay), by+(bx-ax)
	// d: (dx, dy) = cx-(bx-ax), cy-(by-ay)

	p_ax := pillar_list[a][0]
	p_ay := pillar_list[a][1]
	p_bx := pillar_list[b][0]
	p_by := pillar_list[b][1]
	p_cx := p_bx - (p_by - p_ay)
	p_cy := p_by + (p_bx - p_ax)
	p_dx := p_cx - (p_bx - p_ax)
	p_dy := p_cy - (p_by - p_ay)

	if p_cx < 0 || p_cx > 5000 || p_cy < 0 || p_cy > 5000 || p_dx < 0 || p_dx > 5000 || p_dy < 0 || p_dy > 5000 {
		return 0
	}
	if !pillar_location[p_cx][p_cy] || !pillar_location[p_dx][p_dy] {
		return 0
	}

	return calculateArea(p_ax, p_ay, p_bx, p_by, p_cx, p_cy, p_dx, p_dy)
}

func right_turn(a, b int) int {
	// right turn
	// c: (cx, cy) = bx+(by-ay), by-(bx-ax)
	// d: (dx, dy) = cx-(bx-ax), cy-(by-ay)

	p_ax := pillar_list[a][0]
	p_ay := pillar_list[a][1]
	p_bx := pillar_list[b][0]
	p_by := pillar_list[b][1]
	p_cx := p_bx + (p_by - p_ay)
	p_cy := p_by - (p_bx - p_ax)
	p_dx := p_cx - (p_bx - p_ax)
	p_dy := p_cy - (p_by - p_ay)

	if p_cx < 0 || p_cx > 5000 || p_cy < 0 || p_cy > 5000 || p_dx < 0 || p_dx > 5000 || p_dy < 0 || p_dy > 5000 {
		return 0
	}
	if !pillar_location[p_cx][p_cy] || !pillar_location[p_dx][p_dy] {
		return 0
	}

	return calculateArea(p_ax, p_ay, p_bx, p_by, p_cx, p_cy, p_dx, p_dy)
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
