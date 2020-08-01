// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/imageregistry/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImagePrunerLister helps list ImagePruners.
// All objects returned here must be treated as read-only.
type ImagePrunerLister interface {
	// List lists all ImagePruners in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ImagePruner, err error)
	// Get retrieves the ImagePruner from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ImagePruner, error)
	ImagePrunerListerExpansion
}

// imagePrunerLister implements the ImagePrunerLister interface.
type imagePrunerLister struct {
	indexer cache.Indexer
}

// NewImagePrunerLister returns a new ImagePrunerLister.
func NewImagePrunerLister(indexer cache.Indexer) ImagePrunerLister {
	return &imagePrunerLister{indexer: indexer}
}

// List lists all ImagePruners in the indexer.
func (s *imagePrunerLister) List(selector labels.Selector) (ret []*v1.ImagePruner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ImagePruner))
	})
	return ret, err
}

// Get retrieves the ImagePruner from the index for a given name.
func (s *imagePrunerLister) Get(name string) (*v1.ImagePruner, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("imagepruner"), name)
	}
	return obj.(*v1.ImagePruner), nil
}
