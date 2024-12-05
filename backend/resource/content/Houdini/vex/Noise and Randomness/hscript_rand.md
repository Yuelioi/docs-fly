---
title: hscript_rand
order: 14
---
`float|vector|vector4 hscript_rand(float seed)`

Produces the exact same results as the Houdini expression function of
the same name. This function will generate different random values for
every floating point seed. This is different that the
[random](random.html "Generate a random number based on the integer position in 1-4D space.") function which converts the floating point argument
to an integer seed. The `hscript_rand()` function may produce different
results on different hardware or operating systems.
