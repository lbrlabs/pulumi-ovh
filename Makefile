PACK             := ovh
ORG              := pulumiverse
PROJECT          := github.com/${ORG}/pulumi-${PACK}
NODE_MODULE_NAME := @pulumiverse/${PACK}
TF_NAME          := ${PACK}
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version
JAVA_GEN         := pulumi-java-gen
JAVA_GEN_VERSION := v0.5.4
PLUGIN_PATH      := ${HOME}/.pulumi/plugins/

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell pulumictl get version)

TESTPARALLELISM := 10

WORKING_DIR     := $(shell pwd)

.PHONY: development provider build_sdks build_nodejs build_dotnet build_go build_python

development:: install_plugins provider lint_provider build_sdks build_java # Build the provider & SDKs for a development environment

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
build:: install_plugins provider build_sdks install_sdks
only_build:: build

tfgen:: install_plugins
	(cd provider && go build -a -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN})
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}
	(cd provider && VERSION=$(VERSION) go generate cmd/${PROVIDER}/main.go)

provider:: tfgen install_plugins # build the provider binary
	(cd provider && go build -a -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_nodejs build_python build_go build_dotnet build_java  # build all the sdks

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs:: install_plugins tfgen # build the node sdk
	$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/
	cd sdk/nodejs/ && \
        yarn install && \
        yarn run tsc && \
        cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
    	sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

build_python:: PYPI_VERSION := $(shell pulumictl get version --language python)
build_python:: install_plugins tfgen # build the python sdk
	$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

build_dotnet:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
build_dotnet:: install_plugins tfgen # build the dotnet sdk
	pulumictl get version --language dotnet
	$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
        dotnet build /p:Version=${DOTNET_VERSION}

build_go:: install_plugins tfgen # build the go sdk
	$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

build_java:: PACKAGE_VERSION := $(shell pulumictl get version --language generic)
build_java:: bin/pulumi-java-gen
	$(PLUGIN_PATH)/language-java-$(JAVA_GEN_VERSION)/$(JAVA_GEN) generate --schema provider/cmd/$(PROVIDER)/schema.json --out sdk/java  --build gradle-nexus
	cd sdk/java/ && \
		echo "module fake_java_module // Exclude this directory from Go tools\n\ngo 1.17" > go.mod && \
		gradle --console=plain build

bin/pulumi-java-gen:: 
	$(pulumi plugin install language java ${JAVA_GEN_VERSION})

lint_provider:: provider # lint the provider code
	cd provider && golangci-lint run -c ../.golangci.yml

cleanup:: # cleans up the temporary directory
	rm -r $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go

help::
	@grep '^[^.#]\+:\s\+.*#' Makefile | \
 	sed "s/\(.\+\):\s*\(.*\) #\s*\(.*\)/`printf "\033[93m"`\1`printf "\033[0m"`	\3 [\2]/" | \
 	expand -t20

clean::
	rm -rf sdk/{dotnet,nodejs,go,python}

install_plugins::
	[ -x $(shell which pulumi) ] || curl -fsSL https://get.pulumi.com | sh
	pulumi plugin install resource random 4.8.2
	pulumi plugin install resource openstack 3.9.0

install_dotnet_sdk::
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::
   pip3 install -e $(WORKING_DIR)/sdk/python

install_go_sdk::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

install_java_sdk::

install_sdks:: install_dotnet_sdk install_python_sdk install_nodejs_sdk

test::
	cd examples && go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h
