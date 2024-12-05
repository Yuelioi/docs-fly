---
title: usd_hasapi
order: 49
---
| Since | 18.0 |
| --- | --- |

`int  usd_hasapi(<stage>stage, string primpath, string api)`

This function checks whether the given primitive adheres to the given API. I.e., whether that API has been applied to this primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`api`

The name or an alias of the API schema to check.

Returns

1 if the primitive has the given API, or 0 otherwise.

Examples

## examples

```vex
// Check if the primitive has a USD Geometry Model API applied.
int has_geom_model_api_by_name  = usd_hasapi(0, "/geo/sphere", "UsdGeomModelAPI");
int has_geom_model_api_by_alias = usd_hasapi(0, "/geo/sphere", "GeomModelAPI");

```
