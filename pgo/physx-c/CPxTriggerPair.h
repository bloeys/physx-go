#ifndef CPxTriggerPair_H
#define CPxTriggerPair_H

#include "CPxShape.h"
#include "CPxPairFlag.h"
#include "CPxRigidActor.h"
#include "CPxTriggerPairFlag.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxTriggerPair
	{
		CSTRUCT CPxShape triggerShape;			//!< The shape that has been marked as a trigger.
		CSTRUCT CPxRigidActor triggerActor;		//!< The actor to which triggerShape is attached
		CSTRUCT CPxShape otherShape;			//!< The shape causing the trigger event. \deprecated (see #PxSimulationEventCallback::onTrigger()) If collision between trigger shapes is enabled, then this member might point to a trigger shape as well.
		CSTRUCT CPxRigidActor otherActor;		//!< The actor to which otherShape is attached
		CENUM CPxPairFlag status;				//!< Type of trigger event (eNOTIFY_TOUCH_FOUND or eNOTIFY_TOUCH_LOST). eNOTIFY_TOUCH_PERSISTS events are not supported.
		CENUM CPxTriggerPairFlag flags;			//!< Additional information on the pair (see #PxTriggerPairFlag)
	};

#ifdef __cplusplus
}
#endif

#endif