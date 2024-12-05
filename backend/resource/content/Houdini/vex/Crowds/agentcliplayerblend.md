---
title: agentcliplayerblend
order: 9
---
`float  agentcliplayerblend(float values[], float value_weights[], int value_layer_ids[], int layer_blend_modes[], float layer_weights[], int layer_parent_ids[])`

`matrix  agentcliplayerblend(matrix values[], float value_weights[], int value_layer_ids[], int layer_blend_modes[], float layer_weights[], int layer_parent_ids[])`

This function blends the input values according to the blend tree described by the other parameters (the same format used for [layering agent clips](../../crowds/agents.html#currentclips)).
This can be useful for blending custom values that correspond to each of an agent’s assigned clips.

Show/hide arguments

`values`

A list of the input values to blend.

`value_weights`

A list of blend weights for the input values.

`value_layer_ids`

A list containing the layer that each value is an input for.

`layer_blend_modes`

A list of blend modes for each layer. The available blend modes are defined in `$HH/vex/include/crowd_cliplayers.h`.

`layer_weights`

A list of blend weights for each layer. The blend weight is not used for the topmost layer.

`layer_parent_ids`

A list containing the parent layer for each layer (or -1 for the topmost layer). This specifies a tree of animation layers.
