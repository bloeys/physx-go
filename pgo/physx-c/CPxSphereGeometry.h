#ifndef CPxSphereGeometry_H
#define CPxSphereGeometry_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxSphereGeometry
	{
		CPxReal radius;
	};

	CPxAPI CPxInline CSTRUCT CPxSphereGeometry NewCPxSphereGeometry(CPxReal radius);

#ifdef __cplusplus
}
#endif

#endif // !CPxSphereGeometry_H