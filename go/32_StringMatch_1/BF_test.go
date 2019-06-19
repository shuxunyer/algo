package _2_StringMatch_1

import "testing"

func TestStringMatch_BF(t *testing.T) {
	sourceStr1, findStr1 := "123abc321", "221"
	t.Log(StringMatch_BF(sourceStr1, findStr1))

	sourceStr4, findStr4 := "123abc321", "123"
	t.Log(StringMatch_BF(sourceStr4, findStr4))

	sourceStr2, findStr2 := "123abc321", "c32"
	t.Log(StringMatch_BF(sourceStr2, findStr2))

	sourceStr3, findStr3 := "123abc321", "321"
	t.Log(StringMatch_BF(sourceStr3, findStr3))
}
