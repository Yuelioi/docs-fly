---
title: agentrigparent
order: 38
---
`int  agentrigparent(<geometry>geometry, int prim, int transform)`

Returns `-1` if `transform` is a root of the transform hierarchy, `transform` is [out of range](agenttransformcount.html "Returns the number of transforms in an agent primitive’s rig."), `prim` is out of range, or `prim` is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transform`

Index of a transform in the agent’s rig.

Examples

## examples

Find the world transform of the rig’s root, beginning from a given bone.

```vex
int root;
while (true) {
int parent = agentrigparent(0, @primnum, transform);

if (parent < 0)
{
root = transform;
break;
}
else
transform = parent;
}

matrix root_xform = agentworldtransform(0, @primnum, root);

```
