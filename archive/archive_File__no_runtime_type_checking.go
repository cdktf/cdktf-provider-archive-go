//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt archive Provider for Terraform CDK (cdktf)
package archive

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_File) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (f *jsiiProxy_File) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_File) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (f *jsiiProxy_File) validatePutSourceParameters(value interface{}) error {
	return nil
}

func validateFile_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetExcludesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_File) validateSetOutputFileModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetOutputPathParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_File) validateSetSourceContentParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetSourceContentFilenameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetSourceDirParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetSourceFileParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_File) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewFileParameters(scope constructs.Construct, id *string, config *FileConfig) error {
	return nil
}
