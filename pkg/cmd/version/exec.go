// Copyright Contributors to the Open Cluster Management project
package version

import (
	"fmt"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	cmcli "github.com/open-cluster-management/cm-cli"
	"github.com/open-cluster-management/cm-cli/pkg/helpers"
	"github.com/spf13/cobra"
)

func (o *Options) complete(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (o *Options) validate() error {
	return nil
}
func (o *Options) run() (err error) {
	fmt.Printf("client\t\tversion\t:%s\n", cmcli.GetVersion())
	kubeClient, err := o.CMFlags.KubectlFactory.KubernetesClientSet()
	if err != nil {
		return err
	}
	dynamicClient, err := o.CMFlags.KubectlFactory.DynamicClient()
	if err != nil {
		return err
	}
	return o.runWithClient(kubeClient, dynamicClient)
}

func (o *Options) runWithClient(kubeClient kubernetes.Interface, dynamicClient dynamic.Interface) (err error) {
	var version, snapshot string
	switch {
	case helpers.IsRHACM(o.CMFlags.KubectlFactory):
		version, snapshot, err = helpers.GetACMVersion(kubeClient, dynamicClient)
	case helpers.IsMCE(o.CMFlags.KubectlFactory):
		version, snapshot, err = helpers.GetMCEVersion(kubeClient, dynamicClient)
	}
	if version != "" {
		fmt.Printf("server release\tversion\t:%s\n", version)
	}
	if snapshot != "" {
		fmt.Printf("server image\ttag\t:%s\n", snapshot)
	}
	if err != nil {
		return err
	}
	return nil
}
