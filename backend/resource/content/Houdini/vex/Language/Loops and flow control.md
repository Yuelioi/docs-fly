---
title: Loops and flow control
order: 3
---
| On this page | * [{}](#) * [do loop](#do-loop) * [for loop](#for-loop) * [foreach loop](#foreach-loop) * [while loop](#while-loop) * [Other looping statements](#other-looping-statements) * [if](#if) * [return](#return) * [break](#break) * [continue](#continue) |
| --- | --- |
See also [VEX functions](functions/index.html). Most of the work in VEX is done with function calls. Most of the statements are looping constructs, many of which may be familiar from other languages such as C. While `print` is a statement in some languages (such as Python), in VEX you print using [the printf function](functions/printf.html "Prints values to the console which started the VEX program.").
{}

## [¶](#)

As in C and many other languages, you can enclose multiple statements inside curly braces to act as a block.

For example, the `if` statement can execute one statement:

```vex
if ( needs_zapping() ) zap()

```

…or a block inside curly braces:

```vex
if ( needs_zapping() ) {
    zap()
    disintegrate()
    remove_dust()
}

```

do loop

## do-loop

`do statement [while (condition)]`

Executes statement, and then loops if condition is true. Unlike `while`, `do` is guaranteed to execute the statement at least once.

for loop

## for-loop

`for (init; condition; change) statement`

Standard C-style `for` loop. Performs the init statement, then executes statement repeatedly while condition is true, executing the change statement at the end of each iteration.

foreach loop

## foreach-loop

`foreach (value; array) statement`

`foreach (index, value; array) statement`

Executes statement for each member of array (optionally sets
index to the current position in the array). See [foreach](functions/foreach.html "Loops over the items in an array, with optional enumeration.").

while loop

## while-loop

`while (condition) statement`

Executes statement repeatedly while condition is true.

Other looping statements

## other-looping-statements

The [forpoints](functions/forpoints.html), [illuminance](functions/illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables."), and [gather](functions/gather.html "Sends rays into the scene and returns information from the shaders of
surfaces hit by the rays.") statements let you loop over data being processed by VEX.

if

## if

`if (condition) statement_if_true [else statement_if_false]`

Executes statement_if_true if condition is true.

If the `else` clause is included, statement_if_false is executed if condition is false.

return

## return

Exits the function with an optional return value.

```vex
int max(int a, b) {
    if (a > b) {
        return a;
    }
    return b;
}

```

break

## break

`break` immediately exits the loop. It is useful with the `if` statement to stop looping early when some condition is reached.

```vex
for (int i = 0; i < sizes; i++)
{
    mixamount += getAmount(roughness);
    if (mixamount > 1) {
        break;
    }
}

```

continue

## continue

`continue` jumps immediately to the next iteration of the loop.

```vex
foreach (x; myarray) {
    if (x < 10) continue;
    ...
}

```
