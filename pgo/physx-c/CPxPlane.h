#ifndef CPxPlane_H
#define CPxPlane_H

#include "CPxVec3.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxPlane
	{
		CSTRUCT CPxVec3 n; //!< The normal to the plane
		float d;  //!< The distance from the origin
	};

	CPxAPI CSTRUCT CPxPlane NewCPxPlane(float nx, float ny, float nz, float distance);

#ifdef __cplusplus
}
#endif

#endif // !CPxPlane_H
