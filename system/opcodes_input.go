package system

func testKeyHeld(insn uint16, cpu *cpu, sys *System) bool {
	v := cpu.variables[(insn&0x0F00)>>8]
	key := Key(v).ToKeySet()

	return sys.keypad&key != 0
}

func skipIfKeyHeld(insn uint16, cpu *cpu, sys *System) {
	if testKeyHeld(insn, cpu, sys) {
		cpu.skip()
	}
}

func skipIfKeyNotHeld(insn uint16, cpu *cpu, sys *System) {
	if !testKeyHeld(insn, cpu, sys) {
		cpu.skip()
	}
}

func awaitKey(insn uint16, cpu *cpu, sys *System) {
	cpu.waitForPress(sys)
}
