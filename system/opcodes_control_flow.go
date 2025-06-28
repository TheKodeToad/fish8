package system

func returnFromCall(insn uint16, cpu *cpu, sys *System) {
	cpu.programCounter = sys.stack.pop()
}

func jump(insn uint16, cpu *cpu, sys *System) {
	cpu.programCounter = insn & 0x0FFF
	cpu.unskip()
}

func call(insn uint16, cpu *cpu, sys *System) {
	sys.stack.push(cpu.programCounter)
	cpu.programCounter = insn & 0x0FFF
	cpu.unskip()
}

func testVarEqualsByte(cpu *cpu, insn uint16) bool {
	v := cpu.variables[(insn&0x0F00)>>8]
	b := byte(insn & 0x00FF)

	return v == b
}

func skipIfVarEqualsByte(insn uint16, cpu *cpu, sys *System) {
	if testVarEqualsByte(cpu, insn) {
		cpu.skip()
	}
}

func skipIfVarNotEqualsByte(insn uint16, cpu *cpu, sys *System) {
	if !testVarEqualsByte(cpu, insn) {
		cpu.skip()
	}
}

func testVar1EqualsVar2(insn uint16, cpu *cpu, sys *System) bool {
	v1 := cpu.variables[(insn&0x0F00)>>8]
	v2 := cpu.variables[(insn&0x00F0)>>4]

	return v1 == v2
}

func skipIfVar1EqualsVar2(insn uint16, cpu *cpu, sys *System) {
	if testVar1EqualsVar2(insn, cpu, sys) {
		cpu.skip()
	}
}

func skipIfVar1NotEqualsVar2(insn uint16, cpu *cpu, sys *System) {
	if !testVar1EqualsVar2(insn, cpu, sys) {
		cpu.skip()
	}
}

func jumpWithOffset(insn uint16, cpu *cpu, sys *System) {
	cpu.programCounter = insn & (0x0FFF)
	cpu.programCounter += uint16(cpu.variables[0])
	cpu.unskip()
}
