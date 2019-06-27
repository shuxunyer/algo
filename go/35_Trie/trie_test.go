package _5_Trie

import "testing"

func TestFind(t *testing.T) {
	Insert("hello")
	Insert("hi")
	Insert("yello")
	Insert("allei")
	Insert("robin")

	t.Log(Find("robi"))
	t.Log(Find("yello"))
	t.Log(Find("pass"))
	t.Log(Find("hi"))
}
