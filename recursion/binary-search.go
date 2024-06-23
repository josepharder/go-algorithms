package recursion

import "errors"

// BinarySearch is implemented recursively. Its time complexity is O(log n). Use it when items are sorted and a fast algorithm is necessary.
func BinarySearch(items []int, item int, low int, high int) (*int, error) {
	if low > high {
		return nil, errors.New("not found")
	}

	if len(items) == 0 {
		return nil, errors.New("empty list")
	}

	mid := (low + high) / 2
	guess := items[mid]

	if guess == item {
		return &mid, nil
	}

	if guess > item {
		high = mid - 1
	} else {
		low = mid + 1
	}

	return BinarySearch(items, item, low, high)
}
