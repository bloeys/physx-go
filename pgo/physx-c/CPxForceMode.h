#ifndef CPxForceMode_H
#define CPxForceMode_H

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxForceMode
	{
		eFORCE,				//!< parameter has unit of mass * distance/ time^2, i.e. a force
		eIMPULSE,			//!< parameter has unit of mass * distance /time
		eVELOCITY_CHANGE,	//!< parameter has unit of distance / time, i.e. the effect is mass independent: a velocity change.
		eACCELERATION		//!< parameter has unit of distance/ time^2, i.e. an acceleration. It gets treated just like a force except the mass is not divided out before integration.
	};

#ifdef __cplusplus
}
#endif

#endif // !CPxForceMode_H
