package system

func setVarToByte(insn uint16, cpu *cpu, sys *System) {
	vp := &cpu.variables[(insn&0x0F00)>>8]
	b := byte(insn & 0x00FF)

	*vp = b
}

func incVarByByte(insn uint16, cpu *cpu, sys *System) {
	vp := &cpu.variables[(insn&0x0F00)>>8]
	b := byte(insn & 0x00FF)

	*vp += b
}

func setVar1ToVar2(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p = v2
}

func setVar1ToORVar2(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p |= v2
}

func setVar1ToANDVar2(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p &= v2
}

func setVar1ToXORVar2(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p ^= v2
}

func incVar1ByVar2(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v1 := *v1p
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p += v2

	if *v1p < v1 {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}
}

func decVar1ByVar2(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v1 := *v1p
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p -= v2

	if *v1p > v1 {
		cpu.clearCarryFlag()
	} else {
		cpu.setCarryFlag()
	}
}

func setVar1ToVar2MinusVar1(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v1 := *v1p
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p = v2 - v1

	if *v1p > v2 {
		cpu.clearCarryFlag()
	} else {
		cpu.setCarryFlag()
	}
}

func shiftVarRight(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p = v2 >> 1

	if v2&1 != 0 {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}
}

func shiftVarLeft(insn uint16, cpu *cpu, sys *System) {
	v1p := &cpu.variables[(insn&0x0F00)>>8]
	v2 := cpu.variables[(insn&0x00F0)>>4]

	*v1p = v2 << 1

	if v2&(1<<7) != 0 {
		cpu.setCarryFlag()
	} else {
		cpu.clearCarryFlag()
	}
}

func setIndexToByte(insn uint16, cpu *cpu, sys *System) {
	cpu.index = insn & 0x0FFF
}

func incIndexByVar(insn uint16, cpu *cpu, sys *System) {
	v := cpu.variables[(insn&0x0F00)>>8]
	cpu.index += uint16(v)
}
