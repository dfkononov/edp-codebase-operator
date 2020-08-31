package handler

import edpv1alpha1 "github.com/epmd-edp/codebase-operator/v2/pkg/apis/edp/v1alpha1"

type CodebaseBranchHandler interface {
	ServeRequest(c *edpv1alpha1.CodebaseBranch) error
}