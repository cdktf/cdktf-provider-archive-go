// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ArchiveProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ArchiveProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateArchiveProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateArchiveProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateArchiveProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateArchiveProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewArchiveProviderParameters(scope constructs.Construct, id *string, config *ArchiveProviderConfig) error {
	return nil
}

