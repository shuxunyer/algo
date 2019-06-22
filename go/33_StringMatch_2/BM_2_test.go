package _3_StringMatch_2

import (
	"fmt"
	"testing"
)

func TestGenerateGS(t *testing.T) {
	findStr := "abcacabcbcbacabc"
	suffix, prefix := GenerateGS(findStr[:])
	for i := 1; i < len(findStr); i++ {
		fmt.Printf("suffix[%d]=%d,  prefix[%d]=%t", i, suffix[i], i, prefix[i])
		fmt.Println()
	}
	/**
	abcacabcbcbacabc
	*/
}

func TestBM_2(t *testing.T) {
	sourceSrc := "abcacabcbcbacabc"
	findStc := "bcbc"
	t.Log(BM_2(sourceSrc, findStc))
}
