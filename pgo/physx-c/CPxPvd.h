#ifndef __CPxPvd_H__
#define __CPxPvd_H__

#include "CPxFoundation.h"

#ifdef __cplusplus
extern "C" {
#endif
	struct CPxPvd
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxPvd* CPxCreatePvd(CSTRUCT CPxFoundation*);
	CPxAPI void CPxPvd_release(CSTRUCT CPxPvd* cpp);
#ifdef __cplusplus
}
#endif

#endif