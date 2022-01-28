package pgo

type ForceMode uint32

const (
	ForceMode_eFORCE           ForceMode = iota //!< parameter has unit of mass * distance/ time^2, i.e. a force
	ForceMode_eIMPULSE                          //!< parameter has unit of mass * distance /time
	ForceMode_eVELOCITY_CHANGE                  //!< parameter has unit of distance / time, i.e. the effect is mass independent: a velocity change.
	ForceMode_eACCELERATION                     //!< parameter has unit of distance/ time^2, i.e. an acceleration. It gets treated just like a force except the mass is not divided out before integration.
)
