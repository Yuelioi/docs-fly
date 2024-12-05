---
title: nuniqueval
order: 33
---
`int  nuniqueval(<geometry>geometry, string attribclass, string attribute_name)`

Returns the number of *unique* values across all values of an attribute.
You can use [uniqueval](uniqueval.html "Returns one of the set of unique values across all values for an int or string attribute.") to iterate though the set of unique values.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`attribclass`

One of `"detail"` (or `"global"`), `"point"`, `"prim"`, or `"vertex"`.

You can also use `"primgroup"`, `"pointgroup"` or `"vertexgroup"` to [read from groups](../groups.html "You can read the contents of primitive/point/vertex groups in VEX as if they were attributes.").

Examples

## examples

Test if all values of the point attribute `foo` are unique

```vex
int test = nuniqueval(0, "point", "foo") == npoints(0)

```
