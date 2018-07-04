package types

import "k8s.io/apimachinery/pkg/api/resource"

type TypeMeta struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

type ObjectMeta struct {
	Name        string            `yaml:"name"`
	Labels      map[string]string `yaml:"labels"`
	Annotations map[string]string `yaml:"annotations"`
}

type LabelSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
}

type PodSpec struct {
	Volumes          []Volume               `yaml:"volumes"`
	InitContainers   []Container            `yaml:"initContainers"`
	Containers       []Container            `yaml:"containers"`
	ImagePullSecrets []LocalObjectReference `yaml:"imagePullSecrets"`
}

type PodTemplate struct {
	ObjectMeta `yaml:"metadata"`
	Spec       PodSpec `yaml:"spec"`
}

type DeploymentSpec struct {
	Replicas int32         `yaml:"replicas"`
	Selector LabelSelector `yaml:"selector"`
	Template PodTemplate   `yaml:"template"`
}

type Deployment struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata"`
	Spec       DeploymentSpec `yaml:"spec"`
}

type Container struct {
	Name            string               `yaml:"name"`
	Image           string               `yaml:"image"`
	Command         []string             `yaml:"command"`
	Args            []string             `yaml:"args"`
	WorkingDir      string               `yaml:"workingDir"`
	Env             []EnvVar             `yaml:"env"`
	EnvFrom         []EnvFromSource      `yaml:"envFrom"`
	ImagePullPolicy PullPolicy           `yaml:"imagePullPolicy"`
	Ports           []ContainerPort      `yaml:"ports"`
	Resources       ResourceRequirements `yaml:"resources"`
	VolumeMounts    []VolumeMount        `yaml:"volumeMounts"`
}

type EnvVar struct {
	Name      string
	Value     string
	ValueFrom *EnvVarSource
}

type ContainerPort struct {
	Name          string
	HostPort      int32    `yaml:"hostPort"`
	ContainerPort int32    `yaml:"containerPort"`
	Protocol      Protocol `yaml:"protocol"`
	HostIP        string   `yaml:"hostIP"`
}

type Protocol string

type Volume struct {
	Name         string `yaml:"name"`
	VolumeSource `yaml:",inline"`
}

type VolumeSource struct{}

type Service struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata"`
	Spec       ServiceSpec `yaml:"spec"`
}

type ServiceSpec struct {
	Type           ServiceType       `yaml:"type"`
	Ports          []ServicePort     `yaml:"ports"`
	Selector       map[string]string `yaml:"selector"`
	ClusterIP      string            `yaml:"clusterIP"`
	LoadBalancerIP string            `yaml:"loadBalancerIP"`
	ExternalIPs    []string          `yaml:"externalIPs"`
}

type ServicePort struct {
	Name       string   `yaml:"name"`
	Protocol   Protocol `yaml:"protocol"`
	Port       int32    `yaml:"port"`
	TargetPort int32    `yaml:"targetPort"`
	NodePort   int32    `yaml:"nodePort"`
}

type ServiceType string

// Ingress is a collection of rules that allow inbound connections to reach the
// endpoints defined by a backend.
type Ingress struct {
	TypeMeta   `yaml:",inline"`
	ObjectMeta `yaml:"metadata"`
	Spec       IngressSpec `yaml:"spec"`
}

// IngressSpec describes the Ingress the user wishes to exist.
type IngressSpec struct {
	Backend *IngressBackend `yaml:"backend,omitempty"`
	TLS     []IngressTLS    `yaml:"tls,omitempty"`
	Rules   []IngressRule   `yaml:"rules"`
}

// IngressTLS describes the transport layer security associated with an Ingress.
type IngressTLS struct {
	Hosts      []string `yaml:"hosts"`
	SecretName string   `yaml:"secretName,omitempty"`
}

// IngressRule represents the rules mapping the paths under a specified host to
// the related backend services.
type IngressRule struct {
	Host             string `yaml:"host"`
	IngressRuleValue `yaml:",inline"`
}

// IngressRuleValue represents a rule to apply against incoming requests.
type IngressRuleValue struct {
	HTTP *HTTPIngressRuleValue `yaml:"http"`
}

// HTTPIngressRuleValue is a list of http selectors pointing to backends.
type HTTPIngressRuleValue struct {
	Paths []HTTPIngressPath
}

// HTTPIngressPath associates a path regex with a backend. Incoming urls matching
// the path are forwarded to the backend.
type HTTPIngressPath struct {
	Path    string
	Backend IngressBackend
}

// IngressBackend describes all endpoints for a given service and port.
type IngressBackend struct {
	ServiceName string `yaml:"serviceName"`
	ServicePort int32  `yaml:"servicePort"`
}

// EnvVarSource represents a source for the value of an EnvVar.
// Only one of its fields may be set.
type EnvVarSource struct {
	// Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations,
	// metadata.uid, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.
	// +optional
	FieldRef *ObjectFieldSelector
	// Selects a resource of the container: only resources limits and requests
	// (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	// +optional
	ResourceFieldRef *ResourceFieldSelector
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *ConfigMapKeySelector
	// Selects a key of a secret in the pod's namespace.
	// +optional
	SecretKeyRef *SecretKeySelector
}

// ObjectFieldSelector selects an APIVersioned field of an object.
type ObjectFieldSelector struct {
	// Required: Version of the schema the FieldPath is written in terms of.
	// If no value is specified, it will be defaulted to the APIVersion of the
	// enclosing object.
	APIVersion string
	// Required: Path of the field to select in the specified API version
	FieldPath string
}

// ResourceFieldSelector represents container resources (cpu, memory) and their output format
type ResourceFieldSelector struct {
	// Container name: required for volumes, optional for env vars
	// +optional
	ContainerName string
	// Required: resource to select
	Resource string
	// Specifies the output format of the exposed resources, defaults to "1"
	// +optional
	Divisor resource.Quantity
}

// Selects a key from a ConfigMap.
type ConfigMapKeySelector struct {
	// The ConfigMap to select from.
	LocalObjectReference
	// The key to select.
	Key string
	// Specify whether the ConfigMap or it's key must be defined
	// +optional
	Optional *bool
}

// SecretKeySelector selects a key of a Secret.
type SecretKeySelector struct {
	// The name of the secret in the pod's namespace to select from.
	LocalObjectReference
	// The key of the secret to select from.  Must be a valid secret key.
	Key string
	// Specify whether the Secret or it's key must be defined
	// +optional
	Optional *bool
}

// LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type LocalObjectReference struct {
	//TODO: Add other useful fields.  apiVersion, kind, uid?
	Name string
}

// EnvFromSource represents the source of a set of ConfigMaps
type EnvFromSource struct {
	// An optional identifier to prepend to each key in the ConfigMap.
	// +optional
	Prefix string
	// The ConfigMap to select from.
	//+optional
	ConfigMapRef *ConfigMapEnvSource
	// The Secret to select from.
	//+optional
	SecretRef *SecretEnvSource
}

// ConfigMapEnvSource selects a ConfigMap to populate the environment
// variables with.
//
// The contents of the target ConfigMap's Data field will represent the
// key-value pairs as environment variables.
type ConfigMapEnvSource struct {
	// The ConfigMap to select from.
	LocalObjectReference
	// Specify whether the ConfigMap must be defined
	// +optional
	Optional *bool
}

// SecretEnvSource selects a Secret to populate the environment
// variables with.
//
// The contents of the target Secret's Data field will represent the
// key-value pairs as environment variables.
type SecretEnvSource struct {
	// The Secret to select from.
	LocalObjectReference
	// Specify whether the Secret must be defined
	// +optional
	Optional *bool
}

// PullPolicy describes a policy for if/when to pull a container image
type PullPolicy string

// VolumeMount describes a mounting of a Volume within a container.
type VolumeMount struct {
	// Required: This must match the Name of a Volume [above].
	Name string
	// Optional: Defaults to false (read-write).
	// +optional
	ReadOnly bool
	// Required. If the path is not an absolute path (e.g. some/path) it
	// will be prepended with the appropriate root prefix for the operating
	// system.  On Linux this is '/', on Windows this is 'C:\'.
	MountPath string
	// Path within the volume from which the container's volume should be mounted.
	// Defaults to "" (volume's root).
	// +optional
	SubPath string
	// mountPropagation determines how mounts are propagated from the host
	// to container and the other way around.
	// When not set, MountPropagationNone is used.
	// This field is beta in 1.10.
	// +optional
	MountPropagation *MountPropagationMode
}

// MountPropagationMode describes mount propagation.
type MountPropagationMode string

const (
	// MountPropagationNone means that the volume in a container will
	// not receive new mounts from the host or other containers, and filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode corresponds to "private" in Linux terminology.
	MountPropagationNone MountPropagationMode = "None"
	// MountPropagationHostToContainer means that the volume in a container will
	// receive new mounts from the host or other containers, but filesystems
	// mounted inside the container won't be propagated to the host or other
	// containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rslave" in Linux terminology).
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	// MountPropagationBidirectional means that the volume in a container will
	// receive new mounts from the host or other containers, and its own mounts
	// will be propagated from the container to the host or other containers.
	// Note that this mode is recursively applied to all mounts in the volume
	// ("rshared" in Linux terminology).
	MountPropagationBidirectional MountPropagationMode = "Bidirectional"
)

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed.
	// +optional
	Limits ResourceList
	// Requests describes the minimum amount of compute resources required.
	// If Request is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value
	// +optional
	Requests ResourceList
}

// ResourceList is a set of (resource name, quantity) pairs.
type ResourceList map[ResourceName]resource.Quantity

// ResourceName is the name identifying various resources in a ResourceList.
type ResourceName string
