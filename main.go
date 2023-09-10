package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("Usage: go-levd <string1> <string2>")
	}
	strA, strB := flag.Arg(0), flag.Arg(1)
	log.Printf(
		"Levenshtein distance between %s and %s is %d",
		strA,
		strB,
		levenshteinDistance(strA, strB),
	)
}

func levenshteinDistance(strA, strB string) int {
	d := make([][]int, len(strA)+1)
	for i := range d {
		d[i] = make([]int, len(strB)+1)
		d[i][0] = i
		for j := range d[i] {
			d[0][j] = j
		}
	}
	for i := 1; i <= len(strA); i++ {
		for j := 1; j <= len(strB); j++ {
			var cost = 1
			if strA[i-1] == strB[j-1] {
				cost = 0
			}
			d[i][j] = min(
				d[i-1][j]+1,      // deletion
				d[i][j-1]+1,      // insertion
				d[i-1][j-1]+cost, // substitution
			)
		}
	}
	return d[len(strA)][len(strB)]
}
