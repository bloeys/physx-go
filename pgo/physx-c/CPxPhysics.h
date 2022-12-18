#ifndef CPxPhysics_H
#define CPxPhysics_H

#include "CPxPvd.h"
#include "CPxFoundation.h"
#include "CPxScene.h"
#include "CPxSceneDesc.h"
#include "CPxMaterial.h"
#include "CPxTolerancesScale.h"
#include "CPxRigidStatic.h"
#include "CPxRigidDynamic.h"
#include "CPxTransform.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxPhysics {
		void* obj;
	};

	CPxAPI CSTRUCT CPxPhysics CPxCreatePhysics(CSTRUCT CPxFoundation cfoundation, CSTRUCT CPxTolerancesScale cscale, bool trackOutstandingAllocations, CSTRUCT CPxPvd* cpvd);
	CPxAPI CSTRUCT CPxScene CPxPhysics_createScene(CSTRUCT CPxPhysics, CSTRUCT CPxSceneDesc);
	CPxAPI CSTRUCT CPxMaterial CPxPhysics_createMaterial(CSTRUCT CPxPhysics, CPxReal staticFriction, CPxReal dynamicFriction, CPxReal restitution);
	CPxAPI CSTRUCT CPxRigidDynamic CPxPhysics_createRigidDynamic(CSTRUCT CPxPhysics cp, CSTRUCT CPxTransform* ctr);
	CPxAPI CSTRUCT CPxRigidStatic CPxPhysics_createRigidStatic(CSTRUCT CPxPhysics cp, CSTRUCT CPxTransform* ctr);

	CPxAPI void CPxPhysics_release(CSTRUCT CPxPhysics);

#ifdef __cplusplus
}
#endif

#endif