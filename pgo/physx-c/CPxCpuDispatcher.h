#ifndef CPxCpuDispatcher_H
#define CPxCpuDispatcher_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxCpuDispatcher
	{
		void* obj;
	};

	/// <summary>
	/// This only frees C representation of the base class (the CPxCpuDispatcher struct). obj is NOT freed.
	/// To release the PhysX resources release must be called on the actual C implementation (e.g. CPxDefaultCpuDispatcher_release)
	/// </summary>
	/// <param name="CPxCpuDispatcher"></param>
	/// <returns></returns>
	CPxAPI void CPxCpuDispatcher_release(CSTRUCT CPxCpuDispatcher*);

#ifdef __cplusplus
}
#endif

#endif // !PxCpuDispatcher_H