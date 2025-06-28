package system

func setVarToDelayTimer(insn uint16, cpu *cpu, sys *System) {
	vp := &cpu.variables[(insn&0x0F00)>>8]
	*vp = cpu.delayTimer
}

func setDelayTimerToVar(insn uint16, cpu *cpu, sys *System) {
	v := cpu.variables[(insn&0x0F00)>>8]
	cpu.delayTimer = v
}

func setSoundTimerToVar(insn uint16, cpu *cpu, sys *System) {
	v := cpu.variables[(insn&0x0F00)>>8]
	cpu.soundTimer = v
}
