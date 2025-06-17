/*
Copyright 2023 The Kubernetes Authors.

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

package apimachinery

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/features"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/dynamic"
	clientfeatures "k8s.io/client-go/features"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/watchlist"
	"k8s.io/component-base/featuregate"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	"k8s.io/kubernetes/test/e2e/framework"
)

var _ = SIGDescribe("API Streaming (aka. WatchList)", framework.WithFeatureGate(features.WatchList), framework.WithSerial(), func() {
	f := framework.NewDefaultFramework("watchlist")
	ginkgo.It("should be requested by informers when WatchListClient is enabled", func(ctx context.Context) {
		featuregatetesting.SetFeatureGateDuringTest(ginkgo.GinkgoTB(), utilfeature.DefaultFeatureGate, featuregate.Feature(clientfeatures.WatchListClient), true)
		stopCh := make(chan struct{})
		defer close(stopCh)

		secretInformer := cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
					return nil, fmt.Errorf("unexpected list call")
				},
				WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
					options.LabelSelector = "watchlist=true"
					return f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Watch(context.TODO(), options)
				},
			},
			&v1.Secret{},
			time.Duration(0),
			nil,
		)

		expectedSecrets := addWellKnownSecrets(ctx, f)

		ginkgo.By("Starting the secret informer")
		go secretInformer.Run(stopCh)

		ginkgo.By("Waiting until the secret informer is fully synchronised")
		err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, false, func(context.Context) (done bool, err error) {
			return secretInformer.HasSynced(), nil
		})
		framework.ExpectNoError(err, "Failed waiting for the secret informer in %s namespace to be synced", f.Namespace.Namespace)

		ginkgo.By("Verifying if the secret informer was properly synchronised")
		verifyStore(ctx, expectedSecrets, secretInformer.GetStore())

		ginkgo.By("Modifying a secret and checking if the update was picked up by the secret informer")
		secret, err := f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Get(ctx, "secret-1", metav1.GetOptions{})
		framework.ExpectNoError(err)
		secret.StringData = map[string]string{"foo": "bar"}
		secret, err = f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Update(ctx, secret, metav1.UpdateOptions{})
		framework.ExpectNoError(err)

		expectedSecrets[0] = secret
		verifyStore(ctx, expectedSecrets, secretInformer.GetStore())
	})
	ginkgo.It("should NOT be requested by client-go's List method when WatchListClient is enabled", func(ctx context.Context) {
		featuregatetesting.SetFeatureGateDuringTest(ginkgo.GinkgoTB(), utilfeature.DefaultFeatureGate, featuregate.Feature(clientfeatures.WatchListClient), true)

		expectedSecrets := addWellKnownSecrets(ctx, f)

		rt, clientConfig := clientConfigWithRoundTripper(f)
		wrappedKubeClient, err := kubernetes.NewForConfig(clientConfig)
		framework.ExpectNoError(err)

		ginkgo.By("Getting secrets from the server")
		secretList, err := wrappedKubeClient.CoreV1().Secrets(f.Namespace.Name).List(ctx, metav1.ListOptions{LabelSelector: "watchlist=true"})
		framework.ExpectNoError(err)

		ginkgo.By("Verifying retrieved secrets")
		actualSecrets := secretList.Items
		gomega.Expect(cmp.Equal(expectedSecrets, toSecretPointerSlice(actualSecrets))).To(gomega.BeTrueBecause("data received via list must match the added data"))

		ginkgo.By("Verifying if expected requests were sent to the server")
		expectedRequestsMadeByKubeClient := []string{expectedListRequestMadeByClient}
		gomega.Expect(rt.actualRequests).To(gomega.Equal(expectedRequestsMadeByKubeClient))
	})
	ginkgo.It("should NOT be requested by dynamic client's List method when WatchListClient is enabled", func(ctx context.Context) {
		featuregatetesting.SetFeatureGateDuringTest(ginkgo.GinkgoTB(), utilfeature.DefaultFeatureGate, featuregate.Feature(clientfeatures.WatchListClient), true)

		ginkgo.By(fmt.Sprintf("Adding 5 secrets to %s namespace", f.Namespace.Name))
		expectedSecrets := addWellKnownUnstructuredSecrets(ctx, f)

		rt, clientConfig := clientConfigWithRoundTripper(f)
		wrappedDynamicClient, err := dynamic.NewForConfig(clientConfig)
		framework.ExpectNoError(err)

		ginkgo.By("Getting secrets from the server")
		secretList, err := wrappedDynamicClient.Resource(v1.SchemeGroupVersion.WithResource("secrets")).Namespace(f.Namespace.Name).List(ctx, metav1.ListOptions{LabelSelector: "watchlist=true"})
		framework.ExpectNoError(err)

		ginkgo.By("verifying retrieved secrets")
		actualSecrets := secretList.Items
		gomega.Expect(cmp.Equal(expectedSecrets, toSecretPointerSlice(actualSecrets))).To(gomega.BeTrueBecause("data received via list must match the added data"))
		gomega.Expect(secretList.GetObjectKind().GroupVersionKind()).To(gomega.Equal(v1.SchemeGroupVersion.WithKind("SecretList")))

		ginkgo.By("Verifying if expected requests were sent to the server")
		expectedRequestsMadeByDynamicClient := []string{expectedListRequestMadeByClient}
		gomega.Expect(rt.actualRequests).To(gomega.Equal(expectedRequestsMadeByDynamicClient))
	})
	ginkgo.It("should NOT be requested by metadata client's List method when WatchListClient is enabled", func(ctx context.Context) {
		featuregatetesting.SetFeatureGateDuringTest(ginkgo.GinkgoTB(), utilfeature.DefaultFeatureGate, featuregate.Feature(clientfeatures.WatchListClient), true)

		metaClient, err := metadata.NewForConfig(f.ClientConfig())
		framework.ExpectNoError(err)
		expectedMetaSecrets := []metav1.PartialObjectMetadata{}
		for _, addedSecret := range addWellKnownSecrets(ctx, f) {
			addedSecretMeta, err := metaClient.Resource(v1.SchemeGroupVersion.WithResource("secrets")).Namespace(f.Namespace.Name).Get(ctx, addedSecret.Name, metav1.GetOptions{})
			framework.ExpectNoError(err)
			expectedMetaSecrets = append(expectedMetaSecrets, *addedSecretMeta)
		}

		rt, clientConfig := clientConfigWithRoundTripper(f)
		wrappedMetaClient, err := metadata.NewForConfig(clientConfig)
		framework.ExpectNoError(err)

		ginkgo.By("Getting secrets metadata from the server")
		secretMetaList, err := wrappedMetaClient.Resource(v1.SchemeGroupVersion.WithResource("secrets")).Namespace(f.Namespace.Name).List(ctx, metav1.ListOptions{LabelSelector: "watchlist=true"})
		framework.ExpectNoError(err)

		ginkgo.By("verifying retrieved secrets")
		actualMetaSecrets := secretMetaList.Items
		gomega.Expect(cmp.Equal(expectedMetaSecrets, actualMetaSecrets)).To(gomega.BeTrueBecause("data received via list must match the added data"))

		ginkgo.By("Verifying if expected requests were sent to the server")
		expectedRequestsMadeByMetaClient := []string{expectedListRequestMadeByClient}
		gomega.Expect(rt.actualRequests).To(gomega.Equal(expectedRequestsMadeByMetaClient))
	})

	// Validates unsupported Accept headers in WatchList.
	// Sets AcceptContentType to "application/json;as=Table", which the API doesn't support, returning a 406 error.
	ginkgo.It("doesn't support receiving resources as Tables", func(ctx context.Context) {
		featuregatetesting.SetFeatureGateDuringTest(ginkgo.GinkgoTB(), utilfeature.DefaultFeatureGate, featuregate.Feature(clientfeatures.WatchListClient), true)

		modifiedClientConfig := dynamic.ConfigFor(f.ClientConfig())
		modifiedClientConfig.AcceptContentTypes = strings.Join([]string{
			fmt.Sprintf("application/json;as=Table;v=%s;g=%s", metav1.SchemeGroupVersion.Version, metav1.GroupName),
		}, ",")
		modifiedClientConfig.GroupVersion = &v1.SchemeGroupVersion
		restClient, err := rest.RESTClientFor(modifiedClientConfig)
		framework.ExpectNoError(err)
		dynamicClient := dynamic.New(restClient)

		opts, hasPreparedOptions, err := watchlist.PrepareWatchListOptionsFromListOptions(metav1.ListOptions{})
		framework.ExpectNoError(err)
		gomega.Expect(hasPreparedOptions).To(gomega.BeTrueBecause("it should be possible to prepare watchlist opts from an empty ListOptions"))

		_, err = dynamicClient.Resource(v1.SchemeGroupVersion.WithResource("secrets")).Namespace("default").Watch(ctx, opts)
		gomega.Expect(err).To(gomega.HaveOccurred())
		gomega.Expect(err.(apierrors.APIStatus)).To(gomega.HaveField("Status().Code", gomega.Equal(int32(406))))
	})

	// Sets AcceptContentType to both "application/json;as=Table" and "application/json".
	// Unlike the previous test, no 406 error occurs, as the API falls back to "application/json" and returns a valid response.
	ginkgo.It("falls backs to supported content type when when receiving resources as Tables was requested", func(ctx context.Context) {
		featuregatetesting.SetFeatureGateDuringTest(ginkgo.GinkgoTB(), utilfeature.DefaultFeatureGate, featuregate.Feature(clientfeatures.WatchListClient), true)

		modifiedClientConfig := dynamic.ConfigFor(f.ClientConfig())
		modifiedClientConfig.AcceptContentTypes = strings.Join([]string{
			fmt.Sprintf("application/json;as=Table;v=%s;g=%s", metav1.SchemeGroupVersion.Version, metav1.GroupName),
			"application/json",
		}, ",")
		modifiedClientConfig.GroupVersion = &v1.SchemeGroupVersion
		restClient, err := rest.RESTClientFor(modifiedClientConfig)
		framework.ExpectNoError(err)
		dynamicClient := dynamic.New(restClient)

		stopCh := make(chan struct{})
		defer close(stopCh)

		secretInformer := cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
					return nil, fmt.Errorf("unexpected list call")
				},
				WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
					options.LabelSelector = "watchlist=true"
					return dynamicClient.Resource(v1.SchemeGroupVersion.WithResource("secrets")).Namespace(f.Namespace.Name).Watch(context.TODO(), options)
				},
			},
			&unstructured.Unstructured{},
			time.Duration(0),
			nil,
		)

		expectedSecrets := addWellKnownUnstructuredSecrets(ctx, f)

		ginkgo.By("Starting the secret informer")
		go secretInformer.Run(stopCh)

		ginkgo.By("Waiting until the secret informer is fully synchronised")
		err = wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, false, func(context.Context) (done bool, err error) {
			return secretInformer.HasSynced(), nil
		})
		framework.ExpectNoError(err, "Failed waiting for the secret informer in %s namespace to be synced", f.Namespace.Namespace)

		ginkgo.By("Verifying if the secret informer was properly synchronised")
		verifyStore[unstructured.Unstructured](ctx, expectedSecrets, secretInformer.GetStore())
	})
})

type roundTripper struct {
	actualRequests         []string
	actualResponseStatuses []string
	delegate               http.RoundTripper
}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	r.actualRequests = append(r.actualRequests, req.URL.RawQuery)
	rsp, err := r.delegate.RoundTrip(req)
	if rsp != nil {
		r.actualResponseStatuses = append(r.actualResponseStatuses, rsp.Status)
	}
	return rsp, err
}

func (r *roundTripper) Wrap(delegate http.RoundTripper) http.RoundTripper {
	r.delegate = delegate
	return r
}

func clientConfigWithRoundTripper(f *framework.Framework) (*roundTripper, *rest.Config) {
	clientConfig := f.ClientConfig()
	rt := &roundTripper{}
	clientConfig.Wrap(rt.Wrap)

	return rt, clientConfig
}

func verifyStore[T any](ctx context.Context, expectedSecrets []*T, store cache.Store) {
	err := wait.PollUntilContextTimeout(ctx, 100*time.Millisecond, 30*time.Second, true, func(ctx context.Context) (done bool, err error) {
		ginkgo.By("Comparing secrets retrieved directly from the server with the ones that have been streamed to the secret informer")

		expectedSecretsAsMetaObject, err := toMetaObjectSlice(expectedSecrets)
		framework.ExpectNoError(err)
		actualSecretsAsMetaObject, err := toMetaObjectSlice(store.List())
		framework.ExpectNoError(err)

		sort.Sort(byName(expectedSecretsAsMetaObject))
		sort.Sort(byName(actualSecretsAsMetaObject))

		return cmp.Equal(expectedSecretsAsMetaObject, actualSecretsAsMetaObject), nil
	})
	framework.ExpectNoError(err)
}

var expectedListRequestMadeByClient = func() string {
	params := url.Values{}
	params.Add("labelSelector", "watchlist=true")
	return params.Encode()
}()

func addWellKnownSecrets(ctx context.Context, f *framework.Framework) []*v1.Secret {
	ginkgo.By(fmt.Sprintf("Adding 5 secrets to %s namespace", f.Namespace.Name))
	var secrets []*v1.Secret
	for i := 1; i <= 5; i++ {
		secret, err := f.ClientSet.CoreV1().Secrets(f.Namespace.Name).Create(ctx, newSecret(fmt.Sprintf("secret-%d", i)), metav1.CreateOptions{})
		framework.ExpectNoError(err)
		secrets = append(secrets, secret)
	}
	return secrets
}

// addWellKnownUnstructuredSecrets exists because secrets from addWellKnownSecrets
// don't have type info and cannot be converted.
func addWellKnownUnstructuredSecrets(ctx context.Context, f *framework.Framework) []*unstructured.Unstructured {
	var secrets []*unstructured.Unstructured
	for i := 1; i <= 5; i++ {
		unstructuredSecret, err := runtime.DefaultUnstructuredConverter.ToUnstructured(newSecret(fmt.Sprintf("secret-%d", i)))
		framework.ExpectNoError(err)
		secret, err := f.DynamicClient.Resource(v1.SchemeGroupVersion.WithResource("secrets")).Namespace(f.Namespace.Name).Create(ctx, &unstructured.Unstructured{Object: unstructuredSecret}, metav1.CreateOptions{})
		framework.ExpectNoError(err)
		secrets = append(secrets, secret)
	}
	return secrets
}

type byName []metav1.Object

func (a byName) Len() int           { return len(a) }
func (a byName) Less(i, j int) bool { return a[i].GetName() < a[j].GetName() }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func newSecret(name string) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: map[string]string{"watchlist": "true"},
		},
	}
}

func toMetaObjectSlice[T any](s []T) ([]metav1.Object, error) {
	result := make([]metav1.Object, len(s))
	for i, v := range s {
		m, err := meta.Accessor(v)
		if err != nil {
			return nil, err
		}
		result[i] = m
	}
	return result, nil
}

func toSecretPointerSlice[T any](items []T) []*T {
	result := make([]*T, 0, len(items))
	for i := range items {
		result = append(result, &items[i])
	}
	return result
}
