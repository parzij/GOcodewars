package kata

import "fmt"

func CountBits(n uint) int {
  result := ""
  answer := 0
  num := 0
  for n > 0 {
    num = int(n % 2)
    n /= 2
    result += fmt.Sprint(num)
  }
  
  for i := 0; i < len(result); i++ {
    if result[i] == '1' {
      answer++
    }
  }
  
  return answer
}
