---
title: chrate
order: 16
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`float  chrate()`

Uses `-1` as the `opinput`.

`float  chrate(int opinput)`

Returns the sample rate of the given input.

Show/hide arguments

`opinput`

The input number to read from, starting from 0. For example, the first input is 0, the second input is 1, and so on.

If you specify `-1`, the function uses the current CHOP node or input `0` if it is connected.
