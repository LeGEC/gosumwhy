package gosumwhy

import (
	"io"
	"strings"
)

func goModGraphSample() io.Reader {
	// simple example of dependency graph
	return strings.NewReader(`root mod1@v2.0.0
root mod2@v1.1.0
mod1@v2.0.0 mod2@v1.0.0
mod1@v2.0.0 mod3@v0.0.0-20220202123456-eacf32b
mod2@v1.1.0 mod1@v1.0.0
mod2@v1.1.0 mod4@v1.0.0
mod1@v1.0.0 mod2@v0.9.0
mod3@v0.0.0-20220202123456-eacf32b testmodule@v1.0.0
mod4@v1.0.0 mod5@v1.0.0
mod5@v1.0.0 testmodule@v1.0.0
`)
}

func goModGraphRscioQuote() io.Reader {
	// example output of 'go mod graph' for module 'rsc.io/quote'
	return strings.NewReader(`rsc.io/quote rsc.io/quote/v3@v3.0.0
rsc.io/quote rsc.io/sampler@v1.3.0
rsc.io/quote/v3@v3.0.0 rsc.io/sampler@v1.3.0
rsc.io/sampler@v1.3.0 golang.org/x/text@v0.0.0-20170915032832-14c0d48ead0c
`)
}

func goModGraphFilippoioAge() io.Reader {
	// example output of 'go mod graph' for module 'filippo.io/age' (at https://github.com/FiloSottile/age/tree/v1.1.0-rc.1)
	return strings.NewReader(`filippo.io/age filippo.io/edwards25519@v1.0.0-rc.1
filippo.io/age golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5
filippo.io/age golang.org/x/sys@v0.0.0-20210903071746-97244b99971b
filippo.io/age golang.org/x/term@v0.0.0-20210615171337-6886f2dfbf5b
golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/net@v0.0.0-20210226172049-e18ecbb05110
golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
golang.org/x/crypto@v0.0.0-20210817164053-32db794688a5 golang.org/x/text@v0.3.3
golang.org/x/term@v0.0.0-20210615171337-6886f2dfbf5b golang.org/x/sys@v0.0.0-20210615035016-665e8c7367d1
`)
}
