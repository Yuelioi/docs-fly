---
title: Shader Calls
order: 4
---
| On this page | * [The import keyword](#the-import-keyword) * [Invoking a shader](#invoking-a-shader) * [Context of the called shader](#context-of-the-called-shader) * [Examples](#examples) |
| --- | --- |
Beginning with Houdini 12.5, VEX shader functions can call other shader
functions. This technique makes it possible to optimize VEX compiler and
optimizer performance for large shaders, since code that is invoked
multiple times within a shader or in other shaders can be built once and
used many times without additional runtime overhead.
The import keyword

## the-import-keyword

The `import` keyword introduces another shader function by name into the
current shader. The imported shader must be accessible in the houdini path
for compilation to succeed - if it isn’t found, shader compilation will
fail. So when constructing shaders that call other shaders, you’ll need to
build the shaders in dependency order - called shaders, then their callers.
Circular calls are possible but you’ll need to add the import keyword to
the callee after building the first caller.

For example, importing the plastic shader:

```vex
import plastic;

```

Shaders can call themselves recursively - in this case, the import keyword
isn’t required.

Invoking a shader

## invoking-a-shader

Shaders are invoked by name and are passed keyword arguments - string/value
pairs that identify the arguments to be passed or received from the invoked
shader. It’s possible to only bind some parameters, in which case the
called shader will use it’s default values for the parameters that are not
bound. Additionally, only a subset of the exports from the called shader
need to be bound. In this case, the VEX optimizer will strip out any dead
code that computes the exports that are not required, leading to improved
performance.

For example, this code invokes the plastic shader asking for the `Cf`
export and providing the `diff` input:

```vex
import plastic;

surface caller(vector diff = {1,0.5,0})
{
    plastic("diff", diff, "Cf", Cf);
}

```

vcc will check all variadic arguments passed to a called shader to ensure
that they correspond with an argument or export that exists in the called
shader’s parameter list - if the type or access mode doesn’t match, an
error will be reported.

Context of the called shader

## context-of-the-called-shader

Shaders can currently only call shaders that have a matching context type.
For contexts with global variables, any global variables that are not
provided explicitly to the shader as keyword arguments are copied unchanged
from the calling shader to the called shader. For contexts that carry
additional opaque state information (such as the surface context, which
maintains state about the hit surface), this information is also maintained
in the called shader so that calling methods like `getraylevel()` will
produce the same result in the caller and callee.

Examples

## examples

The called shader:

```vex
cvex callee(export int mval = 0;
        int rval = 0;
        export int wval = 0;
        float castval = 0)
{
    mval *= 2;
    wval = rval;
}

```

The calling shader:

```vex
import callee;

cvex caller()
{
    int mval = 1;
    int rval = 2;
    int wval = 1;
    callee("mval", mval, "rval", rval, "wval", wval, "castval",
            1);
    printf("%d %d %d\n", mval, rval, wval);
}

```

A recursive shader:

```vex
cvex fib(int i = 0; export int rval = 0)
{
    if (i >= 2)
    {
        int v1, v2;
        fib("i", i-1, "rval", v1);
        fib("i", i-2, "rval", v2);
        rval = v1 + v2;
    }
    else
        rval = i;
    printf("%d: %d\n", i, rval);
}

```
