---
title: hasmetadata
order: 11
---
| Context(s) | [cop2](../contexts/cop2.html) |
| --- | --- |

`int  hasmetadata(int opinput, string name)`

This function checks if metadata named `name` exists on the COP attached to
the VEX COP’s input `opinput`. If it exists then 1 is returned, otherwise 0.

Show/hide arguments

`opinput`

The input number to read from, starting from 0. For example, the first input is 0, the second input is 1, and so on.

`name`

The name of the metadata to check.
