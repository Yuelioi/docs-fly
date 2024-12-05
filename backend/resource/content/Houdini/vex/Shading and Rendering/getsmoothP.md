---
title: getsmoothP
order: 33
---
| On this page | * [Variadic arguments](#variadic-arguments) * [Examples](#examples) |
| --- | --- |
| Context(s) | [shading](../contexts/shading.html) |
Returns a modified surface position based on a smoothing function.

`int  getsmoothP(vector &smoothP, vector ray_origin, ...)`

Overwrites the `smoothP` variable with the modified surface position.
This function is only meaningful for some primitive types (such as polygons).

`vector  getsmoothP(...)`

Uses the global variables `Eye` and `I` to fill in the ray origin and direction.

Variadic arguments

## variadic-arguments

Show/hide arguments

"style",
`string`

`none`

No smoothing.

`shadow`

Apply a smoothing function appropriate to elimination of the shadow
terminator issue for polygons.

Examples

## examples

```vex
shadow
fastshadow()
{
    vector        surfP;
    if (!getsmoothP(surfP, Eye, I))
        surfP = Ps;                // Set to the Ps (surface P) variable
    vector shad = trace(surfP, normalize(L), Time, "raystyle", "shadow");
    Cl *= ({1,1,1} - shad);
}

```
