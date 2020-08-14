package decoder

import (
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

func GetDecoder() runtime.Decoder {
	apiextensionsv1beta1.AddToScheme(scheme.Scheme)
	networkingv1alpha3.AddToScheme(scheme.Scheme)
	networkingv1beta1.AddToScheme(scheme.Scheme)
	securityv1beta1.AddToScheme(scheme.Scheme)

	// TODO: Look at extending UniversalDeserializer Scheme with CRDs.
	return scheme.Codecs.UniversalDeserializer()
}
