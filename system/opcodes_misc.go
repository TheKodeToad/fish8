package system

import (
	"math/rand"
)

func setVarToRandom(insn uint16, cpu *cpu, sys *System) {
	vp := &cpu.variables[(insn&0x0F00)>>8]
	mask := byte(insn & 0x00FF)

	*vp = byte(rand.Intn(256)) & mask
}

func splitVarDigits(insn uint16, cpu *cpu, sys *System) {
	v := cpu.variables[(insn&0x0F00)>>8]

	sys.ram[cpu.index] = v / 100
	sys.ram[cpu.index+1] = (v % 100) / 10
	sys.ram[cpu.index+2] = v % 10
}

func copyFromVarsUntil(insn uint16, cpu *cpu, sys *System) {
	vars := cpu.variables[:((insn&0x0F00)>>8)+1]

	copy(sys.ram[cpu.index:], vars)
}

func copyToVarsUntil(insn uint16, cpu *cpu, sys *System) {
	region := sys.ram[cpu.index : cpu.index+((insn&0x0F00)>>8)+1]

	copy(cpu.variables[:], region)
}
