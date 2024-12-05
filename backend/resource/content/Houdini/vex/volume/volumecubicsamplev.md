---
title: volumecubicsamplev
order: 3
---
`vector  volumecubicsamplev(<geometry>geometry, int primnum, vector pos)`

`vector  volumecubicsamplev(<geometry>geometry, string volumename, vector pos)`

`vector  volumecubicsamplev(<geometry>geometry, int primnum, vector pos, matrix3 &grad)`

`vector  volumecubicsamplev(<geometry>geometry, string volumename, vector pos, matrix3 &grad)`

`vector  volumecubicsamplev(<geometry>geometry, int primnum, vector pos, matrix3 &grad, matrix3 &hessX, matrix3 &hessY, matrix3 &hessZ)`

`vector  volumecubicsamplev(<geometry>geometry, string volumename, vector pos, matrix3 &grad, matrix3 &hessX, matrix3 &hessY, matrix3 &hessZ)`

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

Returns

The volume primitive’s sampled value at the given position. Values between voxels are evaluated with tri-cubic interpolation.

The `grad` is a matrix whose i-th column is the gradient of the i-th component of the volume.

Matrices `hessX`, `hessY`, `hessZ` are second derivatives of x, y and z component respectively.

Returns 0 if `primnum` or `inputnum` is out of range, the geometry is invalid, or the given primitive is not a volume or vdb primitive.

Examples

## examples

Approximating a volume value at the point `P + u` using volume values at the point `P`.

```vex
vector  P = {1.0, 2.0, 3.0};
matrix3 grad, hessX, hessY, hessZ;
vector val1 = volumecubicsamplev(0, "vel", P, grad, hessX, hessY, hessZ));

vector u = {0.1, 0.01, 0.001};
vector val2 = volumecubicsamplev(0, "vel", P + u);

// By Taylor expansion we have:
// `val1 + u * grad` is approximately equal to `val2`

// And the second order approximation:
// `val1 + u * grad + 0.5 * set(dot(u, u*hessX), dot(u, u*hessY), dot(u, u*hessZ))`
// is appriximately equal to `val2`

```
