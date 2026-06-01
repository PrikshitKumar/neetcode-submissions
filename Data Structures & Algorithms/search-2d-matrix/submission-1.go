func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		left, right := 0, len(row) - 1
		
		if target < row[left] || target > row[right] {
			continue
		}

		for left <= right {
			mid := left + (right-left) / 2
			if row[mid] == target {
				return true
			} else if row[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
