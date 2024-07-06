package main

// Brute force
func maxSubArr(arr []int, sumArr int) (interface{}, interface{}) {
	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			if sum == sumArr {
				return i + 1, j + 1
			}
		}
	}

	return -1, nil
}

//O(n)

func maxSubArrOpt(arr []int, sumArr int) (interface{}, interface{}) {
	start, sum := 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum == sumArr {
			return start + 1, i + 1
		} else if sum > sumArr {
			sum -= arr[start]
			start++
			if sum == sumArr {
				return start + 1, i + 1
			}
		}
	}
	return -1, nil
}
