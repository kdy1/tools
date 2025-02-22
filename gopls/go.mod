module golang.org/x/tools/gopls

go 1.18

require (
	github.com/google/go-cmp v0.5.7
	github.com/jba/printsrc v0.2.2
	github.com/jba/templatecheck v0.6.0
	github.com/sergi/go-diff v1.1.0
	golang.org/x/mod v0.5.1
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9
	golang.org/x/tools v0.1.8
	honnef.co/go/tools v0.2.2
	mvdan.cc/gofumpt v0.2.1
	mvdan.cc/xurls/v2 v2.3.0
)

require (
	github.com/BurntSushi/toml v1.0.0 // indirect
	github.com/google/safehtml v0.0.2 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace golang.org/x/tools => ../
