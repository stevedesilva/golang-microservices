package utils_test

import (
	"reflect"
	"testing"

	"github.com/stevedesilva/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	els := []int{9, 8, 7, 6, 5}
	utils.BubbleSort(els)

	expected := []int{5, 6, 7, 8, 9}

	assert.NotNil(t, els)

	if !reflect.DeepEqual(expected, els) {
		t.Errorf("not equal, was expecting %v, but got %v", expected, els)
	}

}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	assert.NotNil(t, els)

	expected := []int{4, 3, 2, 1, 0}
	if !reflect.DeepEqual(expected, els) {
		t.Errorf("not equal, was expecting %v, but got %v", expected, els)
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		utils.BubbleSort(els)
	}
}

func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		utils.Sort(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		utils.BubbleSort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		utils.Sort(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		utils.BubbleSort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		utils.Sort(els)
	}
}
