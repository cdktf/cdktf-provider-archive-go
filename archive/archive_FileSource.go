// Prebuilt archive Provider for Terraform CDK (cdktf)
package archive


type FileSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#content File#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#filename File#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

