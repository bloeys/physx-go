#ifndef CPxScene_H
#define CPxScene_H

#include "CPxPvdSceneClient.h"
#include "CPxActor.h"
#include "CPxVec3.h"
#include "CPxShape.h"
#include "CPxRigidActor.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxRaycastHit
	{
		//CPxHitFlags			flags;		//!< Hit flags specifying which members contain valid values.
		struct CPxVec3				position;	//!< World-space hit position (flag: #PxHitFlag::ePOSITION)
		struct CPxVec3				normal;		//!< World-space hit normal (flag: #PxHitFlag::eNORMAL)

		/**
		\brief	Distance to hit.
		\note	If the eMTD flag is used, distance will be a negative value if shapes are overlapping indicating the penetration depth.
		\note	Otherwise, this value will be >= 0 */
		CPxF32 distance;
		CPxReal u, v;
		CPxU32	faceIndex;
		struct CPxShape shape;
		struct CPxRigidActor actor;
	};

	struct CPxRaycastBuffer
	{
		struct CPxRaycastHit block;
		struct CPxRaycastHit* touches;
		CPxU32 nbTouches;
		bool hasBlock;
	};

	struct CPxScene
	{
		void* obj;
		void* scratchBuffer;
		CPxU32 scratchBufferSize;
	};

	CPxAPI CSTRUCT CPxPvdSceneClient* CPxScene_getScenePvdClient(CSTRUCT CPxScene*);
	CPxAPI void CPxScene_addActor(CSTRUCT CPxScene*, CSTRUCT CPxActor actor);
	CPxAPI void CPxScene_simulate(CSTRUCT CPxScene*, CPxReal elapsedTime);
	CPxAPI void CPxScene_collide(CSTRUCT CPxScene*, CPxReal elapsedTime);
	CPxAPI bool CPxScene_fetchCollision(CSTRUCT CPxScene*, bool block);
	CPxAPI void CPxScene_advance(CSTRUCT CPxScene*);
	CPxAPI bool CPxScene_fetchResults(CSTRUCT CPxScene*, bool block, CPxU32* errorState);
	CPxAPI bool CPxScene_raycast(CSTRUCT CPxScene* cs, CSTRUCT CPxVec3* origin, CSTRUCT CPxVec3* unitDir, CPxReal distance, CSTRUCT CPxRaycastBuffer** hitRet);

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
