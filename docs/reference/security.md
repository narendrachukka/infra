# Security
We take security very seriously. This document provides an overview of Infra's security model.

If you have found a security vulnerability please disclose it privately to us by email via security@infrahq.com.

## General Security
### HTTPS
Infra communicates over HTTPS by default. By default, Infra uses self-signed TLS certificates. By setting a DNS record to point at Infra server, valid TLS certificates will be generated via LetsEncrypt.

### Authentication
When users login to Infra as a valid user they are issued a session token with a 24 character secret that is randomly generated. The SHA256 hash of this token is stored server-side for token validation. This session token is stored locally under `~/.infra`.

## Deployment
When deploying Infra, we recommend Infra be deployed in its own namespace to minimize the deployment scope.

## Sensitive Information

These secrets can be stored in a variety of [secret storage backends](../install/configure/secrets.md), including Kubernetes secrets, Vault, AWS Secrets Manager, AWS SSM (Systems Manager Parameter Store), and some simple options exist for loading secrets from the OS or container, such as: loading secrets from environment variables, loading secrets from files on the file system, and even plaintext secrets directly in the configuration file (though this is not recommended). With all types except for `plaintext`, the respective secret object names are specified in the configuration file and the actual secret is never persisted in Infra's storage. In the case of all secret types (including `plaintext`), the secret data is [encrypted at rest in the db](#Encrypted_At_Rest).

### Okta secrets
Infra uses an Okta application client secret in order to allow users to authenticate via an OpenID Connect (OIDC) authorization code flow.

#### Okta client secret usage:
The client secret is loaded server-side from the specified Kubernetes secret only when a user is logging in via Okta.

## Encrypted At Rest

Sensitive data is always encrypted at rest. See [Keys](../install/configure/encryption.md) for more information.