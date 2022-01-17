#ifndef CPxVec3_H
#define CPxVec3_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxVec3 {
		float x, y, z;
	};

	CPxAPI CSTRUCT CPxVec3 NewCPxVec3(float x, float y, float z);

#ifdef __cplusplus
}
#endif

#endif // !define CPxVec3_H