---
title: Writing PBR shaders in VEX
order: 10
---
| On this page | * [Overview](#overview) * [Primitive BSDFs](#primitive-bsdfs) * [Combining BSDFs](#combining-bsdfs) |
| --- | --- |
To build materials using VOPs instead, see [building materials in VOPs](../shade/build.html "The basics of how to create, combine, and modify materials in Houdini.").
Overview

## overview

Writing a PBR surface shader involves designing a *bidirectional scattering distribution function* (BSDF) for the surface. This function describes how light scatters when it hits the surface. BSDFs are an opaque data type in VEX that you can store and operate on much like other primitive types.

You output the final BSDF for the shader through a new global variable `F`. In VOPs, you connect a BSDF to the `F` connection of the output node (see VOPs below).

The BSDF is a self-contained description of how a surface reflects light. With PBR, you don’t need to compute illumination (for example, using
[illuminance](functions/illuminance.html "Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.")loops or `trace()` calls). You only need to provide a `bsdf`.

You don’t have to write the scattering function yourself. You can create BSDFs by combining primitive BSDFs using addition, multiplication, and scalar multiplication. See “primitive BSDFs” and “combining BSDFs” below. If you want, you can write your own BSDF functions using [CVEX](functions/cvex_bsdf.html "Creates a bsdf object from two CVEX shader strings.").

Depending on the settings selected for the render, mantra can use the BSDFs calculated by your shader in many different ways. For example, it may perform direct lighting, indirect lighting, or evaluate lighting in reverse for algorithms like photon mapping.

You can design shaders that support both traditional and physically based rendering. You just need to assign values to both types of output variables (`Cf`, `Of`, and `Af` for traditional rendering, `F` for PBR rendering) in your shader.

Primitive BSDFs

## primitive-bsdfs

- [phong](functions/phong.html "Returns a Phong BSDF or computes Phong shading.")
- [phonglobe](functions/phonglobe.html)
- [ashikhmin](functions/ashikhmin.html "Returns a specular BSDF using the Ashikhmin shading model.")
- [blinn](functions/blinn.html "Returns a Blinn BSDF or computes Blinn shading.")
- [matchvex_blinn](functions/matchvex_blinn.html "Returns a BSDF that matches the output of the traditional VEX blinn function.")
- [specular](functions/specular.html "Returns a specular BSDF or computes specular shading.")
- [matchvex_specular](functions/matchvex_specular.html "Returns a BSDF that matches the output of the traditional VEX specular function.")
- [cone](functions/cone.html "Returns a cone reflection BSDF.")
- [diffuse](functions/diffuse.html "Returns a diffuse BSDF or computes diffuse shading.")
- [wireblinn](functions/wireblinn.html)
- [wirediffuse](functions/wirediffuse.html)

Combining BSDFs

## combining-bsdfs

- `bsdf +(bsdf, bsdf)`
  Sum together two bsdfs. Summing bsdfs allows you to include more than one type of component, such as summing together a diffuse() and a phong().
- `bsdf *(float scale, bsdf)`
- `bsdf *(vector scale, bsdf)`
  Scale a bsdf by a color or a scale factor. Scaling a bsdf can allow you to incorporate surface colors (such as texturing) into the bsdf.
- `bsdf *(bsdf, bsdf)`
  Create the produce of two bsdfs. This function is most useful for scaling a specular bsdf by a diffuse().
- `vector albedo(bsdf)`
  Return the portion of reflected light for a `bsdf`. The return value should be between 0 and 1 for bsdfs that conserve energy. The albedo for a default diffuse `bsdf` is 0.5, but will change based on vector scale factors and other `bsdf` components.
