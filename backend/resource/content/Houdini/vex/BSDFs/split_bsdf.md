---
title: split_bsdf
order: 22
---
`void  split_bsdf(bsdf &lobes[], bsdf source, float &weights[])`

`void  split_bsdf(bsdf &lobes[], bsdf source, float &weights[], int mask)`

`void  split_bsdf(bsdf &lobes[], bsdf source, float &weights[], int mask, int type)`

`void  split_bsdf(bsdf &lobes[], bsdf source, float &weights[], int mask, int type, float u)`

`void  split_bsdf(bsdf &lobes[], bsdf source, float &weights[], int mask, int type, float u, float cdf[])`

Show/hide arguments

`&lobes`

The function overwrites this array with the BSDFs for the component lobes.

`source`

The BSDF to split.

`weights`

The function fills this array with the weights for the split lobes (same length as the returned `bsdf` array). When you sample Illumination using the returned lobes you must scale it by these weights.

`mask`

A bitmask indicating which types of bounces to evaluate.

See [bouncemask](bouncemask.html) for information on component label bitmasks.

`type`

How to split the lobes. You can `#import "pbr.h"` to get constant values representing the different split types:

- `PBR_SPLIT_FULL = 0`
- `PBR_SPLIT_RANDOM = 1`
- `PBR_SPLIT_ALBEDO = 2`
- `PBR_SPLIT_COMPONENT = 3`
- `PBR_SPLIT_DEFAULT = PBR_SPLIT_ALBEDO`

`u`

Random value to sample the CDF at.

`cdf`

CDF used to control sampling among components of the BSDF.

Returns

An array of `bsdf` objects representing the lobes.

Examples

## examples

```vex
// Split BSDF into component lobes
float weights[];
bsdf lobes[];
split_bsdf(lobes, hitF, weights);

// Get albedos of lobes
float albedos[];
resize(albedos, len(lobes));
for (int i = 0; i < len(lobes); i++)
{
    albedos[i] = luminance(albedo(lobes[i], -hitnI)) * weights[i];
}

// Compute CDF
float cdf[] = compute_cdf(albedos);

// Randomly select a BSDF based on albedo distribution
int index = 0;
sample_cdf(cdf, s.x, index);

// Do something with the selected BSDF
// lobes[index] ...

```
