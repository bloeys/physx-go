#ifndef CPxRaycastBuffer_H
#define CPxRaycastBuffer_H

#include "CPxVec3.h"
#include "CPxShape.h"
#include "CPxRigidActor.h"

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxHitFlag
	{
		CPxHitFlag_ePOSITION = (1 << 0),	//!< "position" member of #PxQueryHit is valid
		CPxHitFlag_eNORMAL = (1 << 1),	//!< "normal" member of #PxQueryHit is valid
		CPxHitFlag_eUV = (1 << 3),	//!< "u" and "v" barycentric coordinates of #PxQueryHit are valid. Not applicable to sweep queries.
		CPxHitFlag_eASSUME_NO_INITIAL_OVERLAP = (1 << 4),	//!< Performance hint flag for sweeps when it is known upfront there's no initial overlap.
		//!< NOTE: using this flag may cause undefined results if shapes are initially overlapping.
		CPxHitFlag_eMESH_MULTIPLE = (1 << 5),	//!< Report all hits for meshes rather than just the first. Not applicable to sweep queries.
		CPxHitFlag_eMESH_ANY = (1 << 6),	//!< Report any first hit for meshes. If neither eMESH_MULTIPLE nor eMESH_ANY is specified,
		//!< a single closest hit will be reported for meshes.
		CPxHitFlag_eMESH_BOTH_SIDES = (1 << 7),	//!< Report hits with back faces of mesh triangles. Also report hits for raycast
		//!< originating on mesh surface and facing away from the surface normal. Not applicable to sweep queries.
		//!< Please refer to the user guide for heightfield-specific differences.
		CPxHitFlag_ePRECISE_SWEEP = (1 << 8),	//!< Use more accurate but slower narrow phase sweep tests.
		//!< May provide better compatibility with PhysX 3.2 sweep behavior.
		CPxHitFlag_eMTD = (1 << 9),	//!< Report the minimum translation depth, normal and contact point.
		CPxHitFlag_eFACE_INDEX = (1 << 10),	//!< "face index" member of #PxQueryHit is valid
		CPxHitFlag_eDEFAULT = CPxHitFlag_ePOSITION | CPxHitFlag_eNORMAL | CPxHitFlag_eFACE_INDEX,
		/** \brief Only this subset of flags can be modified by pre-filter. Other modifications will be discarded. */
		CPxHitFlag_eMODIFIABLE_FLAGS = CPxHitFlag_eMESH_MULTIPLE | CPxHitFlag_eMESH_BOTH_SIDES | CPxHitFlag_eASSUME_NO_INITIAL_OVERLAP | CPxHitFlag_ePRECISE_SWEEP
	};

	struct CPxRaycastHit
	{
		CENUM CPxHitFlag flags;			//!< Hit flags specifying which members contain valid values.
		CSTRUCT CPxVec3	position;	//!< World-space hit position (flag: #PxHitFlag::ePOSITION)
		CSTRUCT CPxVec3	normal;		//!< World-space hit normal (flag: #PxHitFlag::eNORMAL)

		/**
		\brief	Distance to hit.
		\note	If the eMTD flag is used, distance will be a negative value if shapes are overlapping indicating the penetration depth.
		\note	Otherwise, this value will be >= 0 */
		CPxF32 distance;
		CPxReal u, v;
		CPxU32	faceIndex;
		CSTRUCT CPxShape shape;
		CSTRUCT CPxRigidActor actor;
	};

	struct CPxRaycastBuffer
	{
		CSTRUCT CPxRaycastHit block;
		CSTRUCT CPxRaycastHit* touches;
		CPxU32 nbTouches;
		bool hasBlock;
	};

	CPxAPI CSTRUCT CPxRaycastBuffer* NewCPxRaycastBufferWithAlloc(CPxU32 maxTouches);
	CPxAPI void CPxRaycastBuffer_release(CSTRUCT CPxRaycastBuffer* crb);

#ifdef __cplusplus
}
#endif


#endif // !CPxRaycastBuffer_H
