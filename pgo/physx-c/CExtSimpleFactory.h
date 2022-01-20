#ifndef CExtSimpleFactory_H
#define CExtSimpleFactory_H

#include "CPxPhysics.h"
#include "CPxGeometry.h"
#include "CPxTransform.h"
#include "CPxRigidDynamic.h"

#ifdef __cplusplus
extern "C" {
#endif

	CPxAPI CSTRUCT CPxRigidDynamic* CPxCreateDynamic(CSTRUCT CPxPhysics* sdk, CSTRUCT CPxTransform* transform, CSTRUCT CPxGeometry geometry, CSTRUCT CPxMaterial* material, CPxReal density, CSTRUCT CPxTransform* shapeOffset);

#ifdef __cplusplus
}
#endif

#endif // !CExtSimpleFactory_H
