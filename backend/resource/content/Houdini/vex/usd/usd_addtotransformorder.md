---
title: usd_addtotransformorder
order: 12
---
| Since | 18.0 |
| --- | --- |

`int  usd_addtotransformorder(int stagehandle, string primpath, string name)`

This function appends a transformation to the primitive’s transform order. Transform order is a sequence of transform operations, whose full names are stored in `xformOpOrder` attribute as a string array. Thus, this function appends a new operation name to that attribute.

NOTE: unlike most VEX functions that deal with primitive transforms and take an operation suffix as a parameter, this function takes the full operation name. Use [usd_transformname](usd_transformname.html "Constructs a full name of a transform operation") to obtain the full name if you know the suffix.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`name`

The full name of the transform operation. Use [usd_transformname](usd_transformname.html "Constructs a full name of a transform operation") to obtain the full name from the operation suffix.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Note, the USD_XFORM_TRANSLATE and USD_AXIS_Z constants used below 
// are defined in "usd.h" VEX header, so include it.
#include <usd.h>

// Make the first step (i.e., translate)
string step_suffix = "step";
usd_addtranslate(0, "/geo/cone", step_suffix, {1, 0, 0});

// Now repeat the same step translation by adding it to the transform order
string step_name   = usd_transformname(USD_XFORM_TRANSLATE, step_suffix);
usd_addrotate(0, "/geo/cone", "first_rotation", USD_AXIS_Z, -30);
usd_addtotransformorder(0, "/geo/cone", step_name);
usd_addrotate(0, "/geo/cone", "second_rotation", USD_AXIS_Z, 45);
usd_addtotransformorder(0, "/geo/cone", step_name);

```
