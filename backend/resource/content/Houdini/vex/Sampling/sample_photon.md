---
title: sample_photon
order: 27
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  sample_photon(light lp, vector &pos, vector &dir, float &scale, float time)`

Spawns a photon from the given light source and returns the information for the
first intersection in the scene. The `pos`, `dir` and `scale` are filled out
with the information about where the photon hit in the scene.

The returned integer represents the bounce type mask (this is determined by the
types of illumination labels on the light source).

The function returns 0 if the photon didn’t intersect any geometry.
