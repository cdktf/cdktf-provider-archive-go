// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package file

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The output of the archive file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#output_path File#output_path}
	OutputPath *string `field:"required" json:"outputPath" yaml:"outputPath"`
	// The type of archive to generate. NOTE: `zip` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#type File#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specify files/directories to ignore when reading the `source_dir`. Supports glob file matching patterns including doublestar/globstar (`**`) patterns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#excludes File#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// Boolean flag indicating whether symbolically linked directories should be excluded during the creation of the archive. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#exclude_symlink_directories File#exclude_symlink_directories}
	ExcludeSymlinkDirectories interface{} `field:"optional" json:"excludeSymlinkDirectories" yaml:"excludeSymlinkDirectories"`
	// String that specifies the octal file mode for all archived files.
	//
	// For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#output_file_mode File#output_file_mode}
	OutputFileMode *string `field:"optional" json:"outputFileMode" yaml:"outputFileMode"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#source File#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Add only this content to the archive with `source_content_filename` as the filename.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#source_content File#source_content}
	SourceContent *string `field:"optional" json:"sourceContent" yaml:"sourceContent"`
	// Set this as the filename when using `source_content`.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#source_content_filename File#source_content_filename}
	SourceContentFilename *string `field:"optional" json:"sourceContentFilename" yaml:"sourceContentFilename"`
	// Package entire contents of this directory into the archive.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#source_dir File#source_dir}
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
	// Package this file into the archive.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.5.0/docs/resources/file#source_file File#source_file}
	SourceFile *string `field:"optional" json:"sourceFile" yaml:"sourceFile"`
}

