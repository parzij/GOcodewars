package kata

import "strconv"

func StringToNumber(str string) int {
  num, _ := strconv.Atoi(str)
  return num
}
