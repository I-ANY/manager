# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Main backend module

- `make build` - generate Swagger docs and build `./bin/k8soperation` from `./cmd/k8soperation`.
- `make run-local` - generate Swagger docs and run the API locally with `APP_CONFIG=$(pwd)/configs/config.yaml` and `GIN_MODE=debug`.
- `make run` - build, then run the compiled binary with `APP_CONFIG=$(pwd)/configs/config.yaml`.
- `make test` - run `go test ./... -v` for the main module.
- `go test ./path/to/package -run TestName -v` - run a single Go test or test package.
- `make fmt` - run `go fmt ./...`.
- `make lint` - run `go vet ./...`.
- `make swag` - generate Swagger artifacts in `docs/` using `swag init -g cmd/k8soperation/main.go -o docs -d ./ --parseInternal`; installs `swag` if missing.
- `make swagger-ui` / `make swagger-ui-stop` - run or stop a local Swagger UI container on port 8081.
- `make docker-build` / `make docker-run` - build and run the application container. Override with `DOCKER=nerdctl` when needed.
- Runtime config defaults to `configs/config.yaml`; copy from `configs/config.yaml.example` when setting up a local environment.

### AppConfig operator submodule

The Kubebuilder operator under `operators/appconfig-operator/` is a separate Go module (`gitee.com/jay-kim/appconfig-operator`, Go 1.24.6):

- `cd operators/appconfig-operator && make build` - generate manifests/code, fmt/vet, and build `bin/manager`.
- `cd operators/appconfig-operator && make test` - run envtest-backed unit tests excluding e2e tests.
- `cd operators/appconfig-operator && go test ./internal/controller -run TestName -v` - run a single operator test.
- `cd operators/appconfig-operator && make run` - run the controller locally against the current kubeconfig.
- `cd operators/appconfig-operator && make manifests generate` - regenerate CRD/RBAC manifests and deepcopy code.
- `cd operators/appconfig-operator && make install` - install CRDs into the current cluster.
- `cd operators/appconfig-operator && make deploy IMG=<image>` - deploy the controller with kustomize.
- `cd operators/appconfig-operator && make test-e2e` - run e2e tests using Kind.

## Architecture

This repository contains a Gin-based Kubernetes management API plus an embedded/related AppConfig Kubebuilder operator.

### Main backend flow

- `cmd/main.go` defines a Cobra CLI. The default command starts the HTTP API; the `migrate` subcommand is registered from `cmd/migrate`.
- `internal/bootstrap.InitAll` creates the central `*app.App`, then initializes settings, validation, zap loggers, database/session state, Kubernetes clients, the AppConfig controller-runtime client, and Swagger readiness logging.
- `internal/server` wraps the Gin engine in `http.Server` and handles graceful shutdown, including DB close on shutdown.
- `initialize.NewEngine` builds the Gin engine, installs request metadata/logging/recovery/Kubernetes-error/session middleware, injects `*app.App` into each Gin context, registers health checks, and mounts routes.

### Dependency container and request handling

- `pkg/app.App` is the runtime dependency container: parsed settings, GORM/sql DB handles, system and business loggers, session store, default Kubernetes client/config, metrics client, event API capability flag, and AppConfig client.
- `App.Middleware()` stores `*app.App` in every Gin context; controllers retrieve it with `app.FromContext(ctx)` and construct `services.NewServices(ctx, a)`.
- Handlers generally follow `router -> controller -> request validation -> service -> pkg/k8s or DAO -> response`.
- Request DTOs and validation functions live in `internal/app/requests`; standardized responses are in `pkg/app/response`; database models are in `internal/app/models`; DAO methods are in `internal/app/dao`.
- Controllers live under `internal/app/controllers/api/...`; route registration lives separately under `internal/app/routers/...` and is assembled in `initialize/router.go`.

### Kubernetes API organization

- Kubernetes resource APIs are mounted under `/api/v1/k8s`, with subgroups for pod, deployment, statefulset, daemonset, job, cronjob, service, ingress, secret, configmap, storageclass, pv, pvc, node, and namespace.
- `initialize/router.go` wires the cluster router plus each resource-specific router. Add new HTTP endpoints by updating the relevant router and controller package, then regenerate Swagger when annotations change.
- `internal/app/services` contains orchestration logic and DB-backed operations, while `pkg/k8s/<resource>` contains lower-level client-go operations and resource-specific builders/listing/patching helpers.
- `pkg/k8s/dataselect` provides list filtering/sorting/pagination through `cell` adapters; list endpoints commonly fetch raw Kubernetes objects, wrap them in selectors, then return the paginated native objects.
- Kubernetes client initialization prefers DB-stored kubeconfig for `App.DefaultClusterID`, then `App.GlobalKubeConfigPath`, then in-cluster config. REST configs are tuned with QPS/Burst/Timeout/UserAgent, metrics-server is optional, and event listing detects `events.k8s.io/v1` with fallback to core/v1.

### Auth, errors, and observability

- Public routes include hello/auth/logout/registration and debug routes outside release mode. Protected user routes use JWT middleware. Most Kubernetes resource routes are mounted under `/api/v1/k8s` in `initialize/router.go`; verify auth expectations there before adding sensitive routes.
- Middleware in `middlewares/` handles JWT auth, structured request logging, panic recovery, and final Kubernetes error conversion for errors attached to `gin.Context`.
- Error definitions are centralized in `internal/errorcode`; many Kubernetes handler errors are passed through `ctx.Error(err)` for middleware handling.
- The app has separate system and business zap loggers; use `App.BusinessLog` when emitting business audit events.

### AppConfig operator relationship

- `operators/appconfig-operator/` is a standalone Kubebuilder project for the `operation.top/v1alpha1` `AppConfig` CRD. It reconciles AppConfig resources into Deployment/Service resources and maintains status conditions.
- The main backend imports the operator API type and initializes a controller-runtime client with only the AppConfig scheme in `pkg/cluster/client.go`, allowing API code to manage AppConfig CRs without running inside the operator module.
- `pkg/k8s/crd` and `internal/app/services/kube_appconfig.go` are the main backend-side bridge for AppConfig CRUD/list/update behavior.

## Notes from repository docs

- README documents completed Kubernetes management capabilities: Deployment CRUD/scale/image update/restart/rollback, Pod logs/events/deletion, StatefulSet/DaemonSet lifecycle operations, Service/Ingress patching/TLS/events, Job/CronJob management, Secret/PVC/PV/ConfigMap/StorageClass lifecycle operations, Node cordon/drain/evict/metrics, event aggregation, and multi-cluster kubeconfig management.
- Swagger endpoints are documented as `/swagger` and `/swagger-standalone`, while the current Gin route mounts Swagger under `/swagger/*any`; verify the exact URL in the running app when changing docs or routing.
