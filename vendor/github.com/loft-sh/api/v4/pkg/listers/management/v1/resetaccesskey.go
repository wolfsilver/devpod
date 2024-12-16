// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ResetAccessKeyLister helps list ResetAccessKeys.
// All objects returned here must be treated as read-only.
type ResetAccessKeyLister interface {
	// List lists all ResetAccessKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ResetAccessKey, err error)
	// Get retrieves the ResetAccessKey from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ResetAccessKey, error)
	ResetAccessKeyListerExpansion
}

// resetAccessKeyLister implements the ResetAccessKeyLister interface.
type resetAccessKeyLister struct {
	listers.ResourceIndexer[*v1.ResetAccessKey]
}

// NewResetAccessKeyLister returns a new ResetAccessKeyLister.
func NewResetAccessKeyLister(indexer cache.Indexer) ResetAccessKeyLister {
	return &resetAccessKeyLister{listers.New[*v1.ResetAccessKey](indexer, v1.Resource("resetaccesskey"))}
}