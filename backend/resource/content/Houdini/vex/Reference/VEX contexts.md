---
title: VEX contexts
order: 2
---
Unlike C or C++, VEX has different “contexts” for which you write programs. These contexts define how the function is to be used.

For example, one context is the “surface” context. Functions written in this context are used to calculate the color(s) of a surface during rendering. The global variables and functions provided for each context are slightly different.
Subtopics

## subtopics

Shading contexts

## shading-contexts

See [common shading context features](shading_contexts.html) for information specific to the shading contexts.

- [displace](displace.html "Define a displacement shader with a program that moves a point on a
  surface before the surface is rendered.")
  Define a displacement shader with a program that moves a point on a
  surface before the surface is rendered.
- [fog](fog.html "Deprecated. Define a fog shader with a program that modifies the Cf, Of, or Af
  values to simulate atmospheric effects.")
  Deprecated. Define a fog shader with a program that modifies the Cf, Of, or Af
  values to simulate atmospheric effects.
- [light](light.html "Define a light shader with a program called from surface or fog
  shaders to calculate the illumination of a surface.")
  Define a light shader with a program called from surface or fog
  shaders to calculate the illumination of a surface.
- [shadow](shadow.html "Define a shadow shader by defining a program that’s called from
  surface or fog shaders to calculate the occlusion on a surface from a
  light source.")
  Define a shadow shader by defining a program that’s called from
  surface or fog shaders to calculate the occlusion on a surface from a
  light source.
- [surface](surface.html "Define a surface shader with a program that sets the final color,
  opacity, and alpha of a surface being rendered.")
  Define a surface shader with a program that sets the final color,
  opacity, and alpha of a surface being rendered.

Other contexts

## other-contexts

- [cvex](cvex.html)
- [chop](chop.html "Define a custom CHOP operator with a program that edits channel
  values.")
  Define a custom CHOP operator with a program that edits channel
  values.
- [/vex/contexts/cop.html](cop.html)

Obsolete contexts

## obsolete-contexts

- [image3d](image3d.html "Obsolete. Write a program for use with the i3dgen program to generate 3D
  textures.")
  Obsolete. Write a program for use with the i3dgen program to generate 3D
  textures.
- [sop](sop.html "Obsolete. Define a custom SOP operator with a program that edits geometry point
  attributes.")
  Obsolete. Define a custom SOP operator with a program that edits geometry point
  attributes.
