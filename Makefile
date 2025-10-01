# Example 

# CLUSTER_NAME ?= hatch-cluster

.PHONY: setup-claude-code
setup-claude-code:
	bash ./scripts/setup-claude-code.sh

# .PHONY: setup-cluster
# setup-cluster:
# 	@echo "Setting up cluster..."
# 	@if kind get clusters | grep -q "^${CLUSTER_NAME}$$"; then \
# 		echo "Cluster '${CLUSTER_NAME}' already exists. Skipping..."; \
# 	else \
# 		kind create cluster --name ${CLUSTER_NAME} --config ./infra/k8s/cluster/kind.yaml; \
# 	fi

# .PHONY: teardown-cluster
# teardown-cluster:
# 	@echo "Tearing down cluster..."
# 	@if kind get clusters | grep -q "^${CLUSTER_NAME}$$"; then \
# 		kind delete cluster --name ${CLUSTER_NAME}; \
# 	else \
# 		echo " Cluster '${CLUSTER_NAME}' does not exists"; \
# 	fi

# .PHONY: up
# up: setup-cluster

# .PHONY: down
# down: teardown-cluster

# .PHONY: restart
# restart: down up