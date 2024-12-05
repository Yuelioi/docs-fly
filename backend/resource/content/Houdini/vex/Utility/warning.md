---
title: warning
order: 22
---
`void  warning(string format, ...)`

Reports a custom runtime VEX warning. This uses the same format string syntax as [printf](printf.html "Prints values to the console which started the VEX program.").

If something is so problematic that there is no acceptable fallback behavior, it may be worth reporting an [error](error.html "Reports a custom runtime VEX error."), instead of a warning.

Note
It’s quite easy to accidentally report thousands of different warnings.

Examples

## examples

```vex
if (primintrinsic(0,"typeid",@primnum) != 1) {
    warning("Primitives that aren't polygons are being ignored.");
    return;
}
if (primintrinsic(0,"closed",@primnum) == 0 || @numvtx < 3) {
    warning("Open or degenerate polygons are being ignored.");
    return;
}
float minimumValue = chf("min");
float maximumValue = chf("max");
if (minimumValue > maximumValue) {
    warning("Minimum (%f) can't be greater than maximum (%f); replacing minimum with maximum.", minimumValue, maximumValue);
    minimumValue = maximumValue;
}

```
