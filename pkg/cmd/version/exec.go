// Copyright Contributors to the Open Cluster Management project
package version

import (
	"context"
	"fmt"
	"strings"

	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"

	crclient "sigs.k8s.io/controller-runtime/pkg/client"

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
	client, err := helpers.GetControllerRuntimeClientFromFlags(o.ConfigFlags)
	if err != nil {
		return err
	}
	return o.runWithClient(client)
}

func (o *Options) runWithClient(client crclient.Client) (err error) {
	cms := &corev1.ConfigMapList{}
	ls := labels.SelectorFromSet(labels.Set{
		"ocm-configmap-type": "image-manifest",
	})
	err = client.List(context.TODO(), cms, &crclient.ListOptions{LabelSelector: ls})
	if err != nil {
		return err
	}
	if len(cms.Items) > 1 {
		return fmt.Errorf("found more than one configmap with labelset %v", ls)
	}
	if v, ok := cms.Items[0].Labels["ocm-release-version"]; ok {
		fmt.Printf("server release\tversion\t:%s\n", v)
	} else {
		fmt.Printf("server release\tversion\t: not found")
	}
	ns := cms.Items[0].Namespace
	acmRegistryDeployment := &apps.Deployment{}
	err = client.Get(context.TODO(), crclient.ObjectKey{Name: "acm-custom-registry", Namespace: ns}, acmRegistryDeployment)
	if err != nil {
		return nil
	}
	for _, c := range acmRegistryDeployment.Spec.Template.Spec.Containers {
		if strings.Contains(c.Image, "acm-custom-registry") {
			fmt.Printf("server image\ttag\t:%s\n", strings.Split(c.Image, ":")[1])
			break
		}
	}
	return nil
}
