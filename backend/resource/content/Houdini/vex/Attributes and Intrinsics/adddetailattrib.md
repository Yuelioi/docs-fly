---
title: adddetailattrib
order: 2
---
If you don’t know the class of attribute until runtime, use [addattrib](addattrib.html "Adds an attribute to a geometry.").

`int  adddetailattrib(int geohandle, string name, <type>defvalue)`

`int  adddetailattrib(int geohandle, string name, <type>defvalue[])`

Adds a detail attribute to the given geometry.

`int  adddetailattrib(int geohandle, string name, <type>defvalue, string typeinfo)`

`int  adddetailattrib(int geohandle, string name, <type>defvalue[], string typeinfo)`

Adds a detail attribute with the given transformation info. See [attribtypeinfo](attribtypeinfo.html "Returns the transformation metadata of a geometry attribute.") for more details.

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`name`

The name of the attribute to create.

`defvalue`

The default value for the attribute and determines the type of attribute to create. String and array attributes cannot have defaults, so only the type is used in those cases.

Returns

`geohandle` on success, or `-1` on failure.

- If an attribute of the same name already exists, the function will try to convert it to the new type.
