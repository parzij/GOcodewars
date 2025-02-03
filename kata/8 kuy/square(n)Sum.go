package kata

func SquareSum(num []int) int {
  answer := 0
  for i := 0; i < len(num); i++ {
    answer += num[i] * num[i]
  }
  return answer
}
