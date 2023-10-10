package pgo

/*
#include <stdint.h> // Needed for uint32_t

void goOnContactCallback_cgo(void* pairHeader)
{
	void goOnContactCallback(void*);
	goOnContactCallback(pairHeader);
}

void goOnTriggerCallback_cgo(void* triggerPairs, uint32_t count)
{
	void goOnTriggerCallback(void*, uint32_t);
	goOnTriggerCallback(triggerPairs, count);
}
*/
import "C"
