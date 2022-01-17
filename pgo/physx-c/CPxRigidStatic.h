#ifndef CPxRigidStatic_H
#define CPxRigidStatic_H

#include "CPxPhysics.h"
#include "CPxPlane.h"
#include "CPxMaterial.h"
#include "CPxActor.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxRigidStatic
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxRigidStatic* CPxCreatePlane(CSTRUCT CPxPhysics* sdk, CSTRUCT CPxPlane* plane, CSTRUCT CPxMaterial* material);
	CPxAPI CSTRUCT CPxActor* CPxRigidStatic_toCPxActor(CSTRUCT CPxRigidStatic*);

#ifdef __cplusplus
}
#endif

#endif // !CPxRigidStatic_H
