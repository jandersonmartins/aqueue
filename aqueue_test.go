package aqueue

import (
	"reflect"
	"testing"
)

type MockRegister struct {
	Calls []string
}

func (m *MockRegister) Call(str string) {
	m.Calls = append(m.Calls, str)
}

func Test_aqueue(t *testing.T) {
	mock := &MockRegister{}
	strs := []string{"a", "b", "c"}

	q := NewAqueue(len(strs))

	for _, str := range strs {
		func(s string) {
			q.Add(func() {
				mock.Call(s)
			})
		}(str)
	}

	q.Wait()

	if !reflect.DeepEqual(mock.Calls, strs) {
		t.Errorf("expected %q got %q", strs, mock.Calls)
	}
}
