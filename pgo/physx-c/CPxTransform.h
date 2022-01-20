#ifndef CPxTransform_H
#define CPxTransform_H

#include "CPxVec3.h"
#include "CPxQuat.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxTransform
	{
		CSTRUCT CPxVec3 p;
		CSTRUCT CPxQuat q;
	};

	CPxAPI CPxInline CSTRUCT CPxTransform NewCPxTransform(CSTRUCT CPxVec3*, CSTRUCT CPxQuat*);

#ifdef __cplusplus
}
#endif

#endif // !CPxTransform_H