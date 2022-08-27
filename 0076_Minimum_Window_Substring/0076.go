package main

//"ADOBECODEBANC", t = "ABC"
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	tMap := make(map[byte]int, 52)
	for i := range t {
		tMap[t[i]]++
	}

	matched := 0
	shorter := [2]int{0, len(s)}
	notFound := true
	pMap := make(map[byte]int, 52)
	for lp, rp := 0, 0; rp < len(s); rp++ {

		if tMap[s[rp]] > 0 {
			pMap[s[rp]]++
			if pMap[s[rp]] <= tMap[s[rp]] {
				matched++
			}
			if pMap[s[rp]] > tMap[s[rp]] && s[rp] == s[lp] {
				matched--
				pMap[s[lp]]--
				lp++
			}
		}

		if matched == 0 {
			lp++
			continue
		}

		if matched == len(t) {
			if rp-lp <= shorter[1]-shorter[0]-1 {
				notFound = false
				shorter[0] = lp
				shorter[1] = rp + 1
			}

			for lp < rp && (matched == len(t) || tMap[s[lp]] == 0) {
				if tMap[s[lp]] > 0 {
					if pMap[s[lp]] <= tMap[s[lp]] {
						matched--
					}
					pMap[s[lp]]--
				}
				lp++
			}
		}

	}
	if notFound {
		return ""
	}
	return s[shorter[0]:shorter[1]]
}
