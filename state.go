package fidgetlang

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewState() *State {
	return &State{
		registers:         [256]int32{},
		position:          0,
		instruction:       0,
		instructionMemory: []instruction{},
		buffReader:        bufio.NewReader(os.Stdin),
	}
}

type State struct {
	registers         [256]int32
	position          byte
	instruction       uint32
	instructionMemory []instruction
	buffReader        *bufio.Reader
	DebugOut          bool
	SuperDebug        bool
}

func (s *State) CompileDebug(program string) {
	s.instructionMemory = compileDebug(program)
}

func (s *State) Compile(program string) {
	s.instructionMemory = compileProgram(program)
}

func (s *State) GetDebugProgram() string {
	return makeDebug(s.instructionMemory)
}

func (s *State) bug(text string, args ...interface{}) {
	if !s.SuperDebug {
		return
	}
	var sArgs = make([]string, len(args))
	for k := range args {
		sArgs[k] = fmt.Sprint(args[k])
	}
	println(strings.Join(append([]string{text}, sArgs...), " "))
}

// Step returns false if the end of memory is reached
func (s *State) Step() bool {
	if s.overEnd() {
		return false
	}
	s.bug("X at I", s.instruction)
	switch s.curInstr() {
	case iIncrPtr:
		s.bug("I of P")
		s.position++
	case iDecrPtr:
		s.bug("D of P")
		s.position--
	case iIncrVal:
		s.bug("I of V")
		s.setData(s.curData() + 1)
	case iDecrVal:
		s.bug("D of V")
		s.setData(s.curData() - 1)
	case iPutVal:
		s.bug("P of V")
		if s.DebugOut {
			fmt.Printf("0x%04X\n", s.curData())
		} else {
			fmt.Printf("%s", string(rune(s.curData())))
		}
	case iGetVal:
		s.bug("R of V")
		rrune, _, err := s.buffReader.ReadRune()
		if err == io.EOF {
			return false
		}
		if err != nil {
			panic(err)
		}
		s.setData(rrune)
	case iJumpBehind:
		if s.curData() != 0 {
			for s.curInstr() != iJumpAhead {
				s.instruction--
			}
			s.bug("JB to I", s.instruction)
			return true
		}
	case iJumpAhead:
		if s.atEnd() {
			s.bug("JA abort at", s.instruction)
			return false
		}
		if s.curData() == 0 {
			for s.curInstr() != iJumpBehind {
				s.instruction++
			}
			s.bug("JA to I", s.instruction)
			return true
		}
	}
	if s.atEnd() {
		return false
	}
	s.instruction++
	return true
}

func (s *State) atEnd() bool {
	return int(s.instruction) == len(s.instructionMemory)-1
}

func (s *State) overEnd() bool {
	return int(s.instruction) >= len(s.instructionMemory)
}

func (s *State) curInstr() instruction {
	return instruction(s.instructionMemory[s.instruction])
}

func (s *State) curData() int32 {
	return s.registers[s.position]
}

func (s *State) setData(d int32) {
	s.registers[s.position] = d
}
