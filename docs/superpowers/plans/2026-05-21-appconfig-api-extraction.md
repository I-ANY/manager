# AppConfig API Extraction Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Move AppConfig API types into the main module so `k8soperation` no longer imports `gitee.com/jay-kim/appconfig-operator/api/v1alpha1` while keeping AppConfig CR behavior unchanged.

**Architecture:** Create a focused in-repo API package at `pkg/apis/appconfig/v1alpha1` containing the CRD Go types, scheme registration, and deepcopy methods. Update the backend AppConfig client/DAO/service/CRD helpers to import this local package, then remove the external operator module from the main module graph. Do not change HTTP routes or runtime behavior in this pass.

**Tech Stack:** Go 1.25, Kubernetes API machinery, controller-runtime client, Gin backend, client-go.

---

## File structure

- Create `pkg/apis/appconfig/v1alpha1/groupversion_info.go`: local AppConfig API group/version and `AddToScheme` registration.
- Create `pkg/apis/appconfig/v1alpha1/appconfig_types.go`: local `AppConfig`, `AppConfigList`, `AppConfigSpec`, `AppConfigStatus`, and `EnvVarKV` types copied from the current operator API schema.
- Create `pkg/apis/appconfig/v1alpha1/zz_generated.deepcopy.go`: deepcopy implementations required for Kubernetes runtime objects.
- Modify `initialize/appconfig_client.go`: replace operator API import with local API import.
- Modify `pkg/cluster/client.go`: replace operator API import with local API import.
- Modify `internal/app/dao/kube_appconfig.go`: replace operator API import with local API import.
- Modify `internal/app/services/kube_appconfig.go`: replace operator API import with local API import.
- Modify `pkg/k8s/crd/appconfig_create.go`: replace operator API import with local API import.
- Modify `pkg/k8s/crd/appconfig_get.go`: replace operator API import with local API import.
- Modify `pkg/k8s/crd/appconfig_list.go`: replace operator API import with local API import.
- Modify `pkg/k8s/crd/appconfig_update.go`: replace operator API import with local API import.
- Modify `go.mod` / `go.sum`: remove `gitee.com/jay-kim/appconfig-operator` from the main module dependency graph via `go mod tidy`.
- Do not modify `operators/appconfig-operator` in this implementation. It remains a separately buildable operator submodule for now.

## Task 1: Add local AppConfig API package

**Files:**
- Create: `pkg/apis/appconfig/v1alpha1/groupversion_info.go`
- Create: `pkg/apis/appconfig/v1alpha1/appconfig_types.go`
- Create: `pkg/apis/appconfig/v1alpha1/zz_generated.deepcopy.go`

- [ ] **Step 1: Create the API package directory**

Run:
```bash
mkdir -p pkg/apis/appconfig/v1alpha1
```
Expected: command exits 0.

- [ ] **Step 2: Write `groupversion_info.go`**

Create `pkg/apis/appconfig/v1alpha1/groupversion_info.go` with:
```go
// Package v1alpha1 contains API Schema definitions for the operation v1alpha1 API group.
// +kubebuilder:object:generate=true
// +groupName=operation.operation.top
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	GroupVersion = schema.GroupVersion{Group: "operation.operation.top", Version: "v1alpha1"}
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}
	AddToScheme   = SchemeBuilder.AddToScheme
)
```

- [ ] **Step 3: Write `appconfig_types.go`**

Create `pkg/apis/appconfig/v1alpha1/appconfig_types.go` with:
```go
package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EnvVarKV struct {
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// +kubebuilder:validation:MinLength=1
	Value string `json:"value"`
}

type AppConfigSpec struct {
	// +kubebuilder:validation:MinLength=1
	AppName string `json:"appName"`

	// +kubebuilder:validation:MinLength=1
	Image string `json:"image"`

	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// +optional
	Env []EnvVarKV `json:"env,omitempty"`

	// +optional
	EnableMetrics bool `json:"enableMetrics,omitempty"`

	// +kubebuilder:validation:Enum=RollingUpdate;Recreate
	// +optional
	Strategy string `json:"strategy,omitempty"`
}

type AppConfigStatus struct {
	// +optional
	Phase string `json:"phase,omitempty"`

	// +optional
	Message string `json:"message,omitempty"`

	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// +listType=map
	// +listMapKey=types
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type AppConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec AppConfigSpec `json:"spec"`

	// +optional
	Status AppConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type AppConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppConfig{}, &AppConfigList{})
}
```

- [ ] **Step 4: Write `zz_generated.deepcopy.go`**

Create `pkg/apis/appconfig/v1alpha1/zz_generated.deepcopy.go` with:
```go
//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AppConfig) DeepCopyInto(out *AppConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *AppConfig) DeepCopy() *AppConfig {
	if in == nil {
		return nil
	}
	out := new(AppConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *AppConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AppConfigList) DeepCopyInto(out *AppConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *AppConfigList) DeepCopy() *AppConfigList {
	if in == nil {
		return nil
	}
	out := new(AppConfigList)
	in.DeepCopyInto(out)
	return out
}

func (in *AppConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AppConfigSpec) DeepCopyInto(out *AppConfigSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVarKV, len(*in))
		copy(*out, *in)
	}
}

func (in *AppConfigSpec) DeepCopy() *AppConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AppConfigSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *AppConfigStatus) DeepCopyInto(out *AppConfigStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *AppConfigStatus) DeepCopy() *AppConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AppConfigStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *EnvVarKV) DeepCopyInto(out *EnvVarKV) {
	*out = *in
}

func (in *EnvVarKV) DeepCopy() *EnvVarKV {
	if in == nil {
		return nil
	}
	out := new(EnvVarKV)
	in.DeepCopyInto(out)
	return out
}
```

- [ ] **Step 5: Run gofmt on the new package**

Run:
```bash
gofmt -w pkg/apis/appconfig/v1alpha1/groupversion_info.go pkg/apis/appconfig/v1alpha1/appconfig_types.go pkg/apis/appconfig/v1alpha1/zz_generated.deepcopy.go
```
Expected: command exits 0.

- [ ] **Step 6: Verify the package compiles**

Run:
```bash
go test ./pkg/apis/appconfig/v1alpha1
```
Expected: `ok` or `[no test files]`, exit 0.

- [ ] **Step 7: Commit**

Run:
```bash
git add pkg/apis/appconfig/v1alpha1
git commit -m "feat: add local AppConfig API types"
```
Expected: commit succeeds.

## Task 2: Replace backend imports with local AppConfig API

**Files:**
- Modify: `initialize/appconfig_client.go`
- Modify: `pkg/cluster/client.go`
- Modify: `internal/app/dao/kube_appconfig.go`
- Modify: `internal/app/services/kube_appconfig.go`
- Modify: `pkg/k8s/crd/appconfig_create.go`
- Modify: `pkg/k8s/crd/appconfig_get.go`
- Modify: `pkg/k8s/crd/appconfig_list.go`
- Modify: `pkg/k8s/crd/appconfig_update.go`

- [ ] **Step 1: Replace the import path everywhere in the main module**

Run:
```bash
python3 - <<'PY'
from pathlib import Path
old = 'gitee.com/jay-kim/appconfig-operator/api/v1alpha1'
new = 'k8soperation/pkg/apis/appconfig/v1alpha1'
for path in [
    Path('initialize/appconfig_client.go'),
    Path('pkg/cluster/client.go'),
    Path('internal/app/dao/kube_appconfig.go'),
    Path('internal/app/services/kube_appconfig.go'),
    Path('pkg/k8s/crd/appconfig_create.go'),
    Path('pkg/k8s/crd/appconfig_get.go'),
    Path('pkg/k8s/crd/appconfig_list.go'),
    Path('pkg/k8s/crd/appconfig_update.go'),
]:
    text = path.read_text()
    if old not in text:
        raise SystemExit(f'{path}: old import not found')
    path.write_text(text.replace(old, new))
PY
```
Expected: command exits 0.

- [ ] **Step 2: Run gofmt on modified files**

Run:
```bash
gofmt -w initialize/appconfig_client.go pkg/cluster/client.go internal/app/dao/kube_appconfig.go internal/app/services/kube_appconfig.go pkg/k8s/crd/appconfig_create.go pkg/k8s/crd/appconfig_get.go pkg/k8s/crd/appconfig_list.go pkg/k8s/crd/appconfig_update.go
```
Expected: command exits 0.

- [ ] **Step 3: Verify the external operator API import is gone from backend code**

Run:
```bash
grep -R "gitee.com/jay-kim/appconfig-operator/api/v1alpha1" -n --exclude-dir=.git --exclude-dir=operators . || true
```
Expected: no output.

- [ ] **Step 4: Build-check affected packages**

Run:
```bash
go test ./initialize ./pkg/cluster ./internal/app/dao ./internal/app/services ./pkg/k8s/crd
```
Expected: all listed packages pass or report `[no test files]`, exit 0.

- [ ] **Step 5: Commit**

Run:
```bash
git add initialize/appconfig_client.go pkg/cluster/client.go internal/app/dao/kube_appconfig.go internal/app/services/kube_appconfig.go pkg/k8s/crd/appconfig_create.go pkg/k8s/crd/appconfig_get.go pkg/k8s/crd/appconfig_list.go pkg/k8s/crd/appconfig_update.go
git commit -m "refactor: use local AppConfig API package"
```
Expected: commit succeeds.

## Task 3: Remove external operator module dependency from main module

**Files:**
- Modify: `go.mod`
- Modify: `go.sum`

- [ ] **Step 1: Tidy main module dependencies**

Run:
```bash
go mod tidy
```
Expected: command exits 0.

- [ ] **Step 2: Verify main module no longer references the operator module**

Run:
```bash
grep -n "gitee.com/jay-kim/appconfig-operator" go.mod go.sum || true
```
Expected: no output.

- [ ] **Step 3: Verify operator submodule remains untouched**

Run:
```bash
git diff -- operators/appconfig-operator
```
Expected: no output.

- [ ] **Step 4: Commit**

Run:
```bash
git add go.mod go.sum
git commit -m "chore: drop backend operator module dependency"
```
Expected: commit succeeds.

## Task 4: Full verification

**Files:**
- No code changes expected.

- [ ] **Step 1: Run backend tests**

Run:
```bash
go test ./...
```
Expected: exit 0. If this fails due to a pre-existing compile issue unrelated to AppConfig imports, capture the exact package and error before changing code.

- [ ] **Step 2: Run backend vet through the repository target**

Run:
```bash
make lint
```
Expected: exit 0.

- [ ] **Step 3: Build the backend**

Run:
```bash
make build
```
Expected: exit 0 and a binary at `./bin/k8soperation` or the configured `BIN_FILE`. This target regenerates Swagger docs; include resulting `docs/docs.go`, `docs/swagger.json`, or `docs/swagger.yaml` changes only if they are caused by the build and are expected.

- [ ] **Step 4: Verify no backend import dependency remains**

Run:
```bash
grep -R "gitee.com/jay-kim/appconfig-operator/api/v1alpha1" -n --exclude-dir=.git --exclude-dir=operators . || true
```
Expected: no output.

- [ ] **Step 5: Inspect final diff**

Run:
```bash
git diff --stat HEAD~3..HEAD
```
Expected: shows the new local API package, backend import replacements, and dependency cleanup.

- [ ] **Step 6: Do not commit verification-only changes unless generated docs changed**

If `make build` changed generated Swagger docs, inspect them:
```bash
git diff -- docs/docs.go docs/swagger.json docs/swagger.yaml
```
If the changes are unrelated to this refactor, revert only the generated Swagger files before final status. If they are expected, commit them:
```bash
git add docs/docs.go docs/swagger.json docs/swagger.yaml
git commit -m "chore: refresh generated Swagger docs"
```

## Self-review notes

- Spec coverage: the plan adds a local API package, redirects all backend imports, removes the main-module dependency, and verifies behavior without changing runtime AppConfig flow.
- Placeholder scan: no TBD/TODO/fill-in steps remain.
- Type consistency: all AppConfig references use `v1alpha1.AppConfig`, `v1alpha1.AppConfigList`, `v1alpha1.AppConfigSpec`, and `v1alpha1.EnvVarKV` from `k8soperation/pkg/apis/appconfig/v1alpha1` after Task 2.
