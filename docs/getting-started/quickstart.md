---
title: Quickstart
position: 2
---

# Quickstart

In this quickstart we'll set up Infra to manage user access to a Kubernetes cluster.

## Prerequisites

* Install [helm](https://helm.sh/docs/intro/install/) (v3+)
* Install Kubernetes [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) (v1.14+)
* A Kubernetes cluster. For local testing we recommend [Docker Desktop](https://www.docker.com/products/docker-desktop/)

## Install Infra CLI

{% tabs %}
{% tab label="macOS" %}
```
brew install infrahq/tap/infra
```
{% /tab %}
{% tab label="Windows" %}
```powershell
scoop bucket add infrahq https://github.com/infrahq/scoop.git
scoop install infra
```
{% /tab %}

{% tab label="Linux" %}

#### Ubuntu & Debian
```
echo 'deb [trusted=yes] https://apt.fury.io/infrahq/ /' | sudo tee /etc/apt/sources.list.d/infrahq.list
sudo apt update
sudo apt install infra
```

#### Fedora & Red Hat Enterprise Linux
```
sudo dnf config-manager --add-repo https://yum.fury.io/infrahq/
sudo dnf install infra
```
{% /tab %}
{% /tabs %}


## Deploy Infra

Deploy Infra to your Kubernetes cluster via `helm`:

```
helm repo add infrahq https://helm.infrahq.com/
helm repo update
helm install infra infrahq/infra
```

Next, log into your instance of Infra to setup your admin account:

```
infra login localhost --skip-tls-verify
```


{% callout type="info" %}
If you're not using Docker Desktop, you'll be using a different endpoint than `localhost`, which can be found via the following `kubectl` command:

```
kubectl get service infra-server -o jsonpath="{.status.loadBalancer.ingress[*]['ip', 'hostname']}" -w
```

Note: it may take a few minutes for the LoadBalancer to be provisioned.

{% /callout %}



## Connect your first Kubernetes cluster

Generate a connector key:

```
infra keys add connector
```

Next, use this access key to connect your first cluster via `helm`. **Note:** this can be the same cluster used to install Infra in step 2.

Install the Infra connector via `helm`:

```
helm upgrade --install infra-connector infrahq/infra \
  --set connector.config.name=example-cluster \
  --set connector.config.server=localhost \
  --set connector.config.accessKey=<CONNECTOR_KEY> \
  --set connector.config.skipTLSVerify=true
```

{% callout type="info" %}
It may take a few minutes for the cluster to connect. You can verify the connection by running `infra destinations list`
{% /callout %}

## Add a user and grant cluster access

Next, add a user:

```
infra users add user@example.com
```

{% callout type="info" %}
Infra will provide you a one-time password. Please note this password for the next step.
{% /callout %}

Grant this user read-only access to the Kubernetes cluster you just connected to Infra:

```
infra grants add user@example.com example-cluster --role view
```

## Login as the example user

Use the one-time password in the previous step to log in as the user. You'll be prompted to change the user's password since it's this new user's first time logging in.

```
infra login localhost --skip-tls-verify
```

Next, view this user's cluster access. You should see the user has `view` access to the `example-cluster` cluster connected above:

```
infra list
```

Lastly, switch to this Kubernetes cluster and verify the user's access:

```
infra use example-cluster

# Works since the user has view access
kubectl get pods -A

# Does not work
kubectl create namespace test-namespace
```

Congratulations, you've:
* Installed Infra
* Connected your first cluster
* Created a user and granted them `view` access to the cluster

## Next Steps

* [Connect Okta](../guides/identity-providers/okta.md) to onboard & offboard your team automatically
* [Manage & revoke access](../guides/granting-access.md) to users or groups
* [Understand Kubernetes roles](../connectors/kubernetes.md#roles) for understand different access levels Infra supports for Kubernetes
* [Customize your install](../install/install-on-kubernetes.md)

