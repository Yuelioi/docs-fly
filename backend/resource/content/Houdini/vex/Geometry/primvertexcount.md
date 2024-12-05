---
title: primvertexcount
order: 29
---
`int  primvertexcount(<geometry>geometry, int prim_num)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim_num`

The primitive number of the primitive to count vertices on.

Returns

The number of vertices in the given primitive, or `-1` if the primitive does not exist.

Examples

## examples

```vex
int        nvtx;

// Get the number of vertices of primitive 3
nvtx = primvertexcount("defgeo.bgeo", 3);

```
