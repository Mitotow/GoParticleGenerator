package particles

import "container/list"

func (s *System) Remove(e *list.Element) {
	s.Content.Remove(e)
}
