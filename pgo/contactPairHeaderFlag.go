package pgo

type ContactPairHeaderFlag uint16

const (
	ContactPairHeaderFlag_eREMOVED_ACTOR_0 ContactPairHeaderFlag = (1 << 0) //!< The actor with index 0 has been removed from the scene.
	ContactPairHeaderFlag_eREMOVED_ACTOR_1 ContactPairHeaderFlag = (1 << 1) //!< The actor with index 1 has been removed from the scene.
)
