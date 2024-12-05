---
title: getscope
order: 32
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`void  getscope(material mat, string raystyle, string &scope, string &categories)`

Given a material handle, determine the objects which are visible for a given raystyle (“diffuse”, “reflect”, or “refract”). The object selection is given by the scope/categories returned.
