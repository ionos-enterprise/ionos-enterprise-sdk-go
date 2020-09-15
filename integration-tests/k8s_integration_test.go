package integration_tests

import (
	sdk "github.com/ionos-enterprise/ionos-enterprise-sdk-go/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ClusterName = "GO-SDK-K8S-TEST-CLUSTER"
	ClusterNameUpdated = ClusterName + "-UPDATED"
)

func TestCreateKubernetesCluster(t *testing.T) {
	c := setupTestEnv()
	clusterDef := sdk.KubernetesCluster{
		Properties: &sdk.KubernetesClusterProperties{
			Name: ClusterName,
		},
	}
	ret, err := c.CreateKubernetesCluster(clusterDef)
	if err != nil {
		t.Fatal(err)
	}
	cluster = ret
	assert.Equal(t, ClusterName, ret.Properties.Name)
}

func TestListKubernetesClusters(t *testing.T) {
	c := setupTestEnv()
	clusters, err := c.ListKubernetesClusters()
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(clusters.Items) > 0)
}

func TestGetKubernetesCluster(t *testing.T) {
	c := setupTestEnv()
	ret, err := c.GetKubernetesCluster(cluster.ID)
	if err != nil { t.Fatal(err) }

	assert.Equal(t, cluster.Properties.Name, ret.Properties.Name)
}

func TestUpdateKubernetesCluster(t *testing.T) {
	c := setupTestEnv()

	update := sdk.UpdatedKubernetesCluster{
		Properties: &sdk.KubernetesClusterProperties{
			Name: ClusterNameUpdated,
		},
	}
	ret, err := c.UpdateKubernetesCluster(cluster.ID, update)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, ClusterNameUpdated, ret.Properties.Name)
}

func TestDeleteKubernetesCluster(t *testing.T) {
	c := setupTestEnv()
	_, err := c.DeleteKubernetesCluster(cluster.ID)
	if err != nil {
		t.Fatal(err)
	}
}
