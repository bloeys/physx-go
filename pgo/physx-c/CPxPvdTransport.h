#ifndef CPxPvdTransport_H
#define CPxPvdTransport_H

#ifdef __cplusplus
extern "C" {
#endif
	struct CPxPvdTransport
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxPvdTransport CPxDefaultPvdSocketTransportCreate(const char* address, int port, int timeoutMillis);
#ifdef __cplusplus
}
#endif

#endif