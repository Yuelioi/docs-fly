---
title: VEX language reference
order: 7
---
| On this page | * [Contexts](#contexts) * [Statements](#statements) * [Built-in functions](#built-in-functions) * [User-defined functions](#functions)   + [Notes](#notes) * [Main (context) function](#main-context-function)   + [User interface pragmas](#user-interface-pragmas) * [Operators](#operators)   + [Dot operator](#dot-operator)  + [Comparisons](#comparisons)  + [Precedence table](#precedence)  + [Operator type interactions](#operator-type-interactions) * [Data types](#data-types) * [Structs](#structs)   + [Struct functions](#methods) * [Mantra Specific Types](#mantratypes) * [Type casting](#type-casting)   + [Variable casting](#variable-casting)  + [Function casting](#function-casting) * [Comments](#comments) * [Reserved Keywords](#reserved) |
| --- | --- |

Contexts

## contexts

VEX programs are written for a specific *context*. For example, a shader that
controls the surface color of an object is written for the `surface` context.
A shader that determines the illuminance from a light is written for the
`light` context. A VEX program that creates or filters channel data is
written for the `chop` context.

The context affects which functions, statements, and global variables are
available.

See [VEX contexts](contexts/index.html "Guide to the different contexts in which you can write VEX
programs.") for an overview of the ways in which
you can use VEX.

If you are writing for a shading context (surface, displacement, light, etc.),
you should also read the [shading context specific information](contexts/shading_contexts.html).

Statements

## statements

VEX supports the usual [statements](statement.html) familiar from C. It also supports
shading-specific statements such as the [illuminance](functions/illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.") and [gather](functions/gather.html "Sends rays into the scene and returns information from the shaders of
surfaces hit by the rays.") loops
that are only available in certain contexts.

Built-in functions

## built-in-functions

VEX contains a large library of [built-in functions](functions/index.html). Some functions
are only available in certain contexts.

See [VEX functions](functions/index.html).

User-defined functions

## functions

Functions are defined similarly to C: specify the return type,
the function name, and parenthesized list of arguments, followed by
the code block.

Arguments of the same type can be declared in a comma separated list without
re-declaring the type. Other arguments must be separated by a semi-colon.

```vex
int test(int a, b; string c) {
    if (a > b) {
        printf(c);
    }
}

```

You can *overload* functions with the same name but different argument
signatures *and/or return type*.

You can introduce a function definition with the optional `function` keyword to avoid type ambiguity.

```vex
function int test(int a, b; string c) {
    if (a > b) {
        printf(c);
    }
}

```

```vex
void print(basis b) { 
    printf("basis: { i: %s, j: %s, k: %s }\n", b.i, b.j, b.k); 
} 
void print(matrix m) { 
    printf("matrix: %s\n", m); 
} 
void print(bases b) { 
    printf("bases <%s> {\n", b.description); 
    printf("  "); print(b.m); 
    printf("  "); print(b.n); 
    printf("  "); print(b.o); 
    printf("}\n"); 
} 

basis rotate(basis b; vector axis; float amount) { 
    matrix m = 1; 
    rotate(m, amount, axis); 
    basis result = b; 
    result.i *= m; 
    result.j *= m; 
    result.k *= m; 
    return result; 
} 
void rotate(basis b; vector axis; float amount) { 
    b = rotate(b, axis, amount); 
} 

```

### Notes

notes

- User functions must be declared before they are referenced.
- The functions are in-lined automatically by the [compiler](vcc.html "Overview of how to use the VEX language compiler vcc and its
  pre-processor and pragma statements."), so
  **recursion will not work**. To write a recursive algorithm, you should use
  [shader calls](shadercalls.html) instead.
- As in RenderMan Shading Language, parameters to user functions are always
  passed **by reference**, so modifications in a user function affect the variable
  the function was called with. You can force a shader parameter to be read-only
  by prefixing it with the `const` keyword. To ensure that the user function
  writes to an output parameter, prefix it with the `export` keyword.
- There is no limit on the number of user functions.
- You can have more than one return statement in a function.
- You can access global variables directly (unlike RenderMan Shading
  Language, you do not need to declare them with `extern`). However, we
  recommend you avoid accessing global variables, since this limits your
  function to only work in one context (where those globals exist).
  Instead, pass the global(s) to the function as parameters.
- Functions can be defined inside of a function (nested functions).

Main (context) function

## main-context-function

A VEX program must contain one function whose return type is the name of
the context. This is the main function of the program that is called by mantra.
The compiler expects one context function per file.

This function should do the work (by calling out to built-in and/or
user-defined functions) of calculating any required information and modifying
global variables. You do not use the `return` statement to return a value from
the context function. See the specific [context](contexts/index.html "Guide to the different contexts in which you can write VEX
programs.") pages for the global
variables available in each context.

The arguments to the context function, if any, become the user interface for
the program, for example the parameters of a shading node that references
the VEX program.

If a geometry attribute exists with the same name as a parameter of the context
function, the attribute overrides the parameter’s value. This lets you paint
attributes onto geometry to control VEX code.

```vex
surface
noise_surf(vector clr = {1,1,1}; float frequency = 1;
           export vector nml = {0,0,0})
{
    Cf = clr * (float(noise(frequency * P)) + 0.5) * diffuse(normalize(N));
    nml = normalize(N)*0.5 + 0.5;
}

```

Note
Parameters to context functions are dealt with in a special way with
VEX. It is possible to override a parameter’s value using a geometry
attribute with the same name as the variable. Aside from this
special case, parameters should be considered “const” within the
scope of the shader. This means that it is illegal to modify a
parameter value. The compiler will generate errors if this occurs.

You can used the `export` keyword to flag parameters you do wish
to modify on the original geometry.

### User interface pragmas

user-interface-pragmas
The user interface generated from this program by Houdini will be minimal,
basically just the variable name and a generic text field based on the datatype.
For example, you might want to specify that `frequency` should be a slider
with a certain range, and that `clr` should be treated as a color (giving it a
color picker UI). You can do this with [user interface compiler pragmas](pragmas.html).

```vex
#pragma opname        noise_surf
#pragma oplabel        "Noisy Surface"

#pragma label    clr            "Color"
#pragma label    frequency    "Frequency"

#pragma hint    clr            color
#pragma range    frequency    0.1 10

surface noise_surf(vector clr = {1,1,1}; float frequency = 1;
           export vector nml = {0,0,0})
{
    Cf = clr * (float(noise(frequency * P)) + 0.5) * diffuse(normalize(N));
    nml = normalize(N)*0.5 + 0.5;
}

```

Operators

## operators

VEX has the standard C operators with C precedence, with the following
differences.

Multiplication is defined between two vectors or points. The multiplication performs an element by element multiplication (rather than a dot or cross product; see [cross](functions/cross.html "Returns the cross product between the two vectors.") and [dot](functions/dot.html "Returns the dot product between the arguments.")).

Many operators are defined for non-scalar data types (i.e. a vector multiplied by a matrix will transform the vector by the matrix).

In ambiguous situations where you combine two different types with an operator, the result has the type of the second (right hand side) value, for example

```vex
int + vector = vector
```

### Dot operator

dot-operator
You can use the dot operator (`.`) to reference individual
components of a vector, matrix or `struct`.

For vectors, the component names are fixed.

- `.x` or `.u` to reference the first element of `vector2`.
- `.x` or `.r` to reference the first element of `vector` and `vector4`.
- `.y` or `.v` to reference the second element of `vector2`.
- `.y` or `.g` to reference the second element of `vector` and `vector4`.
- `.z` or `.b` to reference the third element. of `vector` and `vector4`
- `.w` or `.a` to reference the fourth element of a `vector4`.

The choice
of the letters u,v/x,y,z/r,g,b is arbitrary; the same letters apply even
if the vector doesn’t hold a point or color.

For matrices, you can use a pair of letters:

- `.xx` to reference the `[0][0]` element
- `.zz` to reference the `[2][2]` element
- `.ax` to reference the `[3][0]` element

In addition, the dot operator can be used to “swizzle” components of a vector. For example

- `v.zyx` is equivalent to `set(v.z, v.y, v.x)`
- `v4.bgab` is equivalent to `set(v4.b, v4.g, v4.a, v4.b)`

Note
You cannot assign to a swizzled vector, only read from them. So you cannot do `v.zyx = b`, but instead must do `v = b.zyx`.

### Comparisons

comparisons
The comparison operators (==, !=, \<, \<=, >, >=) are defined when the
left hand of the operator is the same type as the right hand side,
for string, float and integer types only. The operations result in
integer types.

The string matching operator (~=) is only defined when there is a string
on both sides of the operator, and is equivalent to calling the
[match](functions/match.html "This function returns 1 if the subject matches the pattern specified,
or 0 if the subject doesn’t match.") function with those two values.

The logical (&&, ||, and !) and bitwise (& |, ^, and ~) operators
are only defined for integers.

### Precedence table

precedence
Operators higher in the table have higher precedence.

| Order | Operator | Associativity | Description |
| --- | --- | --- | --- |
| 15 | `()` | LtR | Function call, expression grouping, structure member. |
| 13 | `!` | LtR | Logical negation |
| 13 | `~` | LtR | One’s complement |
| 13 | `+` | LtR | Unary plus (for example, `+5`) |
| 13 | `-` | LtR | Unary minus (for example, `-5`) |
| 13 | `++` | LtR | Increment (for example, `x++`) |
| 13 | `--` | LtR | Decrememt (for example, `x--`) |
| 13 | `(type)` | LtR | Type cast (for example, `(int)x`). |
| 12 | `*` | LtR | Multiplication |
| 12 | `/` | LtR | Division |
| 12 | `%` | LtR | Modulus |
| 11 | `+` | LtR | Addition |
| 11 | `-` | LtR | Subtraction |
| 10 | `<` | LtR | Less-than |
| 10 | `>` | LtR | Greater than |
| 10 | `<=` | LtR | Less-than or equal |
| 10 | `>=` | LtR | Greater than or equal |
| 9 | `==` | LtR | Equal |
| 9 | `!=` | LtR | NOT Equal |
| 9 | `~=` | LtR | String matches |
| 8 | `&` | LtR | Bitwise AND |
| 7 | `^` | LtR | Bitwise XOR |
| 6 | `|` | LtR | Bitwise OR |
| 5 | `&&` | LtR | Logical AND |
| 4 | `||` | LtR | Logical OR |
| 3 | `?:` | LtR | Ternary conditional (for example, `x ? "true" : "false"`) |
| 2 | `=` `+=` `-=` `*=` `/=` `%=` `&=` ```vex |=``^= ``` | RtL | Variable assignment |
| 1 | `,` | LtR | Argument separator |

### Operator type interactions

operator-type-interactions

- When you apply an operation to a `float` and an `int`, the result is the type to the left of the operator. That is, `float * int = float`, while `int * float = int`.
- If you add, multiply, divide, or subtract a vector with a scalar value (`int` or `float`), VEX returns a vector of the same size, with the operation applied component-wise. For example:

```
{1.0, 2.0, 3.0} * 2.0 == {2.0, 4.0, 6.0}

```vex

- If you add, multiply, divide, or subtract vectors of different size, VEX returns a vector of the larger size. The operation is applied component-wise.
  **Important**: the “missing” component(s) on the smaller vector are filled in as `{0.0, 0.0, 0.0, 1.0}`

```
{1.0, 2.0, 3.0} * {2.0, 3.0, 4.0, 5.0} == {2.0, 6.0, 12.0, 5.0}

```vex

This can give surprising results if you're not expecting it, for example:

```
// Third element of the vector2 is treated as 0,
// but fourth element is treated as 1.0
{1.0, 2.0} + {1.0, 2.0, 3.0, 4.0} == {2.0, 4.0, 3.0, 5.0}

```vex

If you're combining different-sized vectors, you might want to break out the components and operate on them “manually” to get the results you want without surprises.

Data types

## data-types

Warning
By default, VEX uses 32 bit integers. If you use the [![](../icons/SOP/attribcast.svg)AttribCast SOP](../nodes/sop/attribcast.html "Changes the size/precision Houdini uses to store an attribute.") to cast a geometry attribute to 64 bits, VEX will silently discard the extra bits if you manipulate the attribute in VEX code.

The VEX engine runs in either 32bit or 64bit mode. In 32bit mode, all
floats, vectors, and integers are 32bit. In 64bit mode, they are 64bit.
There is no `double` or `long` type to allow mixed precision math.

You can use underscores to break up long numbers.

| Type | Definition | Example |
| --- | --- | --- |
| `int` | Integer values | `21, -3, 0x31, 0b1001, 0212, 1_000_000` |
| `float` | Floating point scalar values | `21.3, -3.2, 1.0, 0.000_000_1` |
| `vector2` | Two floating point values. You might use this to represent texture coordinates (though usually Houdini uses vectors) or complex numbers | `{0,0}, {0.3,0.5}` |
| `vector` | Three floating point values. You can use this to represent positions, directions, normals or colors (RGB or HSV) | `{0,0,0}, {0.3,0.5,-0.5}` |
| `vector4` | Four floating point values. You can use this to represent positions in homogeneous coordinates, or color with alpha (RGBA). It is often used to represent a quaternion. Quaternions in VEX are laid out in `x/y/z/w` order, not `w/x/y/z`. This applies both to both quaternions and positions with homogeneous coordinates. | `{0,0,0,1}, {0.3,0.5,-0.5,0.2}` |
| `array` | A list of values. See [arrays](arrays.html) for more information. | `{ 1, 2, 3, 4, 5, 6, 7, 8 }` |
| `struct` | A fixed set of named values. See [structs](lang.html#structs) for more information. |
| `matrix2` | Four floating point values representing a 2D rotation matrix | `{ {1,0}, {0,1} }` |
| `matrix3` | Nine floating point values representing a 3D rotation matrix or a 2D transformation matrix | `{ {1,0,0}, {0,1,0}, {0,0,1} }` |
| `matrix` | Sixteen floating point values representing a 3D transformation matrix | `{ {1,0,0,0}, {0,1,0,0}, {0,0,1,0}, {0,0,0,1} }` |
| `string` | A string of characters. See [strings](strings.html) for more information. | `"hello world"` |
| `dict` | A dictionary mapping `string`s to other VEX data types. See [dicts](dicts.html) for more information. |
| `bsdf` | A *bidirectional scattering distribution function*. See [writing PBR shaders](pbr.html) for information on BSDFs. |

Structs

## structs

As of Houdini 12, you can define new structured types using the `struct` keyword.

Member data can be assigned default values in the struct definition similar to C++11 member initialization.

Two implicit constructor functions are created for each struct. The first takes initialization arguments in the order they are declared in the struct, the second takes no arguments but sets all members to their default values.

```
#include <math.h> 

struct basis { 
    vector i, j, k; 
} 

struct bases { 
    basis m, n, o; 
    string description; 
} 

struct values {
    int uninitialized;        // Uninitialized member data
    int        ival = 3;
    float fval = 3.14;
    float aval[] = { 1, 2, 3, 4.5 };
}

basis rotate(basis b; vector axis; float amount) { 
    matrix m = 1; 
    rotate(m, amount, axis); 
    basis result = b; 
    result.i *= m; 
    result.j *= m; 
    result.k *= m; 
    return result; 
}

// Declare struct variables
basis b0;        // Initialize using default values (i.e. 0 in this case)
basis b1 = basis({1,0,0}, {0,1,0}, {0,0,1});        // Initialize using constructor
basis b2 = { {1,0,0}, {0,1,0}, {0,0,1} };         // Initialize as explicit struct

// You can use M_PI or PI
b1 = rotate(b1, {0,0,1}, M_PI/6);

```vex

Note
You must define structs before using them in the source file.

### Struct functions

methods
You can define functions inside structs to organize your code and
allow a limited form of object-oriented programming.

- Inside a struct function, you can use `this` to refer to the struct instance.
- Inside a struct function, you can refer to struct fields by name as if they were variables
  (for example, `basis` is a shortcut for `this.basis`).
- You can call struct functions on a struct instance using the `->` arrow operator,
  for example `sampler->sample()`.
  Note inside a struct function that you can call other methods on the struct using `this->method()`.

```
struct randsampler {
    // Fields
    int        seed;

    // Methods
    float sample()
    {
        // Struct functions can refer to fields by name
        return random(seed++);
    }
} 

cvex shader()
{
    randsampler sampler = randsampler(11);
    for (int i = 0; i < 10; i++)
    {
        // Use -> to call methods on struct instances
        printf("%f\n", sampler->sample());
    }
}

```vex

Mantra Specific Types

## mantratypes

Mantra has some pre-defined struct types that are used in shading-specific functions.

| `light` | Defined in mantra shading contexts only. This is a struct representing a  handle to a light source. The struct has methods:\* illuminate(…)  Invokes the VEX surface shader bound to the `vm_illumshader` property of  the light source.  In an IFD, you may see lines like `ray_property light illumshader diffuselighting`  or `ray_property light illumshader mislighting misbias 1.000000`. These statements define the shader invoked when the `illuminate()` method  is called on a light object. |
| --- | --- |
| `material` | Defined in mantra shading contexts only. This is opaque struct  representing the material assigned to an object. |
| `lpeaccumulator` | Defined in mantra shading contexts only. This is a struct representing an  accumulator for Light Path Expressions. The struct has methods:\* begin() - Construct and initialize accumulator. * end() - Finalize and destroy. * move(string eventtype; string scattertype; string tag, string bsdflabel) - Modifies internal state based on current event. If an empty string is  passed in, it’s assumed to be 'any'. * pushstate() - Push internal state onto stack. * popstate() - Pop internal state from stack. Used to pushstate() to “undo” move(). * int matches() - Returns non-zero if the current internal state matches any of the light  path expressions defined by user. * accum(vector color, …) - Accumulates input color onto intermediate buffer. Also takes in optional prefix strings to be compared against prefixes that were declared with LPE image plane. All prefixes must match in order to accumulate. * flush(vector multiplier) - Multiply the intermediate buffer by the multiplier and add it onto image planes. The intermediate buffer exists to allow for variance anti-aliasing (i.e. the multiplier would be 1/number_of_samples). * int getid() - Returns an integer id assigned to lpeaccumulator. * lpeaccumulator getlpeaccumulator(int id) - Returns lpeaccumulator based on id. Used with getid() to pass lpeaccumulator across shader boundaries. |

Type casting

## type-casting

### Variable casting

variable-casting
This is similar to type casting in C++ or Java: you transform a
value of one type into another (for example, an int into a float).

This is sometimes necessary, as when you have the following:

```
int a, b;
float c;
c = a / b;

```vex

In this example, the compiler will do integer division (see [type
resolution](lang.html#typeres) ). If you wanted to do floating
point division instead, you need to explicitly cast `a` and `b` as
floats:

```
int a, b;
float c;
c = (float)a / (float)b;

```vex

This generates additional instructions to perform the casts. This
may be an issue in performance-sensitive sections of your code.

### Function casting

function-casting
VEX dispatches functions based not only on the types of the
arguments (like C++ or Java), but also on the return type. To
disambiguate calls to functions with the same argument types but
different return types, you can *cast the function*.

For example, the [noise](functions/noise.html "There are two forms of Perlin-style noise: a non-periodic noise which
changes randomly throughout N-dimensional space, and a periodic form
which repeats over a given range of space.") function can take different
parameter types, but can also return different types: `noise` can
return either a float or vector.

In the code:

```
float n;
n = noise(noise(P));

```vex

…VEX could dispatch to either `float noise(vector)` or
`vector noise(vector)`.

To cast a function call, surround it with `typename( ... )`, as
in:

```
n = noise( vector( noise(P) ) );

```vex

While this looks like a function call, it does nothing but disambiguate the
function call inside and has no performance overhead.

Function casting is implied when you assign a function call directly to a
variable of a specified type. So the following expressions are equivalent,
and the function cast may be omitted for more concise code:

```
vector n = vector( noise(P) );        // Unnecessary function cast
vector n = noise(P);

```vex

Note
If VEX is unable to determine which signature of a function you are
trying to call, it will trigger an ambiguity error and print out the
candidate functions. You should then choose the appropriate return
value and add a function cast to select it.

Since function casting does not generate any type conversions (it simply
chooses a function to call), there is no performance penalty to using it.
A good rule of thumb is to use function casting wherever possible and to
use variable casting only when an explicit type conversion is required.

Comments

## comments

VEX uses C++ style comments:

- One-line comments are preceded by `//`
- Freeform comments begin with `/*` and end with `*/`

Reserved Keywords

## reserved

`break`, `bsdf`, `char`, `color`, `const`, `continue`, `do`, `dict`, `else`,
`export`, `false`, `float`, `for`, `forpoints`, `foreach`, `gather`,
`hpoint`, `if`, `illuminance`, `import`, `int`, `integer`, `matrix`,
`matrix2`, `matrix3`, `normal`, `point`, `return`, `string`, `struct`, `true`,
`typedef`, `union`, `vector`, `vector2`, `vector4`, `void`, `while`
