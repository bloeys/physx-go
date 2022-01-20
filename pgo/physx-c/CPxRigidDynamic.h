#ifndef CPxRigidDynamic_H
#define CPxRigidDynamic_H

#include "CPxActor.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxRigidDynamic
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxActor CPxRigidDynamic_toCPxActor(CSTRUCT CPxRigidDynamic*);

#ifdef __cplusplus
}
#endif

#endif // !CPxRigidDynamic_H