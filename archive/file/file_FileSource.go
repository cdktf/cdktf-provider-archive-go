package file


type FileSource struct {
	// Add this content to the archive with `filename` as the filename.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#content File#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// Set this as the filename when declaring a `source`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#filename File#filename}
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

