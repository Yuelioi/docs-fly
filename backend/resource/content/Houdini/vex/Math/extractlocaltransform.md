---
title: extractlocaltransform
order: 28
---
| Since | 18.0 |
| --- | --- |

`matrix  extractlocaltransform(matrix world, matrix parent_world, matrix parent_local, int scale_inherit_mode)`

Returns a new local transform given its world and new parent transforms.

`matrix  extractlocaltransform(matrix world, matrix parent_world, matrix parent_local, int mode, matrix &effective_local_transform)`

Returns a new world transform given its local and parent world transforms. The local transform including any inherited scales is stored in the effective_local_transform matrix;

Show/hide arguments

`scale_inherit_mode`

Specifies how scale inheritance from the parent transform is applied to the result. It is one of the following defines from `math.h`:

- `SCALE_INHERIT_DEFAULT` (0) - simple inheritance:

```vex
world = local * parent_world

```

- `SCALE_INHERIT_OFFSET_ONLY` (1) - child doesn’t scale with the parent local scales, but local translation is scaled:

```vex
world = local_scale_rotates * invert(parent_local_scales) * local_translates * parent_world

```

- `SCALE_INHERIT_OFFSET_AND_SCALE` (2) - local translation is scaled as before but parent local scaling is also reapplied by the child in local space:

```vex
world = parent_local_scales * local_scale_rotates * invert(parent_local_scales) * T * parent_world

```

- `SCALE_INHERIT_SCALE_ONLY` (3) - local translation is not scaled, but parent local scaling is reapplied by the child in local space:

```vex
world = parent_local_scales * local * invert(parent_local_scales) * parent_world

```

- `SCALE_INHERIT_IGNORE` (4) - child completely ignores any parent local scaling:

```vex
world = local * invert(parent_local_scales) * parent_world

```
