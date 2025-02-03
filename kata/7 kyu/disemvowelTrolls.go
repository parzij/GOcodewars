package kata

import "strings"

func Disemvowel(str string) string {
  vowel := "aeoiuAEOIU"
  result := ""
  for _, char := range str {
    if !strings.ContainsRune(vowel, char) {
      result += string(char)
    }
  }
  return result

}
