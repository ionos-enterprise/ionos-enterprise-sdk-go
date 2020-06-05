package profitbricks

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type SuiteKubernetesCluster struct {
	ClientBaseSuite
}

func TestKubernetesCluster(t *testing.T) {
	suite.Run(t, new(SuiteKubernetesCluster))
}

func validateMetadata(t *testing.T, m *Metadata) {
	if !assert.NotNil(t, m) {
		return
	}
	assert.NotEmpty(t, m.State)
	assert.NotEmpty(t, m.CreatedBy)
	assert.NotEmpty(t, m.CreatedDate)
	assert.NotEmpty(t, m.CreatedByUserID)
	assert.NotEmpty(t, m.Etag)
	assert.NotEmpty(t, m.LastModifiedBy)
	assert.NotEmpty(t, m.LastModifiedDate)
	assert.NotEmpty(t, m.LastModifiedByUserID)
}

func (s *SuiteKubernetesCluster) Test_ListKubernetesClusters() {
	rsp := loadTestData(s.T(), "get_kubernetes_clusters.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/k8s`,
		httpmock.ResponderFromResponse(mResp))
	cl, err := s.c.ListKubernetesClusters()
	s.NoError(err)
	s.Len(cl.Items, 2)
}
func (s *SuiteKubernetesCluster) Test_GetKubernetesCluster() {
	rsp := loadTestData(s.T(), "get_kubernetes_cluster.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/k8s/1`,
		httpmock.ResponderFromResponse(mResp))
	cl, err := s.c.GetKubernetesCluster("1")
	s.NoError(err)
	s.Len(cl.Entities.NodePools.Items, 1)
	validateMetadata(s.T(), cl.Metadata)
	s.NotEmpty(cl.Properties.Name)
}

func (s *SuiteKubernetesCluster) Test_GetKubeconfig() {
	rsp := loadTestData(s.T(), "get_kubeconfig.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/k8s/1/kubeconfig`,
		httpmock.ResponderFromResponse(mResp))

	cfg, err := s.c.GetKubeconfig("1")
	s.NoError(err)
	s.Equal("---probably valid config", cfg)
}

func (s *SuiteKubernetesCluster) Test_ListKubernetesNodepools() {
	rsp := loadTestData(s.T(), "get_kubernetes_nodepools.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/k8s/1/nodepools`,
		httpmock.ResponderFromResponse(mResp))

	nps, err := s.c.ListKubernetesNodePools("1")
	s.NoError(err)
	s.Len(nps.Items, 1)
}

func (s *SuiteKubernetesCluster) Test_GetKubernetesNodepool() {
	rsp := loadTestData(s.T(), "get_kubernetes_nodepool.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/k8s/1/nodepools/2`,
		httpmock.ResponderFromResponse(mResp))

	np, err := s.c.GetKubernetesNodePool("1", "2")
	s.NoError(err)
	validateMetadata(s.T(), np.Metadata)
	s.NotEmpty(np.Properties.Name)
	s.NotEmpty(np.Properties.NodeCount)
	s.NotEmpty(np.Properties.DatacenterID)
	s.NotEmpty(np.Properties.AvailabilityZone)
	s.NotEmpty(np.Properties.CoresCount)
	s.NotEmpty(np.Properties.CPUFamily)
	s.NotEmpty(np.Properties.RAMSize)
	s.NotEmpty(np.Properties.StorageSize)
	s.NotEmpty(np.Properties.StorageType)
	s.NotEmpty(np.Properties.Autoscaling)
	s.NotEmpty(np.Properties.MaintenanceWindow)
}

func (s *SuiteKubernetesCluster) Test_ListKubernetesNodes() {
	rsp := loadTestData(s.T(), "list_kubernetes_nodes.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/k8s/1/nodepools/2/nodes`,
		httpmock.ResponderFromResponse(mResp))

	nodes, err := s.c.ListKubernetesNodes("1", "2")
	s.NoError(err)
	s.Len(nodes.Items, 2)
}

func (s *SuiteKubernetesCluster) Test_GetKubernetesNode() {
	rsp := loadTestData(s.T(), "get_kubernetes_node.json")
	mResp := makeJsonResponse(http.StatusOK, rsp)
	httpmock.RegisterResponder(http.MethodGet, `=~/k8s/1/nodepools/2/nodes/3`,
		httpmock.ResponderFromResponse(mResp))

	node, err := s.c.GetKubernetesNode("1", "2", "3")
	s.NoError(err)
	s.Equal(node.Properties.Name, "node2")
	s.Equal(node.Properties.PublicIP, "222.222.222.222")
	s.Equal(node.Properties.K8sVersion, "1.16.4")
}

func (s *SuiteKubernetesCluster) Test_DeleteKubernetesNode() {
	mRsp := makeJsonResponse(http.StatusAccepted, nil)
	httpmock.RegisterResponder(
		http.MethodDelete, "=~/k8s/1/nodepools/2/nodes/3", httpmock.ResponderFromResponse(mRsp))
	rsp, err := s.c.DeleteKubernetesNode("1", "2", "3")
	s.NoError(err)
	s.NotNil(rsp)
}

func (s *SuiteKubernetesCluster) Test_ReplaceKubernetesNode() {
	mRsp := makeJsonResponse(http.StatusAccepted, nil)
	httpmock.RegisterResponder(
		http.MethodPost, "=~/k8s/1/nodepools/2/nodes/3/replace", httpmock.ResponderFromResponse(mRsp))
	rsp, err := s.c.ReplaceKubernetesNode("1", "2", "3")
	s.NoError(err)
	s.NotNil(rsp)
}
