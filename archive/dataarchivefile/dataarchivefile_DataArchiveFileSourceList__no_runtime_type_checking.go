//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataarchivefile

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataArchiveFileSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataArchiveFileSourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataArchiveFileSourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataArchiveFileSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataArchiveFileSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataArchiveFileSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataArchiveFileSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

