---
title: mspace
order: 3
---
| Context(s) | [image3d](../contexts/image3d.html) |
| --- | --- |

`vector  mspace(vector P)`

Transforms the position specified into the “local” space of the
metaball. This function is only valid inside the [forpoints](forpoints.html) loop
construct.

An example use of this function would be to compute noise based on a
“rest” position… For example:

```vex
forpoints(P) {
vector npos = mspace(P) - mattrib("rest", P);
nval += noise(npos);
}

```
