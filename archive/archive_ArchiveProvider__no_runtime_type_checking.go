//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt archive Provider for Terraform CDK (cdktf)
package archive

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ArchiveProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ArchiveProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateArchiveProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewArchiveProviderParameters(scope constructs.Construct, id *string, config *ArchiveProviderConfig) error {
	return nil
}

