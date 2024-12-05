---
title: usd_isarraymetadata
order: 65
---
| Since | 18.0 |
| --- | --- |

`int  usd_isarraymetadata(<stage>stage, string path, string name)`

This function checks whether the given metadata is an array.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

`name`

The metadata name.

The name can be namespaced to acces values inside (possibly nested) VtDictionaries, such as custom data dictionary, e.g., “customData:name” or “customData:name:subname”. For non-namespaced names, the object schema needs to declare a given metadata for it to be accessible, e.g., “active” or “documentation”.

Returns

`1` if the metadata exists and is an array, or `0` otherwise.

Examples

## examples

```vex
// Check if the metadata is an array.
int is_array = usd_isarraymetadata(0, "/geo/sphere", "documentation");
int is_array_too = usd_isarraymetadata(0, "/geo/cube", "customData:foo:bar");

```
