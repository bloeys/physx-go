#ifndef CPxRigidActor_H
#define CPxRigidActor_H

#include "CPxFilterData.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxRigidActor
	{
		void* obj;
	};

	//Sets the CPxFilterData on all the shapes of the actor.
	CPxAPI void CPxRigidActor_setSimFilterData(CSTRUCT CPxRigidActor* cra, CSTRUCT CPxFilterData* cfd);

#ifdef __cplusplus
}
#endif

#endif // !CPxRigidActor_H
