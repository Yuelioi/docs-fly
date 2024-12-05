---
title: VEX
order: 1
---
| On this page | * [Define the node’s behavior with a VOP network](#define-the-node-s-behavior-with-a-vop-network) * [Define the node’s behavior with textual VEX code](#define-the-node-s-behavior-with-textual-vex-code) * [How to wire/write the VOPs/VEX](#how-to-wire-write-the-vops-vex) |
| --- | --- |
You can create new surface and particle node types using VOPs/VEX. These nodes evaluate their internal VOP network or VEX program to manipulate point attributes on incoming geometry/particles.

You can define how the node modifies the incoming geometry using a VOP network or with a textual VEX source code.
Define the node’s behavior with a VOP network

## define-the-node-s-behavior-with-a-vop-network

1. In the network editor, go to the **VEX Builder Networks** level (Right-click `obj` in the path and choose **Other networks ▸ vex**).
1. Create a [![](../icons/VEX/sop.svg)SOP Type](../nodes/vex/sop.html "This network defines a surface node (SOP).") network.
1. In the parameter editor for the SOP Type node, set the **SOP type name** to the name you want the node to appear as in the Tab menu.
1. In the network editor, double-click the SOP Type node to go inside and create the VOP network.

Define the node’s behavior with textual VEX code

## define-the-node-s-behavior-with-textual-vex-code

1. Choose **File ▸ New Operator Type**.
1. Set the name and label of the new node type, set the **Operator style** to **VEX type**, and set the **Network Type** to **Geometry Operator**.
1. Click Accept.
1. Use the **Code** tab of the new digital asset to edit the VEX definition. The function with a return type of `sop` is the “main” function Houdini will run to manipulate the geometry.

How to wire/write the VOPs/VEX

## how-to-wire-write-the-vops-vex

- Parameter expressions in VOP nodes **are not evaluated per-point**. You must use VOPs to manipulate the point data.

- You can use [![](../icons/VOP/parameter.svg)Parameter VOP nodes](../nodes/vop/parameter.html "Represents a user-controllable parameter.") nodes (or function parameters in VEX) to create a user interface for the node (controls for the parameters appear in the VOP SOP node’s parameter editor interface).

- If an attribute exists on the input geometry with the same name as a parameter, the parameter value will be overridden by the attribute. (Note that this refers to the “internal” parameter name, not the human-readable parameter *label*.)

- To access point attributes:

  - In VOPs, create a [![](../icons/VOP/global.svg)Global variables VOP node](../nodes/vop/global.html "Provides outputs that represent all the global variables for the
    current VOP network type.") to access common point attributes (from the VOP SOP node’s first input) and feed it through other VOP nodes. For other attributes, you can use a parameter with the same name and type (so it will be overridden by the attribute) or use the [![](../icons/VOP/importattrib.svg)Import Attribute VOP](../nodes/vop/importattrib.html "Imports attribute data from the OP connected to the given input.").
  - In VEX, global variables are available for common point attributes (from the first input) such as `P`. For other attributes, you can use a parameter with the same name and type (so it will be overridden by the attribute) or use the [import()](functions/import.html) function.

- Only the first input’s geometry can be manipulated by the network/program. You can access information from the other inputs using the [![](../icons/VOP/importattrib.svg)Import Attribute VOP](../nodes/vop/importattrib.html "Imports attribute data from the OP connected to the given input.") or the [import()](functions/import.html) function in VEX.

- To add an attribute to the geometry…
  In VOPs use a Parameter node and set **Export** to **Always** or **When input is connected**.

In VEX use a function parameter marked as `export` (e.g. `export float a = 0;`), or the [addattribute()](functions/addattribute.html) function (especially to add/change an attribute whose name is not known at compile time, i.e. the attribute name is controlled by the user interface).

- In VEX code, the function with a return type of `sop` or `pop` is the “main” function that Houdini will call. You can define other “helper” functions alongside this main function as needed.
