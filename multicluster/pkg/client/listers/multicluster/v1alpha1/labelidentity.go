// Copyright 2024 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// LabelIdentityLister helps list LabelIdentities.
// All objects returned here must be treated as read-only.
type LabelIdentityLister interface {
	// List lists all LabelIdentities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LabelIdentity, err error)
	// Get retrieves the LabelIdentity from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LabelIdentity, error)
	LabelIdentityListerExpansion
}

// labelIdentityLister implements the LabelIdentityLister interface.
type labelIdentityLister struct {
	listers.ResourceIndexer[*v1alpha1.LabelIdentity]
}

// NewLabelIdentityLister returns a new LabelIdentityLister.
func NewLabelIdentityLister(indexer cache.Indexer) LabelIdentityLister {
	return &labelIdentityLister{listers.New[*v1alpha1.LabelIdentity](indexer, v1alpha1.Resource("labelidentity"))}
}
