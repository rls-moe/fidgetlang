package fidgetlang

type instruction byte

const (
	iIncrPtr instruction = 1 << iota
	iDecrPtr
	iIncrVal
	iDecrVal
	iPutVal
	iGetVal
	iJumpAhead
	iJumpBehind
	iNop instruction = 0

	sIncrPtr    = "߷"
	sDecrPtr    = "߷߷"
	sIncrVal    = "߷߷߷"
	sDecrVal    = "߷߷߷߷"
	sPutVal     = "߷߷߷߷߷"
	sGetVal     = "߷߷߷߷߷߷"
	sJumpAhead  = "߷߷߷߷߷߷߷"
	sJumpBehind = "߷߷߷߷߷߷߷߷"
	sNop        = " "

	dIncrPtr    = "IncrPtr"
	dDecrPtr    = "DecrPtr"
	dIncrVal    = "IncrVal"
	dDecrVal    = "DecrVal"
	dPutVal     = "PutVal"
	dGetVal     = "GetVal"
	dJumpAhead  = "JumpAhead"
	dJumpBehind = "JumpBehind"
	dNop        = "Nop"
)

func stringToInstruction(d string) instruction {
	switch d {
	case sIncrPtr:
		return iIncrPtr
	case sDecrPtr:
		return iDecrPtr
	case sIncrVal:
		return iIncrVal
	case sDecrVal:
		return iDecrVal
	case sPutVal:
		return iPutVal
	case sGetVal:
		return iGetVal
	case sJumpAhead:
		return iJumpAhead
	case sJumpBehind:
		return iJumpBehind
	default:
		return iNop
	}
}

func instructionToString(i instruction) string {
	switch i {
	case iIncrPtr:
		return sIncrPtr
	case iDecrPtr:
		return sDecrPtr
	case iIncrVal:
		return sIncrVal
	case iDecrVal:
		return sDecrVal
	case iPutVal:
		return sPutVal
	case iGetVal:
		return sGetVal
	case iJumpAhead:
		return sJumpAhead
	case iJumpBehind:
		return sJumpBehind
	default:
		return sNop
	}
}

func debugToInstruction(d string) instruction {
	switch d {
	case dIncrPtr:
		return iIncrPtr
	case dDecrPtr:
		return iDecrPtr
	case dIncrVal:
		return iIncrVal
	case dDecrVal:
		return iDecrVal
	case dPutVal:
		return iPutVal
	case dGetVal:
		return iGetVal
	case dJumpAhead:
		return iJumpAhead
	case dJumpBehind:
		return iJumpBehind
	default:
		return iNop
	}
}

func instructionToDebug(i instruction) string {
	switch i {
	case iIncrPtr:
		return dIncrPtr
	case iDecrPtr:
		return dDecrPtr
	case iIncrVal:
		return dIncrVal
	case iDecrVal:
		return dDecrVal
	case iPutVal:
		return dPutVal
	case iGetVal:
		return dGetVal
	case iJumpAhead:
		return dJumpAhead
	case iJumpBehind:
		return dJumpBehind
	default:
		return dNop
	}
}
