// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package file


type FileSource struct {
	// Add this content to the archive with `filename` as the filename.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#content File#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// Set this as the filename when declaring a `source`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#filename File#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

