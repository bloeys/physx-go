#ifndef CPxShape_H
#define CPxShape_H

#include "CPxTransform.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxShape
	{
		void* obj;
	};

	CPxAPI void CPxShape_setLocalPose(CSTRUCT CPxShape* cs, CSTRUCT CPxTransform* tr);
	CPxAPI CSTRUCT CPxTransform CPxShape_getLocalPose(CSTRUCT CPxShape* cs);

#ifdef __cplusplus
}
#endif

#endif // !CPxShape_H
