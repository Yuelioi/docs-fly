---
title: usd_setvisibility
order: 142
---
| Since | 19.0 |
| --- | --- |

`int  usd_setvisibility(int stagehandle, string primpath, int code)`

This function makes the primitive visible or invisible, or configures it to inherit the visibility from the parent.

Making a primitive visible may require changing the visibility state of its ancestors, while making it invisible or configuring it to inherit the visibility from the parent only requires setting its attribute.

NOTE: This function is similar to `usd_setvisible()` which, is equivalent to calling this function with either visibility or invisibility code.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`code`

A numeric code for visibility

- 0 - make the primitive invisible
- 1 - configure the primitive to be visible
- 2 - mark the primitive to inherit visibility from the parent

Note, these numeric codes are reflected as defines in the “usd.h” header file, as USD_VISIBILITY_INVISIBLE, USD_VISIBILITY_VISIBLE, and USD_VISIBILITY_INHERIT.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
#include <usd.h>
// Make the sphere primitive visible.
usd_setvisibility(0, "/geo/sphere", USD_VISIBILITY_VISIBLE);

// Configure the cube primitive to inherit visibility from parent.
usd_setvisibility(0, "/geo/cube", USD_VISIBILITY_INHERIT);

```
