package system

type stack struct {
	data []uint16
}

func (s *stack) push(value uint16) {
	s.data = append(s.data, value)
}

func (s *stack) pop() uint16 {
	head := len(s.data) - 1
	removed := s.data[head]

	s.data = s.data[:head]

	return removed
}
