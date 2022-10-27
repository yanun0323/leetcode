package main

// Approach 1 improved: more readable
type StringIterator struct {
	words     []string
	wordIndex int
	charIndex int
}

func (s *StringIterator) HasNext() bool {
	return s.wordIndex < len(s.words) && s.charIndex < len(s.words[s.wordIndex])
}

func (s *StringIterator) Next() {
	s.charIndex++
	if s.charIndex == len(s.words[s.wordIndex]) {
		s.wordIndex++
		s.charIndex = 0
	}
}

func (s *StringIterator) Char() byte {
	return s.words[s.wordIndex][s.charIndex]
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	iter1 := &StringIterator{words: word1}
	iter2 := &StringIterator{words: word2}
	for iter1.HasNext() && iter2.HasNext() {
		if iter1.Char() != iter2.Char() {
			return false
		}
		iter1.Next()
		iter2.Next()
	}
	return !iter1.HasNext() && !iter2.HasNext()
}

// Approach 1
func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	l, li := 0, 0
	r, ri := 0, 0
	for l < len(word1) && li < len(word1[l]) && r < len(word2) && ri < len(word2[r]) {
		if word1[l][li] != word2[r][ri] {
			return false
		}
		li++
		ri++

		if li == len(word1[l]) {
			l++
			li = 0
		}

		if ri == len(word2[r]) {
			r++
			ri = 0
		}
	}
	return l == len(word1) && r == len(word2)
}
