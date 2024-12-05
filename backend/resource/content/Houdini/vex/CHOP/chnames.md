---
title: chnames
order: 10
---
| Context(s) | [chop](../contexts/chop.html) |
| --- | --- |

`string [] chnames()`

Uses `-1` as the `opinput`.

`string [] chnames(int opinput)`

Returns an array of channel names in the given input.

Show/hide arguments

`opinput`

The input number to read from, starting from 0. For example, the first input is 0, the second input is 1, and so on.

If you specify `-1`, the function uses the current CHOP node or input `0` if it is connected.
