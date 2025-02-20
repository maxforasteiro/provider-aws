/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AlertManagerDefinitionDescription struct {
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
}

// +kubebuilder:skipversion
type RuleGroupsNamespaceDescription struct {
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// The list of tags assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type RuleGroupsNamespaceSummary struct {
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	ModifiedAt *metav1.Time `json:"modifiedAt,omitempty"`
	// The list of tags assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type WorkspaceDescription struct {
	// A user-assigned workspace alias.
	Alias *string `json:"alias,omitempty"`
	// An ARN identifying a Workspace.
	ARN *string `json:"arn,omitempty"`

	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	PrometheusEndpoint *string `json:"prometheusEndpoint,omitempty"`
	// Represents the status of a workspace.
	Status *WorkspaceStatus_SDK `json:"status,omitempty"`
	// The list of tags assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty"`
	// A workspace ID.
	WorkspaceID *string `json:"workspaceID,omitempty"`
}

// +kubebuilder:skipversion
type WorkspaceStatus_SDK struct {
	// State of a workspace.
	StatusCode *string `json:"statusCode,omitempty"`
}

// +kubebuilder:skipversion
type WorkspaceSummary struct {
	// A user-assigned workspace alias.
	Alias *string `json:"alias,omitempty"`
	// An ARN identifying a Workspace.
	ARN *string `json:"arn,omitempty"`

	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
	// Represents the status of a workspace.
	Status *WorkspaceStatus_SDK `json:"status,omitempty"`
	// The list of tags assigned to the resource.
	Tags map[string]*string `json:"tags,omitempty"`
	// A workspace ID.
	WorkspaceID *string `json:"workspaceID,omitempty"`
}
