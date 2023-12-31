// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceDataVolume struct {
	DeleteWithInstance *bool  `pulumi:"deleteWithInstance"`
	Size               int    `pulumi:"size"`
	VolumeType         string `pulumi:"volumeType"`
}

// InstanceDataVolumeInput is an input type that accepts InstanceDataVolumeArgs and InstanceDataVolumeOutput values.
// You can construct a concrete instance of `InstanceDataVolumeInput` via:
//
//	InstanceDataVolumeArgs{...}
type InstanceDataVolumeInput interface {
	pulumi.Input

	ToInstanceDataVolumeOutput() InstanceDataVolumeOutput
	ToInstanceDataVolumeOutputWithContext(context.Context) InstanceDataVolumeOutput
}

type InstanceDataVolumeArgs struct {
	DeleteWithInstance pulumi.BoolPtrInput `pulumi:"deleteWithInstance"`
	Size               pulumi.IntInput     `pulumi:"size"`
	VolumeType         pulumi.StringInput  `pulumi:"volumeType"`
}

func (InstanceDataVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDataVolume)(nil)).Elem()
}

func (i InstanceDataVolumeArgs) ToInstanceDataVolumeOutput() InstanceDataVolumeOutput {
	return i.ToInstanceDataVolumeOutputWithContext(context.Background())
}

func (i InstanceDataVolumeArgs) ToInstanceDataVolumeOutputWithContext(ctx context.Context) InstanceDataVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDataVolumeOutput)
}

// InstanceDataVolumeArrayInput is an input type that accepts InstanceDataVolumeArray and InstanceDataVolumeArrayOutput values.
// You can construct a concrete instance of `InstanceDataVolumeArrayInput` via:
//
//	InstanceDataVolumeArray{ InstanceDataVolumeArgs{...} }
type InstanceDataVolumeArrayInput interface {
	pulumi.Input

	ToInstanceDataVolumeArrayOutput() InstanceDataVolumeArrayOutput
	ToInstanceDataVolumeArrayOutputWithContext(context.Context) InstanceDataVolumeArrayOutput
}

type InstanceDataVolumeArray []InstanceDataVolumeInput

func (InstanceDataVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceDataVolume)(nil)).Elem()
}

func (i InstanceDataVolumeArray) ToInstanceDataVolumeArrayOutput() InstanceDataVolumeArrayOutput {
	return i.ToInstanceDataVolumeArrayOutputWithContext(context.Background())
}

func (i InstanceDataVolumeArray) ToInstanceDataVolumeArrayOutputWithContext(ctx context.Context) InstanceDataVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDataVolumeArrayOutput)
}

type InstanceDataVolumeOutput struct{ *pulumi.OutputState }

func (InstanceDataVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceDataVolume)(nil)).Elem()
}

func (o InstanceDataVolumeOutput) ToInstanceDataVolumeOutput() InstanceDataVolumeOutput {
	return o
}

func (o InstanceDataVolumeOutput) ToInstanceDataVolumeOutputWithContext(ctx context.Context) InstanceDataVolumeOutput {
	return o
}

func (o InstanceDataVolumeOutput) DeleteWithInstance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InstanceDataVolume) *bool { return v.DeleteWithInstance }).(pulumi.BoolPtrOutput)
}

func (o InstanceDataVolumeOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceDataVolume) int { return v.Size }).(pulumi.IntOutput)
}

func (o InstanceDataVolumeOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceDataVolume) string { return v.VolumeType }).(pulumi.StringOutput)
}

type InstanceDataVolumeArrayOutput struct{ *pulumi.OutputState }

func (InstanceDataVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceDataVolume)(nil)).Elem()
}

func (o InstanceDataVolumeArrayOutput) ToInstanceDataVolumeArrayOutput() InstanceDataVolumeArrayOutput {
	return o
}

func (o InstanceDataVolumeArrayOutput) ToInstanceDataVolumeArrayOutputWithContext(ctx context.Context) InstanceDataVolumeArrayOutput {
	return o
}

func (o InstanceDataVolumeArrayOutput) Index(i pulumi.IntInput) InstanceDataVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceDataVolume {
		return vs[0].([]InstanceDataVolume)[vs[1].(int)]
	}).(InstanceDataVolumeOutput)
}

type InstanceGpuDevice struct {
	Count               *int    `pulumi:"count"`
	EncryptedMemorySize *int    `pulumi:"encryptedMemorySize"`
	MemorySize          *int    `pulumi:"memorySize"`
	ProductName         *string `pulumi:"productName"`
}

// InstanceGpuDeviceInput is an input type that accepts InstanceGpuDeviceArgs and InstanceGpuDeviceOutput values.
// You can construct a concrete instance of `InstanceGpuDeviceInput` via:
//
//	InstanceGpuDeviceArgs{...}
type InstanceGpuDeviceInput interface {
	pulumi.Input

	ToInstanceGpuDeviceOutput() InstanceGpuDeviceOutput
	ToInstanceGpuDeviceOutputWithContext(context.Context) InstanceGpuDeviceOutput
}

type InstanceGpuDeviceArgs struct {
	Count               pulumi.IntPtrInput    `pulumi:"count"`
	EncryptedMemorySize pulumi.IntPtrInput    `pulumi:"encryptedMemorySize"`
	MemorySize          pulumi.IntPtrInput    `pulumi:"memorySize"`
	ProductName         pulumi.StringPtrInput `pulumi:"productName"`
}

func (InstanceGpuDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGpuDevice)(nil)).Elem()
}

func (i InstanceGpuDeviceArgs) ToInstanceGpuDeviceOutput() InstanceGpuDeviceOutput {
	return i.ToInstanceGpuDeviceOutputWithContext(context.Background())
}

func (i InstanceGpuDeviceArgs) ToInstanceGpuDeviceOutputWithContext(ctx context.Context) InstanceGpuDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGpuDeviceOutput)
}

// InstanceGpuDeviceArrayInput is an input type that accepts InstanceGpuDeviceArray and InstanceGpuDeviceArrayOutput values.
// You can construct a concrete instance of `InstanceGpuDeviceArrayInput` via:
//
//	InstanceGpuDeviceArray{ InstanceGpuDeviceArgs{...} }
type InstanceGpuDeviceArrayInput interface {
	pulumi.Input

	ToInstanceGpuDeviceArrayOutput() InstanceGpuDeviceArrayOutput
	ToInstanceGpuDeviceArrayOutputWithContext(context.Context) InstanceGpuDeviceArrayOutput
}

type InstanceGpuDeviceArray []InstanceGpuDeviceInput

func (InstanceGpuDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceGpuDevice)(nil)).Elem()
}

func (i InstanceGpuDeviceArray) ToInstanceGpuDeviceArrayOutput() InstanceGpuDeviceArrayOutput {
	return i.ToInstanceGpuDeviceArrayOutputWithContext(context.Background())
}

func (i InstanceGpuDeviceArray) ToInstanceGpuDeviceArrayOutputWithContext(ctx context.Context) InstanceGpuDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGpuDeviceArrayOutput)
}

type InstanceGpuDeviceOutput struct{ *pulumi.OutputState }

func (InstanceGpuDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceGpuDevice)(nil)).Elem()
}

func (o InstanceGpuDeviceOutput) ToInstanceGpuDeviceOutput() InstanceGpuDeviceOutput {
	return o
}

func (o InstanceGpuDeviceOutput) ToInstanceGpuDeviceOutputWithContext(ctx context.Context) InstanceGpuDeviceOutput {
	return o
}

func (o InstanceGpuDeviceOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceGpuDevice) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o InstanceGpuDeviceOutput) EncryptedMemorySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceGpuDevice) *int { return v.EncryptedMemorySize }).(pulumi.IntPtrOutput)
}

func (o InstanceGpuDeviceOutput) MemorySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceGpuDevice) *int { return v.MemorySize }).(pulumi.IntPtrOutput)
}

func (o InstanceGpuDeviceOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceGpuDevice) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

type InstanceGpuDeviceArrayOutput struct{ *pulumi.OutputState }

func (InstanceGpuDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceGpuDevice)(nil)).Elem()
}

func (o InstanceGpuDeviceArrayOutput) ToInstanceGpuDeviceArrayOutput() InstanceGpuDeviceArrayOutput {
	return o
}

func (o InstanceGpuDeviceArrayOutput) ToInstanceGpuDeviceArrayOutputWithContext(ctx context.Context) InstanceGpuDeviceArrayOutput {
	return o
}

func (o InstanceGpuDeviceArrayOutput) Index(i pulumi.IntInput) InstanceGpuDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceGpuDevice {
		return vs[0].([]InstanceGpuDevice)[vs[1].(int)]
	}).(InstanceGpuDeviceOutput)
}

type InstanceSecondaryNetworkInterface struct {
	PrimaryIpAddress *string  `pulumi:"primaryIpAddress"`
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	SubnetId         string   `pulumi:"subnetId"`
}

// InstanceSecondaryNetworkInterfaceInput is an input type that accepts InstanceSecondaryNetworkInterfaceArgs and InstanceSecondaryNetworkInterfaceOutput values.
// You can construct a concrete instance of `InstanceSecondaryNetworkInterfaceInput` via:
//
//	InstanceSecondaryNetworkInterfaceArgs{...}
type InstanceSecondaryNetworkInterfaceInput interface {
	pulumi.Input

	ToInstanceSecondaryNetworkInterfaceOutput() InstanceSecondaryNetworkInterfaceOutput
	ToInstanceSecondaryNetworkInterfaceOutputWithContext(context.Context) InstanceSecondaryNetworkInterfaceOutput
}

type InstanceSecondaryNetworkInterfaceArgs struct {
	PrimaryIpAddress pulumi.StringPtrInput   `pulumi:"primaryIpAddress"`
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	SubnetId         pulumi.StringInput      `pulumi:"subnetId"`
}

func (InstanceSecondaryNetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceSecondaryNetworkInterface)(nil)).Elem()
}

func (i InstanceSecondaryNetworkInterfaceArgs) ToInstanceSecondaryNetworkInterfaceOutput() InstanceSecondaryNetworkInterfaceOutput {
	return i.ToInstanceSecondaryNetworkInterfaceOutputWithContext(context.Background())
}

func (i InstanceSecondaryNetworkInterfaceArgs) ToInstanceSecondaryNetworkInterfaceOutputWithContext(ctx context.Context) InstanceSecondaryNetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecondaryNetworkInterfaceOutput)
}

// InstanceSecondaryNetworkInterfaceArrayInput is an input type that accepts InstanceSecondaryNetworkInterfaceArray and InstanceSecondaryNetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `InstanceSecondaryNetworkInterfaceArrayInput` via:
//
//	InstanceSecondaryNetworkInterfaceArray{ InstanceSecondaryNetworkInterfaceArgs{...} }
type InstanceSecondaryNetworkInterfaceArrayInput interface {
	pulumi.Input

	ToInstanceSecondaryNetworkInterfaceArrayOutput() InstanceSecondaryNetworkInterfaceArrayOutput
	ToInstanceSecondaryNetworkInterfaceArrayOutputWithContext(context.Context) InstanceSecondaryNetworkInterfaceArrayOutput
}

type InstanceSecondaryNetworkInterfaceArray []InstanceSecondaryNetworkInterfaceInput

func (InstanceSecondaryNetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceSecondaryNetworkInterface)(nil)).Elem()
}

func (i InstanceSecondaryNetworkInterfaceArray) ToInstanceSecondaryNetworkInterfaceArrayOutput() InstanceSecondaryNetworkInterfaceArrayOutput {
	return i.ToInstanceSecondaryNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i InstanceSecondaryNetworkInterfaceArray) ToInstanceSecondaryNetworkInterfaceArrayOutputWithContext(ctx context.Context) InstanceSecondaryNetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceSecondaryNetworkInterfaceArrayOutput)
}

type InstanceSecondaryNetworkInterfaceOutput struct{ *pulumi.OutputState }

func (InstanceSecondaryNetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceSecondaryNetworkInterface)(nil)).Elem()
}

func (o InstanceSecondaryNetworkInterfaceOutput) ToInstanceSecondaryNetworkInterfaceOutput() InstanceSecondaryNetworkInterfaceOutput {
	return o
}

func (o InstanceSecondaryNetworkInterfaceOutput) ToInstanceSecondaryNetworkInterfaceOutputWithContext(ctx context.Context) InstanceSecondaryNetworkInterfaceOutput {
	return o
}

func (o InstanceSecondaryNetworkInterfaceOutput) PrimaryIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceSecondaryNetworkInterface) *string { return v.PrimaryIpAddress }).(pulumi.StringPtrOutput)
}

func (o InstanceSecondaryNetworkInterfaceOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceSecondaryNetworkInterface) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

func (o InstanceSecondaryNetworkInterfaceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceSecondaryNetworkInterface) string { return v.SubnetId }).(pulumi.StringOutput)
}

type InstanceSecondaryNetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (InstanceSecondaryNetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceSecondaryNetworkInterface)(nil)).Elem()
}

func (o InstanceSecondaryNetworkInterfaceArrayOutput) ToInstanceSecondaryNetworkInterfaceArrayOutput() InstanceSecondaryNetworkInterfaceArrayOutput {
	return o
}

func (o InstanceSecondaryNetworkInterfaceArrayOutput) ToInstanceSecondaryNetworkInterfaceArrayOutputWithContext(ctx context.Context) InstanceSecondaryNetworkInterfaceArrayOutput {
	return o
}

func (o InstanceSecondaryNetworkInterfaceArrayOutput) Index(i pulumi.IntInput) InstanceSecondaryNetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceSecondaryNetworkInterface {
		return vs[0].([]InstanceSecondaryNetworkInterface)[vs[1].(int)]
	}).(InstanceSecondaryNetworkInterfaceOutput)
}

type InstanceTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// InstanceTagInput is an input type that accepts InstanceTagArgs and InstanceTagOutput values.
// You can construct a concrete instance of `InstanceTagInput` via:
//
//	InstanceTagArgs{...}
type InstanceTagInput interface {
	pulumi.Input

	ToInstanceTagOutput() InstanceTagOutput
	ToInstanceTagOutputWithContext(context.Context) InstanceTagOutput
}

type InstanceTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (InstanceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTag)(nil)).Elem()
}

func (i InstanceTagArgs) ToInstanceTagOutput() InstanceTagOutput {
	return i.ToInstanceTagOutputWithContext(context.Background())
}

func (i InstanceTagArgs) ToInstanceTagOutputWithContext(ctx context.Context) InstanceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTagOutput)
}

// InstanceTagArrayInput is an input type that accepts InstanceTagArray and InstanceTagArrayOutput values.
// You can construct a concrete instance of `InstanceTagArrayInput` via:
//
//	InstanceTagArray{ InstanceTagArgs{...} }
type InstanceTagArrayInput interface {
	pulumi.Input

	ToInstanceTagArrayOutput() InstanceTagArrayOutput
	ToInstanceTagArrayOutputWithContext(context.Context) InstanceTagArrayOutput
}

type InstanceTagArray []InstanceTagInput

func (InstanceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceTag)(nil)).Elem()
}

func (i InstanceTagArray) ToInstanceTagArrayOutput() InstanceTagArrayOutput {
	return i.ToInstanceTagArrayOutputWithContext(context.Background())
}

func (i InstanceTagArray) ToInstanceTagArrayOutputWithContext(ctx context.Context) InstanceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTagArrayOutput)
}

type InstanceTagOutput struct{ *pulumi.OutputState }

func (InstanceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTag)(nil)).Elem()
}

func (o InstanceTagOutput) ToInstanceTagOutput() InstanceTagOutput {
	return o
}

func (o InstanceTagOutput) ToInstanceTagOutputWithContext(ctx context.Context) InstanceTagOutput {
	return o
}

func (o InstanceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o InstanceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceTag) string { return v.Value }).(pulumi.StringOutput)
}

type InstanceTagArrayOutput struct{ *pulumi.OutputState }

func (InstanceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceTag)(nil)).Elem()
}

func (o InstanceTagArrayOutput) ToInstanceTagArrayOutput() InstanceTagArrayOutput {
	return o
}

func (o InstanceTagArrayOutput) ToInstanceTagArrayOutputWithContext(ctx context.Context) InstanceTagArrayOutput {
	return o
}

func (o InstanceTagArrayOutput) Index(i pulumi.IntInput) InstanceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceTag {
		return vs[0].([]InstanceTag)[vs[1].(int)]
	}).(InstanceTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceDataVolumeInput)(nil)).Elem(), InstanceDataVolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceDataVolumeArrayInput)(nil)).Elem(), InstanceDataVolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGpuDeviceInput)(nil)).Elem(), InstanceGpuDeviceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGpuDeviceArrayInput)(nil)).Elem(), InstanceGpuDeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecondaryNetworkInterfaceInput)(nil)).Elem(), InstanceSecondaryNetworkInterfaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceSecondaryNetworkInterfaceArrayInput)(nil)).Elem(), InstanceSecondaryNetworkInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTagInput)(nil)).Elem(), InstanceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTagArrayInput)(nil)).Elem(), InstanceTagArray{})
	pulumi.RegisterOutputType(InstanceDataVolumeOutput{})
	pulumi.RegisterOutputType(InstanceDataVolumeArrayOutput{})
	pulumi.RegisterOutputType(InstanceGpuDeviceOutput{})
	pulumi.RegisterOutputType(InstanceGpuDeviceArrayOutput{})
	pulumi.RegisterOutputType(InstanceSecondaryNetworkInterfaceOutput{})
	pulumi.RegisterOutputType(InstanceSecondaryNetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(InstanceTagOutput{})
	pulumi.RegisterOutputType(InstanceTagArrayOutput{})
}
