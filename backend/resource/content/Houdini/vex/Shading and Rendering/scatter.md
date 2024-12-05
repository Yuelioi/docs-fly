---
title: scatter
order: 67
---
| Context(s) | [displace](../contexts/displace.html)  [fog](../contexts/fog.html)  [light](../contexts/light.html)  [shadow](../contexts/shadow.html)  [surface](../contexts/surface.html) |
| --- | --- |

`int  scatter(vector ipoint, vector inormal, vector idirection, vector idistribution, float time, float maxdist, vector &opoint, vector &onormal, vector &odirection)`

Evaluates a single scattering event through geometry. A return value of 1 means the scattering was successful.

Show/hide arguments

`ipoint`

Entry point of scattering.

`inormal`

Surface normal at the entry point (surfaces only). Required to orient the scattering plane.

`idirection`

Primary scattering direction at the entry point. Required to orient the scattering plane.

`idistribution`

Initial scatter distribution at the entry point. If a zero-valued vector is passed a random scatter distribution will be used.

`maxdist`

Maximum distance to scatter.

`opoint`

Exit point of scattering.

`onormal`

Normal at the scattering exit point (surfaces only).

`odirection`

Outgoing direction at the scattering exit point (surfaces only).

```vex
// Trace for intersection with scene
vector hitP = 0;
vector hitN = 0;
int hit = trace(P, I, Time, "P", hitP, "N", hitN);

// Scatter a random distance from the intersection
vector idistribution = 0;
int sid = israytrace ? SID : newsampler();
vector s;
nextsample(sid, s.x, s.y, "mode", "nextpixel");
float maxdist = 2.0 * s.x;
vector opoint = 0;
vector onormal = 0;
vector odirection = 0;
hit = scatter(hitP, hitN, I, idistribution, Time, maxdist, opoint, onormal, odirection);

// Trace again from the exit point of the scattering
hit = trace(opoint, odirection, Time, "P", hitP, "N", hitN);

```
