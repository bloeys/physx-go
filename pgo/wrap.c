#include <stdint.h>
#include <stdbool.h>
#define CPxAPI
#define CPxInternalAPI
#define CPxInline
#define CSTRUCT struct
#define CENUM enum
#define CPxU32 uint32_t
#define CPxReal float

#include <CPxFoundation.h>
#include <CPxPvd.h>
#include <CPxCpuDispatcher.h>
#include <CPxDefaultCpuDispatcher.h>
#include <CPxPhysics.h>
#include <CPxPvdTransport.h>
#include <CPxScene.h>
#include <CPxSceneDesc.h>
#include <CPxTolerancesScale.h>
#include <CPxPlane.h>
#include <CPxActor.h>
#include <CPxRigidStatic.h>
#include <CPxRigidDynamic.h>
#include <CPxTransform.h>
#include <CPxPvdSceneClient.h>
#include <CExtSimpleFactory.h>
#include <CPxGeometry.h>
#include <CPxBoxGeometry.h>
#include <CPxSphereGeometry.h>
#include <CPxCapsuleGeometry.h>