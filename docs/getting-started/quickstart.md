# Quickstart

In this quickstart we'll get up and running using Infra to manage access to Kubernetes.

### Prerequisites

* Install [helm](https://helm.sh/docs/intro/install/) (v3+)
* Install Kubernetes [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) (v1.14+)
* A Kubernetes cluster. For local testing we recommend [Docker Desktop](https://www.docker.com/products/docker-desktop/)

### 1. Install Infra CLI

First, install `infra`:

<details>
  <summary><strong>macOS</strong></summary>

  ```bash
  brew install infrahq/tap/infra
  ```

  You may need to perform `brew link` if your symlinks are not working.
  ```bash
  brew link infrahq/tap/infra
  ```
</details>

<details>
  <summary><strong>Windows</strong></summary>

  ```powershell
  scoop bucket add infrahq https://github.com/infrahq/scoop.git
  scoop install infra
  ```

</details>

<details>
  <summary><strong>Linux</strong></summary>

  ```bash
  # Ubuntu & Debian
  echo 'deb [trusted=yes] https://apt.fury.io/infrahq/ /' | sudo tee /etc/apt/sources.list.d/infrahq.list
  sudo apt update
  sudo apt install infra
  ```
  ```bash
  # Fedora & Red Hat Enterprise Linux
  sudo dnf config-manager --add-repo https://yum.fury.io/infrahq/
  sudo dnf install infra
  ```
</details>

### 2. Run the infra server

We'll deploy the Infra server locally:

```
helm repo add infrahq https://helm.infrahq.com/
helm repo update
helm install infra infrahq/infra
```

### 3. Log in

Once the Infra server is running, login to the server to complete the setup.

```
infra login localhost --skip-tls-verify
```

### 4. Connect your first cluster

Next, we'll connect a Kubernetes cluster. We recommend [Docker Desktop](https://www.docker.com/products/docker-desktop/). Once your cluster is running locally (make sure Kubernetes is enabled), you can add a cluster via `infra destinations add`:

```
# switch to the kubernetes context for the cluster you want to add
kubectl config use-context docker-desktop

# add the cluster to infra
infra destinations add docker-desktop
```

Your cluster will be added. To verify it's connected:

```
infra destinations list
```

### 4. Set up your first user with view access to the cluster

Now that the Infra server is setup you can create a user.  The `infra id add` command creates a one-time password for the user.

```
infra id add name@example.com
```

Grant the user view (read-only) privileges.

```
infra grants add name@example.com kubernetes.docker-desktop --role view
```

### 5. Log in

Now, log in as the new user

```
infra login localhost
```

Switch to our desktop cluster

```
infra use docker-desktop
```

Finally, access resources in your cluster:

```
# works
kubectl get pods

# does not work
kubectl create namespace test
```

Congratulations, you've:
* Installed Infra
* Connected your first cluster
* Created a user and granted them `view` access to the cluster

### Next Steps

* [Connect Okta](../guides/identity-providers/okta.md) to onboard & offboard your team automatically
* [Manage & revoke access](../guides/granting-access.md) to users or groups
* [Understand Kubernetes roles](../connectors/kubernetes.md#roles) for understand different access levels Infra supports for Kubernetes
* [Customize your install](../install/install-on-kubernetes.md)

