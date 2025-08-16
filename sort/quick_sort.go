package sort

func QuickSort(arr []int, start, end int) {
	if end-start+1 <= 3 {
		InsertionSort(arr, start, end)
	} else {
		pivot := partision(arr, start, end)
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
}

func partision(arr []int, start, end int) int {
	middle := start + getMidPoint(start, end)
	sortFirstMidLast(arr, start, middle, end)

	arr[end-1], arr[middle] = arr[middle], arr[end-1] // swap middle to a safe spot position

	pivotIndex := end - 1
	pivotValue := arr[end-1]

	fromLeft := start + 1
	fromRight := end - 2

	done := false
	for !done {
		for arr[fromLeft] < pivotValue {
			fromLeft++
		}

		for arr[fromRight] > pivotValue {
			fromRight--
		}

		if fromLeft < fromRight {
			arr[fromLeft], arr[fromRight] = arr[fromRight], arr[fromLeft]
			fromLeft++
			fromRight--
		} else {
			done = true
		}
	}

	arr[fromLeft], arr[pivotIndex] = arr[pivotIndex], arr[fromLeft] // return pivot to its middle position

	return fromLeft

}

func getMidPoint(start, end int) int {
	return (end - start) / 2
}

func sortFirstMidLast(arr []int, start, middle, end int) {
	if arr[start] > arr[middle] {
		arr[start], arr[middle] = arr[middle], arr[start]
	}
	if arr[start] > arr[end] {
		arr[start], arr[end] = arr[end], arr[start]
	}
	if arr[middle] > arr[end] {
		arr[middle], arr[end] = arr[end], arr[middle]
	}
}
