# 🚀 **Memcached Operator**

> This project implements a Kubernetes operator for managing Memcached instances. The operator automates the deployment, scaling, and management of Memcached clusters in a Kubernetes environment. It leverages the Operator SDK to create admission webhooks for validating and mutating custom resources.

![Memcached Operator](./docs/images/mem-op.jpg)

## 📚 **Table of Contents**

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Kind Cluster Setup](#kind-cluster-setup)
  - [Docker Build and Push](#docker-build-and-push)
  - [Deploying the Operator](#deploying-the-operator)
  - [Testing](#testing-the-operator)
- [Usage](#usage)
- [Contributing](#contributing)
- [Code of Conduct](#code-of-conduct)
- [License](#license)
- [Please Star the Repo](#please-star-the-repo)

---

## 🏗️ **Project Structure**

- **Dockerfile, Makefile, README.md**: Project build, configuration, and documentation files. 🛠️
- **LICENSE**: License file for the project. 📄
- **PROJECT**: The project-specific metadata and configuration details. 📋
- **api**: Defines the custom resources, types, and webhooks for the operator.
- **bin**: Contains compiled binaries and tools for the operator.
- **cmd**: Main application entry point for starting the operator.
- **config**: Kubernetes and operator configurations.
- **docs**: Documentation related to the operator.
- **go.mod, go.sum**: Go module definition files for managing dependencies.
- **hack**: Helper scripts or files for project setup or boilerplate code.
- **internal**: Internal logic for the operator.
- **test**: Tests for the operator.

---

## 🏁 **Getting Started**

### 📋 **Prerequisites**

Before you start, make sure you have the following tools installed:

- **Kind** 🛠️
- **Docker** 🐳
- **kubectl** 🛑
- **kustomize** 🧩
- **Go** (for development) 🖥️

### 🌍 **Kind Cluster Setup**

1. **Kind cluster configuration**:

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  image: kindest/node:v1.32.2
- role: worker
  image: kindest/node:v1.32.2
- role: worker
  image: kindest/node:v1.32.2
```

2. **Apply the cluster configuration**:

```bash
kind create cluster --config kind.yaml
```

Follow the [instructions](docs/certs.md) to generate the necessary certificates for the operator.

### 🏗️ **Docker Build and Push**

Build and push the operator image to your Docker registry

### 🚀 **Deploying the Operator**

```bash
make deploy
```

### 🧪 **Testing the Operator**

Test the operator by creating a Memcached custom resource:

```bash
kubectl apply -f manifests/examples/memcached-cr.yaml
```

### 🧹 **Cleanup**

To clean up:

```bash
make undeploy
kind delete cluster
```

---

## 🛠️ **Usage**

The operator manages the lifecycle of Memcached resources. It ensures that the desired state (such as number of replicas and configuration) for Memcached instances is maintained. The operator will:

- Deploy Memcached resources based on custom resources.
- Automatically scale Memcached instances.
- Monitor the health of Memcached instances.

---

## 🙋‍♂️ **Contributing**

Pull requests are welcome! For major changes, please open an issue first to discuss what you'd like to change.

---

## 🧑‍🤝‍🧑 **Code of Conduct**

This project and everyone participating in it is governed by the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/). By participating, you are expected to uphold this code. Please report unacceptable behavior to the project team.

The goal of this code of conduct is to ensure a welcoming, respectful, and productive environment for all contributors. The expectations are:

- **Be respectful**: Treat everyone with kindness and empathy.
- **Be inclusive**: Embrace diversity of thought, background, and perspective.
- **Be responsible**: Take ownership of your contributions and interactions.
- **Be collaborative**: Foster a cooperative environment where knowledge sharing is encouraged.

For more details, please refer to the full [Code of Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct.html).

---

## 📄 **License**

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

---

## ⭐ **Please Star the Repo**

If you find this project useful, please consider giving it a ⭐️ on GitHub. Your star helps us reach more developers and contributors and supports the continued improvement of the project. It’s a great way to show appreciation! 😊
