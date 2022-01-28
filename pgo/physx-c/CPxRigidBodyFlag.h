#ifndef CPxRigidBodyFlag_H
#define CPxRigidBodyFlag_H

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxRigidBodyFlag
	{
		/**
		\brief Enables kinematic mode for the actor.

		Kinematic actors are special dynamic actors that are not
		influenced by forces (such as gravity), and have no momentum. They are considered to have infinite
		mass and can be moved around the world using the setKinematicTarget() method. They will push
		regular dynamic actors out of the way. Kinematics will not collide with static or other kinematic objects.

		Kinematic actors are great for moving platforms or characters, where direct motion control is desired.

		You can not connect Reduced joints to kinematic actors. Lagrange joints work ok if the platform
		is moving with a relatively low, uniform velocity.

		<b>Sleeping:</b>
		\li Setting this flag on a dynamic actor will put the actor to sleep and set the velocities to 0.
		\li If this flag gets cleared, the current sleep state of the actor will be kept.

		\note kinematic actors are incompatible with CCD so raising this flag will automatically clear eENABLE_CCD

		@see PxRigidDynamic.setKinematicTarget()
		*/
		CPxRigidBodyFlag_eKINEMATIC = (1 << 0),		//!< Enable kinematic mode for the body.

		/**
		\brief Use the kinematic target transform for scene queries.

		If this flag is raised, then scene queries will treat the kinematic target transform as the current pose
		of the body (instead of using the actual pose). Without this flag, the kinematic target will only take
		effect with respect to scene queries after a simulation step.

		@see PxRigidDynamic.setKinematicTarget()
		*/
		CPxRigidBodyFlag_eUSE_KINEMATIC_TARGET_FOR_SCENE_QUERIES = (1 << 1),

		/**
		\brief Enables swept integration for the actor.

		If this flag is raised and CCD is enabled on the scene, then this body will be simulated by the CCD system to ensure that collisions are not missed due to
		high-speed motion. Note individual shape pairs still need to enable PxPairFlag::eDETECT_CCD_CONTACT in the collision filtering to enable the CCD to respond to
		individual interactions.

		\note kinematic actors are incompatible with CCD so this flag will be cleared automatically when raised on a kinematic actor

		*/
		CPxRigidBodyFlag_eENABLE_CCD = (1 << 2),		//!< Enable CCD for the body.

		/**
		\brief Enabled CCD in swept integration for the actor.

		If this flag is raised and CCD is enabled, CCD interactions will simulate friction. By default, friction is disabled in CCD interactions because
		CCD friction has been observed to introduce some simulation artifacts. CCD friction was enabled in previous versions of the SDK. Raising this flag will result in behavior
		that is a closer match for previous versions of the SDK.

		\note This flag requires PxRigidBodyFlag::eENABLE_CCD to be raised to have any effect.
		*/
		CPxRigidBodyFlag_eENABLE_CCD_FRICTION = (1 << 3),

		/**
		\brief Register a rigid body for reporting pose changes by the simulation at an early stage.

		Sometimes it might be advantageous to get access to the new pose of a rigid body as early as possible and
		not wait until the call to fetchResults() returns. Setting this flag will schedule the rigid body to get reported
		in #PxSimulationEventCallback::onAdvance(). Please refer to the documentation of that callback to understand
		the behavior and limitations of this functionality.

		@see PxSimulationEventCallback::onAdvance()
		*/
		CPxRigidBodyFlag_eENABLE_POSE_INTEGRATION_PREVIEW = (1 << 4),

		/**
		\brief Register a rigid body to dynamicly adjust contact offset based on velocity. This can be used to achieve a CCD effect.
		*/
		CPxRigidBodyFlag_eENABLE_SPECULATIVE_CCD = (1 << 5),

		/**
		\brief Permit CCD to limit maxContactImpulse. This is useful for use-cases like a destruction system but can cause visual artefacts so is not enabled by default.
		*/
		CPxRigidBodyFlag_eENABLE_CCD_MAX_CONTACT_IMPULSE = (1 << 6),

		/**
		\brief Carries over forces/accelerations between frames, rather than clearning them
		*/
		CPxRigidBodyFlag_eRETAIN_ACCELERATIONS = (1 << 7),

		/**
		\brief Forces kinematic-kinematic pairs notifications for this actor.

		This flag overrides the global scene-level PxPairFilteringMode setting for kinematic actors.
		This is equivalent to having PxPairFilteringMode::eKEEP for pairs involving this actor.

		A particular use case is when you have a large amount of kinematic actors, but you are only
		interested in interactions between a few of them. In this case it is best to use set
		PxSceneDesc.kineKineFilteringMode = PxPairFilteringMode::eKILL, and then raise the
		eFORCE_KINE_KINE_NOTIFICATIONS flag on the small set of kinematic actors that need
		notifications.

		\note This has no effect if PxRigidBodyFlag::eKINEMATIC is not set.

		\warning Changing this flag at runtime will not have an effect until you remove and re-add the actor to the scene.

		@see PxPairFilteringMode PxSceneDesc.kineKineFilteringMode
		*/
		CPxRigidBodyFlag_eFORCE_KINE_KINE_NOTIFICATIONS = (1 << 8),

		/**
		\brief Forces static-kinematic pairs notifications for this actor.

		Similar to eFORCE_KINE_KINE_NOTIFICATIONS, but for static-kinematic interactions.

		\note This has no effect if PxRigidBodyFlag::eKINEMATIC is not set.

		\warning Changing this flag at runtime will not have an effect until you remove and re-add the actor to the scene.

		@see PxPairFilteringMode PxSceneDesc.staticKineFilteringMode
		*/
		CPxRigidBodyFlag_eFORCE_STATIC_KINE_NOTIFICATIONS = (1 << 9),

		/**
		\brief Reserved for internal usage
		*/
		CPxRigidBodyFlag_eRESERVED = (1 << 15)
	};

#ifdef __cplusplus
}
#endif
#endif // !CPxRigidBodyFlag_H
