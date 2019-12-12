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

	// Slice of items in that collection
	// Read Only: true
	Items []KubernetesCluster `json:"items"`

	// The type of resource within a collection
	// Read Only: true
	// Enum: [collection]
	PBType string `json:"type,omitempty"`
}

type KubernetesCluster struct {
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
	Properties *KubernetesClusterProperties `json:"properties"`

	// The type of object
	// Read Only: true
	// Enum: [k8s]
	PBType string `json:"type,omitempty"`

	// Entities of a cluster
	Entities KubernetesClusterEntities `json:"entities,omitempty"`
}

type UpdatedKubernetesCluster struct {
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
	Properties *KubernetesClusterProperties `json:"properties"`

	// The type of object
	// Read Only: true
	// Enum: [k8s]
	PBType string `json:"type,omitempty"`

	// Entities of a cluster
	Entities KubernetesClusterEntities `json:"-"`
}

type KubernetesClusterEntities struct {
	// NodePools of a cluster
	NodePools *KubernetesNodePools `json:"nodepools,omitempty"`
}

type KubernetesClusterProperties struct {
	// A Kubernetes Cluster Name. Valid Kubernetes Cluster name must be 63 characters or less and must not be empty
	// and begin and end with an alphanumeric character ([a-z0-9]) with dashes (-), dots (.) and alphanumerics
	// between.
	// Required: true
	Name string `json:"name"`
}

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
	Properties KubernetesConfigProperties `json:"properties"`

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

type KubernetesNodePoolProperties struct {
	// The availability zone in which the servers should exist
	// Required: true
	// Enum: [AUTO ZONE_1 ZONE_2]
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// Number of cores for node
	// Required: true
	CoresCount uint32 `json:"coresCount,omitempty"`

	// A valid cpu family name
	// Required: true
	CPUFamily string `json:"cpuFamily,omitempty"`

	// The unique identifier of the data center where the worker nodes of the node pool will be provisioned.
	// Required: true
	DatacenterID string `json:"datacenterId,omitempty"`

	// A Kubernetes Node Pool Name. Valid Kubernetes Node Pool name must be 63 characters or less and must not be
	// empty or begin and end with an alphanumeric character ([a-z0-9]) with dashes (-), dots (.) and alphanumerics
	// between.
	// Required: true
	Name string `json:"name,omitempty"`

	// Number of nodes part of the Node Pool
	// Required: true
	NodeCount uint32 `json:"nodeCount,omitempty"`

	// RAM size for node, minimum size 2048MB is recommended
	// Required: true
	RAMSize uint32 `json:"ramSize,omitempty"`

	// The size of the volume in GB. The size should be greater than 10GB.
	// Required: true
	StorageSize uint32 `json:"storageSize,omitempty"`

	// Hardware type of the volume
	// Required: true
	// Enum: [HDD SSD]
	StorageType string `json:"storageType,omitempty"`
}

type KubernetesNodePools struct {
	// URL to the collection representation (absolute path)
	// Read Only: true
	// Format: uri
	Href string `json:"href,omitempty"`

	// Unique representation for Kubernetes Node Pool as a collection on a resource.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Slice of items in that collection
	// Read Only: true
	Items []KubernetesNodePool `json:"items"`

	// The type of resource within a collection
	// Read Only: true
	// Enum: [nodepool]
	Type string `json:"type,omitempty"`
}

// ListKubernetesClusters gets a list of all clusters
func (c *Client) ListKubernetesClusters() (*KubernetesClusters, error) {
	rsp := &KubernetesClusters{}
	return rsp, c.GetOK(kubernetesClustersPath(), rsp)
}

// GetKubernetesCluster gets cluster with given id
func (c *Client) GetKubernetesCluster(clusterID string) (*KubernetesCluster, error) {
	rsp := &KubernetesCluster{}
	return rsp, c.GetOK(kubernetesClusterPath(clusterID), rsp)
}

// CreateKubernetesCluster creates a cluster
func (c *Client) CreateKubernetesCluster(cluster KubernetesCluster) (*KubernetesCluster, error) {
	rsp := &KubernetesCluster{}
	return rsp, c.PostAcc(kubernetesClustersPath(), cluster, rsp)
}

// DeleteKubernetesCluster deletes cluster
func (c *Client) DeleteKubernetesCluster(clusterId string) (*http.Header, error) {
	h := &http.Header{}
	return h, c.Delete(kubernetesClusterPath(clusterId), h, http.StatusOK)
}

// UpdateKubernetesCluster updates cluster
func (c *Client) UpdateKubernetesCluster(clusterID string, cluster UpdatedKubernetesCluster) (*KubernetesCluster, error) {
	rsp := &KubernetesCluster{}
	return rsp, c.Put(kubernetesClusterPath(clusterID), cluster, rsp, http.StatusOK)
}

// GetKubeconfig returns the kubeconfig of cluster
func (c *Client) GetKubeconfig(clusterID string) (string, error) {
	rsp := &KubernetesConfig{}
	if err := c.GetOK(kubeConfigPath(clusterID), rsp); err != nil {
		return "", err
	}
	return rsp.Properties.KubeConfig, nil
}

// GetKubernetesNodePools gets all node pools of cluster
func (c *Client) GetKubernetesNodePools(clusterID string) (*KubernetesNodePools, error) {
	rsp := &KubernetesNodePools{}
	return rsp, c.GetOK(kubernetesNodePoolsPath(clusterID), rsp)
}

// CreateKubernetesNodePool creates a new node pool for cluster
func (c *Client) CreateKubernetesNodePool(clusterID string, nodePool KubernetesNodePool) (*KubernetesNodePool, error) {
	rsp := &KubernetesNodePool{}
	return rsp, c.PostAcc(kubernetesNodePoolsPath(clusterID), nodePool, rsp)
}

// DeleteKubernetesNodePool deletes node pool from cluster
func (c *Client) DeleteKubernetesNodePool(clusterID, nodePoolID string) (*http.Header, error) {
	return c.DeleteAcc(kubernetesNodePoolPath(clusterID, nodePoolID))
}

// GetKubernetesNodePool gets node pool of the cluster
func (c *Client) GetKubernetesNodePool(clusterID, nodePoolID string) (*KubernetesNodePool, error) {
	rsp := &KubernetesNodePool{}
	return rsp, c.GetOK(kubernetesNodePoolPath(clusterID, nodePoolID), rsp)
}

// Update KubernetesNodePool updates node pool
func (c *Client) UpdateKubernetesNodePool(clusterID, nodePoolID string, nodePool KubernetesNodePool) (*KubernetesNodePool, error) {
	rsp := &KubernetesNodePool{}
	return rsp, c.PutAcc(kubernetesNodePoolPath(clusterID, nodePoolID), nodePool, rsp)
}
