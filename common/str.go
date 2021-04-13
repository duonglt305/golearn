package common

type Str struct {
	str string
}

func NewStr(str string) *Str {
	s := &Str{str}
	return s
}

func (s *Str) Slugify() {
}
