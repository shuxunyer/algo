package _3_StringMatch_2

/**
坏字符规则
*/
const SIZE = 123

func GenerateBC(findStr string) []int {
	hashTable := make([]int, SIZE, SIZE)
	m := len(findStr)
	if m == 0 {
		return hashTable
	}

	for i := 0; i < SIZE; i++ {
		hashTable[i] = -1
	}
	for i := 0; i < m; i++ {
		ascii := rune(findStr[i])
		hashTable[ascii] = i
	}
	return hashTable
}

func BM_1(sourceStr, findStr string) int {
	m := len(findStr)
	n := len(sourceStr)
	if m >= n {
		return -1
	}
	hashTable := GenerateBC(findStr)
	i := 0
	var j int
	for {
		for j = m - 1; j >= 0; j-- {
			if sourceStr[i+j] != findStr[j] {
				break
			}
		}

		if j < 0 {
			return i
		}
		//将模式串往后移动	j-hashTable[index]
		index := rune(sourceStr[i+j])
		i += j - hashTable[index]
		if i > n-m {
			return -1
		}
	}
}
