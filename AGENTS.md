# Hatch Project Guidelines

Hatch is a **Go project template** with clear structure and tooling to accelerate development from day one. It's **battle-tested** and suitable for different contexts, from **corporate environments** and **production-grade products** to **quick PoCs**.

## Project Structure

### Infrastructure
- When the project requires infrastructure management, create an `infra/` directory at the root
- Organize by resource type in subdirectories: `charts/`, `terraform/`, `manifests/`
- Maintain consistency between local, staging, and production setups

### Monorepo & Multiple Apps
- Place all apps under the root-level `apps/` directory (e.g., `api/`, `web/`, `worker/`)
- Each app should be self-contained with its own code, configurations, and tests
- For Go projects, use Go workspace (`go.work`) to manage dependencies across apps
- Support multiple components or services under a single repository

### Project Identity
- Replace all instances of the placeholder `HATCH_APP` with the actual project module path
- Update Docker image names in GitHub Actions and docker compose files
- Configure repository secrets for automated workflows
- GitHub Actions includes a workflow to push container images to DockerHub

## Module Organization

### Base Structure
- Use `internal/` as the main module root
- Create a barrel file (`module_name.go`) for initialization at each module root
- Keep modules cohesive with clear separation of concerns
- Use the example module structure as a base

### Shared Contracts
- **Internal Shared**: For contracts reused across modules but private to the repo, use `internal/shared/`
- **Public API**: For libraries reusable across projects, use `pkg/`
- Ensure modules depend only on contracts, not directly on third-party SDKs

### External Integrations
- Within each module, use `external/` for contracts unique to that module
- Maintain clean separation between contracts and third-party implementations
- Example:
  ```
  internal/
    billing/
      external/
        payment_gateway.go    # Contract
        stripe/
          gateway.go         # Implementation
        paypal/
          gateway.go         # Implementation
  ```

## Communication Patterns

### HTTP (Synchronous)
- **REST APIs**: `http/rest/`
- **gRPC Services**: `http/grpc/`
- **GraphQL APIs**: `http/gql/`

### Messaging (Asynchronous)
- **Per-Module Events**:
  - `internal/MODULE/messaging/event/` — Defines module-specific events
  - `internal/MODULE/messaging/subscriber/` — Event listeners
  - Convention: event `transaction.created` → handler `onTransactionCreated`

- **Shared Contracts**:
  - `pkg/messaging/` — If reusable outside this repo
  - `internal/shared/messaging/` — If private to this repo

- **Implementations**: NATS, Kafka, etc. live in the same directory as contracts

## Persistence & Data

### Database
- PostgreSQL configured as the default database
- Includes SQL migrations and database setup scripts
- Implement repository patterns for data access

### Testing
- Test setup is integrated, including support for **mockery**
- Use Makefile test commands
- Follow the test scenario convention:
  ```
  <action> <result> when <condition>
  ```
- Examples:
  - `should return stash with default values when no options are provided`
  - `should return ErrExpired when key is expired`
  - `should store new entry successfully`

## Commands & Automation

### Makefile
- `make up` — Sets up development environment
- `make down` — Resets environment
- Integrated test commands
- Other utility commands for development

### CI/CD
- GitHub Actions configured with Docker workflow
- Configuration for different environments (local, staging, production)
- Repository secrets must be properly configured

## Architectural Principles

### Modularity
- Each module should have well-defined responsibilities
- Avoid circular dependencies between modules
- Use interfaces to decouple implementations

### Separation of Concerns
- Presentation layer (HTTP, gRPC, GraphQL)
- Business layer (domain, use cases)
- Persistence layer (repositories, migrations)
- Integration layer (external services)

### Testability
- All modules must be testable
- Use mocks for external dependencies
- Maintain adequate test coverage
- Tests should  be fast and reliable

This template is designed to grow with your project, from quick prototypes to enterprise-scale systems, always maintaining organization, scalability, and code quality. 