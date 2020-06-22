/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	namespaceresourcesinjectorv1alpha1 "github.com/blakelead/nsinjector/pkg/apis/namespaceresourcesinjector/v1alpha1"
	versioned "github.com/blakelead/nsinjector/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/blakelead/nsinjector/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/blakelead/nsinjector/pkg/generated/listers/namespaceresourcesinjector/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NamespaceResourcesInjectorInformer provides access to a shared informer and lister for
// NamespaceResourcesInjectors.
type NamespaceResourcesInjectorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NamespaceResourcesInjectorLister
}

type namespaceResourcesInjectorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNamespaceResourcesInjectorInformer constructs a new informer for NamespaceResourcesInjector type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNamespaceResourcesInjectorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNamespaceResourcesInjectorInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNamespaceResourcesInjectorInformer constructs a new informer for NamespaceResourcesInjector type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNamespaceResourcesInjectorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BlakeleadV1alpha1().NamespaceResourcesInjectors().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BlakeleadV1alpha1().NamespaceResourcesInjectors().Watch(context.TODO(), options)
			},
		},
		&namespaceresourcesinjectorv1alpha1.NamespaceResourcesInjector{},
		resyncPeriod,
		indexers,
	)
}

func (f *namespaceResourcesInjectorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNamespaceResourcesInjectorInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *namespaceResourcesInjectorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&namespaceresourcesinjectorv1alpha1.NamespaceResourcesInjector{}, f.defaultInformer)
}

func (f *namespaceResourcesInjectorInformer) Lister() v1alpha1.NamespaceResourcesInjectorLister {
	return v1alpha1.NewNamespaceResourcesInjectorLister(f.Informer().GetIndexer())
}
