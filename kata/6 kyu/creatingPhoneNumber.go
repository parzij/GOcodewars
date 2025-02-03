package kata

import "fmt"

func CreatePhoneNumber(num [10]uint) string {
  return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", num[0], num[1], num[2], num[3], num[4], num[5], num[6], num[7], num[8], num[9])
}
