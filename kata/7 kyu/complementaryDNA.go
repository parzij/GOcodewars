package kata

import "strings"

func DNAStrand(dna string) string {
  result := strings.NewReplacer("A", "T", "T", "A", "C", "G", "G", "C")
  return result.Replace(dna)
}
