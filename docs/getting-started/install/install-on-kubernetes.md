---
position: 1
---

# Install Infra on Kubernetes

## Prerequisites

* Install [Helm](https://helm.sh/) (v3+)
* Install [Kubernetes](https://kubernetes.io/) (v1.14+)

## Install via Helm

Install infra via `helm`:

```
helm repo add infrahq https://helm.infrahq.com
helm repo update
helm install infra infrahq/infra
```

## Customize your install

```yaml
# example values.yaml
---
server:
  # Add an Identity Provider
  additionalProviders:
    - name: Okta
      url: example.okta.com
      clientID: example_jsldf08j23d081j2d12sd
      clientSecret:  example_plain_secret # see note above

  # Add an admin user
  additionalUsers:
    - name: admin
      password: password

  additionalGrants:
  # 1. Grant user(s) or group(s) as Infra administrator
  # Setup an user as Infra administrator
    - user: admin
      role: admin
      resource: infra

  # 2. Grant user(s) or group(s) access to a resources
  # Example of granting access to an individual user the `cluster-admin` role. The name of a resource is specified when installing the Infra Engine at that location.
    - user: admin
      role: cluster-admin                  # cluster_roles required
      resource: example-cluster            # limit access to the `example-cluster` Kubernetes cluster

  # Example of granting access to an individual user through assigning them to the 'edit' role in the `web` namespace.
  # In this case, Infra will automatically scope the access to a namespace.
    - user: admin
      role: edit                            # cluster_roles required
      resource: example-cluster.web         # limit access to only the `web` namespace in the `example-cluster` Kubernetes cluster

  # Example of granting access to a group the `view` role.
    - group: Everyone
      role: view                           # cluster_roles required
      resource: example-cluster            # limit access to the `example-cluster` Kubernetes cluster
```

### Ingress

Follow the [Ingress documentation](../reference/helm-chart.md#advanced-ingress-configuration) to configure your Infra server with a Kubernetes ingress.
Once configured, get the Infra API server URL.

```bash
kubectl get ingress infra-server -o jsonpath="{.status.loadBalancer.ingress[*]['ip', 'hostname']}"
```

```bash
infra login <INFRA_API_SERVER>
```

### Add connectors

#### Step 1: Create an Access Key

In order to add connectors to Infra, you will need to generate an access key. If you already have an access key, proceed to step 2.

> Using the Infra admin access key is _not_ recommended as it provides more privileges than is necessary for a connector and may pose a security risk.

```bash
infra keys add <keyName> connector
```

### Step 2: Install Infra Connector

Now that you have an access key, install Infra into your Kubernetes cluster.

```bash
helm upgrade --install infra-connector infrahq/infra --set connector.config.name=<clusterName> --set connector.config.server=<serverAddress> --set connector.config.accessKey=<accessKey>
```

> If the connector will live in the same cluster and namespace as the server, you can set `connector.config.server=localhost`.

> You may also need to set `connector.config.skipTLSVerify=true` if the server is using a self-signed certificate.

## Upgrade Infra

See [Upgrading Infra](./upgrading.md)
