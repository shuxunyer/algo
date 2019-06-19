package _2_StringMatch_1

import (
	"testing"
)

func init() {
	Init()
}

func TestHash(t *testing.T) {
	/**
	1a3 = 1 + 11*12 +3*12*12=1+132+432=565
	a31 = 11 + 3*12 +1*12*12=191
	*/
	t.Log(Hash("1a3"))
	t.Log(Hash("a31"))
}

func TestGetHashFromPrev(t *testing.T) {
	/**
	1a3 = 1 + 11*12 +3*12*12=1+132+432=565
	a3f = 11 + 3*12 +1*12*12=191
	*/
	nowHash := getHashFromPrev("1a3", "a3f", Hash("1a3"))
	if nowHash == Hash("a3f") {
		t.Log("TestgetHashFromPrev correct.")
	} else {
		t.Log("TestgetHashFromPrev error.")
	}
}

func TestStringMatch_RK(t *testing.T) {
	sourceStr1, findStr1 := "123ab6321", "221"
	t.Log(StringMatch_RK(sourceStr1, findStr1))

	sourceStr4, findStr4 := "123ab6321", "123"
	t.Log(StringMatch_RK(sourceStr4, findStr4))

	sourceStr2, findStr2 := "123ab6321", "632"
	t.Log(StringMatch_RK(sourceStr2, findStr2))

	sourceStr5, findStr5 := "123ab6321", "ab6"
	t.Log(StringMatch_RK(sourceStr5, findStr5))

	sourceStr3, findStr3 := "123ab6321", "321"
	t.Log(StringMatch_RK(sourceStr3, findStr3))
}
