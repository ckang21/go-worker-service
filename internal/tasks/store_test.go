package tasks

import "testing"

func TestStoreCreateAndList(t *testing.T) {
	s := NewStore()
	id := NewID()
	s.Create(Task{ID: id, Title: "t1", Status: StatusQueued})
	got := s.List()
	if len(got) != 1 || got[0].ID != id {
		t.Fatalf("want 1 task with id=%s; got=%v", id, got)
	}
}
