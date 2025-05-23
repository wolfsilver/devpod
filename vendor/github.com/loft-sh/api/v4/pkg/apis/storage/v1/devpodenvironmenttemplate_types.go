package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPodWorkspaceEnvironmentSource
// +k8s:openapi-gen=true
type DevPodEnvironmentTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevPodEnvironmentTemplateSpec   `json:"spec,omitempty"`
	Status DevPodEnvironmentTemplateStatus `json:"status,omitempty"`
}

// DevPodEnvironmentTemplateStatus holds the status
type DevPodEnvironmentTemplateStatus struct {
}

func (a *DevPodEnvironmentTemplate) GetVersions() []VersionAccessor {
	var retVersions []VersionAccessor
	for _, v := range a.Spec.Versions {
		b := v
		retVersions = append(retVersions, &b)
	}

	return retVersions
}

func (a *DevPodEnvironmentTemplateVersion) GetVersion() string {
	return a.Version
}

func (a *DevPodEnvironmentTemplate) GetOwner() *UserOrTeam {
	return a.Spec.Owner
}

func (a *DevPodEnvironmentTemplate) SetOwner(userOrTeam *UserOrTeam) {
	a.Spec.Owner = userOrTeam
}

func (a *DevPodEnvironmentTemplate) GetAccess() []Access {
	return a.Spec.Access
}

func (a *DevPodEnvironmentTemplate) SetAccess(access []Access) {
	a.Spec.Access = access
}

type DevPodEnvironmentTemplateSpec struct {
	// DisplayName is the name that should be displayed in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description describes the environment template
	// +optional
	Description string `json:"description,omitempty"`

	// Owner holds the owner of this object
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`

	// Access to the DevPod machine instance object itself
	// +optional
	Access []Access `json:"access,omitempty"`

	// Template is the inline template to use for DevPod environments
	// +optional
	Template *DevPodEnvironmentTemplateDefinition `json:"template,omitempty"`

	// Versions are different versions of the template that can be referenced as well
	// +optional
	Versions []DevPodEnvironmentTemplateVersion `json:"versions,omitempty"`
}

type DevPodEnvironmentTemplateDefinition struct {
	// Git holds configuration for git environment spec source
	// +optional
	Git *GitEnvironmentTemplate `json:"git,omitempty"`

	// Inline holds an inline devcontainer.json definition
	// +optional
	Inline string `json:"inline,omitempty"`

	// WorkspaceRepositoryCloneStrategy determines how the workspaces git repository will be checked out in the pod if the workspace is git based
	// +optional
	WorkspaceRepositoryCloneStrategy GitCloneStrategy `json:"workspaceRepositoryCloneStrategy,omitempty"`

	// WorkspaceRepositorySkipLFS specifies if git lfs will be skipped when cloning the repository into the workspace
	// +optional
	WorkspaceRepositorySkipLFS bool `json:"workspaceRepositorySkipLFS,omitempty"`
}

// GitEnvironmentTemplate stores configuration of Git environment template source
type GitEnvironmentTemplate struct {
	// Repository stores repository URL for Git environment spec source
	Repository string `json:"repository"`

	// Revision stores revision to checkout in repository
	// +optional
	Revision string `json:"revision,omitempty"`

	// SubPath stores subpath within Repositor where environment spec is
	// +optional
	SubPath string `json:"subpath,omitempty"`

	// UseProjectGitCredentials specifies if the project git credentials should be used instead of local ones for this environment
	// +optional
	UseProjectGitCredentials bool `json:"useProjectGitCredentials,omitempty"`
}

type DevPodEnvironmentTemplateVersion struct {
	// Template holds the environment template definition
	// +optional
	Template DevPodEnvironmentTemplateDefinition `json:"template,omitempty"`

	// Version is the version. Needs to be in X.X.X format.
	// +optional
	Version string `json:"version,omitempty"`
}

// +enum
type GitCloneStrategy string

// WARN: Need to match https://github.com/loft-sh/devpod/pkg/git/clone.go
const (
	FullCloneStrategy     GitCloneStrategy = ""
	BloblessCloneStrategy GitCloneStrategy = "blobless"
	TreelessCloneStrategy GitCloneStrategy = "treeless"
	ShallowCloneStrategy  GitCloneStrategy = "shallow"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DevPodEnvironmentTemplateList contains a list of DevPodEnvironmentTemplate objects
type DevPodEnvironmentTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevPodEnvironmentTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DevPodEnvironmentTemplate{}, &DevPodEnvironmentTemplateList{})
}
