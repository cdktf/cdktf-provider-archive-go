// Prebuilt archive Provider for Terraform CDK (cdktf)
package archive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-archive-go/archive/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-archive-go/archive/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/archive archive}.
type ArchiveProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ArchiveProvider
type jsiiProxy_ArchiveProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_ArchiveProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArchiveProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/archive archive} Resource.
func NewArchiveProvider(scope constructs.Construct, id *string, config *ArchiveProviderConfig) ArchiveProvider {
	_init_.Initialize()

	j := jsiiProxy_ArchiveProvider{}

	_jsii_.Create(
		"@cdktf/provider-archive.ArchiveProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/archive archive} Resource.
func NewArchiveProvider_Override(a ArchiveProvider, scope constructs.Construct, id *string, config *ArchiveProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.ArchiveProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ArchiveProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
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
func ArchiveProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-archive.ArchiveProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ArchiveProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-archive.ArchiveProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ArchiveProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ArchiveProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ArchiveProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArchiveProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArchiveProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArchiveProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArchiveProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArchiveProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ArchiveProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive#alias ArchiveProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

// Represents a {@link https://www.terraform.io/docs/providers/archive/d/file archive_file}.
type DataArchiveFile interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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

func (j *jsiiProxy_DataArchiveFile) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_DataArchiveFile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
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


// Create a new {@link https://www.terraform.io/docs/providers/archive/d/file archive_file} Data Source.
func NewDataArchiveFile(scope constructs.Construct, id *string, config *DataArchiveFileConfig) DataArchiveFile {
	_init_.Initialize()

	j := jsiiProxy_DataArchiveFile{}

	_jsii_.Create(
		"@cdktf/provider-archive.DataArchiveFile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/archive/d/file archive_file} Data Source.
func NewDataArchiveFile_Override(d DataArchiveFile, scope constructs.Construct, id *string, config *DataArchiveFileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.DataArchiveFile",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetExcludes(val *[]*string) {
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetOutputFileMode(val *string) {
	_jsii_.Set(
		j,
		"outputFileMode",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetOutputPath(val *string) {
	_jsii_.Set(
		j,
		"outputPath",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetSourceContent(val *string) {
	_jsii_.Set(
		j,
		"sourceContent",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetSourceContentFilename(val *string) {
	_jsii_.Set(
		j,
		"sourceContentFilename",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetSourceDir(val *string) {
	_jsii_.Set(
		j,
		"sourceDir",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetSourceFile(val *string) {
	_jsii_.Set(
		j,
		"sourceFile",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFile) SetType(val *string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-archive.DataArchiveFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataArchiveFile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-archive.DataArchiveFile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataArchiveFile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataArchiveFile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataArchiveFile) PutSource(value interface{}) {
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

func (d *jsiiProxy_DataArchiveFile) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
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

type DataArchiveFileConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#output_path DataArchiveFile#output_path}.
	OutputPath *string `field:"required" json:"outputPath" yaml:"outputPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#type DataArchiveFile#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#excludes DataArchiveFile#excludes}.
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#id DataArchiveFile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#output_file_mode DataArchiveFile#output_file_mode}.
	OutputFileMode *string `field:"optional" json:"outputFileMode" yaml:"outputFileMode"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#source DataArchiveFile#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#source_content DataArchiveFile#source_content}.
	SourceContent *string `field:"optional" json:"sourceContent" yaml:"sourceContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#source_content_filename DataArchiveFile#source_content_filename}.
	SourceContentFilename *string `field:"optional" json:"sourceContentFilename" yaml:"sourceContentFilename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#source_dir DataArchiveFile#source_dir}.
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#source_file DataArchiveFile#source_file}.
	SourceFile *string `field:"optional" json:"sourceFile" yaml:"sourceFile"`
}

type DataArchiveFileSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#content DataArchiveFile#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/d/file#filename DataArchiveFile#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

type DataArchiveFileSourceList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataArchiveFileSourceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataArchiveFileSourceList
type jsiiProxy_DataArchiveFileSourceList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataArchiveFileSourceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataArchiveFileSourceList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataArchiveFileSourceList {
	_init_.Initialize()

	j := jsiiProxy_DataArchiveFileSourceList{}

	_jsii_.Create(
		"@cdktf/provider-archive.DataArchiveFileSourceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataArchiveFileSourceList_Override(d DataArchiveFileSourceList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.DataArchiveFileSourceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataArchiveFileSourceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceList) Get(index *float64) DataArchiveFileSourceOutputReference {
	var returns DataArchiveFileSourceOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataArchiveFileSourceOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataArchiveFileSourceOutputReference
type jsiiProxy_DataArchiveFileSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataArchiveFileSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataArchiveFileSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataArchiveFileSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-archive.DataArchiveFileSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataArchiveFileSourceOutputReference_Override(d DataArchiveFileSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.DataArchiveFileSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) SetContent(val *string) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) SetFilename(val *string) {
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataArchiveFileSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataArchiveFileSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/archive/r/file archive_file}.
type File interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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

type FileConfig struct {
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#output_path File#output_path}.
	OutputPath *string `field:"required" json:"outputPath" yaml:"outputPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#type File#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#excludes File#excludes}.
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#id File#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#output_file_mode File#output_file_mode}.
	OutputFileMode *string `field:"optional" json:"outputFileMode" yaml:"outputFileMode"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source File#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_content File#source_content}.
	SourceContent *string `field:"optional" json:"sourceContent" yaml:"sourceContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_content_filename File#source_content_filename}.
	SourceContentFilename *string `field:"optional" json:"sourceContentFilename" yaml:"sourceContentFilename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_dir File#source_dir}.
	SourceDir *string `field:"optional" json:"sourceDir" yaml:"sourceDir"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#source_file File#source_file}.
	SourceFile *string `field:"optional" json:"sourceFile" yaml:"sourceFile"`
}

type FileSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#content File#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/archive/r/file#filename File#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

type FileSourceList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) FileSourceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FileSourceList
type jsiiProxy_FileSourceList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_FileSourceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewFileSourceList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) FileSourceList {
	_init_.Initialize()

	j := jsiiProxy_FileSourceList{}

	_jsii_.Create(
		"@cdktf/provider-archive.FileSourceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewFileSourceList_Override(f FileSourceList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.FileSourceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FileSourceList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FileSourceList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FileSourceList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FileSourceList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (f *jsiiProxy_FileSourceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceList) Get(index *float64) FileSourceOutputReference {
	var returns FileSourceOutputReference

	_jsii_.Invoke(
		f,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FileSourceOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FileSourceOutputReference
type jsiiProxy_FileSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FileSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFileSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FileSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FileSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-archive.FileSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFileSourceOutputReference_Override(f FileSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-archive.FileSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FileSourceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FileSourceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FileSourceOutputReference) SetContent(val *string) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_FileSourceOutputReference) SetFilename(val *string) {
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_FileSourceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FileSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FileSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FileSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

