package gosumwhy

import (
	"testing"
)

const verbose = false

func TestVersion(t *testing.T) {
	specs := []string{
		"v0.0.0-20130726002347-9928fa2ce45e",
		"v0.0.0-20140211213409-345aaac39713",
		"v0.0.0-20141211081046-ead28755d3c0",
		"v0.0.0-20150114235600-33e0aa1cb7c0",
		"v0.0.0-20150225183255-68e9c0620927",
		"v0.0.0-20161110224834-b9bea6b2b173",
		"v0.0.0-20161228173917-9ac251b645a2",
		"v0.0.0-20170609003504-e2365dfdc4a0",
		"v0.0.0-20171010120322-cdade1c07385",
		"v0.0.0-20180105133128-c9caef953efb",
		"v0.0.0-20180428030007-95032a82bc51",
		"v0.0.0-20190221195224-5a805980a5f3",
		"v0.0.0-20190314090850-75b78581bebe",
		"v0.0.0-20190410193231-58a59202ab31",
		"v0.0.0-20191104211930-d1553a71de50",
		"v0.0.0-20200714003250-2b9c44734f2b",
		"v0.0.0-20200804184101-5ec99f83aff1",
		"v0.0.0-20200903113550-03f5f0333e1f",
		"v0.0.0-20201027041543-1326539a0a0a",
		"v0.0.0-20210105120005-909beea2cc74",
		"v0.0.0-20210112085537-c389da54e794",
		"v0.0.0-20210218163916-a377121e959e",
		"v0.0.0-20210331224755-41bb18bfe9da",
		"v0.0.0-20210819152912-ad72663a72ab",
		"v0.0.0-20211012122336-39d0f177ccd0",
		"v0.0.0-20220214123719-b09a6bfa842f",
		"v0.0.0-20220222213610-43724f9ea8cf",
		"v0.0.0-20220225172249-27dd8689420f",
		"v0.0.0-20220302094943-723b81ca9867",
		"v0.0.0-20220309155454-6242fa91716a",
		"v0.0.0-20220315194320-039c03cc5b86",
		"v0.0.0-20220328175248-053ad81199eb",
		"v0.0.0-20220507011949-2cf3adece122",
		"v0.0.1",
		"v0.0.13",
		"v0.1.0",
		"v0.1.8",
		"v0.1.10",
		"v0.3.0",
		"v0.3.4",
		"v0.3.7",
		"v0.3.9",
		"v0.5.1",
		"v0.5.3",
		"v0.5.7",
		"v0.6.0-dev.0.20220106191415-9b9b3d81d5e3",
		"v0.9.1",
		"v0.10.2",
		"v0.14.0-beta3",
		"v0.20.0",
		"v0.23.0",
		"v0.70.0",
		"v0.100.2",
		"v1.0.0",
		"v1.0.1",
		"v1.0.3",
		"v1.0.5",
		"v1.0.6",
		"v1.1.1",
		"v1.1.4",
		"v1.1.8",
		"v1.1.15",
		"v1.2.1",
		"v1.2.6",
		"v1.3.0",
		"v1.4.1",
		"v1.4.1-0.20201116162257-a2a8dda75c91",
		"v1.4.2",
		"v1.4.7",
		"v1.5.0",
		"v1.5.1",
		"v1.5.2",
		"v1.5.3",
		"v1.6.2",
		"v1.6.4",
		"v1.6.7",
		"v1.7.0",
		"v1.8.0",
		"v1.8.1",
		"v1.8.7",
		"v1.9.3",
		"v1.11.0",
		"v1.12.1",
		"v1.13.5",
		"v1.16.0",
		"v1.27.1",
		"v1.44.0",
		"v1.62.0",
		"v2.0.0",
		"v2.0.1",
		"v2.0.2+incompatible",
		"v2.1.1",
		"v2.1.3+incompatible",
		"v2.3.0",
		"v2.3.1",
		"v2.3.1+incompatible",
		"v2.3.4+incompatible",
		"v2.3.6+incompatible",
		"v2.4.0",
		"v2.6.0",
		"v3.0.0",
		"v3.0.0-20150716171945-2caba252f4dc",
		"v3.0.0-20210107192922-496545a6307b",
		"v3.0.1",
		"v3.2.0+incompatible",
		"v3.13.0+incompatible",
		"v3.21.10",
		"v4.0.0+incompatible",
		"v4.1.0+incompatible",
		"v4.16.1",
		"v5.3.4",
	}

	for _, s := range specs {
		if !(v000rx.MatchString(s) || vXYZrx.MatchString(s)) {
			t.Logf("could not match '%s'", s)
			t.Fail()
		}

		if verbose {
			chunks := v000rx.FindStringSubmatch(s)
			if chunks == nil {
				chunks = vXYZrx.FindStringSubmatch(s)
			}
			t.Logf("%d %v", len(chunks), chunks[1:])
		}
	}

	for i := range specs {
		for j := range specs {
			a, b := specs[i], specs[j]

			less := i < j

			got := lessVersionsString(a, b)
			if got != less {
				t.Logf("bad comparison for '%s' < '%s': expected %v got %v", a, b, less, got)
			}
		}
	}
}