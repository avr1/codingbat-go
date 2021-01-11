package list2

// Return the number of even ints in the given array. Note: the % "mod" operator
// computes the remainder, e.g. 5 % 2 is 1.
func CountEvens(a []int) int {
	var num int
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			num++
		}
	}
	return num
}

// Return the sum of the numbers in the array, returning 0 for an empty array.
// Except the number 13 is very unlucky, so it does not count and numbers that
// come immediately after a 13 also do not count.
func Sum13(a []int) int {
	var sum int
	for i := 0; i < len(a); i++ {
		if a[i] != 13 {
			sum += a[i]
		} else {
			i++
		}
	}
	return sum
}

// Given an array length 1 or more of ints, return the difference between the
// largest and smallest values in the array.
func BigDiff(a []int) int {
	max, min := a[0], a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		} else if a[i] < min {
			min = a[i]
		}
	}
	return max - min
}

// Return the sum of the numbers in the array, except ignore sections of numbers
//starting with a 6 and extending to the next 7 (every 6 will be followed by at
// least one 7). Return 0 for no numbers.
func Sum67(a []int) int {
	isRange, sum := false, 0
	for i := 0; i < len(a); i++ {
		if a[i] == 6 {
			isRange = true
		}
		if !isRange {
			sum += a[i]
		}
		if a[i] == 7 && isRange {
			isRange = false
		}
	}
	return sum
}

// Return the "centered" average of an array of ints, which we'll say is the mean
// average of the values, except ignoring the largest and smallest values in the array.
// If there are multiple copies of the smallest value, ignore just one copy, and
// likewise for the largest value. Use int division to produce the final average.
// You may assume that the array is length 3 or more.
func CenteredAverage(a []int) int {
	max, min, sum := a[0], a[0], a[0]
	for i := 1; i < len(a); i++ {
		sum += a[i]
		if a[i] > max {
			max = a[i]
		} else if a[i] < min {
			min = a[i]
		}
	}
	return (sum - max - min) / (len(a) - 2)
}

// Given an array of ints, return true if the array contains a 2 next to a 2 somewhere.
func Has22(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == 2 && a[i+1] == 2 {
			return true
		}
	}
	return false
}
