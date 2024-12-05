---
title: rawcolormap
order: 9
---
`vector|vector4 rawcolormap(string filename, vector uvw, ...)`

`vector|vector4 rawcolormap(string filename, float u, float v, ...)`

`vector|vector4 rawcolormap(string filename, vector uv, vector du, vector dv, int samples, ...)`

`vector|vector4 rawcolormap(string filename, vector uv0, vector uv1, vector uv2, vector uv3, ...)`

`vector|vector4 rawcolormap(string filename, vector uv0, vector uv1, vector uv2, vector uv3, int samples, ...)`

`vector|vector4 rawcolormap(string filename, float u0, float v0, float u1, float v1, float u2, float v2, float u3, float v3, int samples, ...)`

This function has the same arguments as [colormap](colormap.html "Looks up a (filtered) color from a texture file."), but does not do bilinear interpolation of the pixel values. See [colormap](colormap.html "Looks up a (filtered) color from a texture file.") for information on the arguments.

If you call the function with a `vector4` return type, the fourth component is the alpha channel of the texture. If the image does not have alpha, the fourth component is always `1`.
