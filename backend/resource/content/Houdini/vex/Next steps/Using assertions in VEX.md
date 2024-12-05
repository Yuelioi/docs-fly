---
title: Using assertions in VEX
order: 6
---
When writing/debugging VEX code, you can set the `HOUDINI_VEX_ASSERT` environment variable, and then use the `assert(condition)` macro to test invariants in your code.

The benefit of using `assert()` over [printf](functions/printf.html "Prints values to the console which started the VEX program.") is that when the environment variable is not set, the condition will not run and the assertions will likely be optimized away. You can leave the asserts in your code and they won’t affect the production performance of your VEX code.

Note that the `assert()` macro simply prints a message to the console. It does not stop execution of the program.

You must include `assert.h` at the top of the file to import the `assert()` macro:
Include `assert.h` and then use the `assert` macro to test invariants

```vex
#include "assert.h"

cvex
test_assert(string texture_map="")
{
    assert(1 == 2);
    assert(0 == 0);
    assert(texture_map != "");
}

```

Set `HOUDINI_VEX_ASSERT` to enable the assertions

```vex
% export HOUDINI_VEX_ASSERT=1

% vexexec test_assert.vfl
VEX Assertion Failed ./test_assert.vfl:6 - (1 == 2)
VEX Assertion Failed ./test_assert.vfl:8 - (texture_map != "")

% vexexec test_assert.vfl texture_map Mandril.rat
VEX Assertion Failed ./test_assert.vfl:6 - (1 == 2)

```

The [assert_enabled](functions/assert_enabled.html "Returns 1 if the VEX assertions are enabled (see HOUDINI_VEX_ASSERT) or 0 if assertions are disabled. Used the implement the assert macro.") function tests whether the `HOUDINI_VEX_ASSERT` environment variable is set. You could use that function to write your own assert macro (for example, you might write a macro that used your studio’s logging infrastructure).
