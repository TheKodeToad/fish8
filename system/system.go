package system

import (
	"fmt"
	"time"
)

const ramSize = 0x1000
const programOffset = 0x200

type System struct {
	ram   [ramSize]byte
	stack stack
	cpu   cpu

	display display
	keypad  KeySet

	cpuTicker     *time.Ticker
	displayTicker *time.Ticker
}

const maxProgramSize = ramSize - programOffset

func New(program []byte) (System, error) {
	if len(program) > maxProgramSize {
		return System{}, fmt.Errorf("program is too large - max is %d B", maxProgramSize)
	}

	result := System{}

	copy(result.ram[programOffset:], program)
	result.cpu.programCounter = programOffset

	copy(result.ram[fontOffset:], font[:])

	return result, nil
}

// Tick pauses until an update is needed and returns true if a repaint is needed.
func (s *System) Tick() bool {
	if s.cpuTicker == nil {
		s.cpuTicker = time.NewTicker(time.Second / 700)
	}

	if s.displayTicker == nil {
		s.displayTicker = time.NewTicker(time.Second / 60)
	}

	select {
	case <-s.cpuTicker.C:
		s.cpu.cycle(s)
		return false
	case <-s.displayTicker.C:
		s.cpu.vblank()
		return true
	}
}

// ReadDisplay invokes the provided function on every pixel.
func (s System) ReadDisplay(yield func(x, y int, on bool)) {
	s.display.visit(yield)
}

func (s *System) UpdateKeypad(keys KeySet) {
	s.keypad = keys
}
