---
title: create_cdf
order: 1
---
| On this page | * [Overview](#overview) * [Usage](#usage) * [Examples](#examples) |
| --- | --- |

Overview

## overview

CDFs are useful when sampling from distributions. For example, you could create a CDF of light source power. This would allow sampling of lights with a probability based on power. This is an example of a discrete CDF, where sampling selects among a fixed set of probabilities. (See the example below.)

Use the [sample_cdf](sample_cdf.html "Samples a cumulative distribution function (CDF).") function to sample values from the returned CDF array.

Usage

## usage

`float [] create_cdf(float pdf[])`

Returns a CDF for the input PDF as an array of floats.

Show/hide arguments

`pdf`

Array of PDF values to create the CDF for.

Examples

## examples

```vex
// Iterate over all lights, sampling their power
int[] li = getlights();
float values[];
resize(values, len(li));
int nsamples = 256;
int sid = israytrace ? SID : newsampler();
vector s, pos, clr;
float scale;
for (int i = 0; i < len(li); i++)
{
    for (int j = 0; j < nsamples; j++)
    {
        nextsample(sid, s.x, s.y, "mode", "nextpixel");
        sample_light(li[i], P, s, Time, pos, clr, scale);
        values[i] += luminance(clr);
    }
    values[i] /= nsamples;
}

// Create a CDF of the power distribution
float cdf[] = create_cdf(values);

// Randomly select a light based on power distribution
nextsample(sid, s.x, s.y, "mode", "nextpixel");
int index = 0;
sample_cdf(cdf, s.x, index);

// Do something with the selected light
// li[index] ...

```
