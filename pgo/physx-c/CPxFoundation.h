#ifndef CPxFoundation_H
#define CPxFoundation_H

#ifdef __cplusplus
extern "C" {
#endif
	struct CPxFoundation
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxFoundation* CPxCreateFoundation();
	CPxAPI void CPxFoundation_release(CSTRUCT CPxFoundation* cpf);
#ifdef __cplusplus
}
#endif

#endif