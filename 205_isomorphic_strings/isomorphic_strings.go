package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	var StoT = make(map[byte]byte)
	var TtoS = make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		letterS := s[i]
		letterT := t[i]

		_, existS := StoT[letterS]
		_, existT := TtoS[letterT]
		if !existS && !existT {
			StoT[letterS] = letterT
			TtoS[letterT] = letterS
		} else if StoT[letterS] != letterT || TtoS[letterT] != letterS {
			return false
		}
	}

	return true
}
