---
title: shr
order: 69
---
`int  shr(int a, int bits)`

Bit-shifts `a` to the right by `bits`.

This is an arithmetic shift, the sign is shifted with it. Thus, `shr(-1, 2)`
will give -1, not zero.
