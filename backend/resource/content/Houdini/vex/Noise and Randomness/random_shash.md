---
title: random_shash
order: 35
---
`int  random_shash(string seed)`

Hashes the given string into an integer value.

Note
Two different strings may hash to the same value.

Note
The returned value may be very large and can cause an overflow when further multiplied or is cast to float.
