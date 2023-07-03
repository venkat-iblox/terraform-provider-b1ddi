TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=github.com
NAMESPACE=infobloxopen
NAME=b1ddi
BINARY=terraform-provider-${NAME}
VERSION=0.1.3
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

build-local:
	go build -ldflags "-X main.terraformVersion=${VERSION}" -o ${BINARY}

install-local: build-local
	rm -rf ~/.terraform.d/plugins/registry.terraform.io/${NAMESPACE}/${NAME}/${VERSION}/darwin_amd64
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/${NAMESPACE}/${NAME}/${VERSION}/darwin_amd64
	mv ${BINARY} ~/.terraform.d/plugins/registry.terraform.io/${NAMESPACE}/${NAME}/${VERSION}/darwin_amd64/

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4 -coverprofile cover.out

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -coverprofile testacc-cover.out
