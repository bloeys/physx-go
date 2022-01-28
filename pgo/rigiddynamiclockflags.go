package pgo

type RigidDynamicLockFlags uint32

const (
	RigidDynamicLockFlags_eLOCK_LINEAR_X  RigidDynamicLockFlags = (1 << 0)
	RigidDynamicLockFlags_eLOCK_LINEAR_Y  RigidDynamicLockFlags = (1 << 1)
	RigidDynamicLockFlags_eLOCK_LINEAR_Z  RigidDynamicLockFlags = (1 << 2)
	RigidDynamicLockFlags_eLOCK_ANGULAR_X RigidDynamicLockFlags = (1 << 3)
	RigidDynamicLockFlags_eLOCK_ANGULAR_Y RigidDynamicLockFlags = (1 << 4)
	RigidDynamicLockFlags_eLOCK_ANGULAR_Z RigidDynamicLockFlags = (1 << 5)
)
