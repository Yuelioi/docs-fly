---
title: solvefbik
order: 28
---
| Since | 17.0 |
| --- | --- |

`matrix [] solvefbik(matrix xforms[], int parents[], dict jointoptions[], matrix targetxforms[], int targets[], dict targetoptions[], int iters, float tolerance, int pinroot)`

`matrix [] solvefbik(matrix xforms[], int parents[], int targets[], matrix targetxforms[], int iters)`

`matrix [] solvefbik(matrix xforms[], int parents[], int targets[], matrix targetxforms[], int iters, float tolerance, int pinroot)`

`matrix [] solvefbik(matrix xforms[], int parents[], int targets[], matrix targetxforms[], int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[])`

`matrix [] solvefbik(matrix xforms[], int parents[], int targets[], matrix targetxforms[], int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[], int targettypes[], matrix targetoffsets[])`

`matrix [] solvefbik(matrix xforms[], int parents[], int targets[], matrix targetxforms[], int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[], matrix goalxforms[], vector4 constrainedxforms[], vector jointlimits[])`

`matrix [] solvefbik(matrix xforms[], int parents[], int targets[], matrix targetxforms[], int iters, float tolerance, int pinroot, float targetweights[], int targetpriorities[], int targetdepths[], int targettypes[], matrix targetoffsets[], matrix goalxforms[], vector4 constrainedxforms[], vector jointlimits[])`

Returns an empty array if:

- `xforms` is not the same size as `parents`
- `targets` is not the same size as `targetxforms`
- The `goalxforms`, `constrainedxforms`, and `jointlimits` arrays aren’t empty, but are not the same size as `xforms`

The `goalxforms`, `constrainedxforms`, and `jointlimits` parameters should be in the form produced by the [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints SOP](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints.").

Show/hide arguments

`xforms`

The world transforms of all the transforms in the rig being solved.

`parents`

The parent transform index for each transform. A value of -1 indicates a root.

`jointoptions`

Specifies optional parameters for the joints. The valid keys are:

`limit_goalxform`

A `matrix` specifying the position and orientation of the cone in the space of the parent transform.
This can be set from the attributes produced by [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints.").

`limit_constrainedxform`

A `vector4` (quaternion) specifying the orientation of the twist axis, up axis, and out axis in the space of the child transform.
This can be set from the attributes produced by [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints.").

`limit_angles`

A `vector` specifying the maximum rotation around each axis, in degrees.
This can be set from the attributes produced by [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints.").

`targets`

A list of the transform indices of the end effectors in the skeleton.

`targetxforms`

A list of the target world transforms for the end effectors, in the same order as `targets`.

`targetoptions`

Specifies optional parameters for the targets. The valid keys are:

`weight`

A `float` specifying the importance of the target.
When multiple targets have the same priority level, targets with a higher relative weight are more likely to be reached.
The default value is `1.0`.

`priority`

An `int` specifying the target’s priority level.
Targets from a lower priority level cannot interfere with targets from a higher priority level.
For example, priority levels can be used to ensure that the feet remain planted when manipulating the upper body of a skeleton.
The default value is `0`.

`depth`

An `int` specifying the number of parent joints that can be adjusted to achieve the goal transform.
A negative depth indicates that the entire chain can be affected.
The default value is `-1`.

`target_type`

An `int` specifying how the end effector matches the position or orientation of its target transform.
A value of `0` (default) indicates a position-only target, `1` indicates an orientation-only target, and `2` matches both position and orientation.
A value of `3` indicates a target that controls the center of mass of the skeleton (the transform index from `targets` is not used).
Only one center of mass target can be provided.

`joint_offset`

A `matrix` specifying a local space transform that is combined with the joint transform to produce the transform that the solver attempts to align with the goal transform.
This can be used to place the target at an offset from the joint (for example, at the end of a bone).

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

`goalxforms`

Part of the joint constraints as produced by [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints.").
An empty array indicates no joint constraints.

`constrainedxforms`

Part of the joint constraints as produced by [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints.").
An empty array indicates no joint constraints.

`jointlimits`

Part of the joint constraints as produced by [![](../../icons/SOP/agentconfigurejoints.svg)Agent Configure Joints](../../nodes/sop/agentconfigurejoints.html "Creates point attributes that specify the rotation limits of an agent’s joints.").
An empty array indicates no joint constraints.
