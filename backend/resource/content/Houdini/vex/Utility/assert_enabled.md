---
title: assert_enabled
order: 1
---
`int  assert_enabled()`

Returns 1 if the environment variable `HOUDINI_VEX_ASSERT` is set or 0 if the variable isn’t set.

The `assert()` macro uses this function to only execute assertions when `HOUDINI_VEX_ASSERT` is set:

```vex
#define assert(EXPR)    \
    if (assert_enabled()) { \
        if (!(EXPR)) print_once(sprintf('VEX Assertion Failed %s:%d - (%s)\n', \
                __FILE__, __LINE__, #EXPR)); \
    }

```

You could use this function to write your own assert macro (for example, you might write a macro that used your studio’s logging infrastructure).

See [using assertions in VEX](../assertions.html "You can use the assert() macro to print information while you are debugging VEX code.") for more information.
