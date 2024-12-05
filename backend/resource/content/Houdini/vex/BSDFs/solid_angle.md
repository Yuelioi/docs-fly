---
title: solid_angle
order: 20
---
`float  solid_angle(bsdf b, int mask)`

Show/hide arguments

`b`

BSDF to sample.

`mask`

A bitmask indicating which types of bounces to evaluate.

See [bouncemask](bouncemask.html) for information on component label bitmasks.

Examples

## examples

```vex
// Split BSDF into component lobes
bsdf lobes[] = split_bsdf(hitF);

// Get solid angle of lobes
float angles[];
resize(angles, len(lobes));
for (int i = 0; i < len(lobes); i++)
{
    angles[i] = solid_angle(lobes[i], PBR_ALL_MASK);
}

// Compute PDF from angles
float pdf[] = compute_pdf(angles);

// Compute CDF from PDF
float cdf[] = compute_cdf(pdf);

// Randomly select a BSDF based on albedo distribution
int id = sample_cdf(cdf, sx);

// Do something with the selected BSDF
// lobes[id] ...

```
