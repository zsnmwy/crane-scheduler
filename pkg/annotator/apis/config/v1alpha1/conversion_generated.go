//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	config "github.com/gocrane/crane-scheduler/pkg/annotator/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AnnotatorConfiguration)(nil), (*config.AnnotatorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AnnotatorConfiguration_To_config_AnnotatorConfiguration(a.(*AnnotatorConfiguration), b.(*config.AnnotatorConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.AnnotatorConfiguration)(nil), (*AnnotatorConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_AnnotatorConfiguration_To_v1alpha1_AnnotatorConfiguration(a.(*config.AnnotatorConfiguration), b.(*AnnotatorConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AnnotatorConfiguration_To_config_AnnotatorConfiguration(in *AnnotatorConfiguration, out *config.AnnotatorConfiguration, s conversion.Scope) error {
	out.BindingHeapSize = in.BindingHeapSize
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.PolicyConfigPath = in.PolicyConfigPath
	out.PrometheusAddr = in.PrometheusAddr
	if err := configv1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_v1alpha1_DebuggingConfiguration_To_config_DebuggingConfiguration(&in.Debugging, &out.Debugging, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_AnnotatorConfiguration_To_config_AnnotatorConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_AnnotatorConfiguration_To_config_AnnotatorConfiguration(in *AnnotatorConfiguration, out *config.AnnotatorConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_AnnotatorConfiguration_To_config_AnnotatorConfiguration(in, out, s)
}

func autoConvert_config_AnnotatorConfiguration_To_v1alpha1_AnnotatorConfiguration(in *config.AnnotatorConfiguration, out *AnnotatorConfiguration, s conversion.Scope) error {
	out.BindingHeapSize = in.BindingHeapSize
	out.ConcurrentSyncs = in.ConcurrentSyncs
	out.PolicyConfigPath = in.PolicyConfigPath
	out.PrometheusAddr = in.PrometheusAddr
	if err := configv1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_config_DebuggingConfiguration_To_v1alpha1_DebuggingConfiguration(&in.Debugging, &out.Debugging, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_AnnotatorConfiguration_To_v1alpha1_AnnotatorConfiguration is an autogenerated conversion function.
func Convert_config_AnnotatorConfiguration_To_v1alpha1_AnnotatorConfiguration(in *config.AnnotatorConfiguration, out *AnnotatorConfiguration, s conversion.Scope) error {
	return autoConvert_config_AnnotatorConfiguration_To_v1alpha1_AnnotatorConfiguration(in, out, s)
}
