---
title: has_udim
order: 5
---
| Since | 17.5 |
| --- | --- |

`int  has_udim(string path)`

Scans the input path for special conversion characters to perform either UDIM or UVTILE style filename expansion. If the string has any UDIM/UVTILE expansion patterns, the function returns 1, otherwise the function returns 0.
