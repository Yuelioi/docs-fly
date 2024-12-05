---
title: usd_addrotate
order: 9
---
| Since | 18.0 |
| --- | --- |

`int  usd_addrotate(int stagehandle, string primpath, string suffix, int axis, float angle)`

`int  usd_addrotate(int stagehandle, string primpath, string suffix, int xyz, vector angles)`

This function applies a rotation to the primitive. It creates and sets a value of a transform operation attribute that defines the rotation, and appends it to the primitive’s transform order.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`suffix`

The transform operation suffix.

USD primitives are transformed in space by a series of transform operations whose full names are sequentially listed in the `xformOpOrder` attribute. Full names are namespaced, encode the operation transform type (e.g., translation or rotation), and can also contain a suffix. If primitive has a few operations of the same type, it’s necessary to specify the suffix to differentiate between them. This parameter specifies such a suffix.

`axis`

The numeric code for the rotation axis. See the usd.h VEX header for definitions of the axis.

`angle`

Angle of rotation in degrees, when rotating around a single principle axis.

`axis`

The numeric code for the axis rotation order. See the usd.h VEX header for definitions of the order.

`angles`

Three angles of rotation in degrees, when rotating sequentially around each of the principle axes.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Include "usd.h" that defines axis and order constants.
#include <usd.h>

// Rotate the cube 30 degrees around the z-axis.
usd_addrotate(0, "/geo/cube", "", USD_AXIS_Z, 30);

// Rotate the mesh 45 degrees counter-clock wise around the y-axis.
usd_addrotate(0, "/geo/mesh", "geo_rotation", USD_AXIS_Y, -45);

// Rotate the cone about Euler angles.
usd_addrotate(0, "/geo/cone", "cone_rotation", USD_ROTATE_XYZ, {0, 30, 45});

```
