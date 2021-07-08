// Copyright Contributors to the Open Cluster Management project
package use

import (
	"github.com/open-cluster-management/cm-cli/pkg/clusterpoolhost"
	genericclioptionscm "github.com/open-cluster-management/cm-cli/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type Options struct {
	//CMFlags: The generic optiosn from the cm cli-runtime.
	CMFlags                 *genericclioptionscm.CMFlags
	Cluster                 clusterpoolhost.ClusterPoolHost
	ServiceAccountName      string
	ServiceAccountNameSpace string
	//The file to output the resources will be sent to the file.
	outputFile string
}

func newOptions(cmFlags *genericclioptionscm.CMFlags, streams genericclioptions.IOStreams) *Options {
	return &Options{
		CMFlags:                 cmFlags,
		ServiceAccountNameSpace: clusterpoolhost.ServiceAccountNameSpace,
	}
}
