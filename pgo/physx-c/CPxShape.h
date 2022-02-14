#ifndef CPxShape_H
#define CPxShape_H

#include "CPxTransform.h"
#include "CPxFilterData.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxShape
	{
		void* obj;
	};

	CPxAPI void CPxShape_setLocalPose(CSTRUCT CPxShape* cs, CSTRUCT CPxTransform* tr);
	CPxAPI CSTRUCT CPxTransform CPxShape_getLocalPose(CSTRUCT CPxShape* cs);
	CPxAPI CSTRUCT CPxFilterData CPxShape_getSimulationFilterData(CSTRUCT CPxShape* cs);
	CPxAPI void CPxShape_setSimulationFilterData(CSTRUCT CPxShape* cs, CSTRUCT CPxFilterData* cfd);

#ifdef __cplusplus
}
#endif

#endif // !CPxShape_H
