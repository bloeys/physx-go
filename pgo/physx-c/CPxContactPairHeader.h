#ifndef CPxContactPairHeader_H
#define CPxContactPairHeader_H

#include "CPxShape.h"
#include "CPxRigidActor.h"

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxContactPairFlag
	{

		/**
		\brief The shape with index 0 has been removed from the actor/scene.
		*/
		CPxContactPairFlag_eREMOVED_SHAPE_0 = (1 << 0),

		/**
		\brief The shape with index 1 has been removed from the actor/scene.
		*/
		CPxContactPairFlag_eREMOVED_SHAPE_1 = (1 << 1),

		/**
		\brief First actor pair contact.

		The provided shape pair marks the first contact between the two actors, no other shape pair has been touching prior to the current simulation frame.

		\note: This info is only available if #PxPairFlag::eNOTIFY_TOUCH_FOUND has been declared for the pair.
		*/
		CPxContactPairFlag_eACTOR_PAIR_HAS_FIRST_TOUCH = (1 << 2),

		/**
		\brief All contact between the actor pair was lost.

		All contact between the two actors has been lost, no shape pairs remain touching after the current simulation frame.
		*/
		CPxContactPairFlag_eACTOR_PAIR_LOST_TOUCH = (1 << 3),

		/**
		\brief Internal flag, used by #PxContactPair.extractContacts()

		The applied contact impulses are provided for every contact point.
		This is the case if #PxPairFlag::eSOLVE_CONTACT has been set for the pair.
		*/
		CPxContactPairFlag_eINTERNAL_HAS_IMPULSES = (1 << 4),

		/**
		\brief Internal flag, used by #PxContactPair.extractContacts()

		The provided contact point information is flipped with regards to the shapes of the contact pair. This mainly concerns the order of the internal triangle indices.
		*/
		CPxContactPairFlag_eINTERNAL_CONTACTS_ARE_FLIPPED = (1 << 5)
	};

	enum CPxPairFlags
	{
		/**
		\brief Process the contacts of this collision pair in the dynamics solver.

		\note Only takes effect if the colliding actors are rigid bodies.
		*/
		CPxPairFlags_eSOLVE_CONTACT = (1 << 0),

		/**
		\brief Call contact modification callback for this collision pair

		\note Only takes effect if the colliding actors are rigid bodies.

		@see PxContactModifyCallback
		*/
		CPxPairFlags_eMODIFY_CONTACTS = (1 << 1),

		/**
		\brief Call contact report callback or trigger callback when this collision pair starts to be in contact.

		If one of the two collision objects is a trigger shape (see #PxShapeFlag::eTRIGGER_SHAPE)
		then the trigger callback will get called as soon as the other object enters the trigger volume.
		If none of the two collision objects is a trigger shape then the contact report callback will get
		called when the actors of this collision pair start to be in contact.

		\note Only takes effect if the colliding actors are rigid bodies.

		\note Only takes effect if eDETECT_DISCRETE_CONTACT or eDETECT_CCD_CONTACT is raised

		@see PxSimulationEventCallback.onContact() PxSimulationEventCallback.onTrigger()
		*/
		CPxPairFlags_eNOTIFY_TOUCH_FOUND = (1 << 2),

		/**
		\brief Call contact report callback while this collision pair is in contact

		If none of the two collision objects is a trigger shape then the contact report callback will get
		called while the actors of this collision pair are in contact.

		\note Triggers do not support this event. Persistent trigger contacts need to be tracked separately by observing eNOTIFY_TOUCH_FOUND/eNOTIFY_TOUCH_LOST events.

		\note Only takes effect if the colliding actors are rigid bodies.

		\note No report will get sent if the objects in contact are sleeping.

		\note Only takes effect if eDETECT_DISCRETE_CONTACT or eDETECT_CCD_CONTACT is raised

		\note If this flag gets enabled while a pair is in touch already, there will be no eNOTIFY_TOUCH_PERSISTS events until the pair loses and regains touch.

		@see PxSimulationEventCallback.onContact() PxSimulationEventCallback.onTrigger()
		*/
		CPxPairFlags_eNOTIFY_TOUCH_PERSISTS = (1 << 3),

		/**
		\brief Call contact report callback or trigger callback when this collision pair stops to be in contact

		If one of the two collision objects is a trigger shape (see #PxShapeFlag::eTRIGGER_SHAPE)
		then the trigger callback will get called as soon as the other object leaves the trigger volume.
		If none of the two collision objects is a trigger shape then the contact report callback will get
		called when the actors of this collision pair stop to be in contact.

		\note Only takes effect if the colliding actors are rigid bodies.

		\note This event will also get triggered if one of the colliding objects gets deleted.

		\note Only takes effect if eDETECT_DISCRETE_CONTACT or eDETECT_CCD_CONTACT is raised

		@see PxSimulationEventCallback.onContact() PxSimulationEventCallback.onTrigger()
		*/
		CPxPairFlags_eNOTIFY_TOUCH_LOST = (1 << 4),

		/**
		\brief Call contact report callback when this collision pair is in contact during CCD passes.

		If CCD with multiple passes is enabled, then a fast moving object might bounce on and off the same
		object multiple times. Hence, the same pair might be in contact multiple times during a simulation step.
		This flag will make sure that all the detected collision during CCD will get reported. For performance
		reasons, the system can not always tell whether the contact pair lost touch in one of the previous CCD
		passes and thus can also not always tell whether the contact is new or has persisted. eNOTIFY_TOUCH_CCD
		just reports when the two collision objects were detected as being in contact during a CCD pass.

		\note Only takes effect if the colliding actors are rigid bodies.

		\note Trigger shapes are not supported.

		\note Only takes effect if eDETECT_CCD_CONTACT is raised

		@see PxSimulationEventCallback.onContact() PxSimulationEventCallback.onTrigger()
		*/
		CPxPairFlags_eNOTIFY_TOUCH_CCD = (1 << 5),

		/**
		\brief Call contact report callback when the contact force between the actors of this collision pair exceeds one of the actor-defined force thresholds.

		\note Only takes effect if the colliding actors are rigid bodies.

		\note Only takes effect if eDETECT_DISCRETE_CONTACT or eDETECT_CCD_CONTACT is raised

		@see PxSimulationEventCallback.onContact()
		*/
		CPxPairFlags_eNOTIFY_THRESHOLD_FORCE_FOUND = (1 << 6),

		/**
		\brief Call contact report callback when the contact force between the actors of this collision pair continues to exceed one of the actor-defined force thresholds.

		\note Only takes effect if the colliding actors are rigid bodies.

		\note If a pair gets re-filtered and this flag has previously been disabled, then the report will not get fired in the same frame even if the force threshold has been reached in the
		previous one (unless #eNOTIFY_THRESHOLD_FORCE_FOUND has been set in the previous frame).

		\note Only takes effect if eDETECT_DISCRETE_CONTACT or eDETECT_CCD_CONTACT is raised

		@see PxSimulationEventCallback.onContact()
		*/
		CPxPairFlags_eNOTIFY_THRESHOLD_FORCE_PERSISTS = (1 << 7),

		/**
		\brief Call contact report callback when the contact force between the actors of this collision pair falls below one of the actor-defined force thresholds (includes the case where this collision pair stops being in contact).

		\note Only takes effect if the colliding actors are rigid bodies.

		\note If a pair gets re-filtered and this flag has previously been disabled, then the report will not get fired in the same frame even if the force threshold has been reached in the
		previous one (unless #eNOTIFY_THRESHOLD_FORCE_FOUND or #eNOTIFY_THRESHOLD_FORCE_PERSISTS has been set in the previous frame).

		\note Only takes effect if eDETECT_DISCRETE_CONTACT or eDETECT_CCD_CONTACT is raised

		@see PxSimulationEventCallback.onContact()
		*/
		CPxPairFlags_eNOTIFY_THRESHOLD_FORCE_LOST = (1 << 8),

		/**
		\brief Provide contact points in contact reports for this collision pair.

		\note Only takes effect if the colliding actors are rigid bodies and if used in combination with the flags eNOTIFY_TOUCH_... or eNOTIFY_THRESHOLD_FORCE_...

		\note Only takes effect if eDETECT_DISCRETE_CONTACT or eDETECT_CCD_CONTACT is raised

		@see PxSimulationEventCallback.onContact() PxContactPair PxContactPair.extractContacts()
		*/
		CPxPairFlags_eNOTIFY_CONTACT_POINTS = (1 << 9),

		/**
		\brief This flag is used to indicate whether this pair generates discrete collision detection contacts.

		\note Contacts are only responded to if eSOLVE_CONTACT is enabled.
		*/
		CPxPairFlags_eDETECT_DISCRETE_CONTACT = (1 << 10),

		/**
		\brief This flag is used to indicate whether this pair generates CCD contacts.

		\note The contacts will only be responded to if eSOLVE_CONTACT is enabled on this pair.
		\note The scene must have PxSceneFlag::eENABLE_CCD enabled to use this feature.
		\note Non-static bodies of the pair should have PxRigidBodyFlag::eENABLE_CCD specified for this feature to work correctly.
		\note This flag is not supported with trigger shapes. However, CCD trigger events can be emulated using non-trigger shapes
		and requesting eNOTIFY_TOUCH_FOUND and eNOTIFY_TOUCH_LOST and not raising eSOLVE_CONTACT on the pair.

		@see PxRigidBodyFlag::eENABLE_CCD
		@see PxSceneFlag::eENABLE_CCD
		*/
		CPxPairFlags_eDETECT_CCD_CONTACT = (1 << 11),

		/**
		\brief Provide pre solver velocities in contact reports for this collision pair.

		If the collision pair has contact reports enabled, the velocities of the rigid bodies before contacts have been solved
		will be provided in the contact report callback unless the pair lost touch in which case no data will be provided.

		\note Usually it is not necessary to request these velocities as they will be available by querying the velocity from the provided
		PxRigidActor object directly. However, it might be the case that the velocity of a rigid body gets set while the simulation is running
		in which case the PxRigidActor would return this new velocity in the contact report callback and not the velocity the simulation used.

		@see PxSimulationEventCallback.onContact(), PxContactPairVelocity, PxContactPairHeader.extraDataStream
		*/
		CPxPairFlags_ePRE_SOLVER_VELOCITY = (1 << 12),

		/**
		\brief Provide post solver velocities in contact reports for this collision pair.

		If the collision pair has contact reports enabled, the velocities of the rigid bodies after contacts have been solved
		will be provided in the contact report callback unless the pair lost touch in which case no data will be provided.

		@see PxSimulationEventCallback.onContact(), PxContactPairVelocity, PxContactPairHeader.extraDataStream
		*/
		CPxPairFlags_ePOST_SOLVER_VELOCITY = (1 << 13),

		/**
		\brief Provide rigid body poses in contact reports for this collision pair.

		If the collision pair has contact reports enabled, the rigid body poses at the contact event will be provided
		in the contact report callback unless the pair lost touch in which case no data will be provided.

		\note Usually it is not necessary to request these poses as they will be available by querying the pose from the provided
		PxRigidActor object directly. However, it might be the case that the pose of a rigid body gets set while the simulation is running
		in which case the PxRigidActor would return this new pose in the contact report callback and not the pose the simulation used.
		Another use case is related to CCD with multiple passes enabled, A fast moving object might bounce on and off the same
		object multiple times. This flag can be used to request the rigid body poses at the time of impact for each such collision event.

		@see PxSimulationEventCallback.onContact(), PxContactPairPose, PxContactPairHeader.extraDataStream
		*/
		CPxPairFlags_eCONTACT_EVENT_POSE = (1 << 14),

		CPxPairFlags_eNEXT_FREE = (1 << 15),        //!< For internal use only.

	/**
	\brief Provided default flag to do simple contact processing for this collision pair.
	*/
	CPxPairFlags_eCONTACT_DEFAULT = CPxPairFlags_eSOLVE_CONTACT | CPxPairFlags_eDETECT_DISCRETE_CONTACT,

	/**
	\brief Provided default flag to get commonly used trigger behavior for this collision pair.
	*/
	CPxPairFlags_eTRIGGER_DEFAULT = CPxPairFlags_eNOTIFY_TOUCH_FOUND | CPxPairFlags_eNOTIFY_TOUCH_LOST | CPxPairFlags_eDETECT_DISCRETE_CONTACT
	};

	struct CPxContactPairPoint
	{
		/**
		\brief The position of the contact point between the shapes, in world space.
		*/
		CSTRUCT CPxVec3	position;

		/**
		\brief The separation of the shapes at the contact point.  A negative separation denotes a penetration.
		*/
		CPxReal	separation;

		/**
		\brief The normal of the contacting surfaces at the contact point. The normal direction points from the second shape to the first shape.
		*/
		CSTRUCT CPxVec3	normal;

		/**
		\brief The surface index of shape 0 at the contact point.  This is used to identify the surface material.
		*/
		CPxU32   internalFaceIndex0;

		/**
		\brief The impulse applied at the contact point, in world space. Divide by the simulation time step to get a force value.
		*/
		CSTRUCT CPxVec3	impulse;

		/**
		\brief The surface index of shape 1 at the contact point.  This is used to identify the surface material.
		*/
		CPxU32   internalFaceIndex1;
	};

	struct CPxContactPair
	{
		CSTRUCT CPxShape shapes[2];

		CPxU8* contactPatches;
		CPxU8* contactPoints;
		CPxReal* contactImpulses;
		CPxU32 requiredBufferSize;
		CPxU8 contactCount;
		CPxU8 patchCount;
		CPxU16 contactStreamSize;

		CENUM CPxPairFlags events;
		CENUM CPxContactPairFlag flags;

		CSTRUCT	CPxContactPairPoint* extractedContactPoints;
	};

	enum CPxContactPairHeaderFlag
	{
		CPxContactPairHeaderFlag_eREMOVED_ACTOR_0 = (1 << 0),			//!< The actor with index 0 has been removed from the scene.
		CPxContactPairHeaderFlag_eREMOVED_ACTOR_1 = (1 << 1)			//!< The actor with index 1 has been removed from the scene.
	};

	struct CPxContactPairHeader
	{
		CSTRUCT CPxRigidActor actors[2];

		CPxU8* extraDataStream;
		CPxU16 extraDataStreamSize;

		CENUM CPxContactPairHeaderFlag	flags;

		CSTRUCT CPxContactPair* pairs;
		CPxU32 nbPairs;
	};

#ifdef __cplusplus
}
#endif

#endif // !1