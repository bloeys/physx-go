#ifndef CPxPvdSceneClient_H
#define CPxPvdSceneClient_H

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxPvdSceneFlag
	{
		CPxPvdSceneFlag_eTRANSMIT_CONTACTS = (1 << 0), //Transmits contact stream to PVD.
		CPxPvdSceneFlag_eTRANSMIT_SCENEQUERIES = (1 << 1), //Transmits scene query stream to PVD.
		CPxPvdSceneFlag_eTRANSMIT_CONSTRAINTS = (1 << 2)  //Transmits constraints visualize stream to PVD.
	};

	struct CPxPvdSceneClient
	{
		void* obj;
	};

	CPxAPI void CPxPvdSceneClient_setScenePvdFlag(CSTRUCT CPxPvdSceneClient* c, CENUM CPxPvdSceneFlag flag, bool value);

	/// <summary>
	/// This only releases the C struct
	/// </summary>
	/// <param name=""></param>
	/// <returns></returns>
	CPxAPI void CPxPvdSceneClient_release(CSTRUCT CPxPvdSceneClient*);

#ifdef __cplusplus
}
#endif

#endif // !CPxPvdSceneClient_H
