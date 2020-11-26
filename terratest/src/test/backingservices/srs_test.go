package backingservices

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_shouldNotContainSRSResourcesWhenDisabled(t *testing.T) {
	helmChartParser := NewHelmConfigParser(
		NewHelmTest(t, helmChartRelativePath, map[string]string{
			"srs.enabled": "false",
			"srs.elasticsearch.provisionCluster": "false",
			"srs.deploymentName": "test-srs",
		}),
	)

	for _, i := range srsResources {
		require.False(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}

	for _, i := range esResources {
		require.False(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}
}

func Test_shouldContainSRSResourcesWhenEnabled(t *testing.T) {
	helmChartParser := NewHelmConfigParser(
		NewHelmTest(t, helmChartRelativePath, map[string]string{
			"srs.deploymentName": "test-srs",
			"srs.elasticsearch.provisionCluster": "true",
		}),
	)

	for _, i := range srsResources {
		require.True(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}

	for _, i := range esResources {
		require.True(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}
}

func Test_shouldContainSRSandESResourcesWhenEnabled(t *testing.T) {
	helmChartParser := NewHelmConfigParser(
		NewHelmTest(t, helmChartRelativePath, map[string]string{
			"srs.deploymentName": "test-srs",
			"srs.elasticsearch.provisionCluster": "true",
		}),
	)

	for _, i := range srsResources {
		require.True(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}

	for _, i := range esResources {
		require.True(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}
}

func Test_shouldContainSRSWhenEnabledandNotESResourcesWhenDisabled(t *testing.T) {
	helmChartParser := NewHelmConfigParser(
		NewHelmTest(t, helmChartRelativePath, map[string]string{
			"srs.deploymentName": "test-srs",
			"srs.elasticsearch.provisionCluster": "false",
			"srs.elasticsearch.domain": "es.managed.io",
			"srs.elasticsearch.port": "9200",
			"srs.elasticsearch.protocol": "https",
		}),
	)

	for _, i := range srsResources {
		require.True(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}

	for _, i := range esResources {
		require.False(t, helmChartParser.Contains(SearchResourceOption{
			Name: i.Name,
			Kind: i.Kind,
		}))
	}
}

var srsResources = []SearchResourceOption{
	{
		Name: "test-srs",
		Kind: "Deployment",
	},
	{
		Name: "test-srs",
		Kind: "Service",
	},
	{
		Name: "test-srs",
		Kind: "PodDisruptionBudget",
	},
	{
		Name: "test-srs-networkpolicy",
		Kind: "NetworkPolicy",
	},
	{
		Name: "srs-registry-secret",
		Kind: "Secret",
	},
}

var esResources = []SearchResourceOption{
	{
		Name: "elasticsearch-master",
		Kind: "Service",
	},
	{
		Name: "elasticsearch-master-headless",
		Kind: "Service",
	},
	{
		Name: "elasticsearch-master",
		Kind: "StatefulSet",
	},
	{
		Name: "elasticsearch-master-pdb",
		Kind: "PodDisruptionBudget",
	},
}
