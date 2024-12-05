---
title: chwritebuf
order: 29
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`int  chwritebuf(int index, float value)`

`int  chwritebuf(int index, vector t, vector r, vector s)`

Writes a value of a CHOP context temporary buffer at the specified index.
Return 1 if the write succeeded, 0 otherwise.
