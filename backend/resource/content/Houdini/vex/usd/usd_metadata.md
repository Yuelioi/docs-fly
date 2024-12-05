---
title: usd_metadata
order: 92
---
| Since | 18.0 |
| --- | --- |

`<type> usd_metadata(<stage>stage, string path, string name)`

`<type>[] usd_metadata(<stage>stage, string path, string name)`

This function returns a value of a given metadata from a given object.

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

The value of an existing metadata, or zero/empty value if the metadata does not exist. Use [usd_ismetadata](usd_ismetadata.html "Checks if the primitive has metadata by the given name.") if you want to check whether the metadata exists.

Examples

## examples

```vex
// Get the documentation string of the cube primitive.
string docs = usd_metadata(0, "/geo/cube", "documentation");

// Get custom data from a parameter.
string attrib_path = usd_makeattribpath(0, "/geo/cube", "some_attribute");
float custom_val = usd_metadata(0, attrib_path, "customData:foo:bar");

```
