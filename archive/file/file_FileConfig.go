package file

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#output_path File#output_path}
	OutputPath *string `field:"required" json:"outputPath" yaml:"outputPath"`
	// The type of archive to generate. NOTE: `zip` is supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#type File#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Specify files to ignore when reading the `source_dir`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#excludes File#excludes}
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// String that specifies the octal file mode for all archived files.
	//
	// For example: `"0666"`. Setting this will ensure that cross platform usage of this module will not vary the modes of archived files (and ultimately checksums) resulting in more deterministic behavior.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#output_file_mode File#output_file_mode}
	OutputFileMode *string `field:"optional" json:"outputFileMode" yaml:"outputFileMode"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source File#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Add only this content to the archive with `source_content_filename` as the filename.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_content File#source_content}
	SourceContent *string `field:"optional" json:"sourceContent" yaml:"sourceContent"`
	// Set this as the filename when using `source_content`.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_content_filename File#source_content_filename}
	SourceContentFilename *string `field:"optional" json:"sourceContentFilename" yaml:"sourceContentFilename"`
	// Package entire contents of this directory into the archive.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_dir File#source_dir}
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
	// Package this file into the archive.
	//
	// One and only one of `source`, `source_content_filename` (with `source_content`), `source_file`, or `source_dir` must be specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_file File#source_file}
	SourceFile *string `field:"optional" json:"sourceFile" yaml:"sourceFile"`
}

