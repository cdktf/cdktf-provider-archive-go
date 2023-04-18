package dataarchivefile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-archive-go/archive/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-archive-go/archive/v6/dataarchivefile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/archive/2.3.0/docs/data-sources/file archive_file}.
type DataArchiveFile interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Excludes() *[]*string
	SetExcludes(val *[]*string)
	ExcludesInput() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OutputBase64Sha256() *string
	OutputFileMode() *string
	SetOutputFileMode(val *string)
	OutputFileModeInput() *string
	OutputMd5() *string
	OutputPath() *string
	SetOutputPath(val *string)
	OutputPathInput() *string
	OutputSha() *string
	OutputSize() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Source() DataArchiveFileSourceList
	SourceContent() *string
	SetSourceContent(val *string)
	SourceContentFilename() *string
	SetSourceContentFilename(val *string)
	SourceContentFilenameInput() *string
	SourceContentInput() *string
	SourceDir() *string
	SetSourceDir(val *string)
	SourceDirInput() *string
	SourceFile() *string
	SetSourceFile(val *string)
	SourceFileInput() *string
	SourceInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutSource(value interface{})
	ResetExcludes()
	ResetOutputFileMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSource()
	ResetSourceContent()
	ResetSourceContentFilename()
	ResetSourceDir()
	ResetSourceFile()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataArchiveFile
type jsiiProxy_DataArchiveFile struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataArchiveFile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Excludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) ExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputBase64Sha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBase64Sha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputFileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputFileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) OutputSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Source() DataArchiveFileSourceList {
	var returns DataArchiveFileSourceList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceContentFilename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContentFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceContentFilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContentFilenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFile) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/archive/2.3.0/docs/data-sources/file archive_file} Data Source.
func NewDataArchiveFile(scope constructs.Construct, id *string, config *DataArchiveFileConfig) DataArchiveFile {
	_init_.Initialize()

	if err := validateNewDataArchiveFileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataArchiveFile{}

	_jsii_.Create(
		"@cdktf/provider-archive.dataArchiveFile.DataArchiveFile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/archive/2.3.0/docs/data-sources/file archive_file} Data Source.
func NewDataArchiveFile_Override(d DataArchiveFile, scope constructs.Construct, id *string, config *DataArchiveFileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.dataArchiveFile.DataArchiveFile",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetExcludes(val *[]*string) {
	if err := j.validateSetExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetOutputFileMode(val *string) {
	if err := j.validateSetOutputFileModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFileMode",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetOutputPath(val *string) {
	if err := j.validateSetOutputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputPath",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetSourceContent(val *string) {
	if err := j.validateSetSourceContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceContent",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetSourceContentFilename(val *string) {
	if err := j.validateSetSourceContentFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceContentFilename",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetSourceDir(val *string) {
	if err := j.validateSetSourceDirParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDir",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetSourceFile(val *string) {
	if err := j.validateSetSourceFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFile",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataArchiveFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataArchiveFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-archive.dataArchiveFile.DataArchiveFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataArchiveFile_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataArchiveFile_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-archive.dataArchiveFile.DataArchiveFile",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataArchiveFile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataArchiveFile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-archive.dataArchiveFile.DataArchiveFile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataArchiveFile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-archive.dataArchiveFile.DataArchiveFile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataArchiveFile) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataArchiveFile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataArchiveFile) PutSource(value interface{}) {
	if err := d.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetExcludes() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetOutputFileMode() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputFileMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetSource() {
	_jsii_.InvokeVoid(
		d,
		"resetSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetSourceContent() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceContent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetSourceContentFilename() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceContentFilename",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetSourceDir() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) ResetSourceFile() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataArchiveFile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

