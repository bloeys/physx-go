#ifndef CPxSimpleFactory_H
#define CPxSimpleFactory_H

#include "CPxPlane.h"
#include "CPxPhysics.h"
#include "CPxMaterial.h"
#include "CPxRigidStatic.h"

#ifdef __cplusplus
extern "C" {
#endif

	CPxAPI CSTRUCT CPxRigidStatic CPxCreatePlane(CSTRUCT CPxPhysics sdk, CSTRUCT CPxPlane* plane, CSTRUCT CPxMaterial material);

#ifdef __cplusplus
}
#endif

#endif // !CPxSimpleFactory_H
