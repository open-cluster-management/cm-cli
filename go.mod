module github.com/stolostron/cm-cli

go 1.16

replace (
	github.com/go-logr/logr => github.com/go-logr/logr v0.4.0
	github.com/openshift/api => github.com/openshift/api v0.0.0-20211007134530-4cb30f221b89
	k8s.io/api => k8s.io/api v0.22.1
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.22.1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.22.1
	k8s.io/client-go => k8s.io/client-go v0.22.1
	k8s.io/code-generator => k8s.io/code-generator v0.22.1
	k8s.io/component-base => k8s.io/component-base v0.22.1
	k8s.io/kubectl => k8s.io/kubectl v0.22.1
)

require (
	github.com/Masterminds/semver v1.5.0
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/openshift/api v3.9.1-0.20190924102528-32369d4db2ad+incompatible
	github.com/openshift/client-go v0.0.0-20210916133943-9acee1a0fb83
	github.com/openshift/hive/apis v0.0.0-20211012143010-16ef5a35537d
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stolostron/governance-policy-propagator v0.0.0-20220111211854-68e6640dc4bf
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2
	k8s.io/api v0.22.1
	k8s.io/apiextensions-apiserver v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/cli-runtime v0.22.1
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.22.1
	k8s.io/component-base v0.22.1
	k8s.io/klog/v2 v2.9.0
	k8s.io/kubectl v0.22.1
	open-cluster-management.io/api v0.5.1-0.20211109002058-9676c7a1e606
	open-cluster-management.io/clusteradm v0.1.0-alpha.5.0.20211012235309-d275d8270776
	sigs.k8s.io/controller-runtime v0.9.6
)
