---
title: ch
order: 2
---
`float  ch(string channel)`

`float  ch(string channel, float time_in_sec)`

`string  ch(string channel)`

`string  ch(string channel, float time_in_sec)`

`string  ch(string channel, float time_in_sec)`

`float  ch(int op_id, int parm_index, int vector_index)`

`float  ch(int op_id, int parm_index, int vector_index, float time_in_sec)`

Evaluates a channel (or parameter) and return its value. The time is specified in *seconds*, not in frames. If you don’t specify the time, the function returns the value at the current time.

Houdini includes several functions to evaluate channels/parameters of different types.

- To get a float or string without needing to know the parameter type, use [ch](ch.html "Evaluates a channel (or parameter) and return its value.").
- To get a float, use [chf](chf.html "Evaluates a channel (or parameter) and return its value.").
- To get a string, use [chs](chs.html "Evaluates a channel (or parameter) and return its value.").
- For integer parameters, use [chi](chi.html "Evaluates a channel (or parameter) and return its value.")
- For matrix type parameters, use [ch3](ch3.html "Evaluates a channel (or parameter) and return its value.") or [ch4](ch4.html "Evaluates a channel (or parameter) and return its value.").
- For a ramp parameter, use [chramp](chramp.html "Evaluates a ramp parameter and return its value.") or [chrampderiv](chrampderiv.html "Evaluates the derivative of a parm parameter with respect to position.").
- Use [chid](chid.html "Resolves a channel string (or parameter) and return op_id, parm_index and vector_index.") to get an `op_id`, `parm_index` and `vector_index` to evaluate the channel without having to do string resolution.

Examples

## examples

```vex
// Get the X transform of the box1 SOP at 1.5s into the animation
float tx = ch("/obj/geo1/box1/tx", 1.5)

```
