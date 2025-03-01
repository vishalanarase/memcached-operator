# memcached-operator

> This is a simple operator for memcached.

## Init

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

## Create API