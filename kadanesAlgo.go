// kadanes algo

package main

func kadanesAlgo(arr []int) int {
	maxSum, sum := 0, 0

	for _, val := range arr {
		sum += val
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
