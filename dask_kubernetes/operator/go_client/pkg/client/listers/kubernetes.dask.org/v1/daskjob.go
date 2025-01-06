// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/dask/dask-kubernetes/v2025/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DaskJobLister helps list DaskJobs.
// All objects returned here must be treated as read-only.
type DaskJobLister interface {
	// List lists all DaskJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DaskJob, err error)
	// DaskJobs returns an object that can list and get DaskJobs.
	DaskJobs(namespace string) DaskJobNamespaceLister
	DaskJobListerExpansion
}

// daskJobLister implements the DaskJobLister interface.
type daskJobLister struct {
	indexer cache.Indexer
}

// NewDaskJobLister returns a new DaskJobLister.
func NewDaskJobLister(indexer cache.Indexer) DaskJobLister {
	return &daskJobLister{indexer: indexer}
}

// List lists all DaskJobs in the indexer.
func (s *daskJobLister) List(selector labels.Selector) (ret []*v1.DaskJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaskJob))
	})
	return ret, err
}

// DaskJobs returns an object that can list and get DaskJobs.
func (s *daskJobLister) DaskJobs(namespace string) DaskJobNamespaceLister {
	return daskJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DaskJobNamespaceLister helps list and get DaskJobs.
// All objects returned here must be treated as read-only.
type DaskJobNamespaceLister interface {
	// List lists all DaskJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DaskJob, err error)
	// Get retrieves the DaskJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DaskJob, error)
	DaskJobNamespaceListerExpansion
}

// daskJobNamespaceLister implements the DaskJobNamespaceLister
// interface.
type daskJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DaskJobs in the indexer for a given namespace.
func (s daskJobNamespaceLister) List(selector labels.Selector) (ret []*v1.DaskJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaskJob))
	})
	return ret, err
}

// Get retrieves the DaskJob from the indexer for a given namespace and name.
func (s daskJobNamespaceLister) Get(name string) (*v1.DaskJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("daskjob"), name)
	}
	return obj.(*v1.DaskJob), nil
}
