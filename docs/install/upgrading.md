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

## Upgrading Infra CLI

### macOS

1. Update Homebrew

    ```bash
    brew update
    ```

2. Upgrade Infra CLI

    ```bash
    brew upgrade infra
    ```

3. Check Infra CLI version

    ```bash
    infra version
    ```

### Windows

```powershell
scoop update infra
```

### Linux

```bash
# Ubuntu & Debian
sudo apt update
sudo apt upgrade infra
```

```bash
# Fedora & Red Hat Enterprise Linux
sudo dnf update infra
```

### Other Distributions

Binary releases can be downloaded and installed directly from the repository.

<details>
  <summary><strong>x86_64</strong></summary>

<!-- {x-release-please-start-version} -->
  ```bash
  LATEST=0.12.2
  curl -sSL https://github.com/infrahq/infra/releases/download/v$LATEST/infra_${LATEST}_linux_x86_64.zip
  unzip -d /usr/local/bin infra_${LATEST}_linux_x86_64.zip
  ```
<!-- {x-release-please-end} -->
</details>

<details>
  <summary><strong>ARM</strong></summary>

<!-- {x-release-please-start-version} -->
  ```bash
  LATEST=0.12.2
  curl -sSL https://github.com/infrahq/infra/releases/download/v$LATEST/infra_${LATEST}_linux_arm64.zip
  unzip -d /usr/local/bin infra_${LATEST}_linux_arm64.zip
  ```
<!-- {x-release-please-end} -->
</details>

[1]: https://github.com/infrahq/infra/releases/latest
