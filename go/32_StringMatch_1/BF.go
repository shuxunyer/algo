package _2_StringMatch_1

import "github.com/ethereum/go-ethereum/log"

/**
BF算法 Brute Force 暴力匹配算法，也叫朴素匹配算法

*/
func StringMatch_BF(sourceStr, findStr string) (bool, int) {
	if len(sourceStr) <= len(findStr) {
		log.Warn("findStr.len must less than sourceStr.len",
			"sourceStr.len", len(sourceStr), "findStr.len", len(findStr))
		return false, -1
	}
	m := len(findStr)
	n := len(sourceStr)

	for i := 0; i <= n-m; i++ {
		if findStr == sourceStr[i:i+m] {
			return true, i
		}
	}
	return false, -1
}
