#ifndef CPxPvdTransport_H
#define CPxPvdTransport_H

#ifdef __cplusplus
extern "C" {
#endif
	struct CPxPvdTransport
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxPvdTransport* CPxDefaultPvdSocketTransportCreate(const char* address, int port, int timeoutMillis);
	CPxAPI void CPxPvdTransport_release(CSTRUCT CPxPvdTransport* cppt);
#ifdef __cplusplus
}
#endif

#endif