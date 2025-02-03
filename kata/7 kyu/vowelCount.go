package kata

import "strings"

func GetCount(str string) (count int) {
    vowels := "aeiou"
    for _, char := range str {
        if strings.ContainsRune(vowels, char) {
            count++
        }
    }
    return count
}
