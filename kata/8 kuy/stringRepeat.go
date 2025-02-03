package kata

func RepeatStr(num int, str string) string {
  result := ""
  for i := 0; i < num; i++ {
    result += str
  }
  return result
}
