package system

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type cpu struct {
	index     uint16
	variables [16]byte

	programCounter uint16

	delayTimer byte
	soundTimer byte

	waitFor  waitState
	lastKeys KeySet
}

type waitState int

const (
	waitForNothing waitState = iota
	waitForKeyPress
	waitForKeyRelease
)

func (cpu *cpu) cycle(sys *System) {
	if cpu.handleWaitState(sys) {
		return
	}

	msb, lsb := sys.ram[cpu.programCounter], sys.ram[cpu.programCounter+1]
	insn := (uint16(msb) << 8) | uint16(lsb)

	opcode, err := matchOpcode(insn)
	if err != nil {
		rl.TraceLog(rl.LogWarning, "%s", fmt.Sprintf("%s (pc: %d)", err, cpu.programCounter))
		return
	}

	cpu.skip()
	opcode.interpret(insn, cpu, sys)
}

func (cpu *cpu) waitForPress(sys *System) {
	cpu.waitFor = waitForKeyPress
	cpu.lastKeys = sys.keypad
}

func (cpu *cpu) handleWaitState(sys *System) bool {
	switch sys.cpu.waitFor {
	case waitForNothing:
		return false

	case waitForKeyPress:
		pressedKeys := sys.keypad.KeysNotIn(cpu.lastKeys)

		if pressedKeys != 0 {
			cpu.lastKeys = pressedKeys
			cpu.waitFor = waitForKeyRelease
		}

		return true

	case waitForKeyRelease:
		releasedKey, ok := cpu.lastKeys.KeysNotIn(sys.keypad).FirstKey()

		if ok {
			cpu.variables[0] = byte(releasedKey)
			cpu.lastKeys = 0
			cpu.waitFor = waitForNothing
			return false
		}

		return true

	default:
		panic(fmt.Sprintf("unhandled wait state: %d", cpu.waitFor))
	}
}

func (cpu *cpu) vblank() {
	if cpu.delayTimer != 0 {
		cpu.delayTimer--
	}

	if cpu.soundTimer != 0 {
		cpu.soundTimer--
	}
}

func (cpu *cpu) skip() {
	cpu.programCounter += 2
}

func (cpu *cpu) setCarryFlag() {
	cpu.variables[0xF] = 1
}

func (cpu *cpu) clearCarryFlag() {
	cpu.variables[0xF] = 0
}
