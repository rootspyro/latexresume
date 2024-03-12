bin = latexresume
version = 1.0.2

local:
	go build -ldflags "-X main.appVersion=$(version)" -o $(bin) .

release:
	
