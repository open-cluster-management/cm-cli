module github.com/open-cluster-management/cm-cli

go 1.16

replace (
	github.com/go-logr/logr => github.com/go-logr/logr v0.4.0
	k8s.io/client-go => k8s.io/client-go v0.20.4
	open-cluster-management.io/clusteradm => /Users/dvernier/acm-tools/clusteradm
)

require (
	github.com/ghodss/yaml v1.0.0
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/open-cluster-management/applier v0.0.0-20210422205113-6c10f923726b
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
	k8s.io/cli-runtime v0.20.5
	k8s.io/client-go v1.5.2
	k8s.io/component-base v0.20.1
	k8s.io/kubectl v0.20.1
	sigs.k8s.io/controller-runtime v0.6.2
)
