# memcached-operator

> This is a simple operator for memcached.

## Init Operator

```bash
$ operator-sdk init --domain devspace.com --repo github.com/vishalanarase/memcached-operator
INFO[0000] Writing kustomize manifests for you to edit...
INFO[0000] Writing scaffold for you to edit...
INFO[0000] Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.19.0
INFO[0006] Update dependencies:
$ go mod tidy
Next: define a resource with:
$ operator-sdk create api
```

## Create API and Controller

```bash
$ operator-sdk create api --group cache --version v1 --kind Memcached --resource --controller
INFO[0000] Writing kustomize manifests for you to edit...
INFO[0000] Writing scaffold for you to edit...
INFO[0000] api/v1/memcached_types.go
INFO[0000] api/v1/groupversion_info.go
INFO[0000] internal/controller/suite_test.go
INFO[0000] internal/controller/memcached_controller.go
INFO[0000] internal/controller/memcached_controller_test.go
INFO[0000] Update dependencies:
$ go mod tidy
INFO[0000] Running make:
$ make generate
mkdir -p /Users/vishal/workspace/vishalanarase/memcached-operator/bin
Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.16.1
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests
```

## Generate CRDs and RBAC manifests

```bash
$ make generate
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."

$ make manifests
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
```