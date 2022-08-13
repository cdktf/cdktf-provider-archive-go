// Prebuilt archive Provider for Terraform CDK (cdktf)
package archive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-archive-go/archive/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-archive-go/archive/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/archive/r/file archive_file}.
type File interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	SetId(val *string)
	IdInput() *string
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Source() FileSourceList
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
	ResetId()
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

// The jsii proxy struct for File
type jsiiProxy_File struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_File) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Excludes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) ExcludesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputBase64Sha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBase64Sha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputFileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputFileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFileModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) OutputSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Source() FileSourceList {
	var returns FileSourceList
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceContentFilename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContentFilename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceContentFilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContentFilenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_File) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/archive/r/file archive_file} Resource.
func NewFile(scope constructs.Construct, id *string, config *FileConfig) File {
	_init_.Initialize()

	j := jsiiProxy_File{}

	_jsii_.Create(
		"@cdktf/provider-archive.File",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/archive/r/file archive_file} Resource.
func NewFile_Override(f File, scope constructs.Construct, id *string, config *FileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.File",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_File) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_File) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_File) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_File) SetExcludes(val *[]*string) {
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_File) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_File) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_File) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_File) SetOutputFileMode(val *string) {
	_jsii_.Set(
		j,
		"outputFileMode",
		val,
	)
}

func (j *jsiiProxy_File) SetOutputPath(val *string) {
	_jsii_.Set(
		j,
		"outputPath",
		val,
	)
}

func (j *jsiiProxy_File) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_File) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_File) SetSourceContent(val *string) {
	_jsii_.Set(
		j,
		"sourceContent",
		val,
	)
}

func (j *jsiiProxy_File) SetSourceContentFilename(val *string) {
	_jsii_.Set(
		j,
		"sourceContentFilename",
		val,
	)
}

func (j *jsiiProxy_File) SetSourceDir(val *string) {
	_jsii_.Set(
		j,
		"sourceDir",
		val,
	)
}

func (j *jsiiProxy_File) SetSourceFile(val *string) {
	_jsii_.Set(
		j,
		"sourceFile",
		val,
	)
}

func (j *jsiiProxy_File) SetType(val *string) {
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
func File_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-archive.File",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func File_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-archive.File",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_File) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_File) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_File) PutSource(value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"putSource",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_File) ResetExcludes() {
	_jsii_.InvokeVoid(
		f,
		"resetExcludes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetOutputFileMode() {
	_jsii_.InvokeVoid(
		f,
		"resetOutputFileMode",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetSource() {
	_jsii_.InvokeVoid(
		f,
		"resetSource",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetSourceContent() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceContent",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetSourceContentFilename() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceContentFilename",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetSourceDir() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceDir",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) ResetSourceFile() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceFile",
		nil, // no parameters
	)
}

func (f *jsiiProxy_File) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_File) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

