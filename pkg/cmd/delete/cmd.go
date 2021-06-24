// Copyright Contributors to the Open Cluster Management project
package delete

import (
	"github.com/open-cluster-management/cm-cli/pkg/cmd/delete/cluster"
	genericclioptionscm "github.com/open-cluster-management/cm-cli/pkg/genericclioptions"
	clusteradmdeletetoken "open-cluster-management.io/clusteradm/pkg/cmd/delete/token"
	genericclioptionsclusteradm "open-cluster-management.io/clusteradm/pkg/genericclioptions"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewCmd provides a cobra command wrapping NewCmdImportCluster
func NewCmd(clusteradmFlags *genericclioptionsclusteradm.ClusteradmFlags, cmFlags *genericclioptionscm.CMFlags, streams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a resource",
	}

	cmd.AddCommand(cluster.NewCmd(cmFlags, streams))
	cmd.AddCommand(clusteradmdeletetoken.NewCmd(clusteradmFlags, streams))
	return cmd
}
