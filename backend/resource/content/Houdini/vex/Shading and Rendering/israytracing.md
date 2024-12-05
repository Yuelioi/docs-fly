---
title: israytracing
order: 44
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`int  israytracing()`

Returns true when the current shader is shading ray hits. For micropolygon
rendering, the return value will be false. This function can be used to
disambiguate shading styles for renders that use the vm_rayshade property -
where only some objects are ray traced.
