// Copyright Contributors to the Open Cluster Management project
package cluster

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	clusteradmhelpers "open-cluster-management.io/clusteradm/pkg/helpers"

	"github.com/open-cluster-management/cm-cli/pkg/cmd/delete/cluster/scenario"
	genericclioptionscm "github.com/open-cluster-management/cm-cli/pkg/genericclioptions"
	"github.com/open-cluster-management/cm-cli/pkg/helpers"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var example = `
# Delete a cluster
%[1]s delete cluster --values values.yaml

# Delete a cluster with overwritting the cluster name
%[1]s delete cluster  --cluster mycluster --values values.yaml

# Delete a cluster with overwritting the cluster name with arg
%[1]s delete cluster mycluster--values values.yaml
`

const (
	scenarioDirectory = "delete"
)

var valuesTemplatePath = filepath.Join(scenarioDirectory, "values-template.yaml")

// NewCmd ...
func NewCmd(cmFlags *genericclioptionscm.CMFlags, streams genericclioptions.IOStreams) *cobra.Command {
	o := newOptions(cmFlags, streams)

	cluster := &cobra.Command{
		Use:          "cluster",
		Short:        "Delete a cluster",
		Example:      fmt.Sprintf(example, helpers.GetExampleHeader()),
		SilenceUsage: true,
		PreRunE: func(c *cobra.Command, args []string) error {
			if !helpers.IsRHACM(cmFlags.KubectlFactory) && !helpers.IsMCE(cmFlags.KubectlFactory) {
				return fmt.Errorf("this command '%s %s' is only available on %s or %s",
					helpers.GetExampleHeader(),
					strings.Join(os.Args[1:], " "),
					helpers.RHACM,
					helpers.MCE)
			}
			clusteradmhelpers.DryRunMessage(cmFlags.DryRun)
			return nil
		},
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.complete(c, args); err != nil {
				return err
			}
			if err := o.validate(); err != nil {
				return err
			}
			if err := o.run(); err != nil {
				return err
			}

			return nil
		},
	}

	cluster.SetUsageTemplate(clusteradmhelpers.UsageTempate(cluster, scenario.GetScenarioResourcesReader(), valuesTemplatePath))
	cluster.Flags().StringVar(&o.clusterName, "cluster", "", "Name of the cluster")
	cluster.Flags().StringVar(&o.valuesPath, "values", "", "The files containing the values")

	return cluster
}
