package main

func removeDuplicateLetters(s string) string {
	charPos := make(map[uint8][]int)

	for i := 0; i < len(s); i++ {
		charPos[s[i]] = append(charPos[s[i]], i)
	}

	ret := make([]int, 0, 26)
	lastPos := 0
	for c := uint8('a'); c < 'z'; c++ {
		ps := charPos[c]
		if len(ps) == 0 {
			continue
		}
		i := 0

		for ;i < len(ps); i++ {
			if ps[i] > lastPos {

			}
		}
	}
}
