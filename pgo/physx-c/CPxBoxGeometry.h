#ifndef CPxBoxGeometry_H
#define CPxBoxGeometry_H

#ifdef __cplusplus
extern "C" {
#endif

	struct CPxBoxGeometry
	{
		float hx, hy, hz;
	};

	CPxAPI CPxInline CSTRUCT CPxBoxGeometry NewCPxBoxGeometry(float hx, float hy, float hz)
	{
		CSTRUCT CPxBoxGeometry cbg;
		cbg.hx = hx;
		cbg.hy = hy;
		cbg.hz = hz;

		return cbg;
	}

#ifdef __cplusplus
}
#endif

#endif // !CPxBoxGeometry_H