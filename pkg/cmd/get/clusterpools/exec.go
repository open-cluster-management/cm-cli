// Copyright Contributors to the Open Cluster Management project
package clusterpools

import (
	"fmt"

	printclusterpoolv1alpha1 "github.com/open-cluster-management/cm-cli/api/cm-cli/v1alpha1"
	"github.com/open-cluster-management/cm-cli/pkg/clusterpoolhost"
	"github.com/open-cluster-management/cm-cli/pkg/helpers"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/spf13/cobra"
)

func (o *Options) complete(cmd *cobra.Command, args []string) (err error) {
	if len(*o.PrintFlags.OutputFormat) == 0 {
		o.PrintFlags.OutputFormat = &clusterpoolhost.ClusterPoolsColumns
	}
	return nil
}

func (o *Options) validate() error {
	if o.ClusterPoolHost != "" && o.AllClusterPoolHosts {
		return fmt.Errorf("clusterpoolhost and all-cphs are imcompatible")
	}
	return nil
}

func (o *Options) run() (err error) {
	var cphs, allcphs *clusterpoolhost.ClusterPoolHosts
	allcphs, err = clusterpoolhost.GetClusterPoolHosts()
	if err != nil {
		return err
	}
	currentCph, err := allcphs.GetCurrentClusterPoolHost()
	if err != nil {
		return err
	}

	if o.AllClusterPoolHosts {
		cphs, err = clusterpoolhost.GetClusterPoolHosts()
		if err != nil {
			return err
		}
	} else {
		var cph *clusterpoolhost.ClusterPoolHost
		if o.ClusterPoolHost != "" {
			cph, err = clusterpoolhost.GetClusterPoolHost(o.ClusterPoolHost)
		} else {
			cph, err = clusterpoolhost.GetCurrentClusterPoolHost()
		}
		if err != nil {
			return err
		}
		cphs = &clusterpoolhost.ClusterPoolHosts{
			ClusterPoolHosts: map[string]*clusterpoolhost.ClusterPoolHost{
				cph.Name: cph,
			},
		}
	}

	printClusterPoolLists := &printclusterpoolv1alpha1.PrintClusterPoolList{}
	printClusterPoolLists.GetObjectKind().
		SetGroupVersionKind(
			schema.GroupVersionKind{
				Group:   printclusterpoolv1alpha1.GroupName,
				Kind:    "PrintClusterPool",
				Version: printclusterpoolv1alpha1.GroupVersion.Version})
	for k := range cphs.ClusterPoolHosts {
		err = allcphs.SetActive(allcphs.ClusterPoolHosts[k])
		if err != nil {
			return err
		}
		clusterPools, err := clusterpoolhost.GetClusterPools(o.AllClusterPoolHosts, o.CMFlags.DryRun)
		if err != nil {
			fmt.Printf("Error while retrieving clusterpools from %s\n", cphs.ClusterPoolHosts[k].Name)
			continue
		}
		printClusterPoolList := clusterpoolhost.ConvertToPrintClusterPoolList(cphs.ClusterPoolHosts[k], clusterPools)
		printClusterPoolLists.Items = append(printClusterPoolLists.Items, printClusterPoolList.Items...)
	}
	err = helpers.Print(printClusterPoolLists, o.PrintFlags)
	allcphs.SetActive(currentCph)
	if err != nil {
		return err
	}
	return nil
}
