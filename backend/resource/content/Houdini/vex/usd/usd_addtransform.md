---
title: usd_addtransform
order: 13
---
| Since | 17.5 |
| --- | --- |

`int  usd_addtransform(int stagehandle, string primpath, string suffix, matrix xform)`

This function applies a transformation to the primitive. It creates and sets a value of a transform operation attribute that defines the transformation, and appends it to the primitive’s transform order.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`suffix`

The transform operation suffix.

USD primitives are transformed in space by a series of transform operations whose full names are sequentially listed in the `xformOpOrder` attribute. Full names are namespaced, encode the operation transform type (e.g., translation or rotation), and can also contain a suffix. If primitive has a few operations of the same type, it’s necessary to specify the suffix to differentiate between them. This parameter specifies such a suffix.

`xform`

The the matrix that encodes the space transformation.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Transform the cube.
#include <math.h>
matrix xform = maketransform(XFORM_SRT, XFORM_XYZ, {1,2,3}, {3,45,60}, {0.5,0.25,2});
usd_addtransform(0, "/geo/cube", "my_xform", xform);

```

```vex
// Get the cube's world transform.
matrix xform = usd_worldtransform(0, "/src/cube");

// Match the cone's spacial location to the cube.
// First need to clear current transforms on the cube,
// and also need to block the transform inheritance from the parent.
usd_cleartransformorder(0, "/dst/cone");
usd_settransformreset(0, "/dst/cone", 1);
usd_addtransform(0, "/dst/cone", "", xform);

```
