package kata

func DigitalRoot(num int) int {
  for num >= 10 {
    var sum int
    for num > 0 {
      sum += num % 10
      num /= 10
    }
    num = sum
  }
  return num
}
