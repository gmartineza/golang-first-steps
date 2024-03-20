package sum

type SumService struct{}

func (s *SumService) CalculateSum(numbers []int) int {
    sum := 0
    for _,num := range numbers {
        sum += num
    }
    return sum
}
