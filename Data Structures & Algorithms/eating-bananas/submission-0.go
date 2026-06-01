func minEatingSpeed(piles []int, h int) int {

	maxPile := 0

	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	left, right := 1, maxPile
	answer := maxPile

	for left <= right {
		mid := left + (right - left) / 2
		totalHours := 0

		for _, pile := range piles {
			totalHours += (pile + mid - 1) / mid
		}

		if totalHours <= h {
			answer = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return answer
}
