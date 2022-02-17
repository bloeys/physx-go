#ifndef CPxCapsuleGeometry_H
#define CPxCapsuleGeometry_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxCapsuleGeometry
	{
		CPxReal radius, halfHeight;
	};

	CPxAPI CPxInline CSTRUCT CPxCapsuleGeometry NewCPxCapsuleGeometry(CPxReal radius, CPxReal halfHeight)
	{
		CSTRUCT CPxCapsuleGeometry c;
		c.radius = radius;
		c.halfHeight = halfHeight;
		return c;
	}

#ifdef __cplusplus
}
#endif

#endif // !CPxCapsuleGeometry_H
