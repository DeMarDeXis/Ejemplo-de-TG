package hasher

type HashSet map[string]struct{}

// NewHashSet creates a new HashSet and adds elements to it.
func NewHashSet(elements ...string) HashSet {
	s := HashSet{}
	for _, e := range elements {
		s.Add(e)
	}
	return s
}

// Add adds element to the set.
func (s HashSet) Add(element string) {
	s[element] = struct{}{}
}

// Contains checks if element is in the set.
func (s HashSet) Contains(element string) bool {
	_, ok := s[element]
	return ok
}

// Remove removes element from the set.
func (s HashSet) Remove(e string) {
	delete(s, e)
}

// Slice return elements of the set as slice with the same type. Order of elements is not guaranteed.
func (s HashSet) Slice() []string {
	slice := make([]string, 0, len(s))
	for e := range s {
		slice = append(slice, e)
	}
	return slice
}
