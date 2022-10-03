package dataarchivefile


type DataArchiveFileSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#content DataArchiveFile#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#filename DataArchiveFile#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

