//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package file

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FileSourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FileSourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FileSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FileSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FileSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFileSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

