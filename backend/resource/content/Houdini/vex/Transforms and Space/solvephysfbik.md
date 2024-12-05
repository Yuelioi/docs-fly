---
title: solvephysfbik
order: 30
---
| Since | 18.5 |
| --- | --- |

`matrix [] solvephysfbik(matrix xforms[], int parents[], dict jointoptions[], matrix targetxforms[], int targets[], dict targetoptions[], int iters, float damping, float tolerance)`

This solver uses a different algorithm than [solvefbik](solvefbik.html "Applies a full-body inverse kinematics algorithm to a skeleton.") - it typically performs somewhat slower, but provides more control over the skeleton’s behavior and can produce higher-quality results.
The solver can also use the optional per-joint masses and center of mass positions to achieve a target position for the skeleton’s center of mass, allowing for physics-based behavior such as maintaining balance.

Compared to [solvefbik](solvefbik.html "Applies a full-body inverse kinematics algorithm to a skeleton."), this solver has more stable behavior when joint limits are enabled, and produces more accurate results when there are targets at different priority levels.
This solver also tends to distribute joint angle changes more evenly along the chain (particularly when orientation targets are being used), rather than producing large joint angle changes for one or two joints.

The rotation and translation weight parameters provide control over how each joint axis behaves during the solve. This can be used to ensure that certain joints rotate more than others, to lock a specific joint axis, or to enable translating (stretchy) joints.

Returns an empty array if:

- `xforms` or `jointoptions` are not the same size as `parents`.
- `targetxforms` or `targetoptions` are not the same size as `targets`.

Show/hide arguments

`xforms`

The world transforms of all the transforms in the rig being solved.

`parents`

The parent transform index for each transform. A value of -1 indicates a root.

`jointoptions`

Specifies optional parameters for the joints. The valid keys are:

`rotation_weights`

A `vector` specifying weights for the joint’s rotation axes.
Given a larger relative weight, the solution will tend to be achieved by rotating around that axis.
A weight of zero will disable the rotation axis.
The default value is `{1,1,1}`.

`translation_weights`

A `vector` specifying weights for the joint’s translation axes.
Given a larger relative weight, the solution will tend to be achieved by translating along that axis.
A weight of zero will disable the translation axis.
To achieve an unpinned root joint, the root’s translation weight should be non-zero (e.g. `{1,1,1}`).
The default value is `{0,0,0}`.

`rotation_order`

An `int` specifying the joint’s rotation order.

One of the rotation order constants listed below, which can be imported from `$HFS/houdini/vex/include/math.h`.

| Constant name | Rotation Order |
| --- | --- |
| XFORM_XYZ | Rotate order X, Y, Z |
| XFORM_XZY | Rotate order X, Z, Y |
| XFORM_YXZ | Rotate order Y, X, Z |
| XFORM_YZX | Rotate order Y, Z, X |
| XFORM_ZXY | Rotate order Z, X, Y |
| XFORM_ZYX | Rotate order Z, Y, X |

`rotation_lower_limits`

A `vector` specifying the lower rotation limits (in radians) for the joint’s rotation axes.
The rotation limits are applied relative to the joint’s local rest transform, if specified.

Note
To prevent a particular axis from rotating, it is much more efficient to set its weight to zero instead of setting its rotation limits to zero.

`rotation_upper_limits`

A `vector` specifying the upper rotation limits (in radians) for the joint’s rotation axes.

`translation_lower_limits`

A `vector` specifying the lower translation limits for the joint’s translation axes.

`translation_upper_limits`

A `vector` specifying the upper translation limits for the joint’s translation axes.

`rest_xform`

A local space `matrix` specifying the joint’s rest pose.
The default value is the identity matrix.

`rest_rotation_weights`

A `vector` specifying how strongly the solver attempts to match the rest transform for the rotation axes.
This has a priority lower than any of the end effector targets.
A value of `0.1` is typically a suitable value when enabling the rest transform constraint, and a value of 0 will disable it.
The default value is `{0,0,0}`.

`rest_translation_weights`

A `vector` specifying how strongly the solver attempts to match the rest transform for the translation axes.
This has a priority lower than any of the end effector targets.
A value of `0.1` is typically a suitable value when enabling the rest transform constraint, and a value of 0 will disable it.
The default value is `{0,0,0}`.

`mass`

A `float` specifying the mass of the body associated with the joint.
This parameter is only used when a center of mass target is provided.
The default value is `1.0`.

`local_com`

A `vector` specifying the position of the joint’s center of mass, in local space.
An value of `{0,0,0}` (default) will position the center of mass at the joint’s position.

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

`damping`

Damping factor for the solver.
Larger values will produce more stable results when, for example, a target is unreachable.
A value that is too large, however, will require more iterations to converge.
Around 0.5 is typically a suitable initial value.

`tolerance`

The tolerance to use when checking for convergence, defaults to 1e-5.
If positions converge to within this tolerance, the algorithm will stop.
If 0, the solver will always perform exactly `iters` iterations.
