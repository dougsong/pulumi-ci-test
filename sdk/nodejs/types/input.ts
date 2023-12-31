// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace ecs {
    export interface InstanceDataVolume {
        deleteWithInstance?: pulumi.Input<boolean>;
        size: pulumi.Input<number>;
        volumeType: pulumi.Input<string>;
    }

    export interface InstanceGpuDevice {
        count?: pulumi.Input<number>;
        encryptedMemorySize?: pulumi.Input<number>;
        memorySize?: pulumi.Input<number>;
        productName?: pulumi.Input<string>;
    }

    export interface InstanceSecondaryNetworkInterface {
        primaryIpAddress?: pulumi.Input<string>;
        securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
        subnetId: pulumi.Input<string>;
    }

    export interface InstanceTag {
        key: pulumi.Input<string>;
        value: pulumi.Input<string>;
    }
}
