package plugins

import (
	jenkinsv1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
)

const (
	// GitOpsVersion the version of the jx gitops plugin
	GitOpsVersion = "0.0.54"

	// ProjectVersion the version of the jx project plugin
	ProjectVersion = "0.0.13"

	// SecretVersion the version of the jx secret plugin
	SecretVersion = "0.0.24"
)

var (
	// Plugins default plugins
	Plugins = []jenkinsv1.Plugin{
		CreateJXPlugin("gitops", GitOpsVersion),
		CreateJXPlugin("project", ProjectVersion),
		CreateJXPlugin("secret", SecretVersion),
	}
)