#ifndef CPxRigidDynamic_H
#define CPxRigidDynamic_H

#include "CPxActor.h"
#include "CPxRigidActor.h"
#include "CPxVec3.h"
#include "CPxTransform.h"
#include "CPxForceMode.h"
#include "CPxRigidBodyFlag.h"
#include "CPxRigidDynamicLockFlag.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxRigidDynamic
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxActor CPxRigidDynamic_toCPxActor(CSTRUCT CPxRigidDynamic*);
	CPxAPI CSTRUCT CPxRigidActor CPxRigidDynamic_toCPxRigidActor(CSTRUCT CPxRigidDynamic*);

	CPxAPI void CPxRigidDynamic_addForce(CSTRUCT CPxRigidDynamic* crd, CSTRUCT CPxVec3* force, CENUM CPxForceMode fmode, bool autoAwake);
	CPxAPI void CPxRigidDynamic_addTorque(CSTRUCT CPxRigidDynamic* crd, CSTRUCT CPxVec3* torque, CENUM CPxForceMode fmode, bool autoAwake);

	CPxAPI void CPxRigidDynamic_setLinearDamping(CSTRUCT CPxRigidDynamic* crd, CPxReal damping);
	CPxAPI void CPxRigidDynamic_setAngularDamping(CSTRUCT CPxRigidDynamic* crd, CPxReal damping);
	CPxAPI CPxReal CPxRigidDynamic_getLinearDamping(CSTRUCT CPxRigidDynamic* crd);
	CPxAPI CPxReal CPxRigidDynamic_getAngularDamping(CSTRUCT CPxRigidDynamic* crd);

	CPxAPI CSTRUCT CPxVec3 CPxRigidDynamic_getLinearVelocity(CSTRUCT CPxRigidDynamic* crd);
	CPxAPI CSTRUCT CPxVec3 CPxRigidDynamic_getAngularVelocity(CSTRUCT CPxRigidDynamic* crd);

	CPxAPI void CPxRigidDynamic_setMass(CSTRUCT CPxRigidDynamic* crd, CPxReal mass);
	CPxAPI CPxReal CPxRigidDynamic_getMass(CSTRUCT CPxRigidDynamic* crd);

	CPxAPI void CPxRigidDynamic_setRigidBodyFlag(CSTRUCT CPxRigidDynamic* crd, CENUM CPxRigidBodyFlag flag, bool value);
	CPxAPI void CPxRigidDynamic_setRigidBodyFlags(CSTRUCT CPxRigidDynamic* crd, CENUM CPxRigidBodyFlag flags);
	CPxAPI CENUM CPxRigidBodyFlag CPxRigidDynamic_getRigidBodyFlags(CSTRUCT CPxRigidDynamic* crd);

	CPxAPI void CPxRigidDynamic_setRigidDynamicLockFlag(CSTRUCT CPxRigidDynamic* crd, CENUM CPxRigidDynamicLockFlag flag, bool value);
	CPxAPI void CPxRigidDynamic_setRigidDynamicLockFlags(CSTRUCT CPxRigidDynamic* crd, CENUM CPxRigidDynamicLockFlag flags);
	CPxAPI CENUM CPxRigidDynamicLockFlag CPxRigidDynamic_getRigidDynamicLockFlags(CSTRUCT CPxRigidDynamic* crd);

	CPxAPI void CPxRigidDynamic_putToSleep(CSTRUCT CPxRigidDynamic* crd);
	CPxAPI CSTRUCT CPxTransform CPxRigidDynamic_getGlobalPose(CSTRUCT CPxRigidDynamic* crd);
	CPxAPI void CPxRigidDynamic_setGlobalPose(CSTRUCT CPxRigidDynamic* crd, CSTRUCT CPxTransform* tr, bool autoAwake);

	CPxAPI CSTRUCT CPxTransform CPxRigidDynamic_getCMassLocalPose(CSTRUCT CPxRigidDynamic* crd);
	CPxAPI void CPxRigidDynamic_setCMassLocalPose(CSTRUCT CPxRigidDynamic* crd, CSTRUCT CPxTransform* tr);

#ifdef __cplusplus
}
#endif

#endif // !CPxRigidDynamic_H