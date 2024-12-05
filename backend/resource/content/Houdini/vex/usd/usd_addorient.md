---
title: usd_addorient
order: 5
---
| Since | 18.0 |
| --- | --- |

`int  usd_addorient(int stagehandle, string primpath, string suffix, vector4 orient)`

This function applies a quaternion orientation to the primitive. It creates and sets a value of a transform operation attribute that defines the orientation, and appends it to the primitive’s transform order.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`suffix`

The transform operation suffix.

USD primitives are transformed in space by a series of transform operations whose full names are sequentially listed in the `xformOpOrder` attribute. Full names are namespaced, encode the operation transform type (e.g., translation or rotation), and can also contain a suffix. If primitive has a few operations of the same type, it’s necessary to specify the suffix to differentiate between them. This parameter specifies such a suffix.

`orient`

The quaternion (in a vector4 format) representing the orientation.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Orient the cube
vector4 quat = eulertoquaternion(radians({30,0,0}), XFORM_XYZ);
usd_addorient(0, "/dst/cone", "my_orientation", quat);

```
