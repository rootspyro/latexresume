bin = latexresume
version = 1.0.3

local:
	go build -ldflags "-X main.version=$(version)" -o $(bin) .

local_release:
	goreleaser build --clean --snapshot	

release:
	goreleaser release
