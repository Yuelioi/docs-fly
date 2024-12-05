---
title: agentaddclip
order: 1
---
Warning
This function has been deprecated. Use [hou.AgentDefinition.addClip](../../hom/hou/AgentDefinition.html#addClip) and [hou.crowds.replaceAgentDefinitions](../../hom/hou/crowds.html#replaceAgentDefinitions) instead to edit agent definitions.

This function adds a `.clip` or `.bclip` file saved from a CHOP (or generated
by an [![](../../icons/CROWDS/agent.svg)Agent](../../nodes/out/agent.html "This output operator is used to write agent definition files.") ROP) to the definition of the given agent primitive.
Clips in an agent definition contains transform animation for driving the
agent’s skeleton.

The channels within the clip should of the form `transform_name:channel_name`
where \_transform_name\_ is a string matching the values returned by
[agenttransformnames](agenttransformnames.html "Returns the name of each transform in an agent primitive’s rig.") and \_channel_name\_ is one of `tx`, `ty`, `tz`, `rx`,
`ry`, `rz`, `sx`, `sy`, or `sz`. The channels starting with `t` denote
translation, `r` denote rotation, and `s` denote scale. The resulting
transforms will be treated as *local* transforms, such as those returned by
[agentlocaltransform](agentlocaltransform.html "Returns the current local space transform of an agent primitive’s bone.") (ie. they are relative to the corresponding parent
transform in the agent’s skeleton).

Show/hide arguments

`geohandle`

A handle to the geometry to write to. Currently the only valid value is `0` or [geoself](geoself.html "Returns a handle to the current geometry."), which means the current geometry in a node. (This argument may be used in the future to allow writing to other geometries.)

`prim`

The primitive number of the agent primitive whose definition is to be
modified.

`clipname`

The name to identify the clip. All clips in an agent definition must have
unique names.

`clippath`

The filename of the `.clip` or `.bclip` file saved from a CHOP or generated
by the [![](../../icons/CROWDS/agent.svg)Agent](../../nodes/out/agent.html "This output operator is used to write agent definition files.") ROP. Use `op:full_path_to_chop` to directly refer
to a CHOP in the scene.

`keepref`

When `clippath` refers to a filename on disk, this boolean flag indicates
whether the external reference should be maintained when the geometry will
be saved. If the reference is maintained, then the original source of the
clip needs to be available when the saved geometry is used. Otherwise, a
copy of the clip will be inlined when saving out the geometry so that the
original clip is no longer needed.
