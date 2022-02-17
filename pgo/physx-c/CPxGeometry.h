#ifndef CPxGeometry_H
#define CPxGeometry_H

#include "CPxBoxGeometry.h"
#include "CPxSphereGeometry.h"
#include "CPxCapsuleGeometry.h"

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxGeometryType
	{
		CPxGeometryType_eSPHERE,
		CPxGeometryType_ePLANE,
		CPxGeometryType_eCAPSULE,
		CPxGeometryType_eBOX,
		CPxGeometryType_eCONVEXMESH,
		CPxGeometryType_eTRIANGLEMESH,
		CPxGeometryType_eHEIGHTFIELD,
		CPxGeometryType_eGEOMETRY_COUNT,	//!< internal use only!
		CPxGeometryType_eINVALID = -1		//!< internal use only!
	};

	struct CPxGeometry
	{
		void* obj;
		CENUM CPxGeometryType type;
	};

	//
	// CPxSphereGeometry
	//
	CPxAPI CPxInline CSTRUCT CPxSphereGeometry CPxGeometry_toCPxSphereGeometry(CSTRUCT CPxGeometry cg)
	{
		return *(CSTRUCT CPxSphereGeometry*)(cg.obj);
		//return *static_cast<CSTRUCT CPxSphereGeometry*>(cg.obj);
	}

	CPxAPI CPxInline CSTRUCT CPxGeometry CPxSphereGeometry_toCPxGeometry(CSTRUCT CPxSphereGeometry* csg)
	{
		CSTRUCT CPxGeometry g;
		g.obj = csg;
		g.type = CPxGeometryType_eSPHERE;
		return g;
	}

	//
	// CPxBoxGeometry
	//
	CPxAPI CPxInline CSTRUCT CPxBoxGeometry CPxGeometry_toCPxBoxGeometry(CSTRUCT CPxGeometry cg)
	{
		return *(CSTRUCT CPxBoxGeometry*)(cg.obj);
		//return *static_cast<CSTRUCT CPxBoxGeometry*>(cg.obj);
	}

	CPxAPI CPxInline CSTRUCT CPxGeometry CPxBoxGeometry_toCPxGeometry(CSTRUCT CPxBoxGeometry* cbg)
	{
		CSTRUCT CPxGeometry g;
		g.obj = cbg;
		g.type = CPxGeometryType_eBOX;
		return g;
	}

	//
	// CPxCapsuleGeometry
	//
	CPxAPI CPxInline CSTRUCT CPxCapsuleGeometry CPxGeometry_toCPxCapsuleGeometry(CSTRUCT CPxGeometry cg)
	{
		return *(CSTRUCT CPxCapsuleGeometry*)(cg.obj);
	}

	CPxAPI CPxInline CSTRUCT CPxGeometry CPxCapsuleGeometry_toCPxGeometry(CSTRUCT CPxCapsuleGeometry* ccg)
	{
		CSTRUCT CPxGeometry g;
		g.obj = ccg;
		g.type = CPxGeometryType_eCAPSULE;
		return g;
	}

#ifdef __cplusplus
}
#endif

#endif // !CPxGeometry_H