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

	typedef void (*CPxonContactCallback)(void* pairHeader);

	/// <summary>
	/// Creates a SceneDesc with a custom filterShader that uses word0/word1 as groups shapes can belong to, word2 as a mask on word0, and word3 as a mask on word1.
	/// </summary>
	/// <param name="CPxTolerancesScale"></param>
	/// <returns></returns>
	CPxAPI CSTRUCT CPxSceneDesc NewCPxSceneDesc(CSTRUCT CPxTolerancesScale);
	CPxAPI void CPxSceneDesc_set_gravity(CSTRUCT CPxSceneDesc*, CSTRUCT CPxVec3);
	CPxAPI void CPxSceneDesc_set_cpuDispatcher(CSTRUCT CPxSceneDesc*, CSTRUCT CPxCpuDispatcher*);


	//CPxSceneDesc_set_onContactCallback sets the contact callback handler of the given scene descriptor.
	//The callback is sent an object of type 'CPxContactPairHeader*'. This object is only valid for the duration of the callback handler.
	//Therefore, the callback handler MUST copy data it wishes to keep for longer than the lifetime of the callback handler, as the memory it was handed might be reused/freed.
	//
	//NOTE: This function assumes you are using the default physx-c callback handler. Do NOT use this function if you set 'sceneDesc->simulationEventCallback' with your own custom implementation.
	CPxAPI void CPxSceneDesc_set_onContactCallback(CSTRUCT CPxSceneDesc*, CPxonContactCallback cb);
	CPxAPI void FreeCPxSceneDesc(CSTRUCT CPxSceneDesc*);

#ifdef __cplusplus
}
#endif

#endif