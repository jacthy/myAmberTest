package common

import "testing"

func TestIterator(t *testing.T) {
	ids := make([]int, 9999)
	sum := 0
	Iterator(DefaultPageSiz, len(ids), func(offset, limit int) {
		sum += len(ids[offset:limit])
	})
	if sum != len(ids) {
		t.Errorf("sum should equal to len of ids")
	}
}
