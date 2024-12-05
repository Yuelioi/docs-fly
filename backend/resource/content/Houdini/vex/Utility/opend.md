---
title: opend
order: 8
---
| Context(s) | [shading](../contexts/shading.html) |
| --- | --- |

`void  opend(int handle)`

Informs mantra that a long operation begun with [opstart](opstart.html "Start a long operation.") has completed. Pass in the value returned by [opstart](opstart.html "Start a long operation.").

```vex
int op_handle = opstart("Performing long operation");
perform_long_operation();
if (op_handle >= 0)
    opend(op_handle);

```
