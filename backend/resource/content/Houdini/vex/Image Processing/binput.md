---
title: binput
order: 3
---
| Context(s) | [cop2](../contexts/cop2.html) |
| --- | --- |

`vector  binput(int u, int v, ...)`

`vector4  binput(int u, int v, ...)`

`vector  binput(float u, float v, ...)`

`vector4  binput(float u, float v, ...)`

Sample from the currently cooking plane on input 0 at the current frame.

`float  binput(int comp, int u, int v, ...)`

`float  binput(int comp, float u, float v, ...)`

Sample one component from the given component index of the currently cooking plane on input 0 at the current frame.

`vector  binput(int opinput, int plane, int u, int v, ...)`

`vector4  binput(int opinput, int plane, int u, int v, ...)`

`vector  binput(int opinput, int plane, float u, float v, ...)`

`vector4  binput(int opinput, int plane, float u, float v, ...)`

Sample from the given input/plane at the current frame.

`float  binput(int opinput, int plane, int comp, int u, int v, ...)`

`float  binput(int opinput, int plane, int comp, float u, float v, ...)`

Sample one component from the given input/plane/component at the current frame.

`float  binput(int opinput, int plane, int array_index, int comp, int u, int v, int frame, ...)`

`float  binput(int opinput, int plane, int array_index, int comp, float u, float v, int frame, ...)`

Sample one component from the given input/plane/component at the given frame.

`vector  binput(int opinput, int plane, int array_index, int u, int v, int frame, ...)`

`vector4  binput(int opinput, int plane, int array_index, int u, int v, int frame, ...)`

`vector  binput(int opinput, int plane, int array_index, float u, float v, int frame, ...)`

`vector4  binput(int opinput, int plane, int array_index, float u, float v, int frame, ...)`

Sample from the given input/plane/component at the given frame.

Show/hide arguments

`opinput`

The input number to read the pixels from. Versions that don’t specify this always use the first input (0).

`plane`

The index of a plane in the input.
Versions that don’t specify this always use the currently cooking plane.

`array_index`

For use if the plane has array values. Generally, just pass `0` here.

`comp`

The index of a component within the plane. For example, 0 for red, 1 for green, 2 for blue in an RGB plane.
The versions that return vectors do not take this argument and return all components at once.

`u`, `v`

If you give floating point UVs, the values are interpreted as unit (0-1) values. For example, `0.5, 0.5` would be the center of the image.
If you give integer UVs, the values are in pixels, ranging from `0,0` to `XRES-1, YRES-1`.

`frame`

Frame number to sample at.
Versions that don’t specify this always use the current frame.

Returns

A float, vector or vector4 value. If the channel does not exist, returns 0.
Whenever possible, use the vector versions rather than reading individual components separately.

See [COP pixel sampling functions](../cop2_sample_suite.html) for more information.
