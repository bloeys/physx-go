#ifndef CPxTriggerPairFlag_H
#define CPxTriggerPairFlag_H

#ifdef __cplusplus
extern "C"
{
#endif

	enum CPxTriggerPairFlag
	{
		eREMOVED_SHAPE_TRIGGER = (1 << 0), //!< The trigger shape has been removed from the actor/scene.
		eREMOVED_SHAPE_OTHER = (1 << 1),   //!< The shape causing the trigger event has been removed from the actor/scene.
		eNEXT_FREE = (1 << 2)			   //!< For internal use only.
	};

#ifdef __cplusplus
}
#endif

#endif