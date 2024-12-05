---
title: VEX cookbook
order: 7
---
| On this page | * [VEX examples](#vex_cbook_eg) * [Constant Color COP](#vex_cbook_ccolor) * [User noise functions](#vex_cbook_unoise) |
| --- | --- |

VEX examples

## vex_cbook_eg

To see the VEX implementation of a VEX-based operator:

1. Add the operator to your network.
1. Open the tile menu and choose **Type Properties**.
   Other example VEX source code is available on the web.

Constant Color COP

## vex_cbook_ccolor

You could write a “Constant Color COP” function this way:

```vex
cop2
constant(vector clr=1; float alpha=1)
{
R = clr.r;
G = clr.g;
B = clr.b;
A = alpha;
}

```

However, you should use the high-efficiency [assign](functions/assign.html "An efficient way of extracting the components of a vector or matrix into float variables.")
function whenever possible.

```vex
cop2
constant(vector clr=1; float alpha=1)
{
assign(R, G, B, clr);
A = alpha;
}

```

User noise functions

## vex_cbook_unoise

This example defines two user functions which generate normalized
multiple octaves of Perlin noise.

```vex
float
myfperlin3d(vector pos; int octaves; float rough)
{
int    i;
float    nval;
vector    pp;
float    result, sum, scale;
// Because parameters are passed by reference, we don't really
// want to modify the parameters value.  Therefore, we copy it
// to a temporary variable.
pp = pos;
scale = 1;
sum = 0;
result = 0;
for (i = 0; i <= octaves; i++)
{
    result += noise(pp);
    sum += scale;
    pp *= 2;
    scale *= rough;
}
return result / sum;
}
float
myfperlin1d(float pos; int octaves; float rough)
{
int    i;
float    nval;
float    pp;
float    result, sum, scale;
// Because parameters are passed by reference, we don't really
// want to modify the parameters value.  Therefore, we copy it
// to a temporary variable.
pp = pos;
scale = 1;
sum = 0;
result = 0;
for (i = 0; i <= octaves; i++)
{
    result += noise(pp);
    sum += scale;
    pp *= 2;
    scale *= rough;
}
return result / sum;
}

```

It is possible to use #define macros to encode the contents of the
function and simply call the macro to generate vector or other
dimensions of noise.
