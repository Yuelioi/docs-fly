---
title: forpoints
order: 4
---
In the [image3d context](../contexts/image3d.html "Obsolete. Write a program for use with the i3dgen program to generate 3D
textures."), when geometry is specified (i.e. metaball geometry or particles), you can iterate over the metaballs which affect a point in space.

```vex
forpoints ( position [, distance] ) {

}

```

…where the position is a vector representing a point in space. The statement will be executed once for each metaball/particle at the position passed in.

If you specify the distance, all metaballs/particles within the distance of the point specified will be iterated through. The distance parameter is optional and may result in slower execution of the shader.

Inside the loop, you can call the [mdensity](mdensity.html "Returns the density of the metaball field if metaball geometry is
specified to i3dgen.") and [mattrib](mattrib.html "Returns the value of the point attribute for the metaballs if
metaball geometry is specified to i3dgen.") functions to query the contribution of the current point instead of getting a “blended” value.

For example, the following code will take the point color of the metaball which contributes the maximum weight to the point in space:

```vex
float d = 0, max = 0;
vector clr = 0;
vector blended_color;

forpoints ( P ) {
    d = mdensity(P);
    if (d > max) {
        clr = mattrib("Cd", P);
        max = d;
    }
    blended_color = d * clr;
}

```

Note that when you call [mattrib](mattrib.html "Returns the value of the point attribute for the metaballs if
metaball geometry is specified to i3dgen.") inside a `forpoints` loop, the attribute is not pre-blended by the density of the metaball.
