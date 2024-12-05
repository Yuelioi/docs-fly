---
title: chid
order: 12
---
| Since | 17.5 |
| --- | --- |

`void  chid(string channel_path, int &op_id, int &parm_index, int &vector_index)`

`void  chid(string op_path, string channel_name, int &op_id, int &parm_index, int &vector_index)`

`int  chid(int op_id, int parm_index, int vector_index)`

Resolves a channel given by a channel path or operator path and returns

its corresponding operator id, parameter id and vector_index through output
arguments. Return -1 values on failure. You can also use the last overloaded
function that doesn’t take a channel_path to test for validity of the keys.
Returns 1 if the ids are valid, otherwise returns 0.
