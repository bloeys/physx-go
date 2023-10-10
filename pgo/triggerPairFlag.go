package pgo

type TriggerPairFlag uint32

const (
	TriggerPairFlag_eREMOVED_SHAPE_TRIGGER TriggerPairFlag = (1 << 0) //!< The trigger shape has been removed from the actor/scene.
	TriggerPairFlag_eREMOVED_SHAPE_OTHER   TriggerPairFlag = (1 << 1) //!< The shape causing the trigger event has been removed from the actor/scene.
	TriggerPairFlag_eNEXT_FREE             TriggerPairFlag = (1 << 2) //!< For internal use only.
)
