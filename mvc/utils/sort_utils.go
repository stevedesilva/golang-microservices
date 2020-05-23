package utils

import "sort"

// BubbleSort alg
func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}

		}
	}

}

// Sort alg
func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
		return
	}
	sort.Ints(elements)
}
