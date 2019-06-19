package _2_StringMatch_1

import (
	"github.com/ethereum/go-ethereum/common/math"
)

/**
RK算法的全称为Rabin-Karp 算法，由两位发明人名字来命名，
是BF算法的升级版
*/

var (
	ElementMap = make(map[string]uint8)
	CoutArray  = make([]uint64, strNum, strNum)
)

const (
	strNum = 10
	strMax = 12
)

func initElementMap() {
	ElementMap["1"] = 1
	ElementMap["2"] = 2
	ElementMap["3"] = 3
	ElementMap["4"] = 4
	ElementMap["5"] = 5
	ElementMap["6"] = 6
	ElementMap["7"] = 7
	ElementMap["8"] = 8
	ElementMap["9"] = 9
	ElementMap["10"] = 10

	ElementMap["a"] = 11
	ElementMap["b"] = 12
}

func initCoutArray() {
	for i := 0; i < strNum; i++ {
		CoutArray[i] = math.BigPow(strMax, int64(i)).Uint64()
	}
}

func Init() {
	initElementMap()
	initCoutArray()
}

/**
sourceStr 主串 findStr 模式串
从1-10和a-j共20个元素组成的串
*/
func StringMatch_RK(sourceStr, findStr string) (bool, int) {
	m := len(sourceStr)
	n := len(findStr)
	if n == 0 || m == 0 || n >= m {
		return false, -1
	}
	findStrHash := Hash(findStr)

	var (
		preHash uint64
		preStr  string
		nowStr  string
	)

	for i := 0; i <= m-n; i++ {
		preStr = nowStr
		nowStr = sourceStr[i : n+i]
		nowHash := getHashFromPrev(preStr, nowStr, preHash)
		//fmt.Printf("nowHash=%v,nowStr=%v\n", nowHash,nowStr)
		if nowHash == findStrHash {
			return true, i
		}
		preHash = nowHash

	}
	return false, -1
}

func getHashFromPrev(preStr, str string, prevHash uint64) uint64 {
	if preStr == "" {
		return Hash(str)
	}
	if len(str) == 1 {
		return Hash(str)
	}
	temp := prevHash - getNum(0, preStr[:1])
	sum := temp/strMax + getNum(len(str)-1, str[len(str)-1:len(str)])
	return sum
}

/**
Hash逻辑： 1a3f4d5bc6
假设3个一组：1a3 = 1 + 11*36 +3*36*36
	a3f = 11 + 3*36 +16*36*36
*/
func Hash(subStr string) uint64 {
	m := len(subStr)
	if m == 0 {
		return 0
	}
	var sum uint64
	for i := 0; i < m; i++ {
		sum += getNum(i, subStr[i:i+1])
	}
	return sum
}

func getNum(index int, s string) uint64 {
	return CoutArray[index] * uint64(ElementMap[s])
}
