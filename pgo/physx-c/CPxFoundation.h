#ifndef __CPxFoundation_H__
#define __CPxFoundation_H__

#ifdef __cplusplus
extern "C" {
#endif
	struct CPxFoundation
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxFoundation* NewCPxFoundation();
	CPxAPI void FreeCPxFoundation(CSTRUCT CPxFoundation* cpf);
#ifdef __cplusplus
}
#endif

#endif