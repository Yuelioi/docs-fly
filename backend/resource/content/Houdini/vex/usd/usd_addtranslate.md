---
title: usd_addtranslate
order: 14
---
| Since | 18.0 |
| --- | --- |

`int  usd_addtranslate(int stagehandle, string primpath, string suffix, vector amount)`

This function applies a translation to the primitive. It creates and sets a value of a transform operation attribute that defines the translation, and appends it to the primitive’s transform order.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`primpath`

The path to the primitive.

`suffix`

The transform operation suffix.

USD primitives are transformed in space by a series of transform operations whose full names are sequentially listed in the `xformOpOrder` attribute. Full names are namespaced, encode the operation transform type (e.g., translation or rotation), and can also contain a suffix. If primitive has a few operations of the same type, it’s necessary to specify the suffix to differentiate between them. This parameter specifies such a suffix.

`amount`

The the distance to move along each of the principle axes.

Returns

The value of `stagehandle` on success or `-1` on failure.

Examples

## examples

```vex
// Translate the cube
usd_addtranslate(0, "/geo/cube", "my_translation", {10, 0, -2.5});

```
