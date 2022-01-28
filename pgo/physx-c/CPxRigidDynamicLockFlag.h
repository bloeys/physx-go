#ifndef CPxRigidDynamicLockFlag_H
#define CPxRigidDynamicLockFlag_H

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxRigidDynamicLockFlag
	{
		CPxRigidDynamicLockFlag_eLOCK_LINEAR_X = (1 << 0),
		CPxRigidDynamicLockFlag_eLOCK_LINEAR_Y = (1 << 1),
		CPxRigidDynamicLockFlag_eLOCK_LINEAR_Z = (1 << 2),
		CPxRigidDynamicLockFlag_eLOCK_ANGULAR_X = (1 << 3),
		CPxRigidDynamicLockFlag_eLOCK_ANGULAR_Y = (1 << 4),
		CPxRigidDynamicLockFlag_eLOCK_ANGULAR_Z = (1 << 5)
	};

#ifdef __cplusplus
}
#endif

#endif // !CPxRigidDynamicLockFlag_H
