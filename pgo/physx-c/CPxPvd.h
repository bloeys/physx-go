#ifndef CPxPvd_H
#define CPxPvd_H

#include "CPxFoundation.h"
#include "CPxPvdTransport.h"

#ifdef __cplusplus
extern "C" {
#endif

	enum CPxPvdInstrumentationFlag
	{
		/**
			\brief Send debugging information to PVD.

			This information is the actual object data of the rigid statics, shapes,
			articulations, etc.  Sending this information has a noticeable impact on
			performance and thus this flag should not be set if you want an accurate
			performance profile.
		 */
		CPxPvdInstrumentationFlag_eDEBUG = 1 << 0,

		/**
			\brief Send profile information to PVD.

			This information populates PVD's profile view.  It has (at this time) negligible
			cost compared to Debug information and makes PVD *much* more useful so it is quite
			highly recommended.

			This flag works together with a PxCreatePhysics parameter.
			Using it allows the SDK to send profile events to PVD.
		*/
		CPxPvdInstrumentationFlag_ePROFILE = 1 << 1,

		/**
			\brief Send memory information to PVD.

			The PVD sdk side hooks into the Foundation memory controller and listens to
			allocation/deallocation events.  This has a noticable hit on the first frame,
			however, this data is somewhat compressed and the PhysX SDK doesn't allocate much
			once it hits a steady state.  This information also has a fairly negligible
			impact and thus is also highly recommended.

			This flag works together with a PxCreatePhysics parameter,
			trackOutstandingAllocations.  Using both of them together allows users to have
			an accurate view of the overall memory usage of the simulation at the cost of
			a hashtable lookup per allocation/deallocation.  Again, PhysX makes a best effort
			attempt not to allocate or deallocate during simulation so this hashtable lookup
			tends to have no effect past the first frame.

			Sending memory information without tracking outstanding allocations means that
			PVD will accurate information about the state of the memory system before the
			actual connection happened.
		*/
		CPxPvdInstrumentationFlag_eMEMORY = 1 << 2,

		eALL = (CPxPvdInstrumentationFlag_eDEBUG | CPxPvdInstrumentationFlag_ePROFILE | CPxPvdInstrumentationFlag_eMEMORY)
	};

	struct CPxPvd
	{
		void* obj;
	};

	CPxAPI CSTRUCT CPxPvd CPxCreatePvd(CSTRUCT CPxFoundation);
	CPxAPI bool CPxPvd_connect(CSTRUCT CPxPvd, CSTRUCT CPxPvdTransport, CENUM CPxPvdInstrumentationFlag);
	CPxAPI void CPxPvd_release(CSTRUCT CPxPvd cpp);
#ifdef __cplusplus
}
#endif

#endif