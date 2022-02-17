#ifndef CPxPlane_H
#define CPxPlane_H

#ifdef __cplusplus
extern "C" {
#endif

	//NOTE: Maybe convert this into a value type like CPxSphereGeometry?
	struct CPxPlane
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxPlane* NewCPxPlane(float nx, float ny, float nz, float distance);
	CPxAPI void CPxPlane_release(CSTRUCT CPxPlane*);

#ifdef __cplusplus
}
#endif

#endif // !CPxPlane_H
