package kata

func SmallestIntegerFinder(num []int) int {
  var minEl int
  minEl = num[0]
  
  for i := 1; i < len(num); i++ {
    if num[i] < minEl { 
      minEl = num[i]
    }
  }
  return minEl 
}
