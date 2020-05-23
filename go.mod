module github.com/mkrufky/webchaind

go 1.14

require (
	github.com/Microsoft/go-winio v0.4.14
	github.com/boltdb/bolt v1.3.1
	github.com/davecgh/go-spew v1.1.1
	github.com/denisbrodbeck/machineid v0.8.0 //mark
	github.com/ethereumproject/benchmark v0.0.0-20180113190147-8eff34efba25
	github.com/fatih/color v1.7.0
	github.com/gizak/termui v2.3.0+incompatible
	github.com/hashicorp/golang-lru v0.5.1
	github.com/huin/goupnp v1.0.0
	github.com/jackpal/go-nat-pmp v1.0.1
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983
	github.com/maruel/panicparse v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/nsf/termbox-go v0.0.0-20190624072549-eeb6cd0a1762 // indirect
	github.com/omeid/go-resources v0.0.0-20200113210624-eb442c910d63
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/peterh/liner v1.1.0
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/rjeczalik/notify v0.9.1
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/rs/cors v0.0.0-20170727213201-7af7a1e09ba3
	github.com/spf13/afero v1.2.2
	github.com/stretchr/testify v1.3.0
	github.com/syndtr/goleveldb v0.0.0-20171214120811-34011bf325bc
	github.com/mkrufky/cryptonight v0.0.0-20200523201908-55517619d7d9
	github.com/webchain-network/sputnikvm-ffi v0.0.0-20180625070200-47e0403cc9b4
	golang.org/x/crypto v0.0.0-20190418165655-df01cb2cc480
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299
	gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
	gopkg.in/fatih/set.v0 v0.1.0
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951
	gopkg.in/urfave/cli.v1 v1.17.0
)

replace github.com/mkrufky/webchaind/accounts/abi/bind v0.3.1-0.20200523200816-82ad8c02d13e => ./accounts/abi/bind
