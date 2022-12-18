#ifndef CPxRigidActorExt_H
#define CPxRigidActorExt_H

#include "CPxRigidActor.h"
#include "CPxShape.h"
#include "CPxGeometry.h"
#include "CPxMaterial.h"
#include "CPxShapeFlags.h"

#ifdef __cplusplus
extern "C" {
#endif

	CPxAPI CSTRUCT CPxShape createExclusiveShape(CSTRUCT CPxRigidActor actor, CSTRUCT CPxGeometry geometry, CSTRUCT CPxMaterial material, CENUM CPxShapeFlags shapeFlags);

#ifdef __cplusplus
}
#endif

#endif // !CPxRigidActorExt_H
