---
title: shrz
order: 70
---
`int  shrz(int a, int bits)`

Bit-shifts `a` to the right by `bits`. This is a zero-extend shift, so
new bits are always zero. Thus, `shrz(-1, 2)` is zero, not -1.
