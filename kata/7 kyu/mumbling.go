package kata

import "strings"

func Accum(s string) string {
    result := ""
    for i := 0; i < len(s); i++ {
        result += strings.ToUpper(string(s[i])) + strings.Repeat(strings.ToLower(string(s[i])), i) + "-"
    }
    return result[:len(result)-1]
}
