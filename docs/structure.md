# Structure

## Project Layout

Hatch enforces a clean, standardized directory structure that scales from single applications to complex monorepos.

### Root Directory

```
.
├── app/               # Main application (single app projects)
├── apps/              # Multiple applications (monorepo)
├── infra/             # Infrastructure
├── docs/              # Project documentation
├── .github/           # CI/CD workflows and GitHub configs
├── Makefile           # Orchestration and development tasks
└── docker-compose.yml # Local environment setup
```

## Single Application

For projects with a single application, use the `app/` directory at the root:

```
app/
├── cmd/               # Application entrypoints
├── config/            # Configuration management
├── internal/          # Private application code
├── pkg/               # Public libraries (reusable)
├── test/              # Test utilities and mocks
├── db/                # Database migrations
├── Dockerfile
└── Makefile
```

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

### Monorepo Guidelines

- **Each app is self-contained**: Own dependencies, configs, and Dockerfile
- **Share code via `pkg/`**: Common libraries, types, utilities
- **Independent deployment**: Each app can be built and deployed separately
- **Consistent structure**: All apps follow the same internal organization
- **Go workspace**: Use `go.work` to manage Go dependencies across apps

### Benefits

- **Isolation**: Develop, test, and deploy apps independently
- **Code reuse**: Share common code without duplication
- **Atomic changes**: Update shared code and all apps in a single commit
- **Simplified CI/CD**: Single repository for all services

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

### Infrastructure Guidelines

- **Environment isolation**: Separate configs for `dev`, `staging`, `prod`
- **Secrets management**: Use external secret stores (Vault, AWS Secrets Manager)
- **Documentation**: Include `README.md` with setup instructions in `infra/`
- **Version control**: Track all infrastructure changes in git
- **State management**: Use remote state for Terraform (S3, Terraform Cloud)

## Orchestration & Scripts

All project orchestration is centralized in the root `Makefile`.

### Purpose

The Makefile provides a consistent interface for: Provisioning, deploying, tearing down.

### Best Practices

1. **Include help target**: Document all available commands
2. **Use `.PHONY`**: Mark non-file targets as phony
3. **Environment variables**: Allow configuration via env vars
4. **Fail fast**: Use proper error handling
5. **Descriptive names**: Use clear, self-documenting target names
