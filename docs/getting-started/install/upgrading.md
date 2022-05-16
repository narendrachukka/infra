---
position: 3
---

# Upgrading

You can also download the [latest Infra release][1] directly from the repository.

## Upgrading Infra

1. Update the Helm repository

    ```bash
    helm repo update infrahq
    ```

2. Upgrade Infra. Ensure when upgrading a Helm chart to pass the same configuration values as during installation.

    ```bash
    helm upgrade [-f values.yaml] infra infrahq/infra
    ```

3. Wait for the pods to finish upgrade

    ```bash
    kubectl wait --for=condition=ready pod --selector app.kubernetes.io/name=infra-server
    ```

4. Check Infra version

    ```bash
    infra version
    ```

## Upgrading Infra Kubernetes Connector

1. Update the Helm repository

    ```bash
    helm repo update infrahq
    ```

2. Upgrade Infra. If using Helm values files, ensure those are passed into the upgrade command.

    ```bash
    helm upgrade -f values.yaml infra-connector infrahq/infra
    ```

    If using output from `infra destinations add`, ensure the same arguments are being passed into the upgrade command.

    ```bash
    helm upgrade --set connector.config.name=... --set connector.config.accessKey=... --set connector.config.server=... infra-connector infrahq/infra
    ```

3. Wait for the pods to finish upgrade

    ```bash
    kubectl wait --for=condition=ready pod --selector app.kubernetes.io/name=infra-connector
    ```

4. Check Infra Kubernetes Connector version

    ```bash
    kubectl logs -l app.kubernetes.io/name=infra-connector | grep 'starting infra'
    ```

## Upgrading Infra CLI

{% tabs %}
{% tab label="macOS" %}
```
brew update
brew upgrade infra
```
{% /tab %}
{% tab label="Windows" %}
```powershell
scoop update infra
```
{% /tab %}
{% tab label="Linux" %}
```bash
# Ubuntu & Debian
sudo apt update
sudo apt upgrade infra
```

```bash
# Fedora & Red Hat Enterprise Linux
sudo dnf update infra
```
{% /tab %}
{% /tabs %}

### Other Distributions

Binary releases can be downloaded and installed directly from [the repository](https://github.com/infrahq/infra/releases).
