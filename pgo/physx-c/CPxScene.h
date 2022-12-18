#ifndef CPxScene_H
#define CPxScene_H

#include "CPxPvdSceneClient.h"
#include "CPxActor.h"
#include "CPxVec3.h"
#include "CPxShape.h"
#include "CPxRigidActor.h"
#include "CPxRaycastBuffer.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxScene
	{
		void* obj;
		void* scratchBuffer;
		CPxU32 scratchBufferSize;
	};

	CPxAPI CSTRUCT CPxPvdSceneClient CPxScene_getScenePvdClient(CSTRUCT CPxScene);
	CPxAPI void CPxScene_addActor(CSTRUCT CPxScene, CSTRUCT CPxActor actor);
	CPxAPI void CPxScene_simulate(CSTRUCT CPxScene, CPxReal elapsedTime);
	CPxAPI void CPxScene_collide(CSTRUCT CPxScene, CPxReal elapsedTime);
	CPxAPI bool CPxScene_fetchCollision(CSTRUCT CPxScene, bool block);
	CPxAPI void CPxScene_advance(CSTRUCT CPxScene);
	CPxAPI bool CPxScene_fetchResults(CSTRUCT CPxScene, bool block, CPxU32* errorState);

	//Does a scene raycast. Allocates memory for hitRet and then reads data into it. It is the callers responsibility to free.
	CPxAPI bool CPxScene_raycast(CSTRUCT CPxScene cs, CSTRUCT CPxVec3* origin, CSTRUCT CPxVec3* unitDir, CPxReal distance, CSTRUCT CPxRaycastBuffer** hitRet);

	//Does a scene raycast. 'hit' must be pre-allocated as NO new allocation will happen in the function.
	//hit->touches will be filled up to 'touchesToRead' and must also be pre-allocated. If the hit produces more touches than 'touchesToRead' then the additional touches will be ignored.
	CPxAPI bool CPxScene_raycastWithHitBuffer(CSTRUCT CPxScene cs, CSTRUCT CPxVec3* origin, CSTRUCT CPxVec3* unitDir, CPxReal distance, CSTRUCT CPxRaycastBuffer* hit, CPxU32 touchesToRead);

	/// <summary>
	/// Creates a scratch buffer thats a multiple of 16K to be used by the scene when running CPxScene_simulate.
	/// The buffer MUST be 16-byte aligned. If a buffer already exists then it is freed and a new one is allocated.
	/// If multiples passed are zero then any existing buffers are cleared
	/// </summary>
	/// <returns></returns>
	CPxAPI void CPxScene_setScratchBuffer(CSTRUCT CPxScene, CPxU32 multiplesOf16k);

	CPxAPI void CPxScene_release(CSTRUCT CPxScene);

#ifdef __cplusplus
}
#endif


#endif // !CPxScene_H
