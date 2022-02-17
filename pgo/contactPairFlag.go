package pgo

type ContactPairFlag uint16

const (

	/**
	\brief The shape with index 0 has been removed from the actor/scene.
	*/
	ContactPairFlag_eREMOVED_SHAPE_0 ContactPairFlag = (1 << 0)

	/**
	\brief The shape with index 1 has been removed from the actor/scene.
	*/
	ContactPairFlag_eREMOVED_SHAPE_1 ContactPairFlag = (1 << 1)

	/**
	\brief First actor pair contact.

	The provided shape pair marks the first contact between the two actors, no other shape pair has been touching prior to the current simulation frame.

	\note: This info is only available if #PxPairFlag::eNOTIFY_TOUCH_FOUND has been declared for the pair.
	*/
	ContactPairFlag_eACTOR_PAIR_HAS_FIRST_TOUCH ContactPairFlag = (1 << 2)

	/**
	\brief All contact between the actor pair was lost.

	All contact between the two actors has been lost, no shape pairs remain touching after the current simulation frame.
	*/
	ContactPairFlag_eACTOR_PAIR_LOST_TOUCH ContactPairFlag = (1 << 3)

	/**
	\brief Internal flag, used by #PxContactPair.extractContacts()

	The applied contact impulses are provided for every contact point.
	This is the case if #PxPairFlag::eSOLVE_CONTACT has been set for the pair.
	*/
	ContactPairFlag_eINTERNAL_HAS_IMPULSES ContactPairFlag = (1 << 4)

	/**
	\brief Internal flag, used by #PxContactPair.extractContacts()

	The provided contact point information is flipped with regards to the shapes of the contact pair. This mainly concerns the order of the internal triangle indices.
	*/
	ContactPairFlag_eINTERNAL_CONTACTS_ARE_FLIPPED ContactPairFlag = (1 << 5)
)
