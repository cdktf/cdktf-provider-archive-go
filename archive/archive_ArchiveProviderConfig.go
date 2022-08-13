// Prebuilt archive Provider for Terraform CDK (cdktf)
package archive


type ArchiveProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive#alias ArchiveProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}
