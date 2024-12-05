---
title: agentrigfind
order: 36
---
`int  agentrigfind(<geometry>geometry, int prim, string transformname)`

Returns `-1` if `transformname` was not found in the rig, `prim` is out of range, or `prim` is not an agent primitive.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`prim`

The primitive number.

`transformname`

The name of a transform in the agent’s rig.

Examples

## examples

Find the current local transform of a given bone.

```vex
int idx = agentrigfind(0, @primnum, "Hips");
if (idx >= 0) {
matrix local_xforms[] = agentlocaltransforms(0, @primnum);
matrix xform = local_xforms[idx];
}

```
