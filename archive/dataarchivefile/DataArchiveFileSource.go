// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataarchivefile


type DataArchiveFileSource struct {
	// Add this content to the archive with `filename` as the filename.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.4.2/docs/data-sources/file#content DataArchiveFile#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// Set this as the filename when declaring a `source`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.4.2/docs/data-sources/file#filename DataArchiveFile#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

