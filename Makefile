ifeq (${IMAGE_ORG}, )
  IMAGE_ORG="sonasingh46"
  export IMAGE_ORG
endif

# Specify the docker arg for repository url
ifeq (${DBUILD_REPO_URL}, )
  DBUILD_REPO_URL="https://github.com/sonasingh46/image-store-service"
  export DBUILD_REPO_URL
endif

# Specify the date of build
DBUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')


# Determine the arch/os
ifeq (${XC_OS}, )
  XC_OS:=$(shell go env GOOS)
endif
export XC_OS

ifeq (${XC_ARCH}, )
  XC_ARCH:=$(shell go env GOARCH)
endif
export XC_ARCH

ARCH:=${XC_OS}_${XC_ARCH}
export ARCH

ifeq (${IMAGE_TAG}, )
  IMAGE_TAG = ci
  export IMAGE_TAG
endif

export DBUILD_ARGS=--build-arg DBUILD_DATE=${DBUILD_DATE} --build-arg DBUILD_REPO_URL=${DBUILD_REPO_URL} --build-arg DBUILD_SITE_URL=${DBUILD_SITE_URL} --build-arg ARCH=${ARCH}

# Specify the name of the docker repo for amd64
IMAGE_STORE_SERVICE_REPO_NAME=image-store-service-amd64


# Specify the directory location of main package after bin directory
# e.g. bin/{DIRECTORY_NAME_OF_APP}
IMAGE_STORE_SERVICE=image-store-service

# list only the source code directories
PACKAGES = $(shell go list ./... | grep -v 'vendor\|pkg/client/generated\|tests')

# deps ensures fresh go.mod and go.sum.
.PHONY: deps
deps:
	@go mod tidy
	@go mod verify

.PHONY: test
test:
	go fmt ./...
	go test $(PACKAGES)

.PHONY: build
build:
	go build ./cmd/...



.PHONY: image-store-service.amd64
image-store-service.amd64 :
	@echo -n "--> Building image store service image <--"
	@echo "${IMAGE_ORG}/${IMAGE_STORE_SERVICE_REPO_NAME}:${IMAGE_TAG}"
	@echo "----------------------------"
	@PNAME=${IMAGE_STORE_SERVICE} CTLNAME=${IMAGE_STORE_SERVICE} sh -c "'$(PWD)/build/build.sh'"
	@cp bin/${IMAGE_STORE_SERVICE}/${IMAGE_STORE_SERVICE} build/image-store-service/
	@cd build/${IMAGE_STORE_SERVICE} && sudo docker build -t ${IMAGE_ORG}/${IMAGE_STORE_SERVICE_REPO_NAME}:${IMAGE_TAG} ${DBUILD_ARGS} .
	@rm build/${IMAGE_STORE_SERVICE}/${IMAGE_STORE_SERVICE}

.PHONY: all.amd64
all.amd64: image-store-service.amd64