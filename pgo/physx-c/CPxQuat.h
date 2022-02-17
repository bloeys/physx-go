#ifndef CPxQuat_H
#define CPxQuat_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxQuat
	{
		float x, y, z, w;
	};

	CPxAPI CSTRUCT CPxQuat NewCPxQuat(float angleRads, float x, float y, float z);
	CPxAPI CPxInline CSTRUCT CPxQuat NewCPxQuatXYZW(float x, float y, float z, float w)
	{
		CSTRUCT CPxQuat q;
		q.x = x;
		q.y = y;
		q.z = z;
		q.w = w;
		return q;
	}


#ifdef __cplusplus
}
#endif

#endif // !CPxQuat_H
