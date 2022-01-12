// Copyright Contributors to the Open Cluster Management project
package clusterclaim

import (
	"fmt"

	clusteradmhelpers "open-cluster-management.io/clusteradm/pkg/helpers"

	genericclioptionscm "github.com/stolostron/cm-cli/pkg/genericclioptions"
	"github.com/stolostron/cm-cli/pkg/helpers"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var example = `
# Delete a clusterclaim in the current clusterpoolhost
%[1]s delete cc <clusterclaim_name>[,<clusterclaim_name>...] <options>

# Delete a clusterclaim on a given clusterpoolhost
%[1]s delete cc <clusterclaim_name>[,<clusterclaim_name>...] -cph <clusterpoolhost_name> <options>
`

// NewCmd ...
func NewCmd(cmFlags *genericclioptionscm.CMFlags, streams genericclioptions.IOStreams) *cobra.Command {
	o := newOptions(cmFlags, streams)
	cmd := &cobra.Command{
		Use:          "clusterclaim",
		Aliases:      []string{"cc", "ccs", "clusterclaims"},
		Short:        "Delete clusterclaims",
		Example:      fmt.Sprintf(example, helpers.GetExampleHeader()),
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			clusteradmhelpers.DryRunMessage(cmFlags.DryRun)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.complete(cmd, args))
			cmdutil.CheckErr(o.validate())
			cmdutil.CheckErr(o.run())
		},
	}

	cmd.Flags().StringVar(&o.ClusterPoolHost, "cph", "", "The clusterpoolhost to use")

	return cmd
}
