---
title: primuvconvert
order: 55
---
`vector2  primuvconvert(<geometry>geometry, vector2 uv, int prim_num, int mode)`

`vector2  primuvconvert(<geometry>geometry, vector2 uv, int prim_num, int mode, float tolerance)`

`float  primuvconvert(<geometry>geometry, float uv, int prim_num, int mode)`

`float  primuvconvert(<geometry>geometry, float uv, int prim_num, int mode, float tolerance)`

Show/hide arguments

`geometry`

A string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`uv`

Curve coordinates to convert. Can be a float or a vector2. The function returns the converted coordinates.

`prim_num`

The primitive number of the curve on which to convert the coordinates.

`mode`

One of the `PRIMUV_space_TO_space` constants listed below. You can import the constants from `$HFS/houdini/vex/include/math.h`.

REAL domain is based on the number of curve segments (0 to `nSegments`). A segment can hold multiple control points based on the curve degree. `UNIT` domain is the REAL domain normalized to fit in 0 to 1. `UNITLEN` domain maps the curve based on its length but normalized (0..1). `LEN` domain maps the curve based on its length (0..`CurveLength`).

| Constant name | Int value |
| --- | --- |
| PRIMUV_REAL_TO_UNIT | 0 |
| PRIMUV_REAL_TO_UNITLEN | 1 |
| PRIMUV_REAL_TO_LEN | 2 |
| PRIMUV_UNIT_TO_REAL | 3 |
| PRIMUV_UNIT_TO_UNITLEN | 4 |
| PRIMUV_UNIT_TO_LEN | 5 |
| PRIMUV_UNITLEN_TO_REAL | 6 |
| PRIMUV_UNITLEN_TO_UNIT | 7 |
| PRIMUV_UNITLEN_TO_LEN | 8 |
| PRIMUV_LEN_TO_REAL | 9 |
| PRIMUV_LEN_TO_UNIT | 10 |
| PRIMUV_LEN_TO_UNITLEN | 11 |
