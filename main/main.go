// https://www.hackerrank.com/challenges/sherlock-and-array/problem

package main

import "fmt"

func sum(arr []int32) int32 {
	var sum = int32(0)

	for _, value := range arr {
		sum += value
	}

	return sum
}

func balancedSums(arr []int32) string {
	// Find the middle value
	var midIndex = int32(len(arr) / 2)
	var midValue = arr[midIndex]

	// Find the sum of items at left and right of the middle value
	var sumLeft = sum(arr[0:midIndex])
	var sumRight int32

	// Handle single item array
	if midIndex+1 < int32(len(arr)) {
		sumRight = sum(arr[midIndex+1:])
	}

	// Determine the direction
	var direction int32
	if sumLeft < sumRight {
		direction = 1
	} else if sumLeft > sumRight {
		direction = -1
	}

	for {
		if sumLeft == sumRight {
			// It's balanced
			return "YES"
		}

		// Move the midIndex pointer
		midIndex += direction

		if midIndex < 0 || midIndex >= int32(len(arr)) {
			// Movement is no longer possible
			return "NO"
		}

		var newMidValue = arr[midIndex]

		switch direction {
		case 1:
			// Move to the right
			sumLeft += midValue
			sumRight -= newMidValue

		case -1:
			// Move to the left
			sumLeft -= newMidValue
			sumRight += midValue
		}

		// Change the midIndex value
		midValue = newMidValue
	}
}

func main() {
	fmt.Println(balancedSums([]int32{1, 1, 4, 1, 1}))
	fmt.Println(balancedSums([]int32{2, 0, 0, 0}))
	fmt.Println(balancedSums([]int32{0, 0, 2, 0}))
	fmt.Println(balancedSums([]int32{1, 2, 3}))
}
