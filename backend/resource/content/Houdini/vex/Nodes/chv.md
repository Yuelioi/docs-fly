---
title: chv
order: 20
---
`vector  chv(string channel)`

`vector  chv(string channel, float time)`

`vector  chv(int opid, int pid)`

`vector  chv(int opid, int pid, float time)`

Evaluates a channel or parameter, and return its value.

If the parameter is a vector parameter, the base parameter name
can be used to return the all components as a vector.

The time is specified in seconds, not in frames.
