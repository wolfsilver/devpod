// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ProjectLister helps list Projects.
// All objects returned here must be treated as read-only.
type ProjectLister interface {
	// List lists all Projects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Project, err error)
	// Get retrieves the Project from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Project, error)
	ProjectListerExpansion
}

// projectLister implements the ProjectLister interface.
type projectLister struct {
	listers.ResourceIndexer[*v1.Project]
}

// NewProjectLister returns a new ProjectLister.
func NewProjectLister(indexer cache.Indexer) ProjectLister {
	return &projectLister{listers.New[*v1.Project](indexer, v1.Resource("project"))}
}
