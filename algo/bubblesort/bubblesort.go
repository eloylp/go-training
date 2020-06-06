package bubblesort

func BubbleSortInt(numbers []int) {
	maxIndex := len(numbers) - 1
	for s := 0; s < len(numbers); s++ {
		sweepWithChanges := false
		for i := 0; i < maxIndex-s; i++ {
			if i < maxIndex && numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				sweepWithChanges = true
			}
		}
		if !sweepWithChanges {
			return
		}
	}
}
