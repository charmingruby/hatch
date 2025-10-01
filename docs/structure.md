# Structure

## Infrastructure

If your project requires infrastructure management, create an `infra/` directory at the root.  
Each type of resource should have its own subdirectory, for example: `helm`, `terraform`, `k8s`.  

## Monorepo & Multiple Apps

Hatch supports multiple components or services under a single repository.  

- Place all apps under the root-level `apps/` directory.  
- Each app is self-contained, with its own code, tests, and configurations.  
- **Go-only projects** can leverage a **Go workspace** (`go.work`) to manage dependencies across apps.  

**Example:**

```
└── apps
    ├── api     
    ├── web     
    └── worker  # Background jobs or consumers
```

## Orchestration & Scripts

All project orchestration and environment lifecycle scripts are centralized in the `Makefile` at the root of the repository.  
This includes tasks such as:  

- Provisioning and tearing down local clusters
- Running setup, cleanup, restart, and log commands for development environments  

Centralizing these scripts in a single `Makefile` provides a consistent, documented entry point for developers to manage the entire environment.