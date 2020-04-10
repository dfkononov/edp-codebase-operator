# How to Install Operator

EDP installation can be applied on two container orchestration platforms: OpenShift and Kubernetes.

_**NOTE:** Installation of operators is platform-independent, that is why there is a unified instruction for deploying._

### Prerequisites
1. Linux machine or Windows Subsystem for Linux instance with [Helm 3](https://helm.sh/docs/intro/install/) installed;
2. Cluster admin access to the cluster;
3. EDP project/namespace is deployed by following one of the instructions: [edp-install-openshift](https://github.com/epmd-edp/edp-install/blob/release-2.3/documentation/openshift_install_edp.md#edp-project) or [edp-install-kubernetes](https://github.com/epmd-edp/edp-install/blob/release-2.3/documentation/kubernetes_install_edp.md#edp-namespace).

### Installation
* Go to the [releases](https://github.com/epmd-edp/codebase-operator/releases) page of this repository, choose a version, download an archive and unzip it;

_**NOTE:** It is highly recommended to use the latest released version._

* Go to the unzipped directory and deploy operator:
```bash
helm install codebase-operator --namespace <edp_cicd_project> --set name=codebase-operator --set namespace=<edp_cicd_project> --set platform=<platform_type> --set image.name=epamedp/codebase-operator --set image.version=<operator_version> deploy-templates
```

- _<edp_cicd_project> - a namespace or a project name (in case of OpenShift) that is created by one of the instructions: [edp-install-openshift](https://github.com/epmd-edp/edp-install/blob/release-2.3/documentation/openshift_install_edp.md#edp-project) or [edp-install-kubernetes](https://github.com/epmd-edp/edp-install/blob/release-2.3/documentation/kubernetes_install_edp.md#edp-namespace);_ 

- _<platform_type> - a platform type that can be "kubernetes" or "openshift";_

- _<operator_version> - a selected release version tag for the operator from Docker Hub;_

* Check the <edp_cicd_project> namespace that should be in a pending state of creating a secret by indicating the following message: "Error: secrets "db-admin-console" not found". Such notification is a normal flow and it will be fixed during the EDP installation.