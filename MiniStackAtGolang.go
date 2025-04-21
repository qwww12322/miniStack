package miniStack

type Stack[t any] struct {
	objs []t
}

func NewStack[t any](arr []t) *Stack[t] {
	return &Stack[t]{
		objs: arr,
	}
}

func (s *Stack[t]) Pop() (res t, ok bool) {
	if len(s.objs) == 0 {
		ok = false
	} else {
		res = s.objs[len(s.objs)-1]
		s.objs = s.objs[:len(s.objs)-1]
	}
	return
}
