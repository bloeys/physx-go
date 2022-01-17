#ifndef CPxScene_H
#define CPxScene_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxScene
	{
		void* obj;
	};

	CPxAPI void CPxScene_release(CSTRUCT CPxScene*);

#ifdef __cplusplus
}
#endif


#endif // !CPxScene_H
