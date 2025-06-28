package system

const DisplayWidth = 64
const DisplayHeight = 32

type display [DisplayHeight]uint64

func (d display) visit(yield func(x, y int, on bool)) {
	for y, row := range d {
		for x := range DisplayWidth {
			mask := uint64(1 << (63 - x))
			on := (row & mask) != 0

			yield(x, y, on)
		}
	}
}

func (d *display) clear() {
	*d = display{}
}
