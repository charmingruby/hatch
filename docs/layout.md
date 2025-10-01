# Layout

Hatch enforces a clean, standardized directory structure that scales from single applications to complex monorepos.

## Single Application

For projects with a single application, use the `app/` directory at the root.

## Monorepo & Multiple Apps

For projects with multiple services, use the `apps/` directory:

```
.
├── apps/
│   ├── api/           # REST API service
│   ├── web/           # Frontend application
│   ├── worker/        # Background worker
│   └── gateway/       # API Gateway
├── pkg/               # Shared libraries across apps
│   ├── events/
│   └── clients/
├── go.work            # Go workspace (optional for Go projects)
├── docker-compose.yml # Local development orchestration
└── Makefile            # Root-level orchestration
```

### Guidelines

- **Each app is self-contained**: Own dependencies, configs, and Dockerfile
- **Share code via `pkg/`**: Common libraries, types, utilities
- **Independent deployment**: Each app can be built and deployed separately
- **Consistent structure**: All apps follow the same internal organization
- **Go workspace**: Use `go.work` to manage Go dependencies across apps

## Infrastructure

Infrastructure code lives in the `infra/` directory, organized by tool:

```
infra/
├── terraform/         # Infrastructure as Code
│   ├── environments/
│   │   ├── dev/
│   │   ├── staging/
│   │   └── prod/
│   ├── modules/       # Reusable Terraform modules
│   └── backend.tf
├── k8s/               # Kubernetes manifests
│   ├── base/          # Base configurations
│   └── overlays/      # Environment-specific overlays
│       ├── dev/
│       ├── staging/
│       └── prod/
├── helm/              # Helm charts
│   └── charts/
├── docker/            # Shared Dockerfiles and configs
│   ├── base/
│   └── nginx/
└── scripts/           # Infrastructure automation
```

### Guidelines

- **Environment isolation**: Separate configs for `dev`, `staging`, `prod`
- **Secrets management**: Use external secret stores (Vault, AWS Secrets Manager)
- **Version control**: Track all infrastructure changes in git
- **State management**: Use remote state for Terraform (S3, Terraform Cloud)

## Orchestration & Scripts

All project orchestration is centralized in the root `Makefile`, such as: provisioning, deploying, tearing down.

### Guidelines

- **Include help target**: Document all available commands
- **Use `.PHONY`**: Mark non-file targets as phony
- **Environment variables**: Allow configuration via env vars
- **Fail fast**: Use proper error handling
- **Descriptive names**: Use clear, self-documenting target names