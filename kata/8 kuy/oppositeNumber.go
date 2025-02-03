package kata

func Opposite(value int) int {
  if value > 0 {
    return -value
  } else if value < 0 {
    return value * -1
  }
  return value
}
