package main

import "fmt"

func climbStairsH(n int, seen map[int]int) int {

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	if count, saw := seen[n]; saw {
		return count
	}

	minusOneWays := climbStairsH(n-1, seen)

	seen[n-1] = minusOneWays

	minusTwoWays := climbStairsH(n-2, seen)

	seen[n-2] = minusTwoWays

	return minusOneWays+minusTwoWays

}

func climbStairs(n int) int {
 seen := make(map[int]int)

 return climbStairsH(n, seen)
}

func main() {
	fmt.Printf("Ways to jump %d\n", climbStairs(6))
}
