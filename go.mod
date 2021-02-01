module go-sample

go 1.15

replace go-sample/add => ./add

require (
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/mattn/go-colorable v0.1.8 // indirect
	go-sample/add v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
)
