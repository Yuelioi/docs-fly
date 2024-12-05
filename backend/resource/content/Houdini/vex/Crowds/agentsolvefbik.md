---
title: agentsolvefbik
order: 39
---
| Since | 17.0 |
| --- | --- |

`int  agentsolvefbik(<geometry>geometry, int outgeo, int prim, int targets[], matrix targetxforms[], int xformgroup, int iters)`

`int  agentsolvefbik(<geometry>geometry, int outgeo, int prim, int targets[], matrix targetxforms[], int xformgroup, int iters, float tolerance, int pinroot)`

`int  agentsolvefbik(<geometry>geometry, int outgeo, int prim, int targets[], matrix targetxforms[], int xformgroup, int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[])`

`int  agentsolvefbik(<geometry>geometry, int outgeo, int prim, int targets[], matrix targetxforms[], int xformgroup, int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[], int targettypes[], matrix targetoffsets[])`

`int  agentsolvefbik(<geometry>geometry, int outgeo, int prim, int targets[], matrix targetxforms[], int xformgroup, int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[], string goalxformattrib, string constrainedxformattrib, string jointlimitsattrib)`

`int  agentsolvefbik(<geometry>geometry, int outgeo, int prim, int targets[], matrix targetxforms[], int xformgroup, int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[], int targettypes[], matrix targetoffsets[], string goalxformattrib, string constrainedxformattrib, string jointlimitsattrib)`

Returns `-1` if `prim` is out of range or is not an agent primitive, or targets and targetxforms are not the same length.

If “agent_jointgoalxforms”, “agent_jointconstrainedxforms”, and “agent_jointlimits” attributes are present on the agent, as produced by the [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints SOP](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints."), then they will be interpreted as joint limits to use while solving.

Show/hide arguments

`<geometry>`

When running in the context of a node (such as a wrangle SOP), this argument can be an integer representing the input number (starting at 0) to read the geometry from.

Alternatively, the argument can be a string specifying a geometry file (for example, a `.bgeo`) to read from. When running inside Houdini, this can be an `op:/path/to/sop` reference.

`outgeo`

Handle to the geometry to write to. `geoself()` can be used to get a handle to the current geometry.

`prim`

The primitive number of the agent primitive.

`targets`

A list of the transform indexes of the end effectors in the agent.

`targetxforms`

A list of the target world transforms for the end effectors, in the same order as `targets`.

`xformgroup`

The index of a transform group, which specifies the joints to be used for the IK solver (all transforms not in the transform group will be ignored).
[agentfindtransformgroup](agentfindtransformgroup.html "Finds the index of a transform group in an agent’s definition.") can be used to look up a transform group by name, and a value of -1 indicates that all the transforms in the agent should be included.
It’s recommended to use a transform group that includes only transforms that correspond to the agent’s bone structure.

`iters`

The maximum number of iterations to perform.
The solver may terminate early if the `tolerance` parameter is used.

`tolerance`

The tolerance to use when checking for convergence, defaults to 1e-5.
If positions converge to within this tolerance, the algorithm will stop.
If 0, the solver will always perform exactly `iters` iterations.

`pinroot`

Whether to pin the root to its start position, instead of allowing it to translate.
This can be useful when, for example, solving a subset of an agent’s skeleton.
Defaults to 0 (off).

`targetweights`

A list containing the weight of each end effector, in the same order as `targets`.
For joints with multiple children, the normalized weights will be used to determine their position - this means that a target with a higher weight than other targets will be more likely to be reached.
The default weight is 1.0.

`targetpriorities`

A list containing the priority level of each end effector, in the same order as `targets`.
Targets from a lower priority level will not influence targets with higher priority.
For example, priority levels can be used to ensure that the targets for the feet are always satisfied, while still controlling the relative weights of the upper body targets.
The default priority is 0 (i.e. all targets are equal priority).

`targetdepths`

For each end effector, specifies how many joints above it in the chain can be adjusted to achieve the target transform.
A negative depth can be used to specify that all joints above the target are affected.
The default depth is -1.

`targettypes`

A list containing the target type for each end effector, in the same order as `targets`.
The target type can be used to specify how the end effector matches the position or orientation of its target transform (from `targetxforms`).
A value of `0` indicates a position-only target, `1` indicates an orientation-only target, and `2` matches both position and orientation (default).

`targetoffsets`

A list containing an additional local space transform for each end effector, in the same order as `targets`.
This transform is combined with the end effector’s joint transform to produce the transform that the solver attempts to align with the target transform.
This can be used to place the target at an offset from the joint (for example, at the end of a bone).

`goalxformattrib`

An optional parameter specifying an alternative attribute to use instead of “agent_jointgoalxforms”.

`constraintedxformattrib`

An optional parameter specifying an alternative attribute to use instead of “agent_jointconstrainedxforms”.

`jointlimitsattrib`

An optional parameter specifying an alternative attribute to use instead of “agent_jointlimits”.
