#ifndef CPxActor_H
#define CPxActor_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxActor
	{
		void* obj;
	};

	CPxAPI void CPxActor_release(CSTRUCT CPxActor);

#ifdef __cplusplus
}
#endif

#endif // !CPxActor_H
