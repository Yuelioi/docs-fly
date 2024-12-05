---
title: ocean_sample
order: 7
---
`vector  ocean_sample(string geometry, int phase, int frequency, int amplitude, float hscale, float time, int mode, int downsample, vector pos)`

Evaluates the input ocean spectrum at the given time and position and returns the value specified by `mode`. The input is typically the output of the [![](../../icons/SHELF/wave.svg)Ocean Spectrum](../../nodes/sop/oceanspectrum.html "Generates volumes containing information for simulating ocean waves.") SOP.

Show/hide arguments

`geometry`

The name of the geometry file to reference. Inside Houdini, this may be `op:full_path_to_sop` to reference a SOP.

`phase`

The primitive number of the volume representing wave phase.

`frequency`

The primitive number of the volume representing wave frequency.

`amplitude`

The primitive number of the volume representing wave amplitude.

`hscale`

The amount to scale any horizontal motion in the ocean waves.

`mode`

The type of value to sample from the ocean spectrum, where 0 means displacement, 1 means velocity, and 2 means horizontal spatial derivatives.

`downsample`

The number of times to downsample the input spectrum before evaluation. Each downsample level halves the resolution of the input.

`pos`

The position in object space at which to sample the evaluated ocean.

Examples

## examples

Displace the point position by an ocean spectrum stored in a file.

```vex
@P += ocean_sample("spectrum.bgeo", 0, 1, 2, 0.7, @Time, 0, 0, @P);

```
