package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := [26]int{}
	s2Map := [26]int{}
	count := 0
	for i := range s1 {
		s1Map[s1[i]-'a']++
		s2Map[s2[i]-'a']++
	}
	for k := 0; k < 26; k++ {
		if s1Map[k] == s2Map[k] {
			count++
		}
	}
	offset := len(s1)
	for lp := 0; lp+offset < len(s2); lp++ {
		if count == 26 {
			return true
		}

		rk := s2[lp+offset] - 'a'
		if s2Map[rk] == s1Map[rk] {
			count--
		}
		s2Map[rk]++
		if s2Map[rk] == s1Map[rk] {
			count++
		}

		lk := s2[lp] - 'a'
		if s2Map[lk] == s1Map[lk] {
			count--
		}
		s2Map[lk]--
		if s2Map[lk] == s1Map[lk] {
			count++
		}
	}

	return count == 26
}
