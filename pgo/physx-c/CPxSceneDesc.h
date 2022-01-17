#ifndef CPxSceneDesc_H
#define CPxSceneDesc_H

#include "CPxTolerancesScale.h"
#include "CPxVec3.h"
#include "CPxCpuDispatcher.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxSceneDesc {
		void* obj;
	};

	/// <summary>
	/// Creates a SceneDesc with filterShader=physx::PxDefaultSimulationFilterShader
	/// </summary>
	/// <param name="CPxTolerancesScale"></param>
	/// <returns></returns>
	CPxAPI CSTRUCT CPxSceneDesc* NewCPxSceneDesc(CSTRUCT CPxTolerancesScale);
	CPxAPI void CPxSceneDesc_set_gravity(CSTRUCT CPxSceneDesc*, CSTRUCT CPxVec3);
	CPxAPI void CPxSceneDesc_set_cpuDispatcher(CSTRUCT CPxSceneDesc*, CSTRUCT CPxCpuDispatcher*);
	CPxAPI void FreeCPxSceneDesc(CSTRUCT CPxSceneDesc*);

#ifdef __cplusplus
}
#endif

#endif