CLUSTER_NAME ?= pack-cluster

.PHONY: setup-cluster
setup-cluster:
	@echo "Setting up cluster..."
	@if kind get clusters | grep -q "^${CLUSTER_NAME}$$"; then \
		echo "Cluster '${CLUSTER_NAME}' already exists. Skipping..."; \
	else \
		kind create cluster --name ${CLUSTER_NAME} --config ./.bootstrap/cluster/manifests/kind.yaml; \
	fi

.PHONY: teardown-cluster
teardown-cluster:
	@echo "Tearing down cluster..."
	@if kind get clusters | grep -q "^${CLUSTER_NAME}$$"; then \
		kind delete cluster --name ${CLUSTER_NAME}; \
	else \
		echo " Cluster '${CLUSTER_NAME}' does not exists"; \
	fi

.PHONY: bootstrap
bootstrap: setup-cluster

.PHONY: cleanup
cleanup: teardown-cluster