---
title: hair
order: 12
---
`bsdf  hair(vector N, vector tip, float lobe_shift, float lobe_width_lon, ...)`

`bsdf  hair(vector N, vector tip, float lobe_shift, float lobe_width_lon, float lobe_with_azi, ...)`

`bsdf  hair(vector N, vector tip, float lobe_shift, float lobe_width_lon, float lobe_with_azi, float glint_shift, float glint_intensity, ...)`

Details of the hair BSDF can be found in the source file (`hair_eval.vfl`).

Any variadic arguments to the functions are passed through to the CVEX evaluation function.

Examples

## examples

These different signatures are equivalent to the following code:

```vex
bsdf hair(vector N; vector tip; float lobe_shift; float lobe_width_lon, ...)
{
    cvex_bsdf("hair_eval", "hair_sample",
        "label", "diffuse",
        "tip", tip,
        "lobe_shift", lobe_shift,
        "lobe_width_lon", lobe_width_lon,
    ...);
}

bsdf hair(vector N; vector tip; float lobe_shift; float lobe_width_lon, float lobe_with_azi, ...)
{
    cvex_bsdf("hair_eval", "hair_sample",
        "label", "refract",
        "tip", tip,
        "lobe_shift", lobe_shift,
        "lobe_width_lon", lobe_width_lon,
        "lobe_width_azi", lobe_width_azi,
    ...);
}

bsdf hair(vector N; vector tip; float lobe_shift; float lobe_width_lon, float glint_shift; float glint_intensity, ...)
{
    cvex_bsdf("hair_eval", "hair_sample",
        "label", "reflect",
        "tip", tip,
        "lobe_shift", lobe_shift,
        "lobe_width_lon", lobe_width_lon,
        "glint_shift", glint_shift,
        "glint_intensity", glint_intensity,
    ...);
}

```
