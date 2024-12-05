---
title: attribsize
order: 10
---
If you know the attribute class ahead of time, using [detailattribsize](detailattribsize.html "Returns the size of a geometry detail attribute."), [primattribsize](primattribsize.html "Returns the size of a geometry prim attribute."), [pointattribsize](pointattribsize.html "Returns the size of a geometry point attribute."), or [vertexattribsize](vertexattribsize.html "Returns the size of a geometry vertex attribute.") may be faster.

`int  attribsize(<geometry>geometry, string attribclass, string attribute_name)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

Returns

The size of an attribute’s *type*.

- For a vector type, this is the number of components.
- For an integer, float, or string, this returns `1`.
- For an array attribute, this returns the size of the tuples in the array. The tuple size is controlled by the **Size** parameter on the [![](../../icons/SOP/attribcreate.svg)Attribute Create node](../../nodes/sop/attribcreate.html "Adds or edits user defined attributes.").

If the attribute does not exist, returns `0`.

- This function works with the attribute’s *type*. It does not return the size of an attribute *value*. You can’t use this function to get the length of a string or array value.

Examples

## examples

```vex
// Get the size of the position attribute of "defgeo.bgeo"
int size = attribsize("defgeo.bgeo", "point", "P");

```
