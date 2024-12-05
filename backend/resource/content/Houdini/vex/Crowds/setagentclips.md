---
title: setagentclips
order: 53
---
`int  setagentclips(int geohandle, int prim, string clip_names[], float clip_times[], float clip_weights[], string clip_transform_groups[], int clip_layer_ids[], int layer_blend_modes[], float layer_weights[], int layer_parent_ids[])`

`int  setagentclips(int geohandle, int prim, int clip_ids[], float clip_times[], float clip_weights[], int clip_transform_group_ids[], int clip_layer_ids[], int layer_blend_modes[], float layer_weights[], int layer_parent_ids[])`

This function can provide improved performance compared to using a combination of [setagentclipnames](setagentclipnames.html "Sets the current animation clips for an agent primitive."), [setagentcliptimes](setagentcliptimes.html "Sets the current times for an agent primitive’s animation clips."), and [setagentclipweights](setagentclipweights.html "Sets the blend weights for an agent primitive’s animation clips."), and also modifies the [primitive intrinsics used for layering animation clips](../../crowds/agents.html#currentclips).

Show/hide arguments

`geohandle`

Handle to the geometry to write to. `geoself()` can be used to get a handle to the current geometry.

`prim`

The primitive number.

`clip_names`

A list of animation clip names.

`clip_ids`

A list of animation clip indices in the agent definition.
A clip’s index can be obtained via [agentfindclip](agentfindclip.html "Finds the index of a clip in an agent’s definition.").

`clip_times`

A list of times that the clips should be sampled at.

`clip_weights`

A list of blend weights for the animation clips.

`clip_transform_groups`

A list of transform groups, which specify the joints that each clip should be applied to.

`clip_transform_group_ids`

A list of transform group indices in the agent definition.
A transform group’s index can be obtained via [agentfindtransformgroup](agentfindtransformgroup.html "Finds the index of a transform group in an agent’s definition.").

`clip_layer_ids`

A list containing the layer that each animation clip is an input for.

`layer_blend_modes`

A list of blend modes for each layer. The available blend modes are defined in `$HH/vex/include/crowd_cliplayers.h`.

`layer_weights`

A list of blend weights for each layer. The blend weight is not used for the topmost layer.

`layer_parent_ids`

A list containing the parent layer for each layer (or -1 for the topmost layer). This specifies a tree of animation layers.
