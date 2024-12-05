---
title: usd_setmetadata
order: 130
---
| Since | 18.0 |
| --- | --- |

`int  usd_setmetadata(int stagehandle, string path, string name, <type>value)`

`int  usd_setmetadata(int stagehandle, string path, string name, <type>value[])`

This function sets the metadata value.

Show/hide arguments

`stagehandle`

A handle to the stage to write to. Currently the only valid value is `0`, which means the current stage in a node. (This argument may be used in the future to allow writing to other stages.)

`path`

The path to the object. I.e, a primitive, an attribute, or a relationship.

`name`

The metadata name.

The name can be namespaced to acces values inside (possibly nested) VtDictionaries, such as custom data dictionary, e.g., “customData:name” or “customData:name:subname”. For non-namespaced names, the object schema needs to declare a given metadata for it to be accessible, e.g., “active” or “documentation”.

Returns

The value of `stagehandle` on success, or `-1` on failure.

Examples

## examples

```vex
// Set a documentation string on the sphere.
usd_setmetadata(0, "/geo/sphere", "documentation", "This is new documentation.");

// Set the value of some custom data on the sphere.
usd_setmetadata(0, "/geo/sphere", "customData:a_float", 0.25);
usd_setmetadata(0, "/geo/sphere", "customData:a_string", "foo bar baz");
usd_setmetadata(0, "/geo/sphere", "customData:a_vector", {1.25, 1.50, 1.75});

float  f_arr[] = {0, 0.25, 0.5, 0.75, 1};
usd_setmetadata(0, "/geo/sphere", "customData:a_float_array", f_arr);

// Set the metadata value on an attribute.
string attrib_path = usd_makeattribpath(0, "/geo/sphere", "attrib_name");
sd_setmetadata(0, attrib_path, "customData:foo", 1.25);

```
