---
title: getgroupid
order: 15
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`int  getgroupid()`

Returns the id of a primitive group containing the current face being shaded.
The id is the index of the group in the detail. If the primitive belongs
to several groups, their indices are added up to calculate the returned id.
