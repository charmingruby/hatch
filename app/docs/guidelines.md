# Guidelines

## Project Identity

* Replace all instances of the placeholder module path with your own:

  ```bash
  PACK_APP
  ```

* GitHub Actions includes a workflow to push container images to DockerHub.
  You must:
  * Change the image name.
  * Configure repository secrets.
    
* Replace docker compose api image.

## Modules

* The `internal/` directory contains a sample module structure.
* Use it as a base to organize your own modules.
* At each module root, include a barrel file (`module_name.go`) for initialization.

## Testing

* Test setup is integrated, including support for **mockery**.
* Refer to the Makefile for available test commands.

**Test Scenario Convention**

```
<action> <result> when <condition>
```

**Examples**

```
should return stash with default values when no options are provided
should return ErrExpired when key is expired
should store new entry successfully
```

## Persistence Layer

* PostgreSQL is configured as the default database.
* Includes:

  * SQL migrations,
  * Database setup scripts,
  * Example repository pattern.

## Communication Layer

### HTTP

Synchronous communication:

* `http/rest/` — REST APIs
* `http/grpc/` — gRPC services
* `http/gql/` — GraphQL APIs

### Event

Asynchronous, event-driven communication.

**Within a Module**

* `internal/MODULE/messaging/event/` — Defines module-specific events (contracts + serialization).
* `internal/MODULE/messaging/subscriber/` — Event listeners.

  * Convention: event `transaction.created` → handler `onTransactionCreated`.

**Shared Contracts**

* If messaging is **cross-cutting**, keep contracts in:

  * `pkg/messaging/` — if reusable outside this repo.
  * `internal/shared/messaging/` — if private to this repo.
* Implementations (e.g., NATS, Kafka) live under the same directory.

**Example:**

```
pkg/
  messaging/
    messaging.go     # Publisher, Subscriber, Event
    nats/
      broker.go
    kafka/
      broker.go
```

This ensures modules depend only on **contracts**, not directly on third-party SDKs.

### DTOs (Data Transfer Objects)

* For use case inputs/outputs, centralize contracts using DTOs.
* This avoids having to change parameters in multiple layers (handler, use case, service).
* Where to place:

internal/MODULE/dto/ — if the contract is specific to a single module.

## External & Shared Guidelines

### External

* Inside each module, `external/` defines contracts unique to that module and their implementations.

**Example**

```
internal/
  billing/
    external/
      payment_gateway.go    # PaymentGateway contract
      stripe/
        gateway.go
      paypal/
        gateway.go
```

### Shared

* For contracts reused across modules but not meant for public API, place them in `internal/shared/`.

**Example**

```
internal/
  shared/
    storage/
      storage.go            # Uploader interface
  user/
    external/
      avatar_storage.go     # depends on shared/storage.Uploader
```

### Public (`pkg/`)

* For contracts or libs reusable across **projects**, use `pkg/`.

**Example**

```
pkg/
  messaging/
    messaging.go
  otel/
    trace.go
```
