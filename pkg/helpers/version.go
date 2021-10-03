// Copyright Contributors to the Open Cluster Management project

package helpers

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"

	"k8s.io/client-go/kubernetes"
)

const (
	RHACM string = "RHACM"
	MCE   string = "MCE"
)

func GetACMVersion(kubeClient kubernetes.Interface, dynamicClient dynamic.Interface) (version, snapshot string, err error) {
	lo := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%v = %v", "ocm-configmap-type", "image-manifest"),
	}
	cms, err := kubeClient.CoreV1().ConfigMaps("").List(context.TODO(), lo)
	if err != nil {
		return "", "", err
	}
	if len(cms.Items) == 0 {
		return "", "", fmt.Errorf("no configmap with labelset %v", lo.LabelSelector)
	}
	ns := cms.Items[0].Namespace

	umch, err := dynamicClient.Resource(GvrMCH).Namespace(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", "", err
	}
	if len(umch.Items) == 0 {
		return "", "", fmt.Errorf("no multiclusterhub found in namespace %s", ns)
	}
	ustatus, ok := umch.Items[0].Object["status"]
	if !ok {
		return "", "", fmt.Errorf("no status found multiclusterhub in %s/%s", ns, umch.Items[0].GetName())
	}
	uversion, ok := ustatus.(map[string]interface{})["currentVersion"]
	if !ok {
		return "", "", fmt.Errorf("no currentVersion found multiclusterhub in %s/%s", ns, umch.Items[0].GetName())
	}
	version = uversion.(string)
	lo = metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%v = %v,%v = %v", "ocm-configmap-type", "image-manifest", "ocm-release-version", version),
	}
	cms, err = kubeClient.CoreV1().ConfigMaps("").List(context.TODO(), lo)
	if err != nil {
		return version, "", err
	}
	if len(cms.Items) == 1 {
		ns := cms.Items[0].Namespace
		if v, ok := cms.Items[0].Labels["ocm-release-version"]; ok {
			version = v
		}
		acmRegistryDeployment, err := kubeClient.AppsV1().Deployments(ns).Get(context.TODO(), "acm-custom-registry", metav1.GetOptions{})
		if err == nil {
			for _, c := range acmRegistryDeployment.Spec.Template.Spec.Containers {
				if strings.Contains(c.Image, "acm-custom-registry") {
					snapshot = strings.Split(c.Image, ":")[1]
					break
				}
			}
		}
	}
	return version, snapshot, nil
}

func GetMCEVersion(kubeClient kubernetes.Interface, dynamicClient dynamic.Interface) (version, snapshot string, err error) {
	lo := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%v = %v", "operators.coreos.com/multicluster-engine.multicluster-engine", ""),
	}
	cms, err := kubeClient.CoreV1().ConfigMaps("").List(context.TODO(), lo)
	if err != nil {
		return "", "", err
	}
	if len(cms.Items) == 0 {
		return "", "", fmt.Errorf("no configmap with labelset %v", lo.LabelSelector)
	}
	ns := cms.Items[0].Namespace
	ucsv, err := dynamicClient.Resource(GvrMCE).Namespace(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", "", err
	}
	if err != nil {
		return "", "", err
	}
	if len(ucsv.Items) == 0 {
		return "", "", fmt.Errorf("no clusterserviceversion found in namespace %s", ns)
	}

	uspec := ucsv.Items[0].Object["spec"]
	spec := uspec.(map[string]interface{})
	version = spec["version"].(string)
	return version, "", nil

}

func GetVersion(f cmdutil.Factory) (version string, platform string, err error) {
	kubeClient, err := f.KubernetesClientSet()
	if err != nil {
		return version, platform, err
	}
	dynamicClient, err := f.DynamicClient()
	if err != nil {
		return version, platform, err
	}
	switch {
	case IsRHACM(f):
		platform = RHACM
		version, _, err = GetACMVersion(kubeClient, dynamicClient)
	case IsMCE(f):
		platform = MCE
		version, _, err = GetMCEVersion(kubeClient, dynamicClient)
	}
	return version, platform, err
}

func IsSupported(f cmdutil.Factory, rhacmConstraint string, mceConstraint string) (isSupported bool, platform string, err error) {
	var version string
	kubeClient, err := f.KubernetesClientSet()
	if err != nil {
		return isSupported, platform, err
	}
	dynamicClient, err := f.DynamicClient()
	if err != nil {
		return isSupported, platform, err
	}
	var c *semver.Constraints
	switch {
	case IsRHACM(f):
		platform = RHACM
		version, _, err = GetACMVersion(kubeClient, dynamicClient)
		if err != nil {
			return isSupported, platform, err
		}
		c, err = semver.NewConstraint(rhacmConstraint)
	case IsMCE(f):
		platform = MCE
		version, _, err = GetMCEVersion(kubeClient, dynamicClient)
		if err != nil {
			return isSupported, platform, err
		}
		c, err = semver.NewConstraint(mceConstraint)
	}
	if err != nil {
		return isSupported, platform, err
	}

	vs, err := semver.NewVersion(version)
	if err != nil {
		return isSupported, platform, err
	}

	return c.Check(vs), platform, nil
}
