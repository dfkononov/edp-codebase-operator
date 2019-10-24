package vcs

import (
	"fmt"
	"github.com/epmd-edp/codebase-operator/v2/pkg/model"
	"github.com/epmd-edp/codebase-operator/v2/pkg/vcs/impl/bitbucket"
	"github.com/epmd-edp/codebase-operator/v2/pkg/vcs/impl/gitlab"
	"log"
)

type VCS interface {
	CheckProjectExist(groupPath, projectName string) (*bool, error)
	CreateProject(groupPath, projectName string) (string, error)
	GetRepositorySshUrl(groupPath, projectName string) (string, error)
}

func CreateVCSClient(vcsToolName model.VCSTool, url string, username string, password string) (VCS, error) {
	switch vcsToolName {
	case model.GitLab:
		log.Print("Creating VCS for GitLab implementation...")
		vcsClient := gitlab.GitLab{}
		err := vcsClient.Init(url, username, password)
		if err != nil {
			return nil, err
		}
		return &vcsClient, nil
	case model.BitBucket:
		log.Print("Creating VCS for BitBucket implementation...")
		vcsClient := bitbucket.BitBucket{}
		err := vcsClient.Init(url, username, password)
		if err != nil {
			return nil, err
		}
		return &vcsClient, nil
	default:
		return nil, fmt.Errorf("invalid VCS tool. Currently we do not support %v", vcsToolName)
	}
}
