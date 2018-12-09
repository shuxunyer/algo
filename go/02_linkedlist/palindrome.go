package _2_linkedlist

func isPalindrome1(list *LinkedList) bool {
	lLen := list.len
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}
	s := make([]string, 0, lLen/2)
	cur := list.head
	for i := uint(1); i < lLen; i++ {
		cur = cur.next
		if lLen%2 != 0 && i == (lLen/2+1) {
			continue
		}

		if i <= lLen/2 {
			s = append(s, cur.val.(string))
		} else {
			if s[lLen-i] != cur.val.(string) {
				return false
			}
		}
	}
	return true

}
