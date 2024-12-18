// Time: O(n); Space: O(n)
func LongestPalindrome(s string) string {
	m := prepare(s)
	p := make([]int, len(m))
	r, c := 0, 0

	for i := 1; i < len(m)-1; i++ {
		if r > i {
			p[i] = min(r-i, p[2*c-i])
		}

		for m[i+1+p[i]] == m[i-1-p[i]] {
			p[i]++
		}

		if i+p[i] > r {
			c, r = i, i+p[i]
		}
	}

	var pCenter, pMax int
	for i := 1; i < len(m)-1; i++ {
		if p[i] > pMax {
			pCenter, pMax = i, p[i]
		}
	}

	start := (pCenter - pMax) / 2
	return s[start : start+pMax]
}

func prepare(s string) string {
	var b strings.Builder
	b.Grow(2*len(s) + 3)

	b.WriteRune('^')
	for _, r := range s {
		b.WriteRune('#')
		b.WriteRune(r)
	}
	b.WriteRune('#')
	b.WriteRune('$')
	return b.String()
}
