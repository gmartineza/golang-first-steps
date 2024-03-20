package sum_test

import (
	"example.com/sum"
	"testing"
)

func TestSumService_CalculateSum(t *testing.T) {
    sumService :=  sum.SumService{}
    numbers := []int{1, 2, 3, 4, 5}
    expected := 15
    result := sumService.CalculateSum(numbers)
    if result != expected {
        t.Errorf("Sum of %v was incorrect, got: %d, want: %d.", numbers, result, expected)
    }
}
