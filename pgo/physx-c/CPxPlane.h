#ifndef CPxPlane_H
#define CPxPlane_H

#ifdef __cplusplus
extern "C" {
#endif

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
