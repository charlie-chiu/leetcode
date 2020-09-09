package dutch_national_flag

const (
	red   = 0
	white = 1
	blue  = 2
)

func sort3Colors(colors []int) {
	var lo, mid, hi = 0, 0, len(colors) - 1

	for mid <= hi {
		switch colors[mid] {
		case red:
			colors[lo], colors[mid] = colors[mid], colors[lo]
			lo++
			mid++
		case white:
			mid++
		case blue:
			colors[hi], colors[mid] = colors[mid], colors[hi]
			hi--
		}
	}
}
