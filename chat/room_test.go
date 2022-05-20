package chat

import "testing"

func TestNewRoom(t *testing.T) {
	r := NewRoom(1, "ABC room")
	if r.Num != 1 {
		t.Fail()
	}
	t.Log(r)
}
