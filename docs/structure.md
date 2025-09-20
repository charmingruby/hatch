# Structure

## Bootstrap

Pack includes a `.bootstrap` directory with everything needed to set up the development environment.  
The project Makefile provides two main commands:  

- **bootstrap**: prepares and configures the environment.  
- **cleanup**: removes all setup and resets the environment.  

## Infrastructure

If your project requires infrastructure management, create an `infra/` directory at the root.  
Each type of resource should have its own subdirectory, for example: `charts`, `terraform`, `manifests`.  

## Monorepo & Multiple Apps

Pack supports multiple components or services under a single repository.  

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