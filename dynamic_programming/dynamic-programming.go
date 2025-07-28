package dynamicprogramming

import (
	"fmt"
)

func demonstrateDP(groups []int, constraint int) {
	fmt.Println("=== DP Learning: Your String Problem Simplified ===")

	fmt.Printf("Groups: %v\n", groups)
	fmt.Printf("Need sum >= %d\n\n", constraint)

	maxSum := 0
	for _, g := range groups {
		maxSum += g
	}

	fmt.Printf("Maximum possible sum: %d\n\n", maxSum)

	dp := make([][]int, len(groups)+1)
	for i := range dp {
		dp[i] = make([]int, maxSum+1)
	}

	dp[0][0] = 1

	fmt.Println("DP Table Evolution:")
	fmt.Println("dp[group][sum] = number of ways")

	printTable := func(step int) {
		fmt.Printf("\nAfter processing %d groups:\n", step)
		fmt.Print("sum:  ")
		for sum := 0; sum <= maxSum; sum++ {
			fmt.Printf("%2d ", sum)
		}
		fmt.Println()

		for i := 0; i <= step; i++ {
			fmt.Printf("g%d:   ", i)
			for sum := 0; sum <= maxSum; sum++ {
				if dp[i][sum] == 0 {
					fmt.Print(" . ")
				} else {
					fmt.Printf("%2d ", dp[i][sum])
				}
			}
			fmt.Println()
		}
	}

	printTable(0)

	for i := 0; i < len(groups); i++ {
		fmt.Printf("\n--- Processing group %d (size %d) ---\n", i+1, groups[i])

		for sum := 0; sum <= maxSum; sum++ {
			if dp[i][sum] == 0 {
				continue
			}

			fmt.Printf("From dp[%d][%d]=%d, we can:\n", i, sum, dp[i][sum])

			for pick := 1; pick <= groups[i] && sum+pick <= maxSum; pick++ {
				newSum := sum + pick
				fmt.Printf("  Pick %d chars -> dp[%d][%d] += %d\n",
					pick, i+1, newSum, dp[i][sum])
				dp[i+1][newSum] += dp[i][sum]
			}
		}

		printTable(i + 1)
	}

	result := 0
	fmt.Printf("\nCounting combinations with sum >= %d:\n", constraint)
	for sum := constraint; sum <= maxSum; sum++ {
		if dp[len(groups)][sum] > 0 {
			fmt.Printf("Sum %d: %d ways\n", sum, dp[len(groups)][sum])
			result += dp[len(groups)][sum]
		}
	}

	fmt.Printf("\nTotal valid combinations: %d\n", result)
}

func simpleRecursiveVersion(groups []int, constraint int) int {
	fmt.Println("\n=== Comparing with Recursive (for small input) ===")

	count := 0
	var backtrack func(groupIndex, currentSum int, path []int)
	backtrack = func(groupIndex, currentSum int, path []int) {
		if groupIndex == len(groups) {
			if currentSum >= constraint {
				fmt.Printf("Valid: %v (sum=%d)\n", path, currentSum)
				count++
			}
			return
		}

		for pick := 1; pick <= groups[groupIndex]; pick++ {
			newPath := append(path, pick)
			backtrack(groupIndex+1, currentSum+pick, newPath)
		}
	}

	backtrack(0, 0, []int{})
	fmt.Printf("Recursive result: %d\n", count)
	return count
}

func visualTrace() {
	groups := []int{2, 2}
	maxSum := 4

	dp := make([][]int, 3)
	for i := range dp {
		dp[i] = make([]int, maxSum+1)
	}
	dp[0][0] = 1

	fmt.Println("=== Visual DP Execution Trace ===")
	fmt.Printf("Groups: %v, MaxSum: %d\n\n", groups, maxSum)

	step := 1

	printStep := func(description string) {
		fmt.Printf("Step %d: %s\n", step, description)
		fmt.Print("    sum: ")
		for sum := 0; sum <= maxSum; sum++ {
			fmt.Printf("%2d ", sum)
		}
		fmt.Println()

		for i := 0; i < len(dp); i++ {
			fmt.Printf("    g%d:   ", i)
			for sum := 0; sum <= maxSum; sum++ {
				if dp[i][sum] == 0 {
					fmt.Print(" . ")
				} else {
					fmt.Printf("%2d ", dp[i][sum])
				}
			}
			fmt.Println()
		}
		fmt.Println()
		step++
	}

	printStep("Initial state - dp[0][0] = 1")

	// Outer loop: i = 0 (Group 1)
	fmt.Println("🔄 OUTER LOOP: i = 0 (Processing Group 1, size 2)")

	for sum := 0; sum <= maxSum; sum++ {
		fmt.Printf("  🔄 MIDDLE LOOP: sum = %d, checking dp[0][%d] = %d\n", sum, sum, dp[0][sum])

		if dp[0][sum] == 0 {
			fmt.Printf("    ⏭️  Skip (no ways to reach this state)\n\n")
			continue
		}

		fmt.Printf("    ✅ Process (has %d ways)\n", dp[0][sum])

		for pick := 1; pick <= groups[0]; pick++ {
			newSum := sum + pick
			if newSum <= maxSum {
				fmt.Printf("      🔄 INNER LOOP: pick = %d\n", pick)
				fmt.Printf("        newSum = %d + %d = %d\n", sum, pick, newSum)
				fmt.Printf("        dp[1][%d] += dp[0][%d] → dp[1][%d] += %d\n",
					newSum, sum, newSum, dp[0][sum])

				oldValue := dp[1][newSum]
				dp[1][newSum] += dp[0][sum]

				fmt.Printf("        Result: dp[1][%d] = %d → %d\n",
					newSum, oldValue, dp[1][newSum])

				printStep(fmt.Sprintf("Updated dp[1][%d]", newSum))
			}
		}
	}

	fmt.Println("✅ Finished processing Group 1")

	// Outer loop: i = 1 (Group 2)
	fmt.Println("🔄 OUTER LOOP: i = 1 (Processing Group 2, size 2)")

	for sum := 0; sum <= maxSum; sum++ {
		fmt.Printf("  🔄 MIDDLE LOOP: sum = %d, checking dp[1][%d] = %d\n", sum, sum, dp[1][sum])

		if dp[1][sum] == 0 {
			fmt.Printf("    ⏭️  Skip (no ways to reach this state)\n\n")
			continue
		}

		fmt.Printf("    ✅ Process (has %d ways)\n", dp[1][sum])

		for pick := 1; pick <= groups[1]; pick++ {
			newSum := sum + pick
			if newSum <= maxSum {
				fmt.Printf("      🔄 INNER LOOP: pick = %d\n", pick)
				fmt.Printf("        newSum = %d + %d = %d\n", sum, pick, newSum)
				fmt.Printf("        dp[2][%d] += dp[1][%d] → dp[2][%d] += %d\n",
					newSum, sum, newSum, dp[1][sum])

				oldValue := dp[2][newSum]
				dp[2][newSum] += dp[1][sum]

				fmt.Printf("        Result: dp[2][%d] = %d → %d",
					newSum, oldValue, dp[2][newSum])

				if oldValue > 0 {
					fmt.Printf(" ⭐ (ACCUMULATION!)")
				}
				fmt.Println()

				printStep(fmt.Sprintf("Updated dp[2][%d]", newSum))
			}
		}
	}

	fmt.Println("✅ Finished processing Group 2")
	fmt.Println("🎯 Final DP table represents all possible ways!")
}

const modulus = 1000000007

func possibleStringCount(word string, k int) int {
	wordLength := len(word)
	currentConsecutiveCount := 1
	var consecutiveGroupLengths []int

	for i := 1; i < wordLength; i++ {
		if word[i] == word[i-1] {
			currentConsecutiveCount++
		} else {
			consecutiveGroupLengths = append(consecutiveGroupLengths, currentConsecutiveCount)
			currentConsecutiveCount = 1
		}
	}
	consecutiveGroupLengths = append(consecutiveGroupLengths, currentConsecutiveCount)

	totalPossibilities := 1
	for _, groupLength := range consecutiveGroupLengths {
		totalPossibilities = totalPossibilities * groupLength % modulus
	}

	if len(consecutiveGroupLengths) >= k {
		return totalPossibilities
	}

	cumulativeDP := make([]int, k)

	for i := range cumulativeDP {
		cumulativeDP[i] = 1
	}

	for groupIndex := 0; groupIndex < len(consecutiveGroupLengths); groupIndex++ {

		currentDP := make([]int, k)

		for segmentCount := 1; segmentCount < k; segmentCount++ {
			currentDP[segmentCount] = cumulativeDP[segmentCount-1]
			if segmentCount-consecutiveGroupLengths[groupIndex]-1 >= 0 {
				currentDP[segmentCount] = (currentDP[segmentCount] - cumulativeDP[segmentCount-consecutiveGroupLengths[groupIndex]-1] + modulus) % modulus
			}
		}

		newCumulativeDP := make([]int, k)
		newCumulativeDP[0] = currentDP[0]

		for segmentCount := 1; segmentCount < k; segmentCount++ {
			newCumulativeDP[segmentCount] = (newCumulativeDP[segmentCount-1] + currentDP[segmentCount]) % modulus
		}

		currentDP, cumulativeDP = currentDP, newCumulativeDP
	}

	return (totalPossibilities - cumulativeDP[k-1] + modulus) % modulus
}
