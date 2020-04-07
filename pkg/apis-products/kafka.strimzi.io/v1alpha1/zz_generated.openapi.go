// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.Kafka":       schema_pkg_apis_kafkastrimziio_v1alpha1_Kafka(ref),
		"github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpec":   schema_pkg_apis_kafkastrimziio_v1alpha1_KafkaSpec(ref),
		"github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaStatus": schema_pkg_apis_kafkastrimziio_v1alpha1_KafkaStatus(ref),
	}
}

func schema_pkg_apis_kafkastrimziio_v1alpha1_Kafka(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Installation is the Schema for the installations API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpec", "github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kafkastrimziio_v1alpha1_KafkaSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InstallationSpec defines the desired state of Installation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kafka": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpecKafka"),
						},
					},
					"zookeeper": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpecZookeeper"),
						},
					},
					"entityOperator": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpecEntityOperator"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpecEntityOperator", "github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpecKafka", "github.com/integr8ly/integreatly-operator/pkg/apis-products/kafka.strimzi.io/v1alpha1.KafkaSpecZookeeper"},
	}
}

func schema_pkg_apis_kafkastrimziio_v1alpha1_KafkaStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InstallationStatus defines the observed state of Installation",
				Type:        []string{"object"},
			},
		},
	}
}