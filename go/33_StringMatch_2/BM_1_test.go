package _3_StringMatch_2

import (
	"fmt"
	"testing"
)

func TestGenerateBC(t *testing.T) {
	hashTB := GenerateBC("abcacabcbcbacabc")
	for i := 0; i < len(hashTB); i++ {
		fmt.Print(rune(hashTB[i]), ",")
	}
}

func TestStringMatch_BM(t *testing.T) {
	t.Log(BM_1("abcacbcdcbaceab", "cba"))
	t.Log(BM_1("abcacbcdcbaceab", "abc"))
	t.Log(BM_1("abcacbcdcbaceab", "baceab"))
	t.Log(BM_1("abcacbcdcbaceab", "bbb"))
	//t.Log(BM_1("abcacbcdcbaceab","ebaceab"))

}
