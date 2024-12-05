---
title: rayimport
order: 60
---
| Context(s) | [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  rayimport(string name, <type>&value)`

`int  rayimport(string name, <type>&value[])`

Extracts information any passed when the surface is hit by a ray fired by the [gather](gather.html "Sends rays into the scene and returns information from the shaders of
surfaces hit by the rays.").

Show/hide arguments

`name`

The variable name, as passed using a `"send:name", value` argument pair in [gather](gather.html "Sends rays into the scene and returns information from the shaders of
surfaces hit by the rays.") (without the `send:` prefix).

`value`

If the function can import the named variable, it overwrites this variable with the value.

Returns

`1` if a value by the given name was successfully imported, or `0` otherwise.

Built-in queryable names

## v3 built-in-queryable-names

You can pass the following values to `name` to query built-in ray information (not sent from `gather()`).

`ray:P` (`vector`)

The origin of the ray.

`ray:D` (`vector`)

The direction vector of the ray.

`ray:time` (`float`)

The shutter time associated with the ray.

`ray:hitstack` (`int[]`)

The hit-stack provided by the intersector.

`ray:element` (`int`)

The element provided by the intersector.

`ray:hituv` (`vector`)

The parametric coordinates provided by the intersector.

`ray:Ng` (`vector`)

The geometric normal from the intersector.

Note
Data provided by the Mantra 3 intersector is raw data and may not be meaningful, or may be different across platforms or versions.
