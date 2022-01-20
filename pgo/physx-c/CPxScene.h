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
	};

	CPxAPI CSTRUCT CPxPvdSceneClient* CPxScene_getScenePvdClient(CSTRUCT CPxScene*);
	CPxAPI void CPxScene_addActor(CSTRUCT CPxScene*, CSTRUCT CPxActor actor);
	CPxAPI void CPxScene_simulate(CSTRUCT CPxScene*, CPxReal elapsedTime);
	CPxAPI bool CPxScene_fetchResults(CSTRUCT CPxScene*, bool block, CPxU32* errorState);

	CPxAPI void CPxScene_release(CSTRUCT CPxScene*);

#ifdef __cplusplus
}
#endif


#endif // !CPxScene_H
