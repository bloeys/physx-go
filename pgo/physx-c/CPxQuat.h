#ifndef CPxQuat_H
#define CPxQuat_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxQuat
	{
		float x, y, z, w;
	};

	CPxAPI CPxInline CSTRUCT CPxQuat NewCPxQuat(float angleRads, float x, float y, float z);

#ifdef __cplusplus
}
#endif

#endif // !CPxQuat_H
