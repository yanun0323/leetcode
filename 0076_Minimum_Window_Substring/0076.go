package main

//"ADOBECODEBANC", t = "ABC"
func minWindow(s string, t string) string {
	if len(s) < len(t) || len(s) == 0 || len(t) == 0 {
		return ""
	}

	tMap := make(map[byte]int, 52)
	for i := range t {
		tMap[t[i]]++
	}

	matched := 0
	shorter := [2]int{0, len(s) - 1}
	notFound := true
	pMap := make(map[byte]int, 52)
	for lp, rp := 0, 0; rp < len(s); rp++ {
		if tMap[s[rp]] > 0 {
			pMap[s[rp]]++
			if pMap[s[rp]] <= tMap[s[rp]] {
				matched++
			}
		}

		for ; matched == len(t) && (NextFulfill(&lp, pMap, tMap, s)); lp++ {
			if tMap[s[lp]] > 0 {
				pMap[s[lp]]--
				if pMap[s[lp]] < tMap[s[lp]] {
					matched--
				}
			}
		}
		if matched == len(t) && rp-lp <= shorter[1]-shorter[0] {
			notFound = false
			shorter[0] = lp
			shorter[1] = rp
		}
	}
	if notFound {
		return ""
	}
	return s[shorter[0] : shorter[1]+1]
}

// 拿掉現在這個，也會符合t的話
func NextFulfill(lp *int, pMap, tMap map[byte]int, s string) bool {
	return tMap[s[*lp]] == 0 || pMap[s[*lp]] > tMap[s[*lp]]
}
