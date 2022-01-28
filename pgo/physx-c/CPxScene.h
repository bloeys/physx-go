#ifndef CPxScene_H
#define CPxScene_H

#include "CPxPvdSceneClient.h"
#include "CPxActor.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxScene
	{
		void* obj;
		void* scratchBuffer;
		CPxU32 scratchBufferSize;
	};

	CPxAPI CSTRUCT CPxPvdSceneClient* CPxScene_getScenePvdClient(CSTRUCT CPxScene*);
	CPxAPI void CPxScene_addActor(CSTRUCT CPxScene*, CSTRUCT CPxActor actor);
	CPxAPI void CPxScene_simulate(CSTRUCT CPxScene*, CPxReal elapsedTime);
	CPxAPI bool CPxScene_fetchResults(CSTRUCT CPxScene*, bool block, CPxU32* errorState);

	/// <summary>
	/// Creates a scratch buffer thats a multiple of 16K to be used by the scene when running CPxScene_simulate.
	/// The buffer MUST be 16-byte aligned. If a buffer already exists then it is freed and a new one is allocated.
	/// If multiples passed are zero then any existing buffers are cleared
	/// </summary>
	/// <returns></returns>
	CPxAPI void CPxScene_setScratchBuffer(CSTRUCT CPxScene*, uint32_t multiplesOf16k);

	CPxAPI void CPxScene_release(CSTRUCT CPxScene*);

#ifdef __cplusplus
}
#endif


#endif // !CPxScene_H
