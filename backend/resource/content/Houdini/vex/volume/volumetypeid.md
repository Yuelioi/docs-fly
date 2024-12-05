---
title: volumetypeid
order: 22
---
`int  volumetypeid(<geometry>geometry, int primnum)`

`int  volumetypeid(<geometry>geometry, string volumename)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The integer typeid for the data stored in the volume or VDB. This
will match the result of the VEX typeid() function with the same
type.

Returns -1 if it isn’t a volume or VDB, or isn’t a VEX compatible type.
