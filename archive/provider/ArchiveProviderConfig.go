package provider


type ArchiveProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/archive/2.4.0/docs#alias ArchiveProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

