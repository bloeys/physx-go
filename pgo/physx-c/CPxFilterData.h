#ifndef CPxFilterData_H
#define CPxFilterData_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxFilterData
	{
		unsigned int word0;
		unsigned int word1;
		unsigned int word2;
		unsigned int word3;

		//TODO: For some reason only this file breaks with CPxU32 (uint32_t). Why?
		/*CPxU32 word0;
		CPxU32 word1;
		CPxU32 word2;
		CPxU32 word3;*/
	};

#ifdef __cplusplus
}
#endif

#endif // !CPxFilterData_H
