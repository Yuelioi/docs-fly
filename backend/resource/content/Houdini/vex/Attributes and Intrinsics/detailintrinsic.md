---
title: detailintrinsic
order: 19
---
Intrinsic values are similar to attributes, but are computed on-demand by Houdini rather than stored.

`<type> detailintrinsic(<geometry>geometry, string intrinsic_name)`

`<type>[] detailintrinsic(<geometry>geometry, string intrinsic_name)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`intrinsic_name`

The name of the intrinsic to read. For example, `"pointattributes"`, `"pointcount"`, or `"bounds"`.

Returns

The intrinsic value, or `0` if the given intrinsic does not exist.
