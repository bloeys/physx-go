#ifndef CPxMaterial_H
#define CPxMaterial_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxMaterial
	{
		void* obj;
	};

	CPxAPI void CPxMaterial_release(CSTRUCT CPxMaterial*);

#ifdef __cplusplus
}
#endif

#endif // !CPxMaterial_H
