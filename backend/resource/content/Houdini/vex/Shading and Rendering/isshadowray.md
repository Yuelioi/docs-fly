---
title: isshadowray
order: 46
---
| Context(s) | [surface](../contexts/surface.html) |
| --- | --- |

`int  isshadowray()`

Returns 1 if the shader is being called to evaluate opacity for shadow
rays, or 0 if the shader is being called to evaluate for surface color.

Use this function to compute different opacity if the surface is
shadowing another surface.
