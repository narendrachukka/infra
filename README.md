<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://user-images.githubusercontent.com/251292/179098559-75b53555-e389-40cc-b910-0e53521efad2.svg">
    <img alt="logo" src="https://user-images.githubusercontent.com/251292/179098561-eaa231c1-5757-40d7-9e5f-628e5d9c3e47.svg">
  </picture>
</div>

<div align="center">

[![YouTube Channel Views](https://img.shields.io/youtube/channel/views/UCft1MzQs2BJdW8BIUu6WJkw?style=social)](https://www.youtube.com/channel/UCft1MzQs2BJdW8BIUu6WJkw) [![GitHub Repo stars](https://img.shields.io/github/stars/infrahq/infra?style=social)](https://github.com/infrahq/infra/stargazers) [![Twitter Follow](https://img.shields.io/twitter/follow/infrahq?style=social)](https://twitter.com/infrahq)

</div>

## Introduction

Infra manages access to infrastructure such as Kubernetes, with support for [more connectors](#connectors) coming soon.

- **Discover & access** infrastructure via a single command: `infra login`
- **No more out-of-sync credentials** for users (e.g. Kubeconfig)
- **Okta, Google, Azure AD** identity provider support for onboarding and offboarding
- **Fine-grained** access to specific resources that works with existing RBAC rules
- **API-first design** for managing access as code or via existing tooling
- **Temporary access** to coordinate access with systems like PagerDuty (coming soon)
- **Audit logs** for who did what, when to stay compliant (coming soon)

## Connectors

| Connector          | Status        | Documentation                                                        |
| ------------------ | ------------- | -------------------------------------------------------------------- |
| Kubernetes         | ✅ Available  | [Get started](https://infrahq.com/docs/manage/connectors/kubernetes) |
| Postgres           | _Coming soon_ | _Coming soon_                                                        |
| SSH                | _Coming soon_ | _Coming soon_                                                        |
| AWS                | _Coming soon_ | _Coming soon_                                                        |
| Container Registry | _Coming soon_ | _Coming soon_                                                        |
| MongoDB            | _Coming soon_ | _Coming soon_                                                        |
| Snowflake          | _Coming soon_ | _Coming soon_                                                        |
| MySQL              | _Coming soon_ | _Coming soon_                                                        |
| RDP                | _Coming soon_ | _Coming soon_                                                        |

## Documentation

- [Deploy Infra](https://infrahq.com/docs/getting-started/deploy)
- [Install Infra CLI](https://infrahq.com/docs/start/install-infra-cli)
- [Helm Chart Reference](https://infrahq.com/docs/reference/helm-reference)
- [What is Infra?](https://infrahq.com/docs/getting-started/what-is-infra)
- [Architecture](https://infrahq.com/docs/reference/architecture)
- [Security](https://infrahq.com/docs/reference/security)

## Community

- [Community Forum](https://github.com/infrahq/infra/discussions) Best for: help with building, discussion about infrastructure access best practices.
- [GitHub Issues](https://github.com/infrahq/infra/issues) Best for: bugs and errors you encounter using Infra.
