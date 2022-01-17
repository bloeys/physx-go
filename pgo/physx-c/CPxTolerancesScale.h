#ifndef CPxTolerancesScale_H
#define CPxTolerancesScale_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxTolerancesScale
	{
		CPxReal length, speed;
	};

	CPxAPI CSTRUCT CPxTolerancesScale NewCPxTolerancesScale(CPxReal length, CPxReal speed);

#ifdef __cplusplus
}
#endif

#endif