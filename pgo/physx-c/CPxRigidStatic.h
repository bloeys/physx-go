#ifndef CPxRigidStatic_H
#define CPxRigidStatic_H

#include "CPxActor.h"
#include "CPxRigidActor.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxRigidStatic
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxActor CPxRigidStatic_toCPxActor(CSTRUCT CPxRigidStatic*);
	CPxAPI CSTRUCT CPxRigidActor CPxRigidStatic_toCPxRigidActor(CSTRUCT CPxRigidStatic*);

#ifdef __cplusplus
}
#endif

#endif // !CPxRigidStatic_H
