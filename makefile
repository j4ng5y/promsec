VERSION := $(shell cat VERSION)

all: lint test build

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: test
test:
	@go test -coverprofile=coverage.out -race ./...

.PHONY: build
build: build_linux_amd64 build_linux_arm64 build_darwin_amd64 build_darwin_arm64

.PHONY: build_linux_amd64
build_linux_amd64: CGO_ENABLED=0
build_linux_amd64: GOOS=linux
build_linux_amd64: GOARCH=amd64
build_linux_amd64: cmd/promsec/promsec.go
	@go build \
		-a \
		-o="bin/promsec" \
		-ldflags="-X 'promsec/pkg/promsec.Version=${VERSION}'" \
		cmd/promsec/promsec.go
	@tar cf - LICENSE README.md -C bin promsec | tar cJf "out/promsec_${VERSION}_${GOOS}_${GOARCH}.tar.xz" @-
	@rm -rf bin

.PHONY: build_linux_arm64
build_linux_arm64: CGO_ENABLED=0
build_linux_arm64: GOOS=linux
build_linux_arm64: GOARCH=arm64
build_linux_arm64: cmd/promsec/promsec.go
	@go build \
		-a \
		-o="bin/promsec" \
		-ldflags="-X 'promsec/pkg/promsec.Version=${VERSION}'" \
		cmd/promsec/promsec.go
	@tar cf - LICENSE README.md -C bin promsec | tar cJf "out/promsec_${VERSION}_${GOOS}_${GOARCH}.tar.xz" @-
	@rm -rf bin

.PHONY: build_darwin_amd64
build_darwin_amd64: CGO_ENABLED=0
build_darwin_amd64: GOOS=darwin
build_darwin_amd64: GOARCH=amd64
build_darwin_amd64: cmd/promsec/promsec.go
	@go build \
		-a \
		-o="bin/promsec" \
		-ldflags="-X 'promsec/pkg/promsec.Version=${VERSION}'" \
		cmd/promsec/promsec.go
	@tar cf - LICENSE README.md -C bin promsec | tar cJf "out/promsec_${VERSION}_${GOOS}_${GOARCH}.tar.xz" @-
	@rm -rf bin

.PHONY: build_darwin_arm64
build_darwin_arm64: CGO_ENABLED=0
build_darwin_arm64: GOOS=darwin
build_darwin_arm64: GOARCH=arm64
build_darwin_arm64: cmd/promsec/promsec.go
	@go build \
		-a \
		-o="bin/promsec" \
		-ldflags="-X 'promsec/pkg/promsec.Version=${VERSION}'" \
		cmd/promsec/promsec.go
	@tar cf - LICENSE README.md -C bin promsec | tar cJf "out/promsec_${VERSION}_${GOOS}_${GOARCH}.tar.xz" @-
	@rm -rf bin