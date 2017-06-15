package fidgetlang

import "strings"

func compileDebug(debug string) []instruction {
	var program = []instruction{}

	instr := strings.Split(debug, " ")

	for k := range instr {
		instr[k] = strings.TrimSpace(instr[k])
		program = append(program, debugToInstruction(instr[k]))
	}

	return program
}

func compileProgram(prg string) []instruction {
	var program = []instruction{}

	instr := strings.Split(prg, " ")

	for k := range instr {
		instr[k] = strings.TrimSpace(instr[k])
		program = append(program, stringToInstruction(instr[k]))
	}

	return program
}

func makeDebug(k []instruction) string {
	if len(k) == 0 {
		return ""
	}
	var out = instructionToDebug(k[0])

	if len(k) > 1 {
		for i := range k[1:] {
			out += " " + instructionToDebug(k[i])
		}
	}

	return out
}