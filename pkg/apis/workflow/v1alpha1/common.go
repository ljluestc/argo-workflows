package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ResourceScope string

const (
	ResourceScopeLocal      ResourceScope = "local"
	ResourceScopeNamespaced ResourceScope = "namespaced"
	ResourceScopeCluster    ResourceScope = "cluster"
)

// TemplateHolder is an object that holds templates; e.g. Workflow, WorkflowTemplate, and ClusterWorkflowTemplate
type TemplateHolder interface {
	GetNamespace() string
	GetName() string
	GroupVersionKind() schema.GroupVersionKind
	GetTemplateByName(name string) *Template
	GetResourceScope() ResourceScope
	GetPodMetadata() *Metadata
}

// WorkflowSpecHolder is an object that holds a WorkflowSpec; e.g., WorkflowTemplate, and ClusterWorkflowTemplate
type WorkflowSpecHolder interface {
	metav1.Object
	GetWorkflowSpec() *WorkflowSpec
}

// TemplateReferenceHolder is an object that holds a reference to other templates; e.g. WorkflowStep, DAGTask, and NodeStatus
type TemplateReferenceHolder interface {
	// GetTemplate returns the template. This maybe nil. This is first precedence.
	GetTemplate() *Template
	// GetTemplateRef returns the template ref. This maybe nil. This is second precedence.
	GetTemplateRef() *TemplateRef
	// GetTemplateName returns the template name. This maybe empty. This is last precedence.
	GetTemplateName() string
	// GetName returns the name of the template reference holder.
	GetName() string
	// IsDAGTask returns true if the template reference is a DAGTask.
	IsDAGTask() bool
	// IsWorkflowStep returns true if the template reference is a WorkflowStep.
	IsWorkflowStep() bool
}

// Task is the behaviour shared by WorkflowStep and DAGTask, so callers can handle either without duplicating logic.
// The two remain distinct API types to preserve wire compatibility; see https://github.com/argoproj/argo-workflows/issues/1041.
type Task interface {
	TemplateReferenceHolder
	// GetArguments returns the arguments passed to the referenced template.
	GetArguments() Arguments
	// GetWhen returns the conditional expression guarding execution. This may be empty.
	GetWhen() string
	// GetWithItems returns the static list of items to expand over. This may be nil.
	GetWithItems() []Item
	// GetWithParam returns the JSON list of items to expand over. This may be empty.
	GetWithParam() string
	// GetWithSequence returns the numeric sequence to expand over. This may be nil.
	GetWithSequence() *Sequence
	// ShouldExpand returns true if the task fans out over withItems, withParam, or withSequence.
	ShouldExpand() bool
	// ContinuesOn returns whether execution should proceed if the task finishes in the given phase.
	ContinuesOn(phase NodePhase) bool
	// HasExitHook returns true if the task has an exit hook, via hooks or the deprecated onExit.
	HasExitHook() bool
	// GetExitHook returns the exit hook with args applied, or nil if there is none.
	GetExitHook(args Arguments) *LifecycleHook
}

// SubmitOpts are workflow submission options
type SubmitOpts struct {
	// Name overrides metadata.name
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// GenerateName overrides metadata.generateName
	GenerateName string `json:"generateName,omitempty" protobuf:"bytes,2,opt,name=generateName"`
	// Entrypoint overrides spec.entrypoint
	Entrypoint string `json:"entryPoint,omitempty" protobuf:"bytes,4,opt,name=entrypoint"`
	// Parameters passes input parameters to workflow
	Parameters []string `json:"parameters,omitempty" protobuf:"bytes,5,rep,name=parameters"`
	// ServiceAccount runs all pods in the workflow using specified ServiceAccount.
	ServiceAccount string `json:"serviceAccount,omitempty" protobuf:"bytes,7,opt,name=serviceAccount"`
	// DryRun validates the workflow on the client-side without creating it. This option is not supported in API
	DryRun bool `json:"dryRun,omitempty" protobuf:"varint,8,opt,name=dryRun"`
	// ServerDryRun validates the workflow on the server-side without creating it
	ServerDryRun bool `json:"serverDryRun,omitempty" protobuf:"varint,9,opt,name=serverDryRun"`
	// Labels adds to metadata.labels
	Labels string `json:"labels,omitempty" protobuf:"bytes,10,opt,name=labels"`
	// OwnerReference creates a metadata.ownerReference
	OwnerReference *metav1.OwnerReference `json:"ownerReference,omitempty" protobuf:"bytes,11,opt,name=ownerReference"`
	// Annotations adds to metadata.labels
	Annotations string `json:"annotations,omitempty" protobuf:"bytes,12,opt,name=annotations"`
	// Set the podPriorityClassName of the workflow
	PodPriorityClassName string `json:"podPriorityClassName,omitempty" protobuf:"bytes,13,opt,name=podPriorityClassName"`
	// Priority is used if controller is configured to process limited number of workflows in parallel, higher priority workflows
	// are processed first.
	Priority *int32 `json:"priority,omitempty" protobuf:"bytes,14,opt,name=priority"`
	// Artifacts overrides input artifact locations. Format: name=key
	Artifacts []string `json:"artifacts,omitempty" protobuf:"bytes,15,rep,name=artifacts"`
}
