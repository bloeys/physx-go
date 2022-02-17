package pgo

type HitFlag uint16

const (
	HitFlag_ePOSITION                  HitFlag = (1 << 0) //!< "position" member of #PxQueryHit is valid
	HitFlag_eNORMAL                    HitFlag = (1 << 1) //!< "normal" member of #PxQueryHit is valid
	HitFlag_eUV                        HitFlag = (1 << 3) //!< "u" and "v" barycentric coordinates of #PxQueryHit are valid. Not applicable to sweep queries.
	HitFlag_eASSUME_NO_INITIAL_OVERLAP HitFlag = (1 << 4) //!< Performance hint flag for sweeps when it is known upfront there's no initial overlap.
	//!< NOTE: using this flag may cause undefined results if shapes are initially overlapping.
	HitFlag_eMESH_MULTIPLE HitFlag = (1 << 5) //!< Report all hits for meshes rather than just the first. Not applicable to sweep queries.
	HitFlag_eMESH_ANY      HitFlag = (1 << 6) //!< Report any first hit for meshes. If neither eMESH_MULTIPLE nor eMESH_ANY is specified,
	//!< a single closest hit will be reported for meshes.
	HitFlag_eMESH_BOTH_SIDES HitFlag = (1 << 7) //!< Report hits with back faces of mesh triangles. Also report hits for raycast
	//!< originating on mesh surface and facing away from the surface normal. Not applicable to sweep queries.
	//!< Please refer to the user guide for heightfield-specific differences.
	HitFlag_ePRECISE_SWEEP HitFlag = (1 << 8) //!< Use more accurate but slower narrow phase sweep tests.
	//!< May provide better compatibility with PhysX 3.2 sweep behavior.
	HitFlag_eMTD        HitFlag = (1 << 9)  //!< Report the minimum translation depth, normal and contact point.
	HitFlag_eFACE_INDEX HitFlag = (1 << 10) //!< "face index" member of #PxQueryHit is valid
	HitFlag_eDEFAULT    HitFlag = HitFlag_ePOSITION | HitFlag_eNORMAL | HitFlag_eFACE_INDEX
	/** \brief Only this subset of flags can be modified by pre-filter. Other modifications will be discarded. */
	HitFlag_eMODIFIABLE_FLAGS HitFlag = HitFlag_eMESH_MULTIPLE | HitFlag_eMESH_BOTH_SIDES | HitFlag_eASSUME_NO_INITIAL_OVERLAP | HitFlag_ePRECISE_SWEEP
)
