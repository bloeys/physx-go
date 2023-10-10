#ifndef CPxContactPairHeader_H
#define CPxContactPairHeader_H

#include "CPxShape.h"
#include "CPxPairFlag.h"
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

		CENUM CPxPairFlag events;
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

#endif