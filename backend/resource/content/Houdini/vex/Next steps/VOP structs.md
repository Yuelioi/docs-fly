---
title: VOP structs
order: 8
---
| On this page | * [Overview](#overview) * [Defined and ad-hoc structs](#defined-and-ad-hoc-structs) * [The nodes](#the-nodes) * [How to](#how-to) * [Tips](#tips) * [Storage of defined types](#json) * [Custom languages](#lang) |
| --- | --- |

Overview

## overview

A [VEX structured type](lang.html#structs) (usually called a “struct”) is a compound data type made up of named pieces of sub-data (such as integers, floats, vectors, matrices, and/or other structs) called *members*.

The *definition* of a struct specifies the names, datatypes, and UI labels of the members, like a template. This defines a new *data type*, similar to vectors and matrices. You can then create *instances* of the template and fill in the member values.

You can use structs in VOP networks through the [![](../icons/VOP/struct.svg)Struct](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype."), [![](../icons/VOP/structpack.svg)Struct Pack](../nodes/vop/structpack.html "Bundles input values into an instance of an ad-hoc struct."), and [![](../icons/VOP/structunpack.svg)Struct Unpack](../nodes/vop/structunpack.html "Extracts one or more values from a struct by member name.") nodes.

Note
In Houdini 13, structs are not used by Houdini. They are mostly useful for bundling wires together as a form of organization. Future versions of Houdini will integrate structs more deeply. For example, entities such as ramps and noise settings that are currently represented as multiple wires could be alternately represented as a struct.

Defined and ad-hoc structs

## defined-and-ad-hoc-structs

Houdini can work with structs in VOPs in two ways.

- “Defined” structs are saved in a [file in the Houdini path](vop_structs.html#json) and are available everywhere in Houdini. This is useful for types you want to re-use and/or use in more than one network or share between networks. You create instances of defined structs using the [![](../icons/VOP/struct.svg)Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype.").
- “Ad-hoc” structs are created “on-the-fly” and are only available in the network containing the node. This is useful when you want to bundle wires together to simplify the network. You create ad-hoc struct instances using the [![](../icons/VOP/structpack.svg)Struct Pack node](../nodes/vop/structpack.html "Bundles input values into an instance of an ad-hoc struct.").

The nodes

## the-nodes

- The ![](../icons/VOP/struct.svg) [![](../icons/VOP/struct.svg)Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype.") can bundle separate wires into a struct wire. The node requires that you specify (or create) a *defined struct type*, and creates inputs/outputs based on the defined members.
  You can also use the Struct node to *un-bundle* a struct wire into separate member wires.
- The ![](../icons/VOP/structpack.svg) [![](../icons/VOP/structpack.svg)Struct Pack node](../nodes/vop/structpack.html "Bundles input values into an instance of an ad-hoc struct.") bundles separate wires into a struct wire. Unlike the [![](../icons/VOP/struct.svg)Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype."), it doesn’t require you to specify a defined struct type. Instead, it uses the types of any incoming wires (and names you specify in the parameters) to create an *ad-hoc struct type*.
- The ![](../icons/VOP/structunpack.svg) [![](../icons/VOP/structunpack.svg)Struct Unpack node](../nodes/vop/structunpack.html "Extracts one or more values from a struct by member name.") individual members out of an input struct wire *by name*. Usually you should use the [![](../icons/VOP/struct.svg)Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype.") to unbundle a struct wire into member wires. However, this node lets you write assets/subnetworks that accept *any struct type*, since it extracts members by name and type, whereas of wiring the outputs of Struct is by position.
  If the input struct has a member of the given name but the wrong type, it will cause an error in the network. Houdini forces you to explicitly set the expected type (instead of Houdini automatically changing the output type based on whatever struct type is wired in) to allow you to wire the outputs consistently.

How to

## how-to

| To... | Do this |
| --- | --- |
| Create a new defined struct type | 1. Create a [Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype."). 2. In the parameter editor, click **Edit**. 3. Use the **Create new** menu to create a new 4. Use the [VOP structure editor](../ref/windows/edit_vop_struct.html "Edits VOP struct data types.") to define the members. 5. Click **Accept**. |
| Create an instance of a defined struct | 1. Create a [Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype."). 2. In the parameter editor, choose the struct type from the **Struct name** menu. 3. Wire member data into the inputs. The first output contains the struct instance. |
| Create an ad-hoc struct | 1. Create a [Struct Pack node](../nodes/vop/structpack.html "Bundles input values into an instance of an ad-hoc struct."). 2. Wire member data into the inputs. 3. In the parameter editor, you can name the members. This will make wiring and debugging easier. If the name is blank, Houdini uses the name of the wire. In some cases this will be what you want anyway (for example, if the input is from a Parameter node, the default name is the parameter name). 4. The output contains the struct instance. |
| Extract the members of a struct as separate wires | 1. Create a [Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype."). 2. Wire the struct into the first input. 3. If you want to require that the input be of a certain defined type, in the parameter editor, set the **Struct name** to the type. The node will error if the input is not of that type. 4. The outputs after the first contain the member data. |
| Extract one or more members of a struct by name | 1. Create a [Struct Unpack node](../nodes/vop/structunpack.html "Extracts one or more values from a struct by member name."). 2. Wire the struct into the input. 3. In the parameter editor, set the **Number of members** you want to extract. 4. For each item, set the **Member n name** and the type you expect the member to be. (The node will error if the input struct does not have *all* the specified members or if they are not of the specified type.) 5. The outputs contain the extracted values. |
| Set the value of a struct member | 1. Create a [Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype.") 2. Wire the struct into the first input. 3. Wire new values into the inputs corresponding to the members you want to set. |
| Modify the value of a struct member | 1. Unbundle the struct into separate wires using [Struct](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype.") or [Struct Unpack](../nodes/vop/structunpack.html "Extracts one or more values from a struct by member name.") (see above). 2. Wire the output(s) you want to change into nodes that modify the value. 3. Create a second [Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype.") to re-bundle the struct. Wire the struct (first) output of the unbundling node to the struct (first) input of this node. 4. Insert nodes between the member outputs of the “unbundling” node and the member inputs of the “rebundling” node to change the values. |
| Create a new instance of an ad-hoc struct given a wire of that type | 1. Create a [Struct node](../nodes/vop/struct.html "Creates, modifies, or de-structures an instance of a structured datatype."). 2. Wire the struct into the first input. 3. Wire member data into the inputs. The first output contains the struct instance. |

Tips

## tips

- There is a default wire color for structs, but when you [edit a struct](../ref/windows/edit_vop_struct.html "Edits VOP struct data types.") you can specify a custom color for wires with that struct type.
- To edit the struct definition files on disk “by hand”, see [how structs are stored on disk](vop_structs.html#json) for information about the file locations and format.
- When you create a VOP digital asset, you can specify defined structs as input and output types.
- Since Houdini looks for the definition files on the path (see below), if you have a shared network store on the path, you can define site-wide structs for use by shader writers.

Storage of defined types

## json

The “defined” struct types available across Houdini are loaded from JSON files. Houdini loads any `*.json` files in a path defined by the `HOUDINI_VOP_DEFINITIONS_PATH` environment variable. The default path is `HOUDINI_PATH/vop`. The default file for struct types created using the editor in Houdini is `$HOME/HoudiniX.Y/vop/vopdefinitions.json`.

Houdini will also load struct definitions from an asset if the asset has a section named `VopTypeDefinitions`. This section should contain a `JSON` string in the same format as the files.

The following example shows the basic structure of a struct type definition file. The `connectorColor` key and `label` on members are optional.

```vex
{
    "typeDefinitions": [
        {
            "type": "struct",
            "typeName": "StructTest",
            "connectorColor": {
                "type": "RGB",
                "data": [0.25, 0.2, 0.9]
            },
            "variables": [
                {
                    "typeName": "int",
                    "name": "a"
                },
                {
                    "typeName": "vector",
                    "name": "b",
                    "label": "Label B"
                },
                {
                    "typeName": "matrix3",
                    "name": "c"
                }
            ]
        }
    ]
}

```

The full format is defined in a “JSON schema” file in `$HFS/houdini/vop/TypeDefinitions.json.schema`.

Custom languages

## lang

The [JSON configuration files](vop_structs.html#json) can also define custom languages for VOPs.

A language defines certain aspects and behaviours for VOPs that belong to that
language. For example, it determines whether VOP node input/output data types
support uniform data types (in addition to varying).

The JSON configuration file allows you to define own custom languages and
specify their aspects.

Once you define a language and give it a name, you can set the VOP operator
render mask to that name, to assign the VOP to that language. VOPs like
Parameter or Constant that don’t have any render mask inherit the language from
the parent, when they are created. To explicitly set the language on the parent
subnet, you can use HOM’s @ref hou.VopNode.setShaderLanguageName() function in
the shelf tool that creates that subnet. You can also use that function directly
on the Parameter VOP node (etc.) too.
