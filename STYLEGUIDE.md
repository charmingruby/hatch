# Styleguide

## Premises

- This project follows a modular architecture with clear separation of responsibilities and use of bounded contexts.
- Each module represents a feature, owning its own context, abstractions, and technologies. Modules must not call each other directly.
- Code should be as simple as possible for its context, avoiding unnecessary abstractions.
- Documentation is critical to ensure clarity for both humans and LLMs.
- Testing is essential. Components must expose interfaces to enable mocking, which is the preferred testing strategy.

---

## Conventions

- Barrel files (e.g., broker.go) serve specific purposes:
  - Provide package-level documentation.
  - Define interfaces or contracts to be implemented.
  - Contain the main constructor or factory function.
- File names, directories, and packages must always use the singular form.
- For packages with multiple implementations, each implementation should live in its own subpackage named after the tool or vendor:

  ```bash
  broker/
  ├── broker.go
  └── kafka/
      └── kafka.go
  ```

---

## Anatomy

### Base Directories

- `cmd/`: Contains executable entry points.

  ```bash
  cmd/
  ├── api
  ├── lambda
  └── worker
  ```

- `config/`: Loads and validates environment variables.
- `internal/`: Contains application logic, organized by bounded contexts.
- `pkg/`: Shared libraries used across modules.
- `test/`: Test utilities such as mocks, fixtures, and test containers.
- `infra/`: Infrastructure files (e.g., IaC, Docker, Kubernetes).

  ```bash
  infra/
  ├── iac
  ├── k8s
  └── docker
      └── kong
          └── config
              └── kong.yaml
  ```

- `db/`: (For SQL persistence only) Contains migration and seed SQL files.

### Modules

- Modules are located under internal/{module}.

> **Note:** If your project consists of a single domain or module, you don't need to create a dedicated subdirectory under `internal/`. In this case:
> - Place the code directly under `internal/`
> - Store mocks directly under `test/gen/`
> - You may omit a module-level barrel file (e.g., `device.go`)
>
> This simplifies the structure while keeping it consistent with the modular design principles.

- For example, the device module is responsible for all device-related logic and is structured as follows:
  - `delivery/` – Entry points to the module (e.g., REST and MQTT):
    - `rest/`: Exposes HTTP endpoints.
    - `broker/`: Handles MQTT publish/subscribe logic.

    ```bash
    broker/
    ├── broker.go
    └── mqtt/
        ├── mqtt.go
        └── publisher.go
    ```

  - `integration/` – Connectors for external systems:
    - `client/`: Consumes external APIs.
    - `provider/`: Exports logic for other modules.
    - `repository/` – Data access layer:
    - repository.go: Interfaces.
      - `postgres/`: PostgreSQL implementation.
        - device_repository.go
  - `service/` – Application service layer (use cases).
  - `model/` – Domain entities and value objects.
  - `device.go` – Barrel file that exposes the module's public API for internal use.

---

## Development

### Testing

- Mocking is the chosen testing strategy. We use mockery to generate mocks from interfaces.
- Whenever you want to test a component, you must expose an interface for it. This allows for mocking and decouples the test from concrete implementations.
- All mocks should be stored under test/gen/{related_module}.
- Follow the standard Go convention for test files: *_test.go.
- For big components, create a setup_test.go file to bootstrap common test setup logic.
- Example Workflow
  1.  Define an interface in a barrel file (e.g., repository.go, broker.go).
  2. Write your tests.
  3. Run the following command (mocks will be generated automatically):

    ```bash
    make test

### Integrations

- Integrations with third-party APIs or services must be implemented inside the `integration/` directory of each module.
- Each integration is responsible for communicating with external systems while keeping domain logic isolated in the module's core.
- All external calls must be abstracted behind interfaces, which should be exposed at the appropriate level to support mocking during tests.
- Never embed third-party logic directly into services. Instead, encapsulate all SDK or I/O-related logic inside the corresponding `integration/` component.

  ```bash
  integration/
  ├── client/       # Consumers of external APIs or SDKs
  │   └── example/
  │       ├── example.go
  │       └── anyimpl/
  ├── provider/     # Components that export data to external modules
  │   └── quota/
  │       └── quota.go  # Implements another module's client interface
  ```
