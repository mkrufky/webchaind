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
	github.com/golang/snappy v0.0.1 // indirect
	github.com/hashicorp/golang-lru v0.5.1
	github.com/huin/goupnp v1.0.0
	github.com/ianlancetaylor/demangle v0.0.0-20200523230325-2e7988d3ae45 // indirect
	github.com/jackpal/go-nat-pmp v1.0.1
	github.com/mailru/easyjson v0.0.0-20190403194419-1ea4449da983
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mkrufky/cryptonight v0.0.0-20200523235742-5a895e3780cf
	github.com/mkrufky/sputnikvm-ffi v0.0.0-20200524000605-0b07ad5840e1
	github.com/mkrufky/webchaind/accounts/abi/bind v0.3.1-0.20200523222546-1f2dc81edd9c
	github.com/omeid/go-resources v0.0.0-20200113210624-eb442c910d63
	github.com/peterh/liner v1.1.0
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/rjeczalik/notify v0.9.1
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/rs/cors v0.0.0-20170727213201-7af7a1e09ba3
	github.com/spf13/afero v1.2.2
	github.com/stretchr/testify v1.3.0
	github.com/syndtr/goleveldb v0.0.0-20171214120811-34011bf325bc
	golang.org/x/crypto v0.0.0-20190418165655-df01cb2cc480
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	golang.org/x/sys v0.0.0-20200519105757-fe76b779f299
	gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
	gopkg.in/fatih/set.v0 v0.1.0
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/urfave/cli.v1 v1.17.0
)

replace github.com/mkrufky/webchaind/accounts/abi/bind v0.3.1-0.20200523222546-4409f898399d => ./accounts/abi/bind