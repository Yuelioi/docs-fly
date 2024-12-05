---
title: usd_findtransformname
order: 37
---
| Since | 18.0 |
| --- | --- |

`string  usd_findtransformname(<stage>stage, string primpath, string suffix)`

This function returns the full name of a transform operation given the suffix, if such an operation attribute exists on the given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`suffix`

The transform operation suffix.

USD primitives are transformed in space by a series of transform operations whose full names are sequentially listed in the `xformOpOrder` attribute. Full names are namespaced, encode the operation transform type (e.g., translation or rotation), and can also contain a suffix. If primitive has a few operations of the same type, it’s necessary to specify the suffix to differentiate between them. This parameter specifies such a suffix.

Returns

The full name of the primitive’s transform operation that has the given suffix, or an empty string if no such operation is found.

There may be a few transform operations with the same suffix (eg, a translation an a rotation), so the first encountered one will be returned.

Examples

## examples

```vex
// Find the transform operation name for the pivot translation, and add an iverse of it.
string xform_name = usd_findtransformname(0, "/geo/cone", "cone_pivot");
usd_addinversetotransformorder(0, "/geo/cone", xform_name);

```
