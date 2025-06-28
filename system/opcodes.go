package system

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type opcode struct {
	pattern   string
	interpret func(insn uint16, cpu *cpu, sys *System)
}

var opcodes = [...]opcode{
	{"00E0", clearScreen},
	{"00EE", returnFromCall},
	{"1---", jump},
	{"2---", call},
	{"3---", skipIfVarEqualsByte},
	{"4---", skipIfVarNotEqualsByte},
	{"5--0", skipIfVar1EqualsVar2},
	{"6---", setVarToByte},
	{"7---", incVarByByte},
	{"8--0", setVar1ToVar2},
	{"8--1", setVar1ToORVar2},
	{"8--2", setVar1ToANDVar2},
	{"8--3", setVar1ToXORVar2},
	{"8--4", incVar1ByVar2},
	{"8--5", decVar1ByVar2},
	{"8--6", shiftVarRight},
	{"8--7", setVar1ToVar2MinusVar1},
	{"8--E", shiftVarLeft},
	{"9--0", skipIfVar1NotEqualsVar2},
	{"A---", setIndexToByte},
	{"B---", jumpWithOffset},
	{"C---", setVarToRandom},
	{"D---", drawToScreen},
	{"E-9E", skipIfKeyHeld},
	{"E-A1", skipIfKeyNotHeld},
	{"F-0A", awaitKey},
	{"F-07", setVarToDelayTimer},
	{"F-15", setDelayTimerToVar},
	{"F-18", setSoundTimerToVar},
	{"F-1E", incIndexByVar},
	{"F-29", setIndexToFontFromVar},
	{"F-33", splitVarDigits},
	{"F-55", copyFromVarsUntil},
	{"F-65", copyToVarsUntil},
}

func matchOpcode(insn uint16) (opcode, error) {
	insnHex := fmt.Sprintf("%04X", insn)

matching:
	for _, opcode := range opcodes {
		if len(opcode.pattern) != len(insnHex) {
			panic(fmt.Sprintf("pattern '%s' is the wrong length - it should be %d chars", opcode.pattern, len(insnHex)))
		}

		for i, char := range []byte(opcode.pattern) {
			if char != '-' && insnHex[i] != char {
				continue matching
			}
		}

		name := runtime.FuncForPC(reflect.ValueOf(opcode.interpret).Pointer()).Name()
		dot := strings.LastIndex(name, ".")

		if dot != -1 {
			name = name[dot+1:]
		}

		rl.TraceLog(rl.LogInfo, "$%s -> %s ($%s)", insnHex, name, opcode.pattern)

		return opcode, nil
	}

	return opcode{}, fmt.Errorf("$%s matched no opcodes", insnHex)
}
