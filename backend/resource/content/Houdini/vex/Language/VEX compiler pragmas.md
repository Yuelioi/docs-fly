---
title: VEX compiler pragmas
order: 6
---
| On this page | * [#pragma c++rawstring](#pragma-c-rawstring) * [#pragma bindhandle](#pragma-bindhandle) * [#pragma bindhandlereserved](#pragma-bindhandlereserved) * [#pragma bindselector](#pragma-bindselector) * [#pragma bindselectorreserved](#pragma-bindselectorreserved) * [#pragma callback](#callback) * [#pragma disablewhen and #pragma hidewhen](#pragma-disablewhen-and-pragma-hidewhen) * [#pragma export](#pragma-export) * [#pragma group](#pragma-group) * [#pragma help and #pragma info](#pragma-help-and-pragma-info) * [#pragma hint](#pragma-hint) * [#pragma inputlabel](#pragma-inputlabel) * [#pragma label](#pragma-label) * [#pragma name “text”](#pragma-name-text) * [#pragma opicon](#opicon) * [#pragma opmininputs and #pragma opmaxinputs](#opinputs) * [#pragma opname and #pragma oplabel](#opname) * [#pragma opscript](#opscript) * [#pragma parmhelp parameter_name “text”](#pragma-parmhelp-parameter_name-text) * [#pragma parmtag](#pragma-parmtag) * [#pragma ramp](#ramp) * [#pragma range](#pragma-range) * [#pragma rendermask](#pragma-rendermask) * [#pragma optable](#pragma-optable) * [#pragma rename](#pragma-rename) * [How to create menus for parameters](#how-to-create-menus-for-parameters) * [Operator list filters](#operator-list-filters) |
| --- | --- |
The [VEX compiler (vcc)](vcc.html "Overview of how to use the VEX language compiler vcc and its
pre-processor and pragma statements.") supports pragmas for automatically building UI
dialog scripts. These pragmas are typically ignored unless the -u option
is specified on the vcc command line. The pragmas let you specify
help, hints for how to represent the parameter, organization, and other
information.

Pragmas can be specified in one of two ways: As a `#pragma` preprocessor
directive, or as a `_Pragma` VEX statement. The second form allows for
bundling up multiple pragmas in a single macro.

The following preprocessor form:

```vex
#pragma label parm "Parameter Label"

```

is functionally equivalent to this VEX statement:

```vex
_Pragma("label parm \"Parameter Label\"");

```

Note that the pragma arguments are enclosed in quotes, with the original set
of quotes escaped.

# pragma c++rawstring

## pragma-c-rawstring

`#pragma c++rawstring 0|1`

VEX supports C++ style raw strings. The `c++rawstring` pragma can be used to
disable or re-enable raw string support. For example:

```vex
string a = R"(Hello world\n)";
string b = "Hello world\\n";
#pragma c++rawstrings 0        // Disable C++ style raw string support
string b = R"(This will generate an error!)";
#pragma c++rawstrings 1        // Re-enable C++ style raw string support
```

# pragma bindhandle

## pragma-bindhandle

`#pragma bindhandle channel_name h_name h_label h_index h_settings`

Binds handles to specific parameters by default (users can override
bindings).

`channel_name`

The name of the channel in the VEX operator to bind to the handle.

`h_name`

Handle name. This is one of the pre-defined Houdini handles (for
example `ladder`). You can use the [omls](../commands/omls.html "Lists the available handles for an operator type.") HScript command
for a full list of available handles.

`h_label`

A brief description of the handle.

`h_index`

Many handles (for example xform) have multiple parameters associated
with them. This allows you to choose which handle parameter gets
bound to the VEX parameter.

`h_settings`

An optional handle-specific string that can be used to set some
default behavior for the handle.

```vex
#pragma bindhandle offset1 xform "Translate" tx "invisible(1)"
#pragma bindhandle offset2 xform "Translate" ty
#pragma bindhandle offset3 xform "Translate" tz
sop translate(vector offset=0) { P += offset; }

```

# pragma bindhandlereserved

## pragma-bindhandlereserved

`#pragma bindhandlereserved reserved_channel_name h_name h_label h_index h_settings`

Each scripted operator type has a number of parameters that are added to
every operator of that type (regardless of the contents of the dialog
script file). To bind a handle to one of these parameters, you must use
the bindhandlereserved pragma. This pragma takes exactly the same
arguments as the bindhandle pragma. The only exception is that the
channel name argument must specify the name of a reserved parameter.

# pragma bindselector

## pragma-bindselector

`#pragma bindselector [parm_name] sel_type sel_name sel_prompt sel_mask allow_dragging group_type_parm asterisk_sel_all [input_index input_required]`

When an operator is created interactively in Houdini, the user can be
prompted for the data to work on. These prompts are handled by
selectors. Selectors can be defined on a per-OP basis, or a
per-parameter basis.

For per-OP selectors, the bindselector pragma expects 7 arguments. For
per-parameter selectors, it expects 10 arguments.

parm_name

The VEX parameter to bind the selector to (for per-parameter
selectors).

sel_type

The entity to select. Use the [omsls](../commands/omsls.html "Lists the available selectors for an operator type.") HScript command to
print a list of possible values.

sel_name

A brief description of the selector.

sel_prompt

the prompt presented to the user to select geometry.

sel_mask

a pattern which allows selection of specific primitive types. The
list of possible primitive types are:

all

all primitive types

face

Polygons, NURBs or Bezier curves.

surface

Mesh, NURBs or Bezier surfaces

quadric

Primitive circles, spheres or tubes.

poly

Polygons

nurbscurve

NURBS curves

bezcurve

Bezier curves

mesh

Meshes

nurbs

NURBS surfaces

bezier

Bezier surfaces

circle

Primitive circles

sphere

Primitive spheres

tube

Primitive tubes

meta

Metaballs

particle

Particle systems

The primitive types can be combined in standard Houdini grouping
mechanisms. For example:

- `all,^p*`: select all primitive types except polygons and
  particles.
- `face,surface`: select face and surface primitives.
- `*,^quad*,^meta`: Select any primitive but quadrics or metaballs.

allow_dragging

If set to 1, the selection can be modified without forcing the user
to click `RMB` to complete the selection.

This lets the user select and modify in one step (dragging with the
mouse finishes selection and passes the mouse movements to the
operator handles).

group_type_parm

The name of a parameter which indicates the geometry type the
selection group will have. Typically this value will name a
parameter with a menu for choosing “Points”, “Primitives”, or “Guess
from group”. See the OMbindings file for the [![](../icons/SOP/blast.svg)Blast
SOP](../nodes/sop/blast.html "Deletes primitives, points, edges or breakpoints.").

asterisk_sel_all

If set to 1, the selector needs to set the selection string to “\*”
to indicate all geometry was selected. If 0, the selector assumes an
empty group parameter means all geometry was selected.

input_index

For per-parameter selectors. When the user selects geometry, the
selector must connect the output from the selected operator to the
input of this operator. This parameter specifies the index of the
input number where the operator should be connected. Use -1 if the
selector needs to connect multiple input operators into this
operator.

input_required

For per-parameter selectors. Set to 1 if the user must select
geometry for this input.

```vex
#pragma bindselector prims "Switch Geometry" \
    "Choose the geometry to switch between" \
    all 0 "" 0
#pragma bindhandle input_number 0 ladder Input parm0
sop switcher(int input_number=0) { import("P", P, input_number) }

```

# pragma bindselectorreserved

## pragma-bindselectorreserved

`#pragma bindselectorreserved reserved_parm_name sel_type sel_name sel_prompt sel_mask allow_dragging group_type_parm asterisk_sel_all input_index input_required`

Similar to the bindhandlereserved pragma, this binds selectors to
reserved parameters in your scripted operators. The arguments to this
pragma are the same as those passed to the bindselector pragma. The only
difference is that the parameter name argument must specify a reserved
parameter.

# pragma callback

## callback

`#pragma callback name "script"`

Binds a callback HScript script or Python function to the name
parameter. When the parameter changes, Houdini executes the script
string.

Because of architectural limitations in Houdini, the parameter and
script must meet certain conditions:

- The dialog script needs to be bound to a Houdini node (e.g. SHOP,
  SOP, etc.).
- The parameter must be either a toggle button, or have a menu bound to
  it (see `#pragma hint` and `#pragma choice`).

To indicate the language of the callback (`hscript` or `python`), use `#pragma parmtag`. If you don’t use a `#pragma parmtag` for a callback the default is `hscript`. However, the recommended method for programming callback scripts is `python`.

```vex
#pragma callback parm1 "message $script_parm"
#pragma parmtag parm1 script_callback_language hscript

#pragma callback parm2 "import hou; hou.ui.displayMessage(kwargs)"
#pragma parmtag initialize_menu script_callback_language hscript
```

- For HScript callbacks, script can be an HScript statement or the
  name of a script on the `$HOUDINI_PATH/scripts` path.
- For Python callbacks, script should be Python source code.

Houdini executes script in a context with certain variables available
indicating which parameter changed.

- For HScript callbacks, Houdini creates variables such as `$node`
  containing the path to the node on which the parameter changed.
- For Python callbacks, Houdini creates a `kwargs` variable containing a
  dictionary of information about the parameter that changed.

See [callback scripts](../hom/locations.html#parameter_callback_scripts) for
more information.

# pragma disablewhen and #pragma hidewhen

## pragma-disablewhen-and-pragma-hidewhen

`#pragma disablewhen parm_name conditional_expression`

Disables parm_name when the conditional_expression evaluates to true.

`#pragma hidewhen parm_name conditional_expression`

Hides parm_name from the UI when the conditional_expression evaluates to true.

The syntax of the conditional_expression is the same as for the **Disable when** and
**Hide when** options on parameters. See the [help for conditional rules](../ref/windows/optype.html#conditionals)
for more information on the conditional syntax.

```vex
// Disable 'samples' parameter when 'enable' parameter is toggled off
#pragma disablewhen samples { enable == 0 }

// Hide 'choice_dep1' parameter when string menu 'choice' is set to 'off'.
#pragma hidewhen choice_dep1 { choice == "off" }

```

# pragma export

## pragma-export

`#pragma export parm_name (none | dialog | all)`

Adds the parameter UI to the operation parameters dialog, and optionally
also the operator toolbar.

`none` or 0

No export. The parameter only appears in the operator’s standard
dialog.

`dialog` or 1

The parameter appears in the operator’s parameter window.

`all` or 2

The parameter appears in the operator’s parameter window and the
operator toolbar.

You can use `#pragma hint hidden` in combination with this pragma.

# pragma group

## pragma-group

`#pragma group group_name parameter_name1 parameter_name2 ...`

Groups the named parameters into a tab in the UI.

```vex
// Group Ka, Kd, Ks, roughness into a folder called BRDF
#pragma group BRDF Ka Kd Ks
#pragma group BRDF roughness

```

Note
You can put `/` in your group name to create nested groups. For example, the following creates a `Group` tab with two sub-tabs.

```vex
#pragma group "Group/Sub Group 1" p1 p2
#pragma group "Group/Sub Group 2" p3 p4

```

# pragma help and #pragma info

## pragma-help-and-pragma-info

`#pragma help "text"`

Adds text to the help in the dialog script. You can use this to
document functions and parameters for end-users.

```vex
#pragma help "This is help for the VEX function."
#pragma help "It gets added automatically to the help text"

```

Like `#pragma help`, `#pragma info` text is added to the help in the
dialog script. However, the info text is in a separate section at the
beginning of the help. You can use this to specify copyrights, version
information, etc.

```vex
#pragma info "Created by Bob Loblaw - (c) 2006"

```

Currently, only SOPs display info text.

# pragma hint

## pragma-hint

`#pragma hint parameter_name hint_type`

Adds information about what type of value a parameter represents (for
example, a VEX vector may represent a point in space, a color, or a
direction). The hint is used to specialize the UI for editing the
value.

The hint_type can be one of the following:

`none`

No hint.

`toggle`

The integer or float represents an on/off switch (where 1 is on and 0
is off).

`color`

The parameter is a color. The UI will provide color sliders.

`direction`

The parameter is a direction. The UI will provide a gadget for
specifying direction.

`vector`

The parameter is a 3D vector in space. The UI is the same as the
default UI for a parameter with 3 floats, but the channel names will
end with x, y, z instead of 1, 2, 3.

`vector`

The parameter is a 4D vector in space. The UI is the same as the
default UI for a parameter with 4 floats, but the channel names will
end with x, y, z, w instead of 1, 2, 3, 4.

`uv`

The vector parameter is UV coordinates. The UI is two entry fields
instead of three (the third component passed to VEX is always 0),
and the channel names will end with u, v instead of 1, 2.

`uvw`

The vector parameter is UV coordinates. The UI is two entry fields
instead of three (the third component passed to VEX is always 0),
and the channel names will end with u, v instead of 1, 2.

`angle`

The parameter is a direction vector. The UI will provide a direction
gadget.

`file`

The string parameter is a filename. The UI adds a browser button for
specifying the file. (See also `image` and `geometry` below.)

`image`

The string parameter is a filename of an image file. The UI adds a
browser button for specifying the file, and the browser only
displays image file types.

`geometry`

The string parameter is a filename of a geometry file. The UI adds a
browser button for specifying the file, and the browser only
displays geometry file types.

`hidden`

Don’t include this parameter in the parameter list. This is useful when
you intend a parameter to be overridden by a geometry attribute. This
is distinct from `invisible`, as no underlying scene parameter will
be created.

`invisible`

The created parameter will exist in the operator but hidden from the
user interface. Unlike `hidden`, an actual scene parameter will get
created, just not shown.

`inputinvisible`

For VOP operators built using `#pragma optable vop`, this hint
indicates that the VOP input connector should be hidden by default when
creating new nodes. The input will still be accessible through the
`more...` connector at the bottom of the VOP. This hint has no effect
when building non-VOP operator types or for export parameters.

`oplist [opfilter]`

The parameter is a list of objects. You can optionally specify
opfilter to limit the types of
operators allowed in the list. Use `#pragma parmtag` to resolve any
bundles or groups in the list of objects.

See the list of possible operator filters below.

`oppath [opfilter]`

The parameter is an object path. You can optionally specify
opfilter to limit the types of
operators that can be chosen.

See the list of possible operator filters below.

`joinnext`

Put the parameter after this one on the same row in the GUI.
For narrow controls this can save space in the parameter editor.

Example: `#pragma hint myParm joinnext`

```vex
#pragma hint __nondiffuse toggle // Define as a toggle button
#pragma hint specularcolor color // This represents a color
#pragma hint rest hidden         // Don't show rest parameter in UI
#pragma hint mapname image       // This represents an image file
#pragma hint nullobject oppath "obj/null" // Only null objects

```

# pragma inputlabel

## pragma-inputlabel

`#pragma inputlabel inputnum "label"`

For VEX operator types, sets the label for an operator input. This label
appears when the user presses the middle mouse button on one of the
operator inputs. The inputnum is the index of the input, starting at
1\.

```vex
#pragma inputlabel 1 "Geometry to Modify"

```

# pragma label

## pragma-label

`#pragma label parameter_name "text"`

Specifies a descriptive label for a parameter.

```vex
#pragma label amp    "Noise Amplitude"
displacement bumpy(float amp=0) {
...
}

```

# pragma name “text”

## pragma-name-text

Sets the label that appears in the UI. This pragma is obsolete since the
label is now defined in the operator table definition.

# pragma opicon

## opicon

`#pragma opicon "text"`

Use this pragma to set the icon to use for this operator type. It can be
a path to an external .icon or .bicon file, or the name of one of the
standard icons.

vcc’s `-C` [command line option](vcc.html#otl)
overrides this pragma.

# pragma opmininputs and #pragma opmaxinputs

## opinputs

`#pragma opmininputs num`

For VEX operator types, this sets the *minimum* number of inputs that
must be connected to the operator. This value is ignored for SHOPs,
which take no inputs. The operator will generate an error if fewer than
this many nodes are connected as inputs. The `-t` [command line
option](vcc.html#otl) overrides this pragma.

`#pragma opmaxinputs num`

For VEX operator types, this sets the *maximum* number of inputs that
can be connected to the operator. This value is ignored for SHOPs, which
take no inputs. The `-T` [command line option](vcc.html#otl)
overrides this pragma.

```vex
#pragma opmininputs 1
#pragma opmaxinputs 4

```

# pragma opname and #pragma oplabel

## opname

`#pragma opname "text"`

Specifies the internal operator name of this operator type. By default
the compiler uses the name of the source file. vcc’s `-n` [command line
option](vcc.html#otl) overrides this pragma.

`#pragma oplabel "text"`

Specifies a descriptive name for this operator type. By default the
compiler uses internal operator name. vcc’s `-N` [command line
option](vcc.html#otl) overrides this pragma.

```vex
#pragma opname "myshop"
#pragma oplabel "My New Shop"

```

# pragma opscript

## opscript

`#pragma opscript "text"`

If this operator type is a shader that can be used from renderers other
than mantra, this pragma lets you set the name of the shader file that
the renderer should look for. If you use this pragma, the operator type
name does not need to match the shader file name. vcc’s `-S`
[command line option](vcc.html#otl) overrides this pragma.

```vex
#pragma opscript "rman_myshader"

```

# pragma parmhelp parameter_name “text”

## pragma-parmhelp-parameter_name-text

`#pragma parmhelp parameter_name "text"`

Sets the tooltip that appears when the user hovers over
parameter_name.

```vex
#pragma parmhelp amp "Increase this value to add more noise."
displacement bumpy(float amp=0) {
...
}

```

# pragma parmtag

## pragma-parmtag

`#pragma parmtag parmName token value`

Each parameter defined in an operator has a set of token/value tags
associated with it. This pragma adds a token/value pair to a parameter.

The following tokens are available:

autoscope

A *string* of ones or zeros (for example “1011”) corresponding
to each component of the parameter. 1 means the parameter is
auto-added to the channel list when the node is selected.

If there aren’t enough characters in the string for the number of
components, the last character will be extended for the remaining
components. This lets you specify “1” to autoscope all components and “0” to autoscope no components.

If this tag is not defined for a parameter, Houdini guesses if a
parameter should be scoped or not.

editor

For string parameters, this can be used to show a multi-line editor
instead of a single-line input field. Specify “1” to enable.

editorlines

If the multi-line editor is shown, this can be used to specify how many
lines of text are shown. The default value is `10`.

opfilter

For parameters referring to object paths, this provides a filter for
which types of OPs will be presented in the OP chooser. If you set
this tag, you should also set the oprelative tag.

This tag is also set by the  `[#pragma hint  oppath|/vex/pragmas]`  pragma, but using a different
format.

Possible values are:

!!OBJ!!

Only show objects

!!OBJ/GEOMETRY!!

Only show geometry objects

!!OBJ/LIGHT!!

Only show light objects

!!OBJ/CAMERA!!

Only show camera objects (lights are considered cameras)

!!OBJ/BONE!!

Only show bone objects

!!OBJ/FORCE!!

Only show force objects

!!SOP!!

Only show SOPs

!!CHOP!!

Only show CHOPs

!!COP2!!

Only show COPs

!!VOP!!

Only show VOPs

!!ROP!!

Only show output drivers (ROPs)

!!DOP!!

Only show DOPs

!!SHOP!!

Only show SHOPs

!!SHOP/ATMOSPHERE!!

Only show atmosphere shader SHOPs

!!SHOP/BACKGROUND!!

Only show background shader SHOPs

!!SHOP/CONTOUR!!

Only show contour shader SHOPs

!!SHOP/CONTOUR_CONTRAST!!

Only show contour-contrast shader SHOPs

!!SHOP/CONTOUR_STORE!!

Only show contour-store shader SHOPs

!!SHOP/DISPLACEMENT!!

Only show displacement shader SHOPs

!!SHOP/EMITTER!!

Only show emitter shader SHOPs

!!SHOP/GEOMETRY!!

Only show geometry shader SHOPs

!!SHOP/IMAGE3D!!

Only show image3d shader SHOPs

!!SHOP/LENS!!

Only show lens shader SHOPs

!!SHOP/LIGHT!!

Only show light shader SHOPs

!!SHOP/LIGHT_SHADOW!!

Only show light-shadow shader SHOPs

!!SHOP/OUTPUT!!

Only show output shader SHOPs

!!SHOP/PHOTON!!

Only show photon shader SHOPs

!!SHOP/PHOTON_VOLUME!!

Only show photon volume shader SHOPs

!!SHOP/SURFACE!!

Only show surface shader SHOPs

!!SHOP/SURFACE_SHADOW!!

Only show surface-shadow shader SHOPs

oprelative

How paths should be resolved. Typically, the value of this token is
“.” so that paths are resolved relative to the current operator.
However, when referencing objects you can set the tag to “/obj”. See
the example below.

opexpand

In SHOPs, causes “oplist” parameters to be expanded to the full path
names. If a bundle or pattern is chosen, the pattern will be
expanded before it’s sent to the renderer. See the example below.

opfullpath

When used in conjunction with the opexpand tag, causes the path
names of the objects to be fully qualified rather than relative to
the value of the oprelative tag. See the example below.

rampshowcontrolsdefault

Allows you to automatically hide ramp parameter controls, using a value of `1` or `0`.

script_callback

The callback script associated with the parameter. See [#pragma
callback](pragmas.html#callback).

script_ritype

Used when generating RIB streams and mapping the parameter to an
appropriate renderman type specification. The value should be a
valid renderman type (for example, “uniform color”).

script_unquoted

In VOP definitions, indicates the string parameter should be used in
an unquoted form. This allows strings from menus to be placed
verbatim in the code block (see the trig VOP).

sop_input

Used internally by SOPs to determine the place to search for groups
when building a group menu of points/primitives.

```vex
#pragma parmtag lightmask opfilter "!!OBJ/LIGHT!!"
#pragma parmtag lightmask oprelative "/obj"
#pragma parmtag lightmask opexpand 1
#pragma parmtag reflectmask opfilter "!!OBJ/GEOMETRY!!"
#pragma parmtag reflectmask opexpand 1
#pragma parmtag reflectmask opfullpath 1

```

# pragma ramp

## ramp

Creates a [ramp parameter](../network/ramps.html "Explains the tricks to working with ramp parameters.") and connects it to shader
function arguments.

- `#pragma ramp_rgb ramp_parm basis_parm keys_parm values_parm`
- `#pragma ramp_flt ramp_parm basis_parm keys_parm values_parm`

```vex
#pragma ramp_rgb color color_the_basis_strings color_the_key_positions color_the_key_values

surface
sky(
    string color_the_basis_strings[] = { "linear", "linear" };
    float color_the_key_positions[] = { 0, 1};
    vector color_the_key_values[] = { {0,0,0}, {0,0,1} };
) { ... }

```

Unlike most other pragmas which set up UI for parameters defined in the shader,
the ramp pragma creates a new UI element in Houdini.

ramp_parm

The name of the ramp parameter in the Houdini UI. It cannot conflict with any
parameters defined on the shader (or other ramps of course).

basis_parm

A parameter containing an array of strings. These strings determine the basis
values for each key (i.e. “linear”, “constant”). If a constant basis is to be
specified, then the basis_parm argument can be given as an empty string
(i.e. “”). However, the UI may not necessarily reflect this fact.

keys_parm

A parameter containing an array of floats. If the spline is uniform
(i.e. has no keys), the keys_parm argument may be given as an empty string
(i.e. “”). However, the UI may not necessarily reflect this fact.

values_parm

A parameter containing an array of floats or an array of vectors.
Unlike the basis or key parameter, this is mandatory.

Default values are currently not supported by Houdini ramp parameters, so the
initial values given are not passed through to the UI.

Although the values aren’t passed, the number of keys given in the values
parameter will be used to set the initial number of keys in the ramp.

See [ramp parameters](../network/ramps.html "Explains the tricks to working with ramp parameters.") for more information.

# pragma range

## pragma-range

`#pragma range parameter_name[!] min_value max_value[!]`

Defines the ideal range for the parameter. If you append `!`
to a value, the parameter value will be clamped at that range.
The UI also uses this information to set the range of the slider.

```vex
#pragma range seed    0 10
#pragma range roughness 0.001! 1!
#pragma range gamma     0.001! 10

```

# pragma rendermask

## pragma-rendermask

`#pragma rendermask (VMantra | RIB | OGL)`

This pragma is only useful for SHOP dialog generation. Each SHOP has a
mask defining which renderers can use the SHOP. It is possible to have a
similar shader written in the RenderMan shading language and also in VEX
(or another shading language). In this case, the rendermask can be
specified to include more than just VMantra.

The rendermask parameter is closely bound to the code which generates
scene descriptions for a renderer. Thus, the renderer names are quite
specific. The renderers which support SHOPs are…

RIB

RIB generation for RenderMan compliant renderers.

VMantra

The version of mantra which uses VEX for shading.

OGL

OpenGL rendering. This is a special renderer which automatically
adds itself to most render masks. There is currently no way to
prevent this.

# pragma optable

## pragma-optable

`#pragma optable vop`

This pragma can be used to indicate which operator type to generate.
Currently the only supported option is `vop` meaning to create a VOP
operator.

VOP otls generated using this option have a very specific structure:

- The inputs on the VOP are generated from non-export shader parameters
- The outputs on the VOP are generated from export shader parameters
- The parameters on the VOP are created for all non-export shader parameters, with UI generation working the same way it does for SHOP operators
- The outer code imports the shader function from an external file using the [import](shadercalls.html) directive
- The inner code calls this shader using the [shader call](shadercalls.html) syntax

Since the implementation of the shader is not actually stored in the
generated VOP, it must be accessible through an external .vfl or .vex file
for the VOP to work correctly. Placing the .vfl or compiled .vex file for
the shader in a location accessible to the VEX search path will meet this
requirement.

# pragma rename

## pragma-rename

`#pragma rename oldname newname`

Specify a different name for the operator parameter than is used in the
shader interface. This directive should occur after all other `#pragmas`
referring to that parameter.

When using `#pragma optable vop`, this can be useful to allow a VOP input
or output to use the same name as a VEX global variable, since such a name
collision is valid in VOPs but not in VEX.

How to create menus for parameters

## how-to-create-menus-for-parameters

You can define choice pragmas to create a menu UI for a parameter.

You build the list for a parameter by doing one of the following:

- Use multiple choice pragmas for a parameter to create menu items
  (`choice`, `choicescript`, `choicereplace`, `choicetoggle`).
- Use one choicescript pragmas for a parameter to create a
  script that creates the menu items (`choicescript`,
  `choicereplacescript`, `choicetogglescript`).
  This script is run whenever Houdini needs to generate the menu entries
  (i.e. the generated values are not cached) so this script should be as
  efficient as possible. The output of the script must be a series of
  value/label pairs, which have the same meaning as the value and label
  fields in the choice pragma.

The plain `choice` pragma creates a UI that only allows the user to
choose values from the menu. The `choicereplace` pragma creates UI with
a menu but also lets the user enter a different value manually. The
`choicetoggle` pragma allows user input, but also provides a menu where
choosing a menu item adds or removes it to/from the input.

`#pragma choice parameter_name "value" "label"`

Adds a menu item for parameter_name. Displays a menu of mutually exclusive choices.

`#pragma choicescript parameter_name language "scriptline"`

Defines a script that will generate menu items for
parameter_name. language can be either `python` or
`hscript`. Displays a menu of mutually exclusive choices.

`#pragma choicereplace parameter_name "value" "label"`

Adds a menu item for parameter_name.
Displays a UI where the user can choose an item from the menu or input their own.

`#pragma choicereplacescript parameter_name "scriptline"`

Defines a line in a script that will generate menu items for parameter_name.
Displays a UI where the user can choose an item from the menu or input their own.

`#pragma choicetoggle parameter_name "value" "label"`

Adds a menu item for parameter_name.
Displays a menu of choices that add/remove items to/from the user input.

`#pragma choicetogglescript parameter_name "scriptline"`

Defines a line in a script that will generate menu items for parameter_name.
Displays a menu of choices that add/remove items to/from the user input.

```vex
#pragma choice operation "over"    "Composite A Over B"
#pragma choice operation "under"    "Composite A Under B"
#pragma choice operation "add"    "Add A and B"
#pragma choice operation "sub"    "Subtract A from B"
cop texture(string operation="over")
{
if (operation == "over") ...    // texture coordinates
if (operation == "under") ...    // parametric coordinates
if (operation == "add")     ...    // orthographic
if (operation == "sub")     ...    // polar
}

```

This would define a menu for the parameter “operation”. The menu would
consist of 4 entries. The values for the string parameter would be one
of “over”, “under”, “add” or “sub”. However, the user would be presented
with more meaningful labels for the operation types.

```vex
#pragma choice operation "0"    "Use texture coordinates"
#pragma choice operation "1"    "Use parametric coordinates"
#pragma choice operation "2"    "Orthographic Projection"
#pragma choice operation "3"    "Polar Projection"
sop texture(int operation=0)
{
if (operation == 0) ...    // texture coordinates
if (operation == 1) ...    // parametric coordinates
if (operation == 2) ...    // orthographic
if (operation == 3) ...    // polar
}

```

Note
You can use pragmas other than choicereplace and choicetoggle to
create menus for integer parameters, but the menu items ignore the
labels and simply number the choices starting at 0. The
choicereplace and choicetoggle pragmas are only valid for string
parameters.

`#pragma ophidewhen 1`

Adding this pragma to a VOP’s outer code makes the VOP input visibility
follow parameter visibility defined by the hidewhen conditionnals. Inputs
must have parameter with matching names.

`#pragma opextraparm type name_format parm_name0 parm_name1 ...`

Add this pragma to a VOP’s outer code to define an external parameter binding.
type is a vex variable type.
name_format is a simple formatting string without spaces that can include %d or %s special character to refer to node parameter values.
parm_name0 and other trailing arguments are parameter names to use to fill the %d and %s special characters in name_format.

Operator list filters

## operator-list-filters

These are the parameters you can use for the opfilter argument of
certain pragmas.

obj

Any object.

sop

Any SOP.

chop

Any CHOP.

cop

Any COP.

vop

Any VOP.

rop

Any ROP.

obj/geo

Any Geometry object.

obj/null

Any Null object.

obj/light

Any light.

obj/camera

Any camera.

obj/fog

Any Fog object.

obj/bone

Any bone.

shop/surface

Any surface SHOP.

shop/displace

Any Displace SHOP.

shop/interior

Any Interior SHOP.

shop/light

Any Light SHOP.

shop/shadow

Any Shadow SHOP.

shop/fog

Any Atmosphere SHOP.

shop/photon

Any Photon SHOP.

shop/image3d

Any Image3D SHOP.
