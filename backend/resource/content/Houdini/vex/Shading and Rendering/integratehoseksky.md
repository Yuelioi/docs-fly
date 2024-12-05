---
title: integratehoseksky
order: 38
---
| Since | 20.0 |
| --- | --- |

`vector  integratehoseksky(float solar_altitude, float solar_azimuth, float turbidity, vector ground_albedo, int number_of_samples)`

Compute the irradiance on a horizontal surface patch from the given Hosek Sky. This is particularly useful for calculating ground colors.

Show/hide arguments

`solar_altitude`

Altitude of the sun in degrees measured from the horizon.

`solar_azimuth`

Azimuth angle of the sun in degrees measured from positive X axis.

`turbidity`

Amount of aerosols in the air. The value range from `1` to `10`. Aerosols scatter the lights and decrease the air’s transparency. Higher values create a more desaturated sky with a blurry sun.

`ground_albedo`

The planet’s reflectance color. This color is used for the Hosek Sky and to calculate the resulting ground color.

`number_of_samples`

Number of samples to integrate the sky. The default is `32`. Use higher values increase accuracy.
