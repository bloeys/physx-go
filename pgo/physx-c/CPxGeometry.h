#ifndef CPxGeometry_H
#define CPxGeometry_H

#include "CPxBoxGeometry.h"
#include "CPxSphereGeometry.h"

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

	CPxAPI CPxInline CSTRUCT CPxSphereGeometry CPxGeometry_toCPxSphere(CSTRUCT CPxGeometry);
	CPxAPI CPxInline CSTRUCT CPxGeometry CPxSphereGeometry_toCPxGeometry(CSTRUCT CPxSphereGeometry*);

	CPxAPI CPxInline CSTRUCT CPxBoxGeometry CPxGeometry_toCPxBox(CSTRUCT CPxGeometry);
	CPxAPI CPxInline CSTRUCT CPxGeometry CPxBoxGeometry_toCPxGeometry(CSTRUCT CPxBoxGeometry*);


#ifdef __cplusplus
}
#endif

#endif // !CPxGeometry_H