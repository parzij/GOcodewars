package kata

func PositiveSum(num []int) int {
  sum := 0
  for i := 0; i < len(num); i++ {
    if num[i] > 0 {
      sum += num[i]
    }
}
  
  return sum
}
