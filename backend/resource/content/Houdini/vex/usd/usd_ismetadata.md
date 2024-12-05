---
title: usd_ismetadata
order: 75
---
| Since | 18.0 |
| --- | --- |

`int  usd_ismetadata(<stage>stage, string path, string name)`

This function checks whether the given object has metadata of a given name.

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

`1` if the primitive has the given metadata, or `0` otherwise.

Examples

## examples

```vex
// Check if the primitives have various metadata:
int has_doc = usd_ismetadata(0, "/geo/sphere", "documentation");
int has_custom_foo_bar = usd_ismetadata(0, "/geo/cube", "customData:foo:bar");

// Check if the attribute has custom data set
string attrib_path = usd_makeattribpath(0, "/geo/sphere", "attrib_name");
int has_attrib_foo = usd_ismetadata(0, attrib_path, "customData:foo");

```
