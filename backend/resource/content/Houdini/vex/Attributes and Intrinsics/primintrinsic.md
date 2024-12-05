---
title: primintrinsic
order: 53
---
Intrinsic values are similar to attributes, but are computed on-demand by Houdini rather than stored.

`<type> primintrinsic(<geometry>geometry, string intrinsic_name, int prim_num)`

`<type>[] primintrinsic(<geometry>geometry, string intrinsic_name, int prim_num)`

Intrinsic values are similar to attributes, but are computed on-demand by Houdini rather than stored.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`intrinsic_name`

The name of the intrinsic to read. For example, `"pointattributes"`, `"pointcount"`, or `"bounds"`.

`prim_num`

The number of the primitive to read the given intrinsic attribute for.

Returns

The value of the intrinsic attribute, or `0` if the intrinsic does not exist.
