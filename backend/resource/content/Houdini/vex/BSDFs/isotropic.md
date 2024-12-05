---
title: isotropic
order: 14
---
`bsdf  isotropic(...)`

The isotropic function scatters light equally in all directions and is suitable for use in rendering dense volumetric materials such as smoke. Note that no normal vector is required to construct an isotropic bsdf since it has no directionality. The default albedo for an isotropic `bsdf` is 1, which means the isotropic() function scatters 100% of the incoming light.
