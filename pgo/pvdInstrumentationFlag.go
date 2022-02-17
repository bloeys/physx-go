package pgo

type PvdInstrumentationFlag uint32

const (
	PvdInstrumentationFlag_eDEBUG   PvdInstrumentationFlag = 1 << 0
	PvdInstrumentationFlag_ePROFILE PvdInstrumentationFlag = 1 << 1
	PvdInstrumentationFlag_eMEMORY  PvdInstrumentationFlag = 1 << 2
	PvdInstrumentationFlag_eALL     PvdInstrumentationFlag = (PvdInstrumentationFlag_eDEBUG | PvdInstrumentationFlag_ePROFILE | PvdInstrumentationFlag_eMEMORY)
)
