// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ClusterAccessLister helps list ClusterAccesses.
// All objects returned here must be treated as read-only.
type ClusterAccessLister interface {
	// List lists all ClusterAccesses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterAccess, err error)
	// Get retrieves the ClusterAccess from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterAccess, error)
	ClusterAccessListerExpansion
}

// clusterAccessLister implements the ClusterAccessLister interface.
type clusterAccessLister struct {
	listers.ResourceIndexer[*v1.ClusterAccess]
}

// NewClusterAccessLister returns a new ClusterAccessLister.
func NewClusterAccessLister(indexer cache.Indexer) ClusterAccessLister {
	return &clusterAccessLister{listers.New[*v1.ClusterAccess](indexer, v1.Resource("clusteraccess"))}
}