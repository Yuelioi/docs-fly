---
title: islpeactive
order: 43
---
| Context(s) | [surface](../contexts/surface.html) |
| --- | --- |

`int  islpeactive()`

Returns 1 if there are any image planes with vex variable that starts with
`"lpe:"`, which indicates that Light Path Expression (LPE) is enabled. Returns 0
if no such image plane exists.

This function is used to optimize out any LPE related function calls if the
current render doesn’t have any AOVs that uses LPE.
