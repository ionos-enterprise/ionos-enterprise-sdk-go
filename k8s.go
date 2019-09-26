package profitbricks

import "net/http"

type KubernetesClusters struct {
	// URL to the collection representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// Unique representation for Kubernetes Cluster as a collection on a resource.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Array of items in that collection
	// Read Only: true
	Items []*KubernetesCluster `json:"items"`

	// The type of resource within a collection
	// Read Only: true
	// Enum: [k8s]
	PBType string `json:"type,omitempty"`
}

// KubernetesCluster kubernetes cluster
type KubernetesCluster struct {
	// URL to the object representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// The resource's unique identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// metadata
	Metadata *NoStateMetaData `json:"metadata,omitempty"`

	// properties
	// Required: true
	Properties *KubernetesClusterProperties `json:"properties"`

	// The type of object
	// Read Only: true
	// Enum: [k8s]
	PBType string `json:"type,omitempty"`

	// Entities of a cluster
	Entities KubernetesClusterEntities `json:"entities"`
}

type KubernetesClusterEntities struct {
	// NodePools of a cluster
	NodePools *KubernetesNodePools `json:"nodepools,omitempty"`
}

// KubernetesClusterProperties kubernetes cluster properties
type KubernetesClusterProperties struct {
	// A Kubernetes Cluster Name. Valid Kubernetes Cluster name must be 63 characters or less and must be empty
	// or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.),
	// and alphanumerics between.
	// Required: true
	Name string `json:"name"`
}

// KubernetesConfig kubernetes config
type KubernetesConfig struct {
	// URL to the object representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// The resource's unique identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// properties
	// Required: true
	Properties *KubernetesConfigProperties `json:"properties"`

	// The type of object
	// Read Only: true
	// Enum: [kubeconfig]
	PBType string `json:"type,omitempty"`
}

type KubernetesConfigProperties struct {
	// A Kubernetes Config file data
	KubeConfig string `json:"kubeconfig,omitempty"`
}

type KubernetesNodePool struct {
	// URL to the object representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// The resource's unique identifier.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// metadata
	Metadata *Metadata `json:"metadata,omitempty"`

	// properties
	// Required: true
	Properties *KubernetesNodePoolProperties `json:"properties"`

	// The type of object
	// Read Only: true
	// Enum: [nodepool]
	PBType string `json:"type,omitempty"`
}

// KubernetesNodePoolProperties kubernetes node pool properties
// swagger:model KubernetesNodePoolProperties
type KubernetesNodePoolProperties struct {
	// The availability zone in which the server should exist
	// Required: true
	// Enum: [AUTO ZONE_1 ZONE_2]
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// Number of cores for node
	// Required: true
	CoresCount int64 `json:"coresCount,omitempty"`

	// A valid cpu family name
	// Required: true
	CPUFamily string `json:"cpuFamily,omitempty"`

	// A valid uuid of the datacenter on which user has access
	// Required: true
	DatacenterID string `json:"datacenterId,omitempty"`

	// A Kubernetes Node Pool Name. Valid Kubernetes Node Pool name must be 63 characters or less and must be
	// empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_),
	// dots (.), and alphanumerics between.
	// Required: true
	Name string `json:"name,omitempty"`

	// Number of nodes part of the Node Pool
	// Required: true
	NodeCount int64 `json:"nodeCount,omitempty"`

	// RAM size for node, minimum size 2048MB is recommended
	// Required: true
	RAMSize int64 `json:"ramSize,omitempty"`

	// The size of the volume in GB. The size should be greater than 10GB.
	// Required: true
	StorageSize int64 `json:"storageSize,omitempty"`

	// Hardware type of the volume
	// Required: true
	// Enum: [HDD SSD]
	StorageType string `json:"storageType,omitempty"`
}

// KubernetesNodePools kubernetes node pools
// swagger:model KubernetesNodePools
type KubernetesNodePools struct {
	// URL to the collection representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// Unique representation for Kubernetes Node Pool as a collection on a resource.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Array of items in that collection
	// Read Only: true
	Items []*KubernetesNodePool `json:"items"`

	// The type of resource within a collection
	// Read Only: true
	// Enum: [nodepool]
	Type string `json:"type,omitempty"`
}

func (c *Client) ListKubernetesClusters() (*KubernetesClusters, error) {
	rsp := &KubernetesClusters{}
	return rsp, c.GetOK(kubernetesClustersPath(), rsp)
}

func (c *Client) GetKubernetesCluster(clusterId string) (*KubernetesCluster, error) {
	rsp := &KubernetesCluster{}
	return rsp, c.GetOK(kubernetesClusterPath(clusterId), rsp)
}

func (c *Client) CreateKubernetesCluster(cluster KubernetesCluster) (*KubernetesCluster, error) {
	rsp := &KubernetesCluster{}
	return rsp, c.PostAcc(kubernetesClustersPath(), cluster, rsp)
}

func (c *Client) DeleteKubernetesCluster(clusterId string) (*http.Header, error) {
	h := &http.Header{}
	return h, c.Delete(kubernetesClusterPath(clusterId), h, http.StatusOK)
}

func (c *Client) UpdateKubernetesCluster(clusterId string, cluster KubernetesCluster) (*KubernetesCluster, error) {
	rsp := &KubernetesCluster{}
	return rsp, c.Put(kubernetesClusterPath(clusterId), cluster, rsp, http.StatusOK)
}

func (c *Client) GetKubeconfig(clusterId string) (string, error) {
	rsp := &KubernetesConfig{}
	if err := c.GetOK(kubeConfigPath(clusterId), rsp); err != nil {
		return "", err
	}
	return rsp.Properties.KubeConfig, nil
}

func (c *Client) GetKubernetesNodePools(clusterId string) (*KubernetesNodePools, error) {
	rsp := &KubernetesNodePools{}
	return rsp, c.GetOK(kubernetesNodePoolsPath(clusterId), rsp)
}

func (c *Client) CreateKubernetesNodePool(clusterId string, nodePool KubernetesNodePool) (*KubernetesNodePool, error) {
	rsp := &KubernetesNodePool{}
	return rsp, c.PostAcc(kubernetesNodePoolsPath(clusterId), nodePool, rsp)
}

func (c *Client) DeleteKubernetesNodePool(clusterId, nodePoolId string) (*http.Header, error) {
	return c.DeleteAcc(kubernetesNodePoolPath(clusterId, nodePoolId))
}

func (c *Client) GetKubernetesNodePool(clusterId, nodePoolId string) (*KubernetesNodePool, error) {
	rsp := &KubernetesNodePool{}
	return rsp, c.GetOK(kubernetesNodePoolPath(clusterId, nodePoolId), rsp)
}

func (c *Client) UpdateKubernetesNodePool(clusterId, nodePoolId string, nodePool KubernetesNodePool) (*KubernetesNodePool, error) {
	rsp := &KubernetesNodePool{}
	return rsp, c.PutAcc(kubernetesNodePoolPath(clusterId, nodePoolId), nodePool, rsp)
}
