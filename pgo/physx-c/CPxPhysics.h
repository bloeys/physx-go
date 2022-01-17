#ifndef CPxPhysics_H
#define CPxPhysics_H

#include "CPxPvd.h"
#include "CPxFoundation.h"
#include "CPxScene.h"
#include "CPxSceneDesc.h"
#include "CPxTolerancesScale.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxPhysics {
		void* obj;
	};

	CPxAPI CSTRUCT CPxPhysics* CPxCreatePhysics(CSTRUCT CPxFoundation* cfoundation, CSTRUCT CPxTolerancesScale cscale, bool trackOutstandingAllocations, CSTRUCT CPxPvd* cpvd);
	CPxAPI CSTRUCT CPxScene* CPxPhysics_createScene(CSTRUCT CPxPhysics*, CSTRUCT CPxSceneDesc*);
	CPxAPI void CPxPhysics_release(CSTRUCT CPxPhysics*);

#ifdef __cplusplus
}
#endif

#endif