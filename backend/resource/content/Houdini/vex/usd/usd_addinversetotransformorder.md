---
title: usd_addinversetotransformorder
order: 4
---
| Since | 18.0 |
| --- | --- |

`int  usd_addinversetotransformorder(int stagehandle, string primpath, string name)`

This function appends an inversed transformation to the primitive’s transform order. Transform order is a sequence of transform operations, whose full names are stored in `xformOpOrder` attribute as a string array. Thus, this function appends a new operation name to that attribute.

Inverse transforms are used primary for rotating (or scaling) around a pivot that does not coincide with the origin. The usual practice is to apply a translation to the pivot point, then perform rotation, and finally apply an inverse of the original translation. This function is used for applying the inverse of the original translation.

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

// Construct the pivot translation operation suffix and name. 
string pivot_xform_suffix = "some_pivot";
string pivot_xform_name   = usd_transformname(USD_XFORM_TRANSLATE, pivot_xform_suffix);

// Rotate about z-axis that goes thru pivot (1,0,0).
usd_addtranslate(0, "/geo/cone", pivot_xform_suffix, {1, 0, 0});
usd_addrotate(0, "/geo/cone", "some_rotation", USD_AXIS_Z, -90);
usd_addinversetotransformorder(0, "/geo/cone", pivot_xform_name);

```
