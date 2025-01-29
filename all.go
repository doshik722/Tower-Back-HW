uniq

package main

import "fmt"

type Three struct {
	value      any
	left_node  *Three
	right_node *Three
}

func (s *Three) Add(value int64) {
	if s.value == nil {
		s.value = value
	} else {
		if value > s.value.(int64) {
			if s.right_node == nil {
				s.right_node = &Three{value: value, right_node: nil, left_node: nil}
			} else {
				s.right_node.Add(value)
			}
		} else if value < s.value.(int64) {
			if s.left_node == nil {
				s.left_node = &Three{value: value, right_node: nil, left_node: nil}
			} else {
				s.left_node.Add(value)
			}
		}

	}
}

func (s *Three) IsExist(value int64) bool {
	if s.value == nil {
		return false
	} else if s.value.(int64) == value {
		return true

	} else {
		if value >= s.value.(int64) {
			if s.right_node == nil {
				return false
			} else {
				return s.right_node.IsExist(value)
			}
		} else {
			if s.left_node == nil {
				return false
			} else {
				return s.left_node.IsExist(value)
			}
		}
	}
}
func (s *Three) Delete(value int64) *Three {
	if s.value == nil { //если дошли, что текущее значение null, то есть не сущ искомого узла
		return nil
	} else if s.value.(int64) == value { //искомый узел найден и это текщий
		if s.right_node == nil && s.left_node == nil {
			s = nil
			return nil
		} else if s.right_node == nil {
			s = s.left_node
			return s
		} else {
			s = s.right_node
			return s
		}
	} else { //продолжение поиска узла
		if value > s.value.(int64) {
			if s.right_node == nil {
				return nil
			} else {
				s.right_node = s.right_node.Delete(value)
				return s
			}
		} else {
			if s.left_node == nil {
				return nil
			} else {
				s.left_node = s.left_node.Delete(value)
				return s
			}
		}
	}
}
func main() {
	var s Three
	fmt.Print(s.IsExist(0))
	s.Add(5)
	s.Add(8)
	s.Add(0)
	s.Add(1)
	fmt.Print(s.IsExist(8))
	s.Add(7)
	s.Add(9)
	s.Add(10)
	fmt.Print(s.Delete(8))
	s.Add(21)
	fmt.Print(s.IsExist(8))

}
main
