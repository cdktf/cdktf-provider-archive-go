package dataarchivefile


type DataArchiveFileSource struct {
	// Add this content to the archive with `filename` as the filename.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#content DataArchiveFile#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// Set this as the filename when declaring a `source`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#filename DataArchiveFile#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

