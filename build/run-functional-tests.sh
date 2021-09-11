#!/bin/bash
# Copyright Contributors to the Open Cluster Management project

# set -x
# set -e
TEST_DIR=test/functional
TEST_RESULT_DIR=$TEST_DIR/tmp
ERROR_REPORT=""
CLUSTER_NAME=$PROJECT_NAME-functional-test
export KUBECONFIG=$TEST_DIR/tmp/kind.yaml

rm -rf $TEST_RESULT_DIR
mkdir -p $TEST_RESULT_DIR

kind create cluster --name ${CLUSTER_NAME}
kind get kubeconfig --name ${CLUSTER_NAME} > ${TEST_DIR}/tmp/kind.yaml

# Configure the kind cluster
kubectl apply -f $TEST_DIR/resources/clustermanager_crd.yaml
kubectl apply -f $TEST_DIR/resources/multiclusterhubs_crd.yaml
kubectl apply -f $TEST_DIR/resources/hive_v1_clusterdeployment_crd.yaml
kubectl apply -f $TEST_DIR/resources/open-cluster-management-ns.yaml
kubectl apply -f $TEST_DIR/resources/openshift-config-ns.yaml
kubectl apply -f $TEST_DIR/resources/clustermanager_cr.yaml
kubectl apply -f $TEST_DIR/resources/multiclusterhubs_cr.yaml
kubectl apply -f $TEST_DIR/resources/pull_secret_cr.yaml
kubectl apply -f $TEST_DIR/resources/acm_config_cm.yaml

go run test/functional/kubestatus.go update multiclusterhub -n default --group operator.open-cluster-management.io --version v1 --resource multiclusterhubs --status "currentVersion: 2.4.0"

echo "Test cm version"
cm version
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm version AWS failed\n"
fi

echo "Test cm create cluster AWS"
cm create cluster --values $TEST_DIR/create/cluster/aws_values.yaml --dry-run --output-file $TEST_RESULT_DIR/aws_result.yaml
diff -u $TEST_DIR/create/cluster/aws_result.yaml $TEST_RESULT_DIR/aws_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm create cluster AWS failed\n"
fi

echo "Test cm create cluster AWS with labels"
cm create cluster --values $TEST_DIR/create/cluster/aws_values_with_labels.yaml --dry-run --output-file $TEST_RESULT_DIR/aws_result_with_labels.yaml
diff -u $TEST_DIR/create/cluster/aws_result_with_labels.yaml $TEST_RESULT_DIR/aws_result_with_labels.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm create cluster AWS failed\n"
fi

echo "Test cm create cluster Azure"
cm create cluster --values $TEST_DIR/create/cluster/azure_values.yaml --dry-run --output-file $TEST_RESULT_DIR/azure_result.yaml
diff -u $TEST_DIR/create/cluster/azure_result.yaml $TEST_RESULT_DIR/azure_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm create cluster Azure failed\n"
fi

echo "Test cm create cluster GCP"
cm create cluster --values $TEST_DIR/create/cluster/gcp_values.yaml --dry-run --output-file $TEST_RESULT_DIR/gcp_result.yaml
diff -u $TEST_DIR/create/cluster/gcp_result.yaml $TEST_RESULT_DIR/gcp_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm create cluster GCP failed\n"
fi

echo "Test cm create cluster OpenStack"
cm create cluster --values $TEST_DIR/create/cluster/openstack_values.yaml --dry-run --output-file $TEST_RESULT_DIR/openstack_result.yaml
diff -u $TEST_DIR/create/cluster/openstack_result.yaml $TEST_RESULT_DIR/openstack_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm create cluster OpenStack failed\n"
fi

echo "Test cm create cluster vSphere"
cm create cluster --values $TEST_DIR/create/cluster/vsphere_values.yaml --dry-run --output-file $TEST_RESULT_DIR/vsphere_result.yaml
diff -u $TEST_DIR/create/cluster/vsphere_result.yaml $TEST_RESULT_DIR/vsphere_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm create cluster vSphere failed\n"
fi

echo "Test cm attach cluster manual"
cm attach cluster --values $TEST_DIR/attach/cluster/manual_values.yaml --dry-run --output-file $TEST_RESULT_DIR/manual_result.yaml --import-file $TEST_RESULT_DIR/manual_import_result.yaml
diff -u $TEST_DIR/attach/cluster/manual_result.yaml $TEST_RESULT_DIR/manual_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm attach cluster manual failed\n"
fi

echo "Test cm attach cluster manual with labels"
cm attach cluster --values $TEST_DIR/attach/cluster/manual_values_with_labels.yaml --dry-run --output-file $TEST_RESULT_DIR/manual_result_with_labels.yaml --import-file $TEST_RESULT_DIR/manual_import_result_with_labels.yaml
diff -u $TEST_DIR/attach/cluster/manual_result_with_labels.yaml $TEST_RESULT_DIR/manual_result_with_labels.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm attach cluster manual failed\n"
fi

echo "Test cm attach cluster no values.yaml"
cm attach cluster --cluster mycluster --dry-run --output-file $TEST_RESULT_DIR/manual_no_values_result.yaml  --import-file $TEST_RESULT_DIR/manual_import_no_values_result.yaml
diff -u $TEST_DIR/attach/cluster/manual_no_values_result.yaml $TEST_RESULT_DIR/manual_no_values_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm attach cluster manual failed without values.yaml\n"
fi


echo "Test cm attach cluster kubeconfig"
cm attach cluster --values $TEST_DIR/attach/cluster/kubeconfig_values.yaml --dry-run --output-file $TEST_RESULT_DIR/kubeconfig_result.yaml
diff -u $TEST_DIR/attach/cluster/kubeconfig_result.yaml $TEST_RESULT_DIR/kubeconfig_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm attach cluster kubeconfig failed\n"
fi

echo "Test cm attach cluster kubeconfig no values.yaml with kubeconfig"
cm attach cluster --cluster mycluster --cluster-kubeconfig $TEST_DIR/attach/cluster/fake-kubeconfig.yaml --dry-run --output-file $TEST_RESULT_DIR/kubeconfig_no_values_result.yaml
diff -u $TEST_DIR/attach/cluster/kubeconfig_no_values_result.yaml $TEST_RESULT_DIR/kubeconfig_no_values_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm attach cluster kubeconfig failed without values.yaml\n"
fi

echo "Test cm attach cluster token"
cm attach cluster --values $TEST_DIR/attach/cluster/token_values.yaml --dry-run --output-file $TEST_RESULT_DIR/token_result.yaml
diff -u $TEST_DIR/attach/cluster/token_result.yaml $TEST_RESULT_DIR/token_result.yaml
if [ $? != 0 ]
then
   ERROR_REPORT=$ERROR_REPORT+"cm attach cluster token failed\n"
fi

if [ -z "$ERROR_REPORT" ]
then
    echo "Success"
else
    echo -e "\n\nErrors\n======\n"$ERROR_REPORT
    exit 1
fi

kind delete cluster --name $CLUSTER_NAME
