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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#output_path File#output_path}.
	OutputPath *string `field:"required" json:"outputPath" yaml:"outputPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#type File#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#excludes File#excludes}.
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#id File#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#output_file_mode File#output_file_mode}.
	OutputFileMode *string `field:"optional" json:"outputFileMode" yaml:"outputFileMode"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source File#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_content File#source_content}.
	SourceContent *string `field:"optional" json:"sourceContent" yaml:"sourceContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_content_filename File#source_content_filename}.
	SourceContentFilename *string `field:"optional" json:"sourceContentFilename" yaml:"sourceContentFilename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_dir File#source_dir}.
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_file File#source_file}.
	SourceFile *string `field:"optional" json:"sourceFile" yaml:"sourceFile"`
}

