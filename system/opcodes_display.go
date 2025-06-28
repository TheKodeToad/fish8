package system

func clearScreen(insn uint16, cpu *cpu, sys *System) {
	sys.display.clear()
}

func drawToScreen(insn uint16, cpu *cpu, sys *System) {
	x := cpu.variables[(insn&0x0F00)>>8]
	y := cpu.variables[(insn&0x00F0)>>4]

	height := insn & (0x000F)
	sprite := sys.ram[cpu.index : cpu.index+height]

	x = x % DisplayWidth

	cpu.clearCarryFlag()

	for _, row := range sprite {
		if int(y) >= len(sys.display) {
			return
		}

		pattern := (uint64(row) << 56) >> x

		sys.display[y] ^= pattern

		if (pattern & ^uint64(sys.display[y])) != 0 {
			cpu.setCarryFlag()
		}

		y++
	}
}

func setIndexToFontFromVar(insn uint16, cpu *cpu, sys *System) {
	v := cpu.variables[(insn&0x0F00)>>8]

	cpu.index = fontOffset + uint16(v&0xF)*5
}
