package main

func largestRange(arr []int) []int {
	var formerLargestRange []int
	currentLargestRange := []int{arr[0]}
	for index := 1; index <= len(arr)-1; index++ {
		currentNum := arr[index]
		previousNum := arr[index-1]
		if currentNum >= previousNum {
			currentLargestRange = append(currentLargestRange, currentNum)
		} else if len(currentLargestRange) > len(formerLargestRange) {
			formerLargestRange = currentLargestRange
			currentLargestRange = []int{}
		} else {
			currentLargestRange = []int{}
		}
	}
	if len(formerLargestRange) > len(currentLargestRange) {
		return formerLargestRange
	}
	return currentLargestRange
}
