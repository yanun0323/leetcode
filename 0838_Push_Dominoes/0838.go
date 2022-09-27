package main

func pushDominoes(dominoes string) string {
	buf := []byte(dominoes)
	lp := 0
	rp := 0
	for buf[lp] == '.' && lp < len(buf)-1 {
		lp++
	}
	if lp == len(buf)-1 && buf[lp] == '.' {
		return dominoes
	}
	if lp > 0 && buf[lp] == 'L' {
		for i := 0; i < lp; i++ {
			buf[i] = 'L'
		}
	}

	rp = lp + 1
	for ; rp < len(buf); rp++ {
		if buf[rp] == '.' {
			continue
		}

		if buf[lp] == 'R' && buf[rp] == 'L' {
			for l, r := lp+1, rp-1; l < r; l, r = l+1, r-1 {
				buf[l] = 'R'
				buf[r] = 'L'
			}
			lp = rp
			continue
		}

		if buf[lp] == 'R' {
			lp++
			for lp < rp {
				buf[lp] = 'R'
				lp++
			}
			continue
		}

		if buf[rp] == 'L' {
			lp++
			for lp < rp {
				buf[lp] = 'L'
				lp++
			}
			continue
		}
		lp = rp
	}
	if lp < len(buf)-1 && buf[lp] == 'R' {
		for i := lp + 1; i < len(buf); i++ {
			buf[i] = 'R'
		}
	}
	return string(buf)
}

func pushDominoes2(dominoes string) string {
	buf := []byte(dominoes)
	queue := []int{}
	for i := range dominoes {
		if dominoes[i] != '.' {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		for count := len(queue); count > 0; count-- {
			switch buf[queue[0]] {
			case 'R':
				if queue[0]+1 < len(buf) {
					switch buf[queue[0]+1] {
					case 'l':
						buf[queue[0]+1] = '+'
					case '.':
						buf[queue[0]+1] = 'r'
						queue = append(queue, queue[0]+1)
					}
				}
			case 'L':
				if queue[0]-1 >= 0 {
					switch buf[queue[0]-1] {
					case 'r':
						buf[queue[0]-1] = '+'
					case '.':
						buf[queue[0]-1] = 'l'
						queue = append(queue, queue[0]-1)
					}
				}
			}
			queue = queue[1:]
		}
		for i := range buf {
			switch buf[i] {
			case 'l':
				buf[i] = 'L'
			case 'r':
				buf[i] = 'R'
			}
		}
	}

	for i := range buf {
		switch buf[i] {
		case '+':
			buf[i] = '.'
		case 'l':
			buf[i] = 'L'
		case 'r':
			buf[i] = 'R'
		}
	}
	return string(buf)
}
