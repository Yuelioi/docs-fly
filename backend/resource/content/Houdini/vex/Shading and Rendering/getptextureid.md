---
title: getptextureid
order: 28
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`int  getptextureid()`

Returns the ptexture id for the current face being shaded. This will typically
be the same as `getprimid()` except in the case of subdivision surfaces. For
subdivision surfaces, mantra splits non-quadrilateral faces into multiple
patches. Each of these split faces is assigned a unique ptexture id.
