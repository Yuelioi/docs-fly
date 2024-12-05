---
title: opid
order: 26
---
| Since | 17.5 |
| --- | --- |

`int  opid(string op_path)`

`int  opid(int op_id)`

Resolves an operator path and returns its corresponding operator id.

Return -1 on failure. The form that takes an operator id as an input returns 1
if it is a valid operator id, otherwise it returns 0. This can be used to test
if an operator id is valid.
