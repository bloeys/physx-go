#ifndef CPxDefaultCpuDispatcher_H
#define CPxDefaultCpuDispatcher_H

#include "CPxCpuDispatcher.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxDefaultCpuDispatcher {
		void* obj;
	};

	CPxAPI CSTRUCT CPxDefaultCpuDispatcher CPxDefaultCpuDispatcherCreate(CPxU32 numThreads, CPxU32* affinityMasks);
	CPxAPI CSTRUCT CPxCpuDispatcher CPxDefaultCpuDispatcher_toCPxCpuDispatcher(CSTRUCT CPxDefaultCpuDispatcher cdcd);
	CPxAPI void CPxDefaultCpuDispatcher_release(CSTRUCT CPxDefaultCpuDispatcher cdcd);

#ifdef __cplusplus
}
#endif

#endif // !CPxDefaultCpuDispatcher_H