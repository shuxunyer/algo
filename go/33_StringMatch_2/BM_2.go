package _3_StringMatch_2

func GenerateGS(findStr string) (suffix []int, prefix map[int]bool) {
	m := len(findStr)
	if m == 0 {
		return nil, nil
	}
	suffix = make([]int, m, m)
	prefix = make(map[int]bool)

	for i := 1; i <= m-1; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < m-1; i++ {
		subStr := findStr[m-1-i : m] //后缀串
		k := m - 2
		if subStr == findStr[0:i+1] {
			prefix[len(subStr)] = true
		}
		for {
			sStr := findStr[k-i : k+1] //k+1 =k-i+len(subStr)
			if k-i >= 0 && subStr == sStr {
				suffix[len(subStr)] = k - i
				/*if (k-i == 0) {
					prefix[len(subStr)] = true
				}*/
				break
			}
			k--
			if k-i < 0 {
				break
			}
		}
	}
	return
}

func BM_2(sourceStr, findStr string) int {
	m := len(findStr)
	n := len(sourceStr)
	if m >= n {
		return -1
	}
	hashTable := GenerateBC(findStr)
	suffix, prefix := GenerateGS(findStr)
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
		x := j - hashTable[index]
		y := moveByGS(suffix, prefix, m, j)
		if x > y {
			i += x
		} else {
			i += y
		}
		if i > n-m {
			return -1
		}
	}
}

func moveByGS(suffix []int, prefix map[int]bool, m, j int) int {
	k := m - 1 - j //好后缀长度
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] == true {
			return r
		}
	}
	return m
}
