# memcached-operator

> This is a simple operator for memcached.

## Init Operator

```bash
operator-sdk init --domain devspace.com --repo github.com/vishalanarase/memcached-operator
```
<details>
<summary>Output</summary>
```bash
INFO[0000] Writing kustomize manifests for you to edit...
INFO[0000] Writing scaffold for you to edit...
INFO[0000] Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.19.0
INFO[0006] Update dependencies:
$ go mod tidy
Next: define a resource with:
$ operator-sdk create api
```
</details>

## Create API and Controller

```bash
operator-sdk create api --group cache --version v1 --kind Memcached --resource --controller
```

<details>
<summary>Output</summary>
```bash
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
</details>

## Generate CRDs and RBAC manifests

```bash
$ make generate
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."

$ make manifests
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
```

## Deploy the controller to the cluster

```bash
make deploy
```

<details>
<summary>Output</summary>
```bash
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
Downloading sigs.k8s.io/kustomize/kustomize/v5@v5.4.3
cd config/manager && /Users/vishal/workspace/vishalanarase/memcached-operator/bin/kustomize edit set image controller=vishalanarase/memcached-operator:0.0.1
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/kustomize build config/default | kubectl apply -f -
namespace/memcached-operator-system created
customresourcedefinition.apiextensions.k8s.io/memcacheds.cache.devspace.com created
serviceaccount/memcached-operator-controller-manager created
role.rbac.authorization.k8s.io/memcached-operator-leader-election-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-manager-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-memcached-editor-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-memcached-viewer-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-metrics-auth-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-metrics-reader created
rolebinding.rbac.authorization.k8s.io/memcached-operator-leader-election-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/memcached-operator-manager-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/memcached-operator-metrics-auth-rolebinding created
service/memcached-operator-controller-manager-metrics-service created
deployment.apps/memcached-operator-controller-manager created
```
</details>

## Admission Webhooks

```bash
operator-sdk create webhook --group cache --version v1 --kind Memcached --defaulting --programmatic-validation
```

<details>
<summary>Output</summary>
```bash
INFO[0000] Writing kustomize manifests for you to edit...
INFO[0000] Writing scaffold for you to edit...
INFO[0000] api/v1/memcached_webhook.go
INFO[0000] api/v1/memcached_webhook_test.go
INFO[0000] api/v1/webhook_suite_test.go
INFO[0003] Update dependencies:
$ go mod tidy
INFO[0003] Running make:
$ make generate
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new Webhook and generate the manifests with:
$ make manifests
```
</details>

- Uncommenting sections in `config/default/kustomization.yaml` to enable webhook and cert-manager configuration through kustomize.

## Generate the manifests

```bash
make manifests
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
```

## Install the Cert Manager if you don't have it already

From [here](https://cert-manager.io/docs/installation/)


## Deploy the operator

```bash
make undeploy
```

```bash
make deploy
```

<details>
<summary>Output</summary>
```bash
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/controller-gen rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
cd config/manager && /Users/vishal/workspace/vishalanarase/memcached-operator/bin/kustomize edit set image controller=vishalanarase/memcached-operator:0.0.1
/Users/vishal/workspace/vishalanarase/memcached-operator/bin/kustomize build config/default | kubectl apply -f -
namespace/memcached-operator-system created
customresourcedefinition.apiextensions.k8s.io/memcacheds.cache.devspace.com created
serviceaccount/memcached-operator-controller-manager created
role.rbac.authorization.k8s.io/memcached-operator-leader-election-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-manager-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-memcached-editor-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-memcached-viewer-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-metrics-auth-role created
clusterrole.rbac.authorization.k8s.io/memcached-operator-metrics-reader created
rolebinding.rbac.authorization.k8s.io/memcached-operator-leader-election-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/memcached-operator-manager-rolebinding created
clusterrolebinding.rbac.authorization.k8s.io/memcached-operator-metrics-auth-rolebinding created
service/memcached-operator-controller-manager-metrics-service created
service/memcached-operator-webhook-service created
deployment.apps/memcached-operator-controller-manager created
certificate.cert-manager.io/memcached-operator-serving-cert created
issuer.cert-manager.io/memcached-operator-selfsigned-issuer created
mutatingwebhookconfiguration.admissionregistration.k8s.io/memcached-operator-mutating-webhook-configuration created
validatingwebhookconfiguration.admissionregistration.k8s.io/memcached-operator-validating-webhook-configuration created
```
<details>