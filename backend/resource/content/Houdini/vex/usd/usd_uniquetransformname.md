---
title: usd_uniquetransformname
order: 150
---
| Since | 18.0 |
| --- | --- |

`string  usd_uniquetransformname(<stage>stage, string primpath, int transformtype, string suffix)`

This function returns a unique full name for a transform operation, given its type and suffix, that is different from any name that currently exists on the given primitive. It can be useful to ensure that the suffix used for the operation name does not stomp on any existing one.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`transformtype`

The numerical code for the transformation type. See the VEX “usd.h” header for defines, such as `USD_XFORM_TRANSLATE`, `USD_XFORM_TRANSFORM`, or `USD_XFORM_ROTATE_XYZ`.

`suffix`

The transform operation suffix.

USD primitives are transformed in space by a series of transform operations whose full names are sequentially listed in the `xformOpOrder` attribute. Full names are namespaced, encode the operation transform type (e.g., translation or rotation), and can also contain a suffix. If primitive has a few operations of the same type, it’s necessary to specify the suffix to differentiate between them. This parameter specifies such a suffix.

Returns

A unique full name of transform operation.

Examples

## examples

```vex
// Construct a unique full name for a translation operation with suffix "cone_pivot"
string unique_xform_name  = usd_uniquetransformname(0, "/geo/cone", USD_XFORM_TRANSLATE, "cone_pivot");

```
