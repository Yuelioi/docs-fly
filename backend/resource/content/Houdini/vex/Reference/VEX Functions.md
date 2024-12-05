---
title: VEX Functions
order: 1
---
See [VEX contexts](../contexts/index.html "Guide to the different contexts in which you can write VEX
programs.") to learn about the different contexts (such as surface shaders or displacement shaders) in which the various functions and [statements](../statement.html) are available.
Functions

## subtopics

Arrays

## array_group

- [append](append.html)
  Adds an item to an array or string.
- [argsort](argsort.html)
  Returns the indices of a sorted version of an array.
- [array](array.html)
  Efficiently creates an array from its arguments.
- [foreach](foreach.html)
  Loops over the items in an array, with optional enumeration.
- [insert](insert.html)
  Inserts an item, array, or string into an array or string.
- [isvalidindex](isvalidindex.html)
  Checks if the index given is valid for the array or string given.
- [len](len.html)
  Returns the length of an array.
- [pop](pop.html)
  Removes the last element of an array and returns it.
- [push](push.html)
  Adds an item to an array.
- [removeindex](removeindex.html)
  Removes an item at the given index from an array.
- [removevalue](removevalue.html)
  Removes an item from an array.
- [reorder](reorder.html)
  Reorders items in an array or string.
- [resize](resize.html)
  Sets the length of an array.
- [reverse](reverse.html)
  Returns an array or string in reverse order.
- [slice](slice.html)
  Slices a sub-string or sub-array of a string or array.
- [sort](sort.html)
  Returns the array sorted in increasing order.
- [upush](upush.html)
  Adds a uniform item to an array.

Attributes and Intrinsics

## attrib_group

- [addattrib](addattrib.html)
  Adds an attribute to a geometry.
- [adddetailattrib](adddetailattrib.html)
  Adds a detail attribute to a geometry.
- [addpointattrib](addpointattrib.html)
  Adds a point attribute to a geometry.
- [addprimattrib](addprimattrib.html)
  Adds a primitive attribute to a geometry.
- [addvertexattrib](addvertexattrib.html)
  Adds a vertex attribute to a geometry.
- [addvisualizer](addvisualizer.html)
  Appends to a geometry’s visualizer detail attribute.
- [attrib](attrib.html)
  Reads the value of an attribute from geometry.
- [attribclass](attribclass.html)
  Returns the class of a geometry attribute.
- [attribdataid](attribdataid.html)
  Returns the data id of a geometry attribute.
- [attribsize](attribsize.html)
  Returns the size of a geometry attribute.
- [attribtype](attribtype.html)
  Returns the type of a geometry attribute.
- [attribtypeinfo](attribtypeinfo.html)
  Returns the transformation metadata of a geometry attribute.
- [curvearclen](curvearclen.html)
  Evaluates the length of an arc on a primitive defined by an array of points using parametric uv coordinates.
- [detail](detail.html)
  Reads the value of a detail attribute value from a geometry.
- [detailattrib](detailattrib.html)
  Reads a detail attribute value from a geometry.
- [detailattribsize](detailattribsize.html)
  Returns the size of a geometry detail attribute.
- [detailattribtype](detailattribtype.html)
  Returns the type of a geometry detail attribute.
- [detailattribtypeinfo](detailattribtypeinfo.html)
  Returns the type info of a geometry attribute.
- [detailintrinsic](detailintrinsic.html)
  Reads the value of a detail intrinsic from a geometry.
- [findattribval](findattribval.html)
  Finds a primitive/point/vertex that has a certain attribute value.
- [findattribvalcount](findattribvalcount.html)
  Returns number of elements where an integer or string attribute has a certain value.
- [getattrib](getattrib.html)
  Reads an attribute value from geometry, with validity check.
- [getattribute](getattribute.html)
  Copies the value of a geometry attribute into a variable and returns a success flag.
- [hasattrib](hasattrib.html)
  Checks whether a geometry attribute exists.
- [hasdetailattrib](hasdetailattrib.html)
  Returns if a geometry detail attribute exists.
- [haspointattrib](haspointattrib.html)
  Returns if a geometry point attribute exists.
- [hasprimattrib](hasprimattrib.html)
  Returns if a geometry prim attribute exists.
- [hasvertexattrib](hasvertexattrib.html)
  Returns if a geometry vertex attribute exists.
- [idtopoint](idtopoint.html)
  Finds a point by its id attribute.
- [idtoprim](idtoprim.html)
  Finds a primitive by its id attribute.
- [nametopoint](nametopoint.html)
  Finds a point by its name attribute.
- [nametoprim](nametoprim.html)
  Finds a primitive by its name attribute.
- [nuniqueval](nuniqueval.html)
  Returns the number of unique values from an integer or string attribute.
- [point](point.html)
  Reads a point attribute value from a geometry.
- [pointattrib](pointattrib.html)
  Reads a point attribute value from a geometry and outputs a success/fail flag.
- [pointattribsize](pointattribsize.html)
  Returns the size of a geometry point attribute.
- [pointattribtype](pointattribtype.html)
  Returns the type of a geometry point attribute.
- [pointattribtypeinfo](pointattribtypeinfo.html)
  Returns the type info of a geometry attribute.
- [pointlocaltransforms](pointlocaltransforms.html)
  Returns an array of point localtransforms from an array of point indices.
- [pointtransform](pointtransform.html)
  Returns a point transform from a point index.
- [pointtransformrigid](pointtransformrigid.html)
  Returns a rigid point transform from a point index.
- [pointtransforms](pointtransforms.html)
  Returns an array of point transforms from an array of point indices.
- [pointtransformsrigid](pointtransformsrigid.html)
  Returns an array of rigid point transforms from an array of point indices.
- [prim](prim.html)
  Reads a primitive attribute value from a geometry.
- [prim_attribute](prim_attribute.html)
  Interpolates the value of an attribute at a certain parametric (u, v) position and copies it into a variable.
- [primarclen](primarclen.html)
  Evaluates the length of an arc on a primitive using parametric uv coordinates.
- [primattrib](primattrib.html)
  Reads a primitive attribute value from a geometry, outputting a success flag.
- [primattribsize](primattribsize.html)
  Returns the size of a geometry prim attribute.
- [primattribtype](primattribtype.html)
  Returns the type of a geometry prim attribute.
- [primattribtypeinfo](primattribtypeinfo.html)
  Returns the type info of a geometry attribute.
- [primduv](primduv.html)
  Returns position derivative on a primitive at a certain parametric (u, v) position.
- [priminteriorweights](priminteriorweights.html)
  Finds the indices and weightings of the vertices that will compute an
  interior point given the UVW coordinates.
- [primintrinsic](primintrinsic.html)
  Reads a primitive intrinsic from a geometry.
- [primuv](primuv.html)
  Interpolates the value of an attribute at a certain parametric (uvw) position.
- [primuvconvert](primuvconvert.html)
  Convert parametric UV locations on curve primitives between different spaces.
- [removedetailattrib](removedetailattrib.html)
  Removes a detail attribute from a geometry.
- [removepointattrib](removepointattrib.html)
  Removes a point attribute from a geometry.
- [removepointgroup](removepointgroup.html)
  Removes a point group from a geometry.
- [removeprimattrib](removeprimattrib.html)
  Removes a primitive attribute from a geometry.
- [removeprimgroup](removeprimgroup.html)
  Removes a primitive group from a geometry.
- [removevertexattrib](removevertexattrib.html)
  Removes a vertex attribute from a geometry.
- [removevertexgroup](removevertexgroup.html)
  Removes a vertex group from a geometry.
- [setattrib](setattrib.html)
  Writes an attribute value to geometry.
- [setattribtypeinfo](setattribtypeinfo.html)
  Sets the meaning of an attribute in geometry.
- [setdetailattrib](setdetailattrib.html)
  Sets a detail attribute in a geometry.
- [setdetailintrinsic](setdetailintrinsic.html)
  Sets the value of a writeable detail intrinsic attribute.
- [setpointattrib](setpointattrib.html)
  Sets a point attribute in a geometry.
- [setpointlocaltransforms](setpointlocaltransforms.html)
  Sets an array of point local transforms at the given point indices.
- [setpointtransform](setpointtransform.html)
  Sets the world space transform of a given point
- [setpointtransforms](setpointtransforms.html)
  Sets an array of point transforms at the given point indices.
- [setprimattrib](setprimattrib.html)
  Sets a primitive attribute in a geometry.
- [setprimintrinsic](setprimintrinsic.html)
  Sets the value of a writeable primitive intrinsic attribute.
- [setvertexattrib](setvertexattrib.html)
  Sets a vertex attribute in a geometry.
- [uniqueval](uniqueval.html)
  Returns one of the set of unique values across all values for an int or string attribute.
- [uniquevals](uniquevals.html)
  Returns the set of unique values across all values for an int or string attribute.
- [uvsample](uvsample.html)
  Interpolates the value of an attribute at certain UV coordinates using a UV attribute.
- [vertex](vertex.html)
  Reads a vertex attribute value from a geometry.
- [vertexattrib](vertexattrib.html)
  Reads a vertex attribute value from a geometry.
- [vertexattribsize](vertexattribsize.html)
  Returns the size of a geometry vertex attribute.
- [vertexattribtype](vertexattribtype.html)
  Returns the type of a geometry vertex attribute.
- [vertexattribtypeinfo](vertexattribtypeinfo.html)
  Returns the type info of a geometry attribute.

BSDFs

## bsdf_group

- [albedo](albedo.html)
  Returns the albedo (percentage of reflected light) for a bsdf given the outgoing light direction.
- [ashikhmin](ashikhmin.html)
  Returns a specular BSDF using the Ashikhmin shading model.
- [blinn](blinn.html)
  Returns a Blinn BSDF or computes Blinn shading.
- [chiang](chiang.html)
  Returns a chiang BSDF.
- [chiang_fur](chiang_fur.html)
  Returns a chiang_fur BSDF.
- [cone](cone.html)
  Returns a cone reflection BSDF.
- [cvex_bsdf](cvex_bsdf.html)
  Creates a bsdf object from two CVEX shader strings.
- [diffuse](diffuse.html)
  Returns a diffuse BSDF or computes diffuse shading.
- [eval_bsdf](eval_bsdf.html)
  Evaluates a bsdf given two vectors.
- [getbounces](getbounces.html)
- [ggx](ggx.html)
  Returns a ggx BSDF.
- [hair](hair.html)
  Returns a BSDF for shading hair.
- [henyeygreenstein](henyeygreenstein.html)
  Returns an anisotropic volumetric BSDF, which can scatter light forward or backward.
- [isotropic](isotropic.html)
  Returns an isotropic BSDF, which scatters light equally in all directions.
- [mask_bsdf](mask_bsdf.html)
  Returns new BSDF that only includes the components specified by the mask.
- [normal_bsdf](normal_bsdf.html)
  Returns the normal for the diffuse component of a BSDF.
- [phong](phong.html)
  Returns a Phong BSDF or computes Phong shading.
- [phonglobe](phonglobe.html)
- [sample_bsdf](sample_bsdf.html)
  Samples a BSDF.
- [solid_angle](solid_angle.html)
  Computes the solid angle (in steradians) a BSDF function subtends.
- [split_bsdf](split_bsdf.html)
  Splits a bsdf into its component lobes.
- [sssapprox](sssapprox.html)
  Creates an approximate SSS BSDF.

BSDFs

## BSDFs_group

- [specular](specular.html)
  Returns a specular BSDF or computes specular shading.

CHOP

## CHOP_group

- [chadd](chadd.html)
  Adds new channels to a CHOP node.
- [chattr](chattr.html)
  Reads from a CHOP attribute.
- [chattrnames](chattrnames.html)
  Reads CHOP attribute names of a given attribute class from a CHOP input.
- [chend](chend.html)
  Returns the sample number of the last sample in a given CHOP input.
- [chendf](chendf.html)
  Returns the frame corresponding to the last sample of the input specified.
- [chendt](chendt.html)
  Returns the time corresponding to the last sample of the input
  specified.
- [chindex](chindex.html)
  Returns the channel index from a input given a channel name.
- [chinput](chinput.html)
  Returns the value of a channel at the specified sample.
- [chinputlimits](chinputlimits.html)
  Computes the minimum and maximum value of samples in an input channel.
- [chnames](chnames.html)
  Returns all the CHOP channel names of a given CHOP input.
- [chnumchan](chnumchan.html)
  Returns the number of channels in the input specified.
- [chop](chop.html)
  Returns the value of a CHOP channel at the specified sample.
- [choplocal](choplocal.html)
  Returns the value of a CHOP local transform channel at the specified sample.
- [choplocalt](choplocalt.html)
  Returns the value of a CHOP local transform channel at the specified sample and evaluation time.
- [chopt](chopt.html)
  Returns the value of a CHOP channel at the specified sample and evaluation time.
- [chrate](chrate.html)
  Returns the sample rate of the input specified.
- [chreadbuf](chreadbuf.html)
  Returns the value of CHOP context temporary buffer at the specified index.
- [chremove](chremove.html)
  Removes channels from a CHOP node.
- [chremoveattr](chremoveattr.html)
  Removes a CHOP attribute.
- [chrename](chrename.html)
  Renames a CHOP channel.
- [chresizebuf](chresizebuf.html)
  Resize the CHOP context temporary buffer
- [chsetattr](chsetattr.html)
  Sets the value of a CHOP attribute.
- [chsetlength](chsetlength.html)
  Sets the length of the CHOP channel data.
- [chsetrate](chsetrate.html)
  Sets the sampling rate of the CHOP channel data.
- [chsetstart](chsetstart.html)
  Sets the CHOP start sample in the channel data.
- [chstart](chstart.html)
  Returns the start sample of the input specified.
- [chstartf](chstartf.html)
  Returns the frame corresponding to the first sample of the input
  specified.
- [chstartt](chstartt.html)
  Returns the time corresponding to the first sample of the input
  specified.
- [chwritebuf](chwritebuf.html)
  Writes a value of CHOP context temporary buffer at the specified index.
- [isframes](isframes.html)
  Returns 1 if the Vex CHOP’s Unit Menu is currently set to 'frames', 0
  otherwise.
- [issamples](issamples.html)
  Returns 1 if the Vex CHOP’s Unit Menu is currently set to 'samples',
  0 otherwise.
- [isseconds](isseconds.html)
  Returns 1 if the Vex CHOP’s Unit Menu is currently set to 'seconds',
  0 otherwise.
- [ninputs](ninputs.html)
  Returns the number of inputs.

Channel Primitives

## chprim_group

- [chprim_clear](chprim_clear.html)
  Clears a channel primitive, removing all keys.
- [chprim_destroykey](chprim_destroykey.html)
  Destroy an existing key from a channel primitive.
- [chprim_end](chprim_end.html)
  Get the end time of a channel primitive.
- [chprim_eval](chprim_eval.html)
  Evaluate a channel primitive at the given time.
- [chprim_insertkey](chprim_insertkey.html)
  Insert a key into a channel primitive.
- [chprim_keycount](chprim_keycount.html)
  Get the number of keys in a channel primitive.
- [chprim_keytimes](chprim_keytimes.html)
  Get the key times of a channel primitive.
- [chprim_keyvalues](chprim_keyvalues.html)
  Get the key values of a channel primitive.
- [chprim_length](chprim_length.html)
  Get the length of a channel primitive.
- [chprim_setkeyaccel](chprim_setkeyaccel.html)
  Set the acceleration of a channel primitive key.
- [chprim_setkeyslope](chprim_setkeyslope.html)
  Set the slope of a channel primitive key.
- [chprim_setkeyvalue](chprim_setkeyvalue.html)
  Set the value of a channel primitive key.
- [chprim_start](chprim_start.html)
  Get the start time of a channel primitive.

color

## color_group

- [blackbody](blackbody.html)
  Compute the color value of an incandescent black body.
- [ctransform](ctransform.html)
  Transforms between color spaces.
- [luminance](luminance.html)
  Compute the luminance of the RGB color specified by the parameters.

Conversion

## convert_group

- [atof](atof.html)
  Converts a string to a float.
- [atoi](atoi.html)
  Converts a string to an integer.
- [cracktransform](cracktransform.html)
  Depending on the value of c, returns the translate (c=0), rotate
  (c=1), scale (c=2), or shears (c=3) component of the transform (xform).
- [degrees](degrees.html)
  Converts the argument from radians into degrees.
- [eulertoquaternion](eulertoquaternion.html)
  Creates a vector4 representing a quaternion from euler angles.
- [hsvtorgb](hsvtorgb.html)
  Convert HSV color space into RGB color space.
- [qconvert](qconvert.html)
  Converts a quaternion represented by a vector4 to a matrix3 representation.
- [quaterniontoeuler](quaterniontoeuler.html)
  Creates a euler angle representing a quaternion.
- [radians](radians.html)
  Converts the argument from degrees into radians.
- [rgbtohsv](rgbtohsv.html)
  Convert RGB color space to HSV color space.
- [rgbtoxyz](rgbtoxyz.html)
  Convert a linear sRGB triplet to CIE XYZ tristimulus values.
- [serialize](serialize.html)
  Flattens an array of vector or matrix types into an array of floats.
- [unserialize](unserialize.html)
  Turns a flat array of floats into an array of vectors or matrices.
- [xyztorgb](xyztorgb.html)
  Convert CIE XYZ tristimulus values to a linear sRGB triplet.

Crowds

## crowd_group

- [agentaddclip](agentaddclip.html)
  Add a clip into an agent’s definition.
- [agentchannelcount](agentchannelcount.html)
  Returns the number of channels in an agent primitive’s rig.
- [agentchannelnames](agentchannelnames.html)
  Returns the names of the channels in an agent primitive’s rig.
- [agentchannelvalue](agentchannelvalue.html)
  Returns the current value of an agent primitive’s channel.
- [agentchannelvalues](agentchannelvalues.html)
  Returns the current values of an agent primitive’s channels.
- [agentclipcatalog](agentclipcatalog.html)
  Returns all of the animation clips that have been loaded for an agent primitive.
- [agentclipchannel](agentclipchannel.html)
  Finds the index of a channel in an agent’s animation clip.
- [agentclipchannelnames](agentclipchannelnames.html)
  Returns the names of the channels in an agent’s animation clip.
- [agentcliplayerblend](agentcliplayerblend.html)
  Blends values according to an agent’s animation layers.
- [agentcliplength](agentcliplength.html)
  Returns the length (in seconds) of an agent’s animation clip.
- [agentclipnames](agentclipnames.html)
  Returns an agent primitive’s current animation clips.
- [agentclipsample](agentclipsample.html)
  Samples a channel of an agent’s clip at a specific time.
- [agentclipsamplelocal](agentclipsamplelocal.html)
  Samples an agent’s animation clip at a specific time.
- [agentclipsamplerate](agentclipsamplerate.html)
  Returns the sample rate of an agent’s animation clip.
- [agentclipsampleworld](agentclipsampleworld.html)
  Samples an agent’s animation clip at a specific time.
- [agentclipstarttime](agentclipstarttime.html)
  Returns the start time (in seconds) of an agent’s animation clip.
- [agentcliptimes](agentcliptimes.html)
  Returns the current times for an agent primitive’s animation clips.
- [agentcliptransformgroups](agentcliptransformgroups.html)
  Returns the transform groups for an agent primitive’s current animation clips.
- [agentclipweights](agentclipweights.html)
  Returns the blend weights for an agent primitive’s animation clips.
- [agentcollisionlayer](agentcollisionlayer.html)
  Returns the name of the collision layer of an agent primitive.
- [agentcollisionlayers](agentcollisionlayers.html)
  Returns the names of an agent primitive’s collision layers.
- [agentcurrentlayer](agentcurrentlayer.html)
  Returns the name of the current layer of an agent primitive.
- [agentcurrentlayers](agentcurrentlayers.html)
  Returns the names of an agent primitive’s current layers.
- [agentfindclip](agentfindclip.html)
  Finds the index of a clip in an agent’s definition.
- [agentfindlayer](agentfindlayer.html)
  Finds the index of a layer in an agent’s definition.
- [agentfindtransformgroup](agentfindtransformgroup.html)
  Finds the index of a transform group in an agent’s definition.
- [agentlayerbindings](agentlayerbindings.html)
  Returns the transform that each shape in an agent’s layer is bound to.
- [agentlayers](agentlayers.html)
  Returns all of the layers that have been loaded for an agent primitive.
- [agentlayershapes](agentlayershapes.html)
  Returns the names of the shapes referenced by an agent primitive’s layer.
- [agentlocaltransform](agentlocaltransform.html)
  Returns the current local space transform of an agent primitive’s bone.
- [agentlocaltransforms](agentlocaltransforms.html)
  Returns the current local space transforms of an agent primitive.
- [agentmetadata](agentmetadata.html)
  Returns the agent definition’s metadata dictionary.
- [agentrestlocaltransform](agentrestlocaltransform.html)
  Returns the local space rest transform for an agent primitive’s joint.
- [agentrestworldtransform](agentrestworldtransform.html)
  Returns the world space rest transform for an agent primitive’s joint.
- [agentrigchildren](agentrigchildren.html)
  Returns the child transforms of a transform in an agent primitive’s rig.
- [agentrigfind](agentrigfind.html)
  Finds the index of a transform in an agent primitive’s rig.
- [agentrigfindchannel](agentrigfindchannel.html)
  Finds the index of a channel in an agent primitive’s rig.
- [agentrigparent](agentrigparent.html)
  Returns the parent transform of a transform in an agent primitive’s rig.
- [agentsolvefbik](agentsolvefbik.html)
  Applies a full-body inverse kinematics algorithm to an agent’s skeleton.
- [agenttransformcount](agenttransformcount.html)
  Returns the number of transforms in an agent primitive’s rig.
- [agenttransformgroupmember](agenttransformgroupmember.html)
  Returns whether a transform is a member of the specified transform group.
- [agenttransformgroupmemberchannel](agenttransformgroupmemberchannel.html)
  Returns whether a channel is a member of the specified transform group.
- [agenttransformgroups](agenttransformgroups.html)
  Returns the names of the transform groups in an agent’s definition.
- [agenttransformgroupweight](agenttransformgroupweight.html)
  Returns the weight of a member of the specified transform group.
- [agenttransformnames](agenttransformnames.html)
  Returns the name of each transform in an agent primitive’s rig.
- [agenttransformtolocal](agenttransformtolocal.html)
  Converts transforms from world space to local space for an agent primitive.
- [agenttransformtoworld](agenttransformtoworld.html)
  Converts transforms from local space to world space for an agent primitive.
- [agentworldtransform](agentworldtransform.html)
  Returns the current world space transform of an agent primitive’s bone.
- [agentworldtransforms](agentworldtransforms.html)
  Returns the current world space transforms of an agent primitive.
- [setagentchannelvalue](setagentchannelvalue.html)
  Overrides the value of an agent primitive’s channel.
- [setagentchannelvalues](setagentchannelvalues.html)
  Overrides the values of an agent primitive’s channels.
- [setagentclipnames](setagentclipnames.html)
  Sets the current animation clips for an agent primitive.
- [setagentclips](setagentclips.html)
  Sets the animation clips that an agent should use to compute its transforms.
- [setagentcliptimes](setagentcliptimes.html)
  Sets the current times for an agent primitive’s animation clips.
- [setagentclipweights](setagentclipweights.html)
  Sets the blend weights for an agent primitive’s animation clips.
- [setagentcollisionlayer](setagentcollisionlayer.html)
  Sets the collision layer of an agent primitive.
- [setagentcollisionlayers](setagentcollisionlayers.html)
  Sets the collision layers of an agent primitive.
- [setagentcurrentlayer](setagentcurrentlayer.html)
  Sets the current layer of an agent primitive.
- [setagentcurrentlayers](setagentcurrentlayers.html)
  Sets the current display layers of an agent primitive.
- [setagentlocaltransform](setagentlocaltransform.html)
  Overrides the local space transform of an agent primitive’s bone.
- [setagentlocaltransforms](setagentlocaltransforms.html)
  Overrides the local space transforms of an agent primitive.
- [setagentworldtransform](setagentworldtransform.html)
  Overrides the world space transform of an agent primitive’s bone.
- [setagentworldtransforms](setagentworldtransforms.html)
  Overrides the world space transforms of an agent primitive.

dict

## dict_group

- [json_dumps](json_dumps.html)
  Converts a VEX dictionary into a JSON string.
- [json_loads](json_loads.html)
  Converts a JSON string into a VEX dictionary.
- [keys](keys.html)
  Returns all the keys in a dictionary.
- [typeid](typeid.html)
  Returns a numeric code identifying a VEX data type.

displace

## displace_group

- [dimport](dimport.html)
  Reads a variable from the displacement shader for the surface.

File I/O

## file_group

- [file_stat](file_stat.html)
  Returns file system status for a given file.

filter

## filter_group

- [filter_remap](filter_remap.html)
  Computes an importance sample based on the given filter type and input uv.

Fuzzy Logic

## fuzzy_group

- [fuzzify](fuzzify.html)
- [fuzzy_and](fuzzy_and.html)
- [fuzzy_defuzz_centroid](fuzzy_defuzz_centroid.html)
- [fuzzy_nand](fuzzy_nand.html)
- [fuzzy_nor](fuzzy_nor.html)
- [fuzzy_not](fuzzy_not.html)
- [fuzzy_nxor](fuzzy_nxor.html)
- [fuzzy_or](fuzzy_or.html)
- [fuzzy_xor](fuzzy_xor.html)

Geometry

## geo_group

- [addpoint](addpoint.html)
  Adds a point to the geometry.
- [addprim](addprim.html)
  Adds a primitive to the geometry.
- [addvertex](addvertex.html)
  Adds a vertex to a primitive in a geometry.
- [clip](clip.html)
  Clip the line segment between p0 and p1.
- [geoself](geoself.html)
  Returns a handle to the current geometry.
- [geounwrap](geounwrap.html)
  Returns an oppath: string to unwrap the geometry in-place.
- [inedgegroup](inedgegroup.html)
  Returns 1 if the edge specified by the point pair is in the group specified by the string.
- [intersect](intersect.html)
  This function computes the first intersection of a ray with geometry.
- [intersect_all](intersect_all.html)
  Computes all intersections of the specified ray with geometry.
- [minpos](minpos.html)
  Given a position in world space, returns the position of the closest point on a given geometry.
- [nearpoint](nearpoint.html)
  Finds the closest point in a geometry.
- [nearpoints](nearpoints.html)
  Finds the all the closest point in a geometry.
- [nedgesgroup](nedgesgroup.html)
  Returns the number of edges in the group.
- [neighbour](neighbour.html)
  Returns the point number of the next point connected to a given point.
- [neighbourcount](neighbourcount.html)
  Returns the number of points that are connected to the specified point.
- [neighbours](neighbours.html)
  Returns an array of the point numbers of the neighbours of a point.
- [npoints](npoints.html)
  Returns the number of points in the input or geometry file.
- [nprimitives](nprimitives.html)
  Returns the number of primitives in the input or geometry file.
- [nvertices](nvertices.html)
  Returns the number of vertices in the input or geometry file.
- [nverticesgroup](nverticesgroup.html)
  Returns the number of vertices in the group.
- [pointprims](pointprims.html)
  Returns the list of primitives containing a point.
- [pointvertex](pointvertex.html)
  Returns a linear vertex number of a point in a geometry.
- [pointvertices](pointvertices.html)
  Returns the list of vertices connected to a point.
- [polyneighbours](polyneighbours.html)
  Returns an array of the primitive numbers of the edge-neighbours of a polygon.
- [primfind](primfind.html)
  Returns a list of primitives potentially intersecting a given bounding box.
- [primpoint](primpoint.html)
  Converts a primitive/vertex pair into a point number.
- [primpoints](primpoints.html)
  Returns the list of points on a primitive.
- [primvertex](primvertex.html)
  Converts a primitive/vertex pair into a linear vertex.
- [primvertexcount](primvertexcount.html)
  Returns number of vertices in a primitive in a geometry.
- [primvertices](primvertices.html)
  Returns the list of vertices on a primitive.
- [removeattrib](removeattrib.html)
  Removes an attribute or group from the geometry.
- [removepoint](removepoint.html)
  Removes a point from the geometry.
- [removeprim](removeprim.html)
  Removes a primitive from the geometry.
- [removevertex](removevertex.html)
  Removes a vertex from the geometry.
- [setedgegroup](setedgegroup.html)
  Sets edge group membership in a geometry.
- [setprimvertex](setprimvertex.html)
  Rewires a vertex in the geometry to a different point.
- [setvertexpoint](setvertexpoint.html)
  Rewires a vertex in the geometry to a different point.
- [uvintersect](uvintersect.html)
  This function computes the intersection of the specified ray with the geometry in uv space.
- [vertexcurveparam](vertexcurveparam.html)
  Returns the parametric coordinate of a vertex along the perimeter of
  its primitive.
- [vertexindex](vertexindex.html)
  Converts a primitive/vertex pair into a linear vertex.
- [vertexnext](vertexnext.html)
  Returns the linear vertex number of the next vertex sharing a point with a given vertex.
- [vertexpoint](vertexpoint.html)
  Returns the point number of linear vertex in a geometry.
- [vertexprev](vertexprev.html)
  Returns the linear vertex number of the previous vertex sharing a point with a given vertex.
- [vertexprim](vertexprim.html)
  Returns the number of the primitive containing a given vertex.
- [vertexprimindex](vertexprimindex.html)
  Converts a linear vertex index into a primitive vertex number.

groups

## groups_group

- [expandedgegroup](expandedgegroup.html)
- [expandpointgroup](expandpointgroup.html)
  Returns an array of point numbers corresponding to a group string.
- [expandprimgroup](expandprimgroup.html)
  Returns an array of prim numbers corresponding to a group string.
- [expandvertexgroup](expandvertexgroup.html)
  Returns an array of linear vertex numbers corresponding to a group string.
- [inpointgroup](inpointgroup.html)
  Returns 1 if the point specified by the point number is in the group specified by the string.
- [inprimgroup](inprimgroup.html)
  Returns 1 if the primitive specified by the primitive number is in the group specified by the string.
- [invertexgroup](invertexgroup.html)
  Returns 1 if the vertex specified by the vertex number is in the group specified by the string.
- [npointsgroup](npointsgroup.html)
  Returns the number of points in the group.
- [nprimitivesgroup](nprimitivesgroup.html)
  Returns the number of primitives in the group.
- [setpointgroup](setpointgroup.html)
  Adds or removes a point to/from a group in a geometry.
- [setprimgroup](setprimgroup.html)
  Adds or removes a primitive to/from a group in a geometry.
- [setvertexgroup](setvertexgroup.html)
  Adds or removes a vertex to/from a group in a geometry.

Half-edges

## hedge_group

- [hedge_dstpoint](hedge_dstpoint.html)
  Returns the destination point of a half-edge.
- [hedge_dstvertex](hedge_dstvertex.html)
  Returns the destination vertex of a half-edge.
- [hedge_equivcount](hedge_equivcount.html)
  Returns the number of half-edges equivalent to a given half-edge.
- [hedge_isequiv](hedge_isequiv.html)
  Determines whether a two half-edges are equivalent (represent the same edge).
- [hedge_isprimary](hedge_isprimary.html)
  Determines whether a half-edge number corresponds to a primary half-edge.
- [hedge_isvalid](hedge_isvalid.html)
  Determines whether a half-edge number corresponds to a valid half-edge.
- [hedge_next](hedge_next.html)
  Returns the half-edge that follows a given half-edge in its polygon.
- [hedge_nextequiv](hedge_nextequiv.html)
  Returns the next half-edges equivalent to a given half-edge.
- [hedge_postdstpoint](hedge_postdstpoint.html)
  Returns the point into which the vertex following the destination vertex of a half-edge in its primitive is wired.
- [hedge_postdstvertex](hedge_postdstvertex.html)
  Returns the vertex following the destination vertex of a half-edge in its primitive.
- [hedge_presrcpoint](hedge_presrcpoint.html)
  Returns the point into which the vertex that precedes the source vertex of a half-edge in its primitive is wired.
- [hedge_presrcvertex](hedge_presrcvertex.html)
  Returns the vertex that precedes the source vertex of a half-edge in its primitive.
- [hedge_prev](hedge_prev.html)
  Returns the half-edge that precedes a given half-edge in its polygon.
- [hedge_prim](hedge_prim.html)
  Returns the primitive that contains a half-edge.
- [hedge_primary](hedge_primary.html)
  Returns the primary half-edge equivalent to a given half-edge.
- [hedge_srcpoint](hedge_srcpoint.html)
  Returns the source point of a half-edge.
- [hedge_srcvertex](hedge_srcvertex.html)
  Returns the source vertex of a half-edge.
- [pointedge](pointedge.html)
  Finds and returns a half-edge with the given endpoints.
- [pointhedge](pointhedge.html)
  Finds and returns a half-edge with a given source point or with given source and destination points.
- [pointhedgenext](pointhedgenext.html)
  Returns the next half-edge with the same source as a given half-edge.
- [primhedge](primhedge.html)
  Returns one of the half-edges contained in a primitive.
- [vertexhedge](vertexhedge.html)
  Returns the half-edge which has a vertex as source.

hex

## hex_group

- [hex_adjacent](hex_adjacent.html)
  Returns primitive number of an adjacent hexahedron.
- [hex_faceindex](hex_faceindex.html)
  Returns vertex indices of each face of a hexahedron.

Image Processing

## image_group

- [accessframe](accessframe.html)
  Tells the COP manager that you need access to the given frame.
- [alphaname](alphaname.html)
  Returns the default name of the alpha plane (as it appears in the
  compositor preferences).
- [binput](binput.html)
  Samples a 2×2 pixel block around the given UV position, and bilinearly interpolates these pixels.
- [bumpname](bumpname.html)
  Returns the default name of the bump plane (as it appears in the
  compositor preferences).
- [chname](chname.html)
  Returns the name of a numbered channel.
- [cinput](cinput.html)
  Samples the exact (unfiltered) pixel color at the given coordinates.
- [colorname](colorname.html)
  Returns the default name of the color plane (as it appears in the
  compositor preferences).
- [depthname](depthname.html)
  Returns the default name of the depth plane (as it appears in the
  compositor preferences).
- [dsmpixel](dsmpixel.html)
  Reads the z-records stored in a pixel of a deep shadow map
  or deep camera map.
- [finput](finput.html)
  Returns fully filtered pixel input.
- [hasmetadata](hasmetadata.html)
  Queries if metadata exists on a composite operator.
- [hasplane](hasplane.html)
  Returns 1 if the plane specified by the parameter exists in this
  COP.
- [iaspect](iaspect.html)
  Returns the aspect ratio of the specified input.
- [ichname](ichname.html)
  Returns the channel name of the indexed plane of the given input.
- [iend](iend.html)
  Returns the last frame of the specified input.
- [iendtime](iendtime.html)
  Returns the end time of the specified input.
- [ihasplane](ihasplane.html)
  Returns 1 if the specified input has a plane named planename.
- [inumplanes](inumplanes.html)
  Returns the number of planes in the given input.
- [iplaneindex](iplaneindex.html)
  Returns the index of the plane named 'planename' in the specified input.
- [iplanename](iplanename.html)
  Returns the name of the plane specified by the planeindex of the given input
- [iplanesize](iplanesize.html)
  Returns the number of components in the plane named planename in
  the specified input.
- [irate](irate.html)
  Returns the frame rate of the specified input.
- [istart](istart.html)
  Returns the starting frame of the specified input.
- [istarttime](istarttime.html)
  Returns the start time of the specified input.
- [ixres](ixres.html)
  Returns the X resolution of the specified input.
- [iyres](iyres.html)
  Returns the Y resolution of the specified input.
- [lumname](lumname.html)
  Returns the default name of the luminaence plane (as it appears in the
  compositor preferences).
- [maskname](maskname.html)
  Returns the default name of the mask plane (as it appears in the
  compositor preferences).
- [metadata](metadata.html)
  Returns a metadata value from a composite operator.
- [ninput](ninput.html)
  Reads a component from a pixel and its eight neighbors.
- [normalname](normalname.html)
  Returns the default name of the normal plane (as it appears in the
  compositor preferences).
- [planeindex](planeindex.html)
  Returns the index of the plane specified by the parameter, starting
  at zero.
- [planename](planename.html)
  Returns the name of the plane specified by the index (e.
- [planesize](planesize.html)
  Returns the number of components in the plane (1 for scalar planes
  and up to 4 for vector planes).
- [pointname](pointname.html)
  Returns the default name of the point plane (as it appears in the
  compositor preferences).
- [velocityname](velocityname.html)
  Returns the default name of the velocity plane (as it appears in the
  compositor preferences).

Interpolation

## interp_group

- [ckspline](ckspline.html)
  Samples a Catmull-Rom (Cardinal) spline defined by position/value keys.
- [clamp](clamp.html)
  Returns value clamped between min and max.
- [cspline](cspline.html)
  Samples a Catmull-Rom (Cardinal) spline defined by uniformly spaced keys.
- [efit](efit.html)
  Takes the value in one range and shifts it to the corresponding value in a new range.
- [fit](fit.html)
  Takes the value in one range and shifts it to the corresponding value in a new range.
- [fit01](fit01.html)
  Takes the value in the range (0, 1) and shifts it to the corresponding value in a new range.
- [fit10](fit10.html)
  Takes the value in the range (1, 0) and shifts it to the corresponding value in a new range.
- [fit11](fit11.html)
  Takes the value in the range (-1, 1) and shifts it to the corresponding value in a new range.
- [invlerp](invlerp.html)
  Inverses a linear interpolation between the values.
- [lerp](lerp.html)
  Performs linear interpolation between the values.
- [lkspline](lkspline.html)
  Samples a polyline between the key points.
- [lspline](lspline.html)
  Samples a polyline defined by linearly spaced values.
- [slerp](slerp.html)
  Quaternion blend between q1 and q2 based on the bias.
- [slerpv](slerpv.html)
  Spherical blends between two vectors based on the bias.
- [smooth](smooth.html)
  Computes ease in/out interpolation between values.

light

## light_group

- [ambient](ambient.html)
  Returns the color of ambient light in the scene.
- [atten](atten.html)
  Computes attenuated falloff.
- [fastshadow](fastshadow.html)
  Sends a ray from the position P along the direction specified by the
  direction D.
- [filtershadow](filtershadow.html)
  Sends a ray from the position P along direction D.

Math

## math_group

- [abs](abs.html)
  Returns the absolute value of the argument.
- [acos](acos.html)
  Returns the inverse cosine of the argument.
- [asin](asin.html)
  Returns the inverse sine of the argument.
- [atan](atan.html)
  Returns the inverse tangent of the argument.
- [atan2](atan2.html)
  Returns the inverse tangent of y/x.
- [avg](avg.html)
  Returns the average value of the input(s)
- [cbrt](cbrt.html)
  Returns the cube root of the argument.
- [ceil](ceil.html)
  Returns the smallest integer greater than or equal to the argument.
- [combinelocaltransform](combinelocaltransform.html)
  Combines Local and Parent Transforms with Scale Inheritance.
- [cos](cos.html)
  Returns the cosine of the argument.
- [cosh](cosh.html)
  Returns the hyperbolic cosine of the argument.
- [cross](cross.html)
  Returns the cross product between the two vectors.
- [determinant](determinant.html)
  Computes the determinant of the matrix.
- [diag](diag.html)
  Extracts diagonal entries or constructs a diagonal matrix.
- [diagonalizesymmetric](diagonalizesymmetric.html)
  Diagonalizes Symmetric Matrices.
- [distance_pointline](distance_pointline.html)
  This function returns the closest distance between the point Q and the
  infinite line going through O parallel to vector D.
- [distance_pointray](distance_pointray.html)
  This function returns the closest distance between the point Q and the
  semi-finite ray starting at O and extending in the direction D.
- [distance_pointsegment](distance_pointsegment.html)
  This function returns the closest distance between the point Q and a
  finite line segment between points P0 and P1.
- [dot](dot.html)
  Returns the dot product between the arguments.
- [Du](Du.html)
  Returns the derivative of the given value with respect to U.
- [Dv](Dv.html)
  Returns the derivative of the given value with respect to V.
- [Dw](Dw.html)
  Returns the derivative of the given value with respect to the 3rd axis (for volume rendering).
- [eigenvalues](eigenvalues.html)
  Computes the eigenvalues of a 3×3 matrix.
- [erf](erf.html)
  Gauss error function.
- [erf_inv](erf_inv.html)
  Inverse Gauss error function.
- [erfc](erfc.html)
  Gauss error function’s complement.
- [exp](exp.html)
  Returns the exponential function of the argument.
- [extractlocaltransform](extractlocaltransform.html)
  Extracts Local Transform from a World Transform with Scale Inheritance.
- [floor](floor.html)
  Returns the largest integer less than or equal to the argument.
- [frac](frac.html)
  Returns the fractional component of the floating point number.
- [ident](ident.html)
  Returns an identity matrix.
- [invert](invert.html)
  Inverts a matrix.
- [isfinite](isfinite.html)
  Checks whether a value is a normal finite number.
- [isinf](isinf.html)
  Checks whether a value is a positive or negative infinity.
- [isnan](isnan.html)
  Checks whether a value is not a number.
- [kspline](kspline.html)
  Returns an interpolated value along a curve defined by a basis and key/position pairs.
- [length](length.html)
  Returns the magnitude of a vector.
- [length2](length2.html)
  Returns the squared distance of the vector or vector4.
- [log](log.html)
  Returns the natural logarithm of the argument.
- [log10](log10.html)
  Returns the logarithm (base 10) of the argument.
- [makebasis](makebasis.html)
  Creates an orthonormal basis given a z-axis vector.
- [max](max.html)
- [min](min.html)
- [norm_1](norm_1.html)
  Returns the induced matrix 1-norm.
- [norm_fro](norm_fro.html)
  Returns the Frobenius norm of a matrix.
- [norm_inf](norm_inf.html)
  Returns the induced matrix infinity-norm.
- [norm_max](norm_max.html)
  Returns the matrix max-norm.
- [norm_spectral](norm_spectral.html)
  Returns the matrix spectral norm.
- [normalize](normalize.html)
  Returns a normalized vector.
- [outerproduct](outerproduct.html)
  Returns the outer product between the arguments.
- [pinvert](pinvert.html)
  Computes the pseudo-inverse of a matrix.
- [planesphereintersect](planesphereintersect.html)
  Computes the intersection of a 3D sphere and an infinite 3D plane.
- [pow](pow.html)
  Raises the first argument to the power of the second argument.
- [predicate_incircle](predicate_incircle.html)
  Determines if a point is inside or outside a triangle circumcircle.
- [predicate_insphere](predicate_insphere.html)
  Determines if a point is inside or outside a tetrahedron circumsphere.
- [predicate_orient2d](predicate_orient2d.html)
  Determines the orientation of a point with respect to a line.
- [predicate_orient3d](predicate_orient3d.html)
  Determines the orientation of a point with respect to a plane.
- [premul](premul.html)
  Pre multiply matrices.
- [product](product.html)
  Returns the product of a list of numbers.
- [ptlined](ptlined.html)
  This function returns the closest distance between the point Q and a
  finite line segment between points P0 and P1.
- [qdistance](qdistance.html)
  Finds distance between two quaternions.
- [qinvert](qinvert.html)
  Inverts a quaternion rotation.
- [qmultiply](qmultiply.html)
  Multiplies two quaternions and returns the result.
- [qrotate](qrotate.html)
  Rotates a vector by a quaternion.
- [quaternion](quaternion.html)
  Creates a vector4 representing a quaternion.
- [resample_linear](resample_linear.html)
- [rint](rint.html)
  Rounds the number to the closest whole number.
- [shl](shl.html)
  Bit-shifts an integer left.
- [shr](shr.html)
  Bit-shifts an integer right.
- [shrz](shrz.html)
  Bit-shifts an integer right.
- [sign](sign.html)
  Returns -1, 0, or 1 depending on the sign of the argument.
- [sin](sin.html)
  Returns the sine of the argument.
- [sinh](sinh.html)
  Returns the hyperbolic sine of the argument.
- [slideframe](slideframe.html)
  Finds the normal component of frame slid along a curve.
- [solvecubic](solvecubic.html)
  Solves a cubic function returning the number of real roots.
- [solvepoly](solvepoly.html)
  Finds the real roots of a polynomial.
- [solvequadratic](solvequadratic.html)
  Solves a quadratic function returning the number of real roots.
- [solvetriangleSSS](solvetriangleSSS.html)
  Finds the angles of a triangle from its sides.
- [spline](spline.html)
  Samples a value along a polyline or spline curve.
- [spline_cdf](spline_cdf.html)
  Generate a cumulative distribution function (CDF) by sampling a spline curve.
- [sqrt](sqrt.html)
  Returns the square root of the argument.
- [sum](sum.html)
  Returns the sum of a list of numbers.
- [svddecomp](svddecomp.html)
  Computes the singular value decomposition of a given matrix.
- [tan](tan.html)
  Returns the trigonometric tangent of the argument
- [tanh](tanh.html)
  Returns the hyperbolic tangent of the argument
- [tr](tr.html)
  Returns the trace of the given matrix.
- [transpose](transpose.html)
  Transposes the given matrix.
- [trunc](trunc.html)
  Removes the fractional part of a floating point number.

measure

## measure_group

- [distance](distance.html)
  Returns the distance between two points.
- [distance2](distance2.html)
  Returns the squared distance between the two points.
- [getbbox](getbbox.html)
  Sets two vectors to the minimum and maximum corners of the bounding box for the geometry.
- [getbbox_center](getbbox_center.html)
  Returns the center of the bounding box for the geometry.
- [getbbox_max](getbbox_max.html)
  Returns the maximum of the bounding box for the geometry.
- [getbbox_min](getbbox_min.html)
  Returns the minimum of the bounding box for the geometry.
- [getbbox_size](getbbox_size.html)
  Returns the size of the bounding box for the geometry.
- [getbounds](getbounds.html)
  Returns the bounding box of the geometry specified by the filename.
- [getpointbbox](getpointbbox.html)
  Sets two vectors to the minimum and maximum corners of the bounding box for the geometry.
- [getpointbbox_center](getpointbbox_center.html)
  Returns the center of the bounding box for the geometry.
- [getpointbbox_max](getpointbbox_max.html)
  Returns the maximum of the bounding box for the geometry.
- [getpointbbox_min](getpointbbox_min.html)
  Returns the minimum of the bounding box for the geometry.
- [getpointbbox_size](getpointbbox_size.html)
  Returns the size of the bounding box for the geometry.
- [planepointdistance](planepointdistance.html)
  Computes the distance and closest point of a point to an infinite plane.
- [relbbox](relbbox.html)
  Returns the relative position of the point given with respect to the bounding box of the geometry.
- [relpointbbox](relpointbbox.html)
  Returns the relative position of the point given with respect to the bounding box of the geometry.
- [surfacedist](surfacedist.html)
  Finds the distance of a point to a group of points along the surface of a geometry.
- [uvdist](uvdist.html)
  Finds the distance of a uv coordinate to a geometry in uv space.
- [windingnumber](windingnumber.html)
  Computes the winding number of a mesh around a point. Winding number indicates how many times a geometry wraps around a point. Useful for inside/outside test, the winding number is equal to one inside of the mesh and zero outside.
- [windingnumber2d](windingnumber2d.html)
  Computes the winding number of a curve in XY plane around a point. Winding number indicates how many times a curve wraps around a point. Useful for inside/outside test, the winding number is equal to one inside of the curve and zero outside.
- [xyzdist](xyzdist.html)
  Finds the distance from a point to the closest location on surface geometry.

metaball

## metaball_group

- [metaimport](metaimport.html)
  Once you get a handle to a metaball using metastart and metanext, you
  can query attributes of the metaball with metaimport.
- [metamarch](metamarch.html)
  Takes the ray defined by p0 and p1 and partitions it into zero or
  more sub-intervals where each interval intersects a cluster of metaballs
  from filename.
- [metanext](metanext.html)
  Iterate to the next metaball in the list of metaballs returned by the metastart() function.
- [metastart](metastart.html)
  Open a geometry file and return a handle for the metaballs of
  interest, at the position p.
- [metaweight](metaweight.html)
  Returns the metaweight of the geometry at position p.

Nodes

## nodes_group

- [addvariablename](addvariablename.html)
  Adds a mapping for an attribute to a local variable.
- [ch](ch.html)
  Evaluates a channel (or parameter) and return its value.
- [ch2](ch2.html)
  Evaluates a channel (or parameter) and return its value.
- [ch3](ch3.html)
  Evaluates a channel (or parameter) and return its value.
- [ch4](ch4.html)
  Evaluates a channel (or parameter) and return its value.
- [chdict](chdict.html)
  Evaluates a key-value dictionary parameter and return its value.
- [chexpr](chexpr.html)
  Evaluates a channel with a new segment expression.
- [chexprf](chexprf.html)
  Evaluates a channel with a new segment expression at a given frame.
- [chexprt](chexprt.html)
  Evaluates a channel with a new segment expression at a given time.
- [chf](chf.html)
  Evaluates a channel (or parameter) and return its value.
- [chi](chi.html)
  Evaluates a channel (or parameter) and return its value.
- [chid](chid.html)
  Resolves a channel string (or parameter) and return op_id, parm_index and vector_index.
- [chp](chp.html)
  Evaluates a channel (or parameter) and return its value.
- [chramp](chramp.html)
  Evaluates a ramp parameter and return its value.
- [chrampderiv](chrampderiv.html)
  Evaluates the derivative of a parm parameter with respect to position.
- [chs](chs.html)
  Evaluates a channel (or parameter) and return its value.
- [chsop](chsop.html)
  Evaluates an operator path parameter and return the path to the operator.
- [chsraw](chsraw.html)
  Returns the raw string channel (or parameter).
- [chu](chu.html)
  Evaluates a channel or parameter, and return its value.
- [chv](chv.html)
  Evaluates a channel or parameter, and return its value.
- [cregioncapturetransform](cregioncapturetransform.html)
  Returns the capture transform associated with a Capture Region SOP.
- [cregiondeformtransform](cregiondeformtransform.html)
  Returns the deform transform associated with a Capture Region SOP.
- [cregionoverridetransform](cregionoverridetransform.html)
  Returns the capture or deform transform associated with a Capture Region SOP based on the global capture override flag.
- [isconnected](isconnected.html)
  Returns 1 if input_number is connected, or 0 if the input is not connected.
- [opfullpath](opfullpath.html)
  Returns the full path for the given relative path
- [opid](opid.html)
  Resolves an operator path string and return its op_id.
- [opparentbonetransform](opparentbonetransform.html)
  Returns the parent bone transform associated with an OP.
- [opparenttransform](opparenttransform.html)
  Returns the parent transform associated with an OP.
- [opparmtransform](opparmtransform.html)
  Returns the parm transform associated with an OP.
- [oppreconstrainttransform](oppreconstrainttransform.html)
  Returns the preconstraint transform associated with an OP.
- [oppreparmtransform](oppreparmtransform.html)
  Returns the pre and parm transform associated with an OP.
- [opprerawparmtransform](opprerawparmtransform.html)
  Returns the pre and raw parm transform associated with an OP.
- [oppretransform](oppretransform.html)
  Returns the pretransform associated with an OP.
- [oprawparmtransform](oprawparmtransform.html)
  Returns the raw parm transform associated with an OP.
- [optransform](optransform.html)
  Returns the transform associated with an OP.

Noise and Randomness

## noise_group

- [anoise](anoise.html)
  Generates alligator noise.
- [curlgxnoise](curlgxnoise.html)
  Computes divergence free noise based on simplex noise.
- [curlgxnoise2d](curlgxnoise2d.html)
  Computes divergence free noise in the plane based on simplex noise.
- [curlnoise](curlnoise.html)
  Computes divergence free noise based on Perlin noise.
- [curlnoise2d](curlnoise2d.html)
  Computes 2d divergence free noise based on Perlin noise.
- [curlxnoise](curlxnoise.html)
  Computes divergence free noise based on Simplex noise.
- [curlxnoise2d](curlxnoise2d.html)
  Computes 2d divergence free noise based on simplex noise.
- [cwnoise](cwnoise.html)
  Generates Worley (cellular) noise using a Chebyshev distance metric.
- [flownoise](flownoise.html)
  Generates 1D and 3D Perlin Flow Noise from 3D and 4D data.
- [flowpnoise](flowpnoise.html)
  There are two forms of Perlin-style noise: a non-periodic noise which
  changes randomly throughout N-dimensional space, and a periodic form
  which repeats over a given range of space.
- [gxnoise](gxnoise.html)
  Evaluates a simplex noise field.
- [gxnoised](gxnoised.html)
  Evaluates a simplex noise field and its derivatives.
- [hscript_noise](hscript_noise.html)
  Generates noise matching the output of the Hscript noise() expression function.
- [hscript_rand](hscript_rand.html)
  Produces the exact same results as the Houdini expression function of
  the same name.
- [hscript_snoise](hscript_snoise.html)
- [hscript_sturb](hscript_sturb.html)
- [hscript_turb](hscript_turb.html)
  Generates turbulence matching the output of the HScript turb() expression function.
- [mwnoise](mwnoise.html)
  Generates Worley (cellular) noise using a Manhattan distance metric.
- [mx_cellnoise](mx_cellnoise.html)
  MaterialX compatible cellnoise
- [mx_perlin](mx_perlin.html)
  MaterialX compatible Perlin noise
- [mx_voronoi](mx_voronoi.html)
  MaterialX compatible Voronoi noise
- [mx_worley](mx_worley.html)
  MaterialX compatible Worley noise
- [noise](noise.html)
  There are two forms of Perlin-style noise: a non-periodic noise which
  changes randomly throughout N-dimensional space, and a periodic form
  which repeats over a given range of space.
- [noised](noised.html)
  Derivatives of Perlin Noise.
- [nrandom](nrandom.html)
  Non-deterministic random number generation function.
- [onoise](onoise.html)
  These functions are similar to wnoise and vnoise.
- [pnoise](pnoise.html)
  There are two forms of Perlin-style noise: a non-periodic noise which
  changes randomly throughout N-dimensional space, and a periodic form
  which repeats over a given range of space.
- [pxnoised](pxnoised.html)
  Periodic derivatives of Simplex Noise.
- [rand](rand.html)
  Creates a random number between 0 and 1 from a seed.
- [random](random.html)
  Generate a random number based on the integer position in 1-4D space.
- [random_brj](random_brj.html)
  Generate a uniformly distributed random number.
- [random_fhash](random_fhash.html)
  Hashes floating point numbers to integers.
- [random_ihash](random_ihash.html)
  Hashes integer numbers to integers.
- [random_poisson](random_poisson.html)
  Generates a random Poisson variable given the mean to the distribution and a seed.
- [random_shash](random_shash.html)
  Hashes a string to an integer.
- [random_sobol](random_sobol.html)
  Generate a uniformly distributed random number.
- [snoise](snoise.html)
  These functions are similar to wnoise.
- [vnoise](vnoise.html)
  Generates Voronoi (cellular) noise.
- [wnoise](wnoise.html)
  Generates Worley (cellular) noise.
- [xnoise](pxnoise.html)
  Simplex noise is very close to Perlin noise, except with the samples on a simplex mesh rather than a grid. This results in less grid artifacts. It also uses a higher order bspline to provide better derivatives. This is the periodic simplex noise
- [xnoise](xnoise.html)
  Simplex noise is very close to Perlin noise, except with the samples on a simplex mesh rather than a grid. This results in less grid artifacts. It also uses a higher order bspline to provide better derivatives.
- [xnoised](xnoised.html)
  Derivatives of Simplex Noise.

normals

## normals_group

- [computenormal](computenormal.html)
  In shading contexts, computes a normal. In the SOP contexts, sets how/whether to recompute normals.
- [prim_normal](prim_normal.html)
  Returns the normal of the primitive (prim_number) at parametric location u, v.

Open Color IO

## ocio_group

- [ocio_activedisplays](ocio_activedisplays.html)
  Returns the names of active displays supported in Open Color IO
- [ocio_activeviews](ocio_activeviews.html)
  Returns the names of active views supported in Open Color IO
- [ocio_import](ocio_import.html)
  Imports attributes from OpenColorIO spaces.
- [ocio_parsecolorspace](ocio_parsecolorspace.html)
  Parse the color space from a string
- [ocio_roles](ocio_roles.html)
  Returns the names of roles supported in Open Color IO
- [ocio_spaces](ocio_spaces.html)
  Returns the names of color spaces supported in Open Color IO.
- [ocio_transform](ocio_transform.html)
  Transform colors using Open Color IO
- [ocio_transformview](ocio_transformview.html)
  Transform colors to a view using Open Color IO

particles

## particles_group

- [filamentsample](filamentsample.html)
  Samples the velocity field defined by a set of vortex filaments.

Point Clouds and 3D Images

## ptcloud_group

- [mattrib](mattrib.html)
  Returns the value of the point attribute for the metaballs if
  metaball geometry is specified to i3dgen.
- [mdensity](mdensity.html)
  Returns the density of the metaball field if metaball geometry is
  specified to i3dgen.
- [mspace](mspace.html)
  Transforms the position specified into the local space of the
  metaball.
- [pcclose](pcclose.html)
  This function closes the handle associated with a pcopen
  function.
- [pccone](pccone.html)
  Returns a list of closest points from a file within a specified cone.
- [pccone_radius](pccone_radius.html)
  Returns a list of closest points from a file in a cone, taking into account their radii
- [pcconvex](pcconvex.html)
- [pcexport](pcexport.html)
  Writes data to a point cloud inside a pciterate or a pcunshaded loop.
- [pcfarthest](pcfarthest.html)
  Returns the distance to the farthest point found in the search
  performed by pcopen.
- [pcfilter](pcfilter.html)
  Filters points found by pcopen using a simple reconstruction filter.
- [pcfind](pcfind.html)
  Returns a list of closest points from a file.
- [pcfind_radius](pcfind_radius.html)
  Returns a list of closest points from a file taking into account their radii.
- [pcgenerate](pcgenerate.html)
  Generates a point cloud.
- [pcimport](pcimport.html)
  Imports channel data from a point cloud inside a pciterate or a pcunshaded loop.
- [pcimportbyidx3](pcimportbyidx3.html)
  Imports channel data from a point cloud outside a pciterate or a pcunshaded loop.
- [pcimportbyidx4](pcimportbyidx4.html)
  Imports channel data from a point cloud outside a pciterate or a pcunshaded loop.
- [pcimportbyidxf](pcimportbyidxf.html)
  Imports channel data from a point cloud outside a pciterate or a pcunshaded loop.
- [pcimportbyidxi](pcimportbyidxi.html)
  Imports channel data from a point cloud outside a pciterate or a pcunshaded loop.
- [pcimportbyidxp](pcimportbyidxp.html)
  Imports channel data from a point cloud outside a pciterate or a pcunshaded loop.
- [pcimportbyidxs](pcimportbyidxs.html)
  Imports channel data from a point cloud outside a pciterate or a pcunshaded loop.
- [pcimportbyidxv](pcimportbyidxv.html)
  Imports channel data from a point cloud outside a pciterate or a pcunshaded loop.
- [pciterate](pciterate.html)
  This function can be used to iterate over all the points which were
  found in the pcopen query.
- [pcline](pcline.html)
  Returns a list of closest points to an infinite line from a specified file
- [pcline_radius](pcline_radius.html)
  Returns a list of closest points to an infinite line from a specified file
- [pcnumfound](pcnumfound.html)
  This node returns the number of points found by pcopen.
- [pcopen](pcopen.html)
  Returns a handle to a point cloud file.
- [pcopenlod](pcopenlod.html)
  Returns a handle to a point cloud file.
- [pcsampleleaf](pcsampleleaf.html)
  Changes the current iteration point to a leaf descendant of the current aggregate point.
- [pcsegment](pcsegment.html)
  Returns a list of closest points to a line segment from a specified file
- [pcsegment_radius](pcsegment_radius.html)
  Returns a list of closest points to a line segment from a specified file
- [pcsize](pcsize.html)
- [pcunshaded](pcunshaded.html)
  Iterate over all of the points of a read-write channel which haven’t
  had any data written to the channel yet.
- [pcwrite](pcwrite.html)
  Writes data to a point cloud file.
- [pgfind](pgfind.html)
  Returns a list of closest points from a file.
- [photonmap](photonmap.html)
  Samples a color from a photon map.
- [texture3d](texture3d.html)
  Returns the value of the 3d image at the position specified by P.
- [texture3dBox](texture3dBox.html)
  This function queries the 3D texture map specified and returns the
  bounding box information of the file.

Sampling

## sampling_group

- [create_cdf](create_cdf.html)
  Creates a cumulative distribution function (CDF) from an array of probability density function (PDF) values.
- [create_pdf](create_pdf.html)
  Creates a probability density function from an array of input values.
- [limit_sample_space](limit_sample_space.html)
  Limits a unit value in a way that maintains uniformity and in-range consistency.
- [newsampler](newsampler.html)
  Initializes a sampling sequence for the nextsample function.
- [nextsample](nextsample.html)
- [sample_cauchy](sample_cauchy.html)
  Samples the Cauchy (Lorentz) distribution.
- [sample_cdf](sample_cdf.html)
  Samples a cumulative distribution function (CDF).
- [sample_circle_arc](sample_circle_arc.html)
  Generates a uniform unit vector2, within maxangle of center, given a uniform number between 0 and 1.
- [sample_circle_edge_uniform](sample_circle_edge_uniform.html)
  Generates a uniform unit vector2, given a uniform number between 0 and 1.
- [sample_circle_ring_uniform](sample_circle_ring_uniform.html)
  Generates a uniform vector2 with alpha \< length \< 1, where 0 \< alpha \< 1, given a vector2 of uniform numbers between 0 and 1.
- [sample_circle_slice](sample_circle_slice.html)
  Generates a uniform vector2 with length \< 1, within maxangle of center, given a vector2 of uniform numbers between 0 and 1.
- [sample_circle_uniform](sample_circle_uniform.html)
  Generates a uniform vector2 with length \< 1, given a vector2 of uniform numbers between 0 and 1.
- [sample_direction_cone](sample_direction_cone.html)
  Generates a uniform unit vector, within maxangle of center, given a vector2 of uniform numbers between 0 and 1.
- [sample_direction_uniform](sample_direction_uniform.html)
  Generates a uniform unit vector, given a vector2 of uniform numbers between 0 and 1.
- [sample_discrete](sample_discrete.html)
  Returns an integer, either uniform or weighted, given a uniform number between 0 and 1.
- [sample_exponential](sample_exponential.html)
  Samples the exponential distribution.
- [sample_geometry](sample_geometry.html)
  Samples geometry in the scene and returns information from the shaders of surfaces that were sampled.
- [sample_hemisphere](sample_hemisphere.html)
  Generates a unit vector, optionally biased, within a hemisphere, given a vector2 of uniform numbers between 0 and 1.
- [sample_hypersphere_cone](sample_hypersphere_cone.html)
  Generates a uniform vector4 with length \< 1, within maxangle of center, given a vector4 of uniform numbers between 0 and 1.
- [sample_hypersphere_uniform](sample_hypersphere_uniform.html)
  Generates a uniform vector4 with length \< 1, given a vector4 of uniform numbers between 0 and 1.
- [sample_light](sample_light.html)
  Samples a 3D position on a light source and runs the light shader at that point.
- [sample_lognormal](sample_lognormal.html)
  Samples the log-normal distribution based on parameters of the underlying normal distribution.
- [sample_lognormal_by_median](sample_lognormal_by_median.html)
  Samples the log-normal distribution based on median and standard deviation.
- [sample_normal](sample_normal.html)
  Samples the normal (Gaussian) distribution.
- [sample_orientation_cone](sample_orientation_cone.html)
  Generates a uniform unit vector4, within maxangle of center, given a vector of uniform numbers between 0 and 1.
- [sample_orientation_uniform](sample_orientation_uniform.html)
  Generates a uniform unit vector4, given a vector of uniform numbers between 0 and 1.
- [sample_photon](sample_photon.html)
  Samples a 3D position on a light source and runs the light shader at that point.
- [sample_sphere_cone](sample_sphere_cone.html)
  Generates a uniform vector with length \< 1, within maxangle of center, given a vector of uniform numbers between 0 and 1.
- [sample_sphere_shell_uniform](sample_sphere_shell_uniform.html)
  Generates a uniform vector with alpha \< length \< 1, where 0 \< alpha \< 1, given a vector of uniform numbers between 0 and 1.
- [sample_sphere_uniform](sample_sphere_uniform.html)
  Generates a uniform vector with length \< 1, given a vector of uniform numbers between 0 and 1.
- [sampledisk](sampledisk.html)
  Warps uniform random samples to a disk.
- [variance](variance.html)
  Computes the mean value and variance for a value.

Sensor Input

## sensor_group

- [sensor_panorama_create](sensor_panorama_create.html)
  Sensor function to render GL scene and query the result.
- [sensor_panorama_getcolor](sensor_panorama_getcolor.html)
  Sensor function query a rendered GL scene.
- [sensor_panorama_getcone](sensor_panorama_getcone.html)
  Sensor function to query average values from rendered GL scene.
- [sensor_panorama_getdepth](sensor_panorama_getdepth.html)
  Sensor function query a rendered GL scene.
- [sensor_save](sensor_save.html)
  Sensor function to save a rendered GL scene.

Shading and Rendering

## shading_group

- [area](area.html)
  Returns the area of the micropolygon containing a variable such as P.
- [blinnBRDF](blinnBRDF.html)
- [bouncelabel](bouncelabel.html)
- [bouncemask](bouncemask.html)
- [diffuseBRDF](diffuseBRDF.html)
- [filterstep](filterstep.html)
  Returns the anti-aliased weight of the step function.
- [fresnel](fresnel.html)
  Computes the fresnel reflection/refraction contributions given an
  incoming vector, surface normal (both normalized), and an index of
  refraction (eta).
- [frontface](frontface.html)
  If dot(I, Nref) is less than zero, N will be negated.
- [gather](gather.html)
  Sends rays into the scene and returns information from the shaders of
  surfaces hit by the rays.
- [getblurP](getblurP.html)
  Returns the blurred point position (P) vector at a fractional time within the motion blur exposure.
- [getcomponents](getcomponents.html)
- [getderiv](getderiv.html)
  Evaluates surface derivatives of an attribute.
- [getfogname](getfogname.html)
  Returns the name of the current object whose shader is being run.
- [getglobalraylevel](getglobalraylevel.html)
  Returns the depth of the ray tree for computing global
  illumination.
- [getgroupid](getgroupid.html)
  Returns group id containing current primitive.
- [getlight](getlight.html)
  Returns a light struct for the specified light identifier.
- [getlightid](getlightid.html)
  Returns the light id for a named light (or -1 for an invalid name).
- [getlightname](getlightname.html)
  Returns the name of the current light when called from within an illuminance loop, or converts an integer light ID into the light’s name.
- [getlights](getlights.html)
  Returns an array of light identifiers for the currently shaded surface.
- [getlightscope](getlightscope.html)
  Returns a selection of lights that illuminate a given material.
- [getlocalcurvature](getlocalcurvature.html)
  Evaluates local curvature of primitive grid, using the same curvature evaluation method as Measure SOPs.
- [getmaterial](getmaterial.html)
  Returns a material struct for the current surface.
- [getmaterialid](getmaterialid.html)
  Returns material id of shaded primitive.
- [getobjectid](getobjectid.html)
  Returns the object id for the current shading context.
- [getobjectname](getobjectname.html)
  Returns the name of the current object whose shader is being run.
- [getphotonlight](getphotonlight.html)
  Returns the integer ID of the light being used for photon shading.
- [getprimid](getprimid.html)
  Returns the number of the current primitive.
- [getptextureid](getptextureid.html)
  Returns the ptexture face id for the current primitive.
- [getraylevel](getraylevel.html)
  Returns the depth of the ray tree for the current shading.
- [getrayweight](getrayweight.html)
  Returns an approximation to the contribution of the ray to the final
  pixel color.
- [getsamplestore](getsamplestore.html)
  Looks up sample data in a channel, referenced by a point.
- [getscope](getscope.html)
  Returns a selection of objects visible to rays for a given material.
- [getsmoothP](getsmoothP.html)
  Returns modified surface position based on a smoothing function.
- [getuvtangents](getuvtangents.html)
  Evaluates UV tangents at a point on an arbitrary object.
- [gradient](gradient.html)
  Returns the gradient of a field.
- [haslight](haslight.html)
  Returns whether a light illuminates the given material.
- [illuminance](illuminance.html)
  Loops through all light sources in the scene, calling the light shader for each light source to set the Cl and L global variables.
- [integratehoseksky](integratehoseksky.html)
  Computes irradiance from the given Hosek Sky on a horizontal surface
- [interpolate](interpolate.html)
  Interpolates a value across the currently shaded micropolygon.
- [intersect_lights](intersect_lights.html)
  Finds the nearest intersection of a ray with any of a list of (area) lights and runs the light shader at the intersection point.
- [irradiance](irradiance.html)
  Computes irradiance (global illumination) at the point P with the normal N.
- [isfogray](isfogray.html)
  Returns 1 if the shader is being called to evaluate illumination for
  fog objects, or 0 if the light or shadow shader is being called to
  evaluate surface illumination.
- [islpeactive](islpeactive.html)
  Returns 1 if Light Path Expressions are enabled. 0 Otherwise.
- [israytracing](israytracing.html)
  Indicates whether a shader is being executed for ray tracing.
- [isshadingRHS](isshadingRHS.html)
  Detects the orientation of default shading space.
- [isshadowray](isshadowray.html)
  Returns 1 if the shader is being called to evaluate opacity for
  shadow rays, or 0 if the shader is being called to evaluate for surface
  color.
- [isuvrendering](isuvrendering.html)
  Indicates whether the shader is being evaluated while doing UV rendering (e.g. texture unwrapping)
- [lightbounces](lightbounces.html)
  Returns the bounce mask for a light struct.
- [lightid](lightid.html)
  Returns the light id for a light struct.
- [lightstate](lightstate.html)
  Queries the renderer for a named property.
- [limport](limport.html)
  Imports a variable from the light shader for the surface.
- [matchvex_blinn](matchvex_blinn.html)
  Returns a BSDF that matches the output of the traditional VEX blinn function.
- [matchvex_specular](matchvex_specular.html)
  Returns a BSDF that matches the output of the traditional VEX specular function.
- [nbouncetypes](nbouncetypes.html)
- [objectstate](objectstate.html)
  Queries the renderer for a named property.
- [occlusion](occlusion.html)
  Computes ambient occlusion.
- [pathtrace](pathtrace.html)
  Computes global illumination using PBR for secondary bounces.
- [phongBRDF](phongBRDF.html)
- [rayhittest](rayhittest.html)
  Sends a ray from the position P along the direction D.
- [rayimport](rayimport.html)
  Imports a value sent by a shader in a gather loop.
- [reflect](reflect.html)
  Returns the vector representing the reflection of the direction
  against the normal.
- [reflectlight](reflectlight.html)
  Computes the amount of reflected light which hits the surface.
- [refract](refract.html)
  Returns the refraction ray given an incoming direction, the
  normalized normal and an index of refraction.
- [refractlight](refractlight.html)
  Computes the illumination of surfaces refracted by the current
  surface.
- [renderstate](renderstate.html)
  Queries the renderer for a named property.
- [resolvemissedray](resolvemissedray.html)
  Returns the background color for rays that exit the scene.
- [scatter](scatter.html)
  Evaluates a scattering event through the domain of a geometric object.
- [setcurrentlight](setcurrentlight.html)
  Sets the current light
- [setsamplestore](setsamplestore.html)
  Stores sample data in a channel, referenced by a point.
- [shadow](shadow.html)
  Calls shadow shaders for the current light source.
- [shadow_light](shadow_light.html)
  Executes the shadow shader for a given light and returns the amount of shadowing as a multiplier of the shaded color.
- [shimport](shimport.html)
  Imports a variable from the shadow shader for the surface.
- [simport](simport.html)
  Imports a variable sent by a surface shader in an illuminance loop.
- [specularBRDF](specularBRDF.html)
  Returns the computed BRDFs for the different lighting models used in VEX shading.
- [storelightexport](storelightexport.html)
  Stores exported data for a light.
- [switch](switch.html)
  Use a different bsdf for direct or indirect lighting.
- [trace](trace.html)
  Sends a ray from P along the normalized vector D.
- [translucent](translucent.html)
  Returns a Lambertian translucence BSDF.
- [uvunwrap](uvunwrap.html)
  Computes the position and normal at given (u, v) coordinates, for use in a lens shader.
- [wireblinn](wireblinn.html)
- [wirediffuse](wirediffuse.html)
- [writepixel](writepixel.html)
  Writes color information to a pixel in the output image

Strings

## string_group

- [abspath](abspath.html)
  Returns the full path of a file.
- [chr](chr.html)
  Converts an unicode codepoint to a UTF8 string.
- [concat](concat.html)
  Concatenate all the strings specified to form a single string.
- [decode](decode.html)
  Decodes a variable name that was previously encoded.
- [decodeattrib](decodeattrib.html)
  Decodes a geometry attribute name that was previously encoded.
- [decodeparm](decodeparm.html)
  Decodes a node parameter name that was previously encoded.
- [decodeutf8](decodeutf8.html)
  Decodes a UTF8 string into a series of codepoints.
- [encode](encode.html)
  Encodes any string into a valid variable name.
- [encodeattrib](encodeattrib.html)
  Encodes any string into a valid geometry attribute name.
- [encodeparm](encodeparm.html)
  Encodes any string into a valid node parameter name.
- [encodeutf8](encodeutf8.html)
  Encodes a UTF8 string from a series of codepoints.
- [endswith](endswith.html)
  Indicates the string ends with the specified string.
- [find](find.html)
  Finds an item in an array or string.
- [isalpha](isalpha.html)
  Returns 1 if all the characters in the string are alphabetic
- [isdigit](isdigit.html)
  Returns 1 if all the characters in the string are numeric
- [itoa](itoa.html)
  Converts an integer to a string.
- [join](join.html)
  Concatenate all the strings of an array inserting a common spacer.
- [lstrip](lstrip.html)
  Strips leading whitespace from a string.
- [makevalidvarname](makevalidvarname.html)
  Forces a string to conform to the rules for variable names.
- [match](match.html)
  This function returns 1 if the subject matches the pattern specified,
  or 0 if the subject doesn’t match.
- [opdigits](opdigits.html)
  Returns the integer value of the last sequence of digits of a string
- [ord](ord.html)
  Converts an UTF8 string into a codepoint.
- [pluralize](pluralize.html)
  Converts an English noun to its plural.
- [re_find](re_find.html)
  Matches a regular expression in a string
- [re_findall](re_findall.html)
  Finds all instances of the given regular expression in the string
- [re_match](re_match.html)
  Returns 1 if the entire input string matches the expression
- [re_replace](re_replace.html)
  Replaces instances of regex_find with regex_replace
- [re_split](re_split.html)
  Splits the given string based on regex match.
- [relativepath](relativepath.html)
  Computes the relative path for two full paths.
- [relpath](relpath.html)
  Returns the relative path to a file.
- [replace](replace.html)
  Replaces occurrences of a substring.
- [replace_match](replace_match.html)
  Replaces the matched string pattern with another pattern.
- [rstrip](rstrip.html)
  Strips trailing whitespace from a string.
- [split](split.html)
  Splits a string into tokens.
- [splitpath](splitpath.html)
  Splits a file path into the directory and name parts.
- [sprintf](sprintf.html)
  Formats a string like printf but returns the result as a string
  instead of printing it.
- [startswith](startswith.html)
  Returns 1 if the string starts with the specified string.
- [strip](strip.html)
  Strips leading and trailing whitespace from a string.
- [strlen](strlen.html)
  Returns the length of the string.
- [titlecase](titlecase.html)
  Returns a string that is the titlecase version of the input string.
- [tolower](tolower.html)
  Converts all characters in string to lower case
- [toupper](toupper.html)
  Converts all characters in string to upper case

Subdivision Surfaces

## subd_group

- [osd_facecount](osd_facecount.html)
- [osd_firstpatch](osd_firstpatch.html)
- [osd_limitsurface](osd_limitsurface.html)
  Evaluates a point attribute at the subdivision limit surface using Open Subdiv.
- [osd_limitsurfacevertex](osd_limitsurfacevertex.html)
  Evaluates a vertex attribute at the subdivision limit surface using Open Subdiv.
- [osd_lookupface](osd_lookupface.html)
  Outputs the Houdini face and UV coordinates corresponding to the given coordinates on an OSD patch.
- [osd_lookuppatch](osd_lookuppatch.html)
  Outputs the OSD patch and UV coordinates corresponding to the given coordinates on a Houdini polygon face.
- [osd_patchcount](osd_patchcount.html)
- [osd_patches](osd_patches.html)
  Returns a list of patch IDs for the patches in a subdivision hull.

Tetrahedrons

## tet_group

- [tet_adjacent](tet_adjacent.html)
  Returns primitive number of an adjacent tetrahedron.
- [tet_faceindex](tet_faceindex.html)
  Returns vertex indices of each face of a tetrahedron.

Texturing

## texture_group

- [colormap](colormap.html)
  Looks up a (filtered) color from a texture file.
- [depthmap](depthmap.html)
  The depthmap functions work on an image which was rendered as a
  z-depth image from mantra.
- [environment](environment.html)
  Returns the color of the environment texture.
- [expand_udim](expand_udim.html)
  Perform UDIM or UVTILE texture filename expansion.
- [has_udim](has_udim.html)
  Test string for UDIM or UVTILE patterns.
- [importance_remap](importance_remap.html)
  Remaps a texture coordinate to another coordinate in the map to optimize sampling of brighter areas.
- [ocean_sample](ocean_sample.html)
  Evaluates an ocean spectrum and samples the result at a given time and location.
- [ptexture](ptexture.html)
  Computes a filtered sample from a ptex texture map. Use texture instead.
- [rawcolormap](rawcolormap.html)
  Looks up an unfiltered color from a texture file.
- [shadowmap](shadowmap.html)
  The shadowmap function will treat the shadow map as if the image were
  rendered from a light source.
- [teximport](teximport.html)
  Imports attributes from texture files.
- [texprintf](texprintf.html)
  Similar to sprintf, but does expansion of UDIM or UVTILE texture filename expansion.
- [texture](texture.html)
  Computes a filtered sample of the texture map specified.

Transforms and Space

## transform_group

- [dihedral](dihedral.html)
  Computes the rotation matrix or quaternion which rotates the vector a onto the vector b.
- [fromNDC](fromNDC.html)
  Transforms a position from normal device coordinates to the
  coordinates in the appropriate space.
- [getpackedtransform](getpackedtransform.html)
  Gets the transform of a packed primitive.
- [getspace](getspace.html)
  Returns a transform from one space to another.
- [instance](instance.html)
  Creates an instance transform matrix.
- [lookat](lookat.html)
  Computes a rotation matrix or angles to orient the negative z-axis along the vector (to-from) under the transformation.
- [maketransform](maketransform.html)
  Builds a 3×3 or 4×4 transform matrix.
- [ndcdepth](ndcdepth.html)
  Returns the camera space z-depth of the NDC z-depth value.
- [ntransform](ntransform.html)
  Transforms a normal vector.
- [orthographic](orthographic.html)
  Create an orthographic projection matrix.
- [ow_nspace](ow_nspace.html)
  Transforms a normal vector from Object to World space.
- [ow_space](ow_space.html)
  Transforms a position value from Object to World space.
- [ow_vspace](ow_vspace.html)
  Transforms a direction vector from Object to World space.
- [packedtransform](packedtransform.html)
  Transforms a packed primitive.
- [perspective](perspective.html)
  Create a perspective projection matrix.
- [polardecomp](polardecomp.html)
  Computes the polar decomposition of a matrix.
- [prerotate](prerotate.html)
  Applies a pre rotation to the given matrix.
- [prescale](prescale.html)
  Prescales the given matrix in three directions simultaneously (X, Y, Z -
  given by the components of the scale_vector).
- [pretranslate](pretranslate.html)
  Pretranslates a matrix by a vector.
- [ptransform](ptransform.html)
  Transforms a vector from one space to another.
- [rotate](rotate.html)
  Applies a rotation to the given matrix.
- [rotate_x_to](rotate_x_to.html)
  Rotates a vector by a rotation that would bring the x-axis to a given direction.
- [scale](scale.html)
  Scales the given matrix in three directions simultaneously (X, Y, Z -
  given by the components of the scale_vector).
- [setpackedtransform](setpackedtransform.html)
  Sets the transform of a packed primitive.
- [smoothrotation](smoothrotation.html)
  Returns the closest equivalent Euler rotations to a reference rotation.
- [solveconstraint](solveconstraint.html)
  Applies an inverse kinematics algorithm to a skeleton.
- [solvecurve](solvecurve.html)
  Applies a curve inverse kinematics algorithm to a skeleton.
- [solvefbik](solvefbik.html)
  Applies a full-body inverse kinematics algorithm to a skeleton.
- [solveik](solveik.html)
  Applies an inverse kinematics algorithm to a skeleton.
- [solvephysfbik](solvephysfbik.html)
  Applies a full-body inverse kinematics algorithm to a skeleton, with optional control over the center of mass.
- [toNDC](toNDC.html)
  Transforms a position into normal device coordinates.
- [translate](translate.html)
  Translates a matrix by a vector.
- [tw_nspace](tw_nspace.html)
  Transforms a normal vector from Texture to World space.
- [tw_space](tw_space.html)
  Transforms a position value from Texture to World space.
- [tw_vspace](tw_vspace.html)
  Transforms a direction vector from Texture to World space.
- [vtransform](vtransform.html)
  Transforms a directional vector.
- [wo_nspace](wo_nspace.html)
  Transforms a normal vector from World to Object space.
- [wo_space](wo_space.html)
  Transforms a position value from World to Object space.
- [wo_vspace](wo_vspace.html)
  Transforms a direction vector from World to Object space.
- [wt_nspace](wt_nspace.html)
  Transforms a normal vector from World to Texture space.
- [wt_space](wt_space.html)
  Transforms a position value from World to Texture space.
- [wt_vspace](wt_vspace.html)
  Transforms a direction vector from World to Texture space.

usd

## usd_group

- [usd_addattrib](usd_addattrib.html)
  Creates an attribute of a given type on a primitive.
- [usd_addcollectionexclude](usd_addcollectionexclude.html)
  Excludes an object from the collection
- [usd_addcollectioninclude](usd_addcollectioninclude.html)
  Includes an object in the collection
- [usd_addinversetotransformorder](usd_addinversetotransformorder.html)
  Appends an inversed transform operation to the primitive’s transform order
- [usd_addorient](usd_addorient.html)
  Applies a quaternion orientation to the primitive
- [usd_addprim](usd_addprim.html)
  Creates a primitive of a given type.
- [usd_addprimvar](usd_addprimvar.html)
  Creates a primvar of a given type on a primitive.
- [usd_addrelationshiptarget](usd_addrelationshiptarget.html)
  Adds a target to the primitive’s relationship
- [usd_addrotate](usd_addrotate.html)
  Applies a rotation to the primitive
- [usd_addscale](usd_addscale.html)
  Applies a scale to the primitive
- [usd_addschemaattrib](usd_addschemaattrib.html)
  Creates an attribute of a given type on a primitive, and sets the custom metadata flag to False.
- [usd_addtotransformorder](usd_addtotransformorder.html)
  Appends a transform operation to the primitive’s transform order
- [usd_addtransform](usd_addtransform.html)
  Applies a transformation to the primitive
- [usd_addtranslate](usd_addtranslate.html)
  Applies a translation to the primitive
- [usd_applyapi](usd_applyapi.html)
  Apply an API schema to a primitive.
- [usd_attrib](usd_attrib.html)
  Reads the value of an attribute from the USD primitive.
- [usd_attribelement](usd_attribelement.html)
  Reads the value of an element from an array attribute.
- [usd_attriblen](usd_attriblen.html)
  Returns the length of the array attribute.
- [usd_attribnames](usd_attribnames.html)
  Returns the names of the attributes available on the primitive.
- [usd_attribsize](usd_attribsize.html)
  Returns the tuple size of the attribute.
- [usd_attribtimesamples](usd_attribtimesamples.html)
  Returns the time codes at which the attribute values are authored.
- [usd_attribtypename](usd_attribtypename.html)
  Returns the name of the attribute type.
- [usd_blockattrib](usd_blockattrib.html)
  Blocks the attribute.
- [usd_blockprimvar](usd_blockprimvar.html)
  Blocks the primvar.
- [usd_blockprimvarindices](usd_blockprimvarindices.html)
  Blocks the primvar.
- [usd_blockrelationship](usd_blockrelationship.html)
  Blocks the primitive’s relationship
- [usd_boundmaterialpath](usd_boundmaterialpath.html)
  Returns the material path bound to a given primitive.
- [usd_childnames](usd_childnames.html)
  Returns the names of the primitive’s children.
- [usd_clearmetadata](usd_clearmetadata.html)
  Clears the value of the metadata.
- [usd_cleartransformorder](usd_cleartransformorder.html)
  Clears the primitive’s transform order
- [usd_collectioncomputedpaths](usd_collectioncomputedpaths.html)
  Obtains the list of all objects that belong to the collection
- [usd_collectioncontains](usd_collectioncontains.html)
  Checks if an object path belongs to the collection
- [usd_collectionexcludes](usd_collectionexcludes.html)
  Obtains the object paths that are in the collection’s exclude list
- [usd_collectionexpansionrule](usd_collectionexpansionrule.html)
  Obtains the collection’s expansion rule
- [usd_collectionincludes](usd_collectionincludes.html)
  Obtains the object paths that are in the collection’s include list
- [usd_drawmode](usd_drawmode.html)
  Returns the primitive’s draw mode.
- [usd_findtransformname](usd_findtransformname.html)
  Retrurns primitive’s transform operation full name for given the transform operation suffix
- [usd_flattenediprimvar](usd_flattenediprimvar.html)
  Reads the value of a flattened primvar directly from the USD primitive or from USD primitive’s ancestor.
- [usd_flattenediprimvarelement](usd_flattenediprimvarelement.html)
  Reads an element value of a flattened array primvar directly from the USD primitive or from its ancestor.
- [usd_flattenedprimvar](usd_flattenedprimvar.html)
  Reads the value of an flattened primvar directly from the USD primitive.
- [usd_flattenedprimvarelement](usd_flattenedprimvarelement.html)
  Reads an element value of a flattened array primvar directly from a USD primitive.
- [usd_getbbox](usd_getbbox.html)
  Sets two vectors to the minimum and maximum corners of the bounding box for the primitive.
- [usd_getbbox_center](usd_getbbox_center.html)
  Returns the center of the bounding box for the primitive.
- [usd_getbbox_max](usd_getbbox_max.html)
  Returns the maximum of the bounding box for the primitive.
- [usd_getbbox_min](usd_getbbox_min.html)
  Returns the minimum of the bounding box for the primitive.
- [usd_getbbox_size](usd_getbbox_size.html)
  Returns the size of the bounding box for the primitive.
- [usd_getbounds](usd_getbounds.html)
  Obtains the primitive’s bounds
- [usd_getpointinstancebounds](usd_getpointinstancebounds.html)
  Obtains the primitive’s bounds
- [usd_hasapi](usd_hasapi.html)
  Checks if the primitive adheres to the given API.
- [usd_haspayload](usd_haspayload.html)
  Checks if the primitive adheres to the given API.
- [usd_iprimvar](usd_iprimvar.html)
  Reads the value of a primvar directly from the USD primitive or from USD primitive’s ancestor.
- [usd_iprimvarelement](usd_iprimvarelement.html)
  Reads the value of an element from the array primvar directly from the USD primitive or from USD primitive’s ancestor.
- [usd_iprimvarelementsize](usd_iprimvarelementsize.html)
  Returns the element size of the primvar directly from the USD primitive or from USD primitive’s ancestor.
- [usd_iprimvarindices](usd_iprimvarindices.html)
  Returns the index array of an indexed primvar directly on the USD primitive or on USD primitive’s ancestor.
- [usd_iprimvarinterpolation](usd_iprimvarinterpolation.html)
  Returns the element size of the primvar directly on the USD primitive or on USD primitive’s ancestor.
- [usd_iprimvarlen](usd_iprimvarlen.html)
  Returns the length of the array primvar directly on the USD primitive or on USD primitive’s ancestor.
- [usd_iprimvarnames](usd_iprimvarnames.html)
  Returns the names of the primvars available directly on the given USD primitive or on USD primitive’s ancestor.
- [usd_iprimvarsize](usd_iprimvarsize.html)
  Returns the tuple size of the primvar directly on the USD primitive or on USD primitive’s ancestor.
- [usd_iprimvartimesamples](usd_iprimvartimesamples.html)
  Returns the time codes at which the primvar values are authored directly on the given primitive or on its ancestor.
- [usd_iprimvartypename](usd_iprimvartypename.html)
  Returns the name of the primvar type found on the given primitive or its ancestor.
- [usd_isabstract](usd_isabstract.html)
  Checks if the primitive is abstract.
- [usd_isactive](usd_isactive.html)
  Checks if the primitive is active.
- [usd_isarray](usd_isarray.html)
  Checks if the attribute is an array.
- [usd_isarrayiprimvar](usd_isarrayiprimvar.html)
  Checks if there is an array primvar directly on the USD primitive or on USD primitive’s ancestor.
- [usd_isarraymetadata](usd_isarraymetadata.html)
  Checks if the given metadata is an array.
- [usd_isarrayprimvar](usd_isarrayprimvar.html)
  Checks if there is an array primvar directly on the USD primitive.
- [usd_isattrib](usd_isattrib.html)
  Checks if the primitive has an attribute by the given name.
- [usd_iscollection](usd_iscollection.html)
  Checks if the collection exists.
- [usd_iscollectionpath](usd_iscollectionpath.html)
  Checks if the path is a valid collection path.
- [usd_isindexediprimvar](usd_isindexediprimvar.html)
  Checks if there is an indexed primvar directly on the USD primitive or on USD primitive’s ancestor.
- [usd_isindexedprimvar](usd_isindexedprimvar.html)
  Checks if there is an indexed primvar directly on the USD primitive.
- [usd_isinstance](usd_isinstance.html)
  Checks if the primitive is an instance.
- [usd_isiprimvar](usd_isiprimvar.html)
  Checks if the primitive or its ancestor has a primvar of the given name.
- [usd_iskind](usd_iskind.html)
  Checks if the primitive is of a given kind.
- [usd_ismetadata](usd_ismetadata.html)
  Checks if the primitive has metadata by the given name.
- [usd_ismodel](usd_ismodel.html)
  Checks if the primitive is a model.
- [usd_isprim](usd_isprim.html)
  Checks if the path refers to a valid primitive.
- [usd_isprimvar](usd_isprimvar.html)
  Checks if the primitive has a primvar of the given name.
- [usd_isrelationship](usd_isrelationship.html)
  Checks if the primitive has a relationship by the given name.
- [usd_isstage](usd_isstage.html)
  Checks if the stage is valid.
- [usd_istransformreset](usd_istransformreset.html)
  Checks if the primitive transform is reset
- [usd_istype](usd_istype.html)
  Checks if the primitive is of a given type.
- [usd_isvisible](usd_isvisible.html)
  Checks if the primitive is visible.
- [usd_kind](usd_kind.html)
  Returns the primitive’s kind.
- [usd_localtransform](usd_localtransform.html)
  Obtains the primitive’s local transform
- [usd_makeattribpath](usd_makeattribpath.html)
  Constructs an attribute path from a primitive path and an attribute name.
- [usd_makecollectionpath](usd_makecollectionpath.html)
  Constructs a collection path from a primitive path and a collection name.
- [usd_makepropertypath](usd_makepropertypath.html)
  Constructs an property path from a primitive path and an property name.
- [usd_makerelationshippath](usd_makerelationshippath.html)
  Constructs an relationship path from a primitive path and a relationship name.
- [usd_makevalidprimname](usd_makevalidprimname.html)
  Forces a string to conform to the rules for naming USD primitives.
- [usd_makevalidprimpath](usd_makevalidprimpath.html)
  Forces a string to conform to the rules for paths to USD primitives.
- [usd_metadata](usd_metadata.html)
  Reads the value of metadata from the USD object.
- [usd_metadataelement](usd_metadataelement.html)
  Reads the value of an element from the array metadata.
- [usd_metadatalen](usd_metadatalen.html)
  Returns the length of the array metadata.
- [usd_metadatanames](usd_metadatanames.html)
  Returns the names of the metadata available on the object.
- [usd_name](usd_name.html)
  Returns the name of the primitive.
- [usd_parentpath](usd_parentpath.html)
  Returns the path of the primitive’s parent.
- [usd_pointinstance_getbbox](usd_pointinstance_getbbox.html)
  Sets two vectors to the minimum and maximum corners of the bounding box for the given instance inside point instancer.
- [usd_pointinstance_getbbox_center](usd_pointinstance_getbbox_center.html)
  Returns the center of the bounding box for the instance inside a point instancer primitive.
- [usd_pointinstance_getbbox_max](usd_pointinstance_getbbox_max.html)
  Returns the maximum position of the bounding box for the instance inside a point instancer primitive.
- [usd_pointinstance_getbbox_min](usd_pointinstance_getbbox_min.html)
  Returns the minimum position of the bounding box for the instance inside a point instancer primitive.
- [usd_pointinstance_getbbox_size](usd_pointinstance_getbbox_size.html)
  Returns the size of the bounding box for the instance inside a point instancer primitive.
- [usd_pointinstance_relbbox](usd_pointinstance_relbbox.html)
  Returns the relative position of the point given with respect to the bounding box of the geometry.
- [usd_pointinstancetransform](usd_pointinstancetransform.html)
  Obtains the transform for the given point instance
- [usd_primvar](usd_primvar.html)
  Reads the value of a primvar directly from the USD primitive.
- [usd_primvarattribname](usd_primvarattribname.html)
  Returns the namespaced attribute name for the given primvar.
- [usd_primvarelement](usd_primvarelement.html)
  Reads the value of an element from the array primvar directly from the USD primitive.
- [usd_primvarelementsize](usd_primvarelementsize.html)
  Returns the element size of the primvar directly from the USD primitive.
- [usd_primvarindices](usd_primvarindices.html)
  Returns the index array of an indexed primvar directly on the USD primitive.
- [usd_primvarinterpolation](usd_primvarinterpolation.html)
  Returns the element size of the primvar directly on the USD primitive.
- [usd_primvarlen](usd_primvarlen.html)
  Returns the length of the array primvar directly on the USD primitive.
- [usd_primvarnames](usd_primvarnames.html)
  Returns the names of the primvars available on the given USD primitive.
- [usd_primvarsize](usd_primvarsize.html)
  Returns the tuple size of the primvar directly on the USD primitive.
- [usd_primvartimesamples](usd_primvartimesamples.html)
  Returns the time codes at which the primvar values are authored directly on
  the given primitive.
- [usd_primvartypename](usd_primvartypename.html)
  Returns the name of the primvar type found on the given primitive.
- [usd_purpose](usd_purpose.html)
  Returns the primitive’s purpose.
- [usd_relationshipforwardedtargets](usd_relationshipforwardedtargets.html)
  Obtains the relationship forwarded targets.
- [usd_relationshipnames](usd_relationshipnames.html)
  Returns the names of the relationships available on the primitive.
- [usd_relationshiptargets](usd_relationshiptargets.html)
  Obtains the relationship targets.
- [usd_relbbox](usd_relbbox.html)
  Returns the relative position of the point given with respect to the bounding box of the geometry.
- [usd_removerelationshiptarget](usd_removerelationshiptarget.html)
  Remove a target from the primitive’s relationship
- [usd_setactive](usd_setactive.html)
  Sets the primitive active state.
- [usd_setattrib](usd_setattrib.html)
  Sets the value of an attribute.
- [usd_setattribelement](usd_setattribelement.html)
  Sets the value of an element in an array attribute.
- [usd_setcollectionexcludes](usd_setcollectionexcludes.html)
  Sets the excludes list on the collection
- [usd_setcollectionexpansionrule](usd_setcollectionexpansionrule.html)
  Sets the expansion rule on the collection
- [usd_setcollectionincludes](usd_setcollectionincludes.html)
  Sets the includes list on the collection
- [usd_setdrawmode](usd_setdrawmode.html)
  Sets the primitive’s draw mode.
- [usd_setkind](usd_setkind.html)
  Sets the primitive’s kind.
- [usd_setmetadata](usd_setmetadata.html)
  Sets the value of an metadata.
- [usd_setmetadataelement](usd_setmetadataelement.html)
  Sets the value of an element in an array metadata.
- [usd_setprimvar](usd_setprimvar.html)
  Sets the value of a primvar.
- [usd_setprimvarelement](usd_setprimvarelement.html)
  Sets the value of an element in an array primvar.
- [usd_setprimvarelementsize](usd_setprimvarelementsize.html)
  Sets the element size of a primvar.
- [usd_setprimvarindices](usd_setprimvarindices.html)
  Sets the indices for the given primvar.
- [usd_setprimvarinterpolation](usd_setprimvarinterpolation.html)
  Sets the interpolation of a primvar.
- [usd_setpurpose](usd_setpurpose.html)
  Sets the primitive’s purpose.
- [usd_setrelationshiptargets](usd_setrelationshiptargets.html)
  Sets the targets in the primitive’s relationship
- [usd_settransformorder](usd_settransformorder.html)
  Sets the primitive’s transform order
- [usd_settransformreset](usd_settransformreset.html)
  Sets/clears the primitive’s transform reset flag
- [usd_setvariantselection](usd_setvariantselection.html)
  Sets the selected variant in the given variant set.
- [usd_setvisibility](usd_setvisibility.html)
  Configures the primitive to be visible, invisible, or to inherit visibility
  from the parent.
- [usd_setvisible](usd_setvisible.html)
  Makes the primitive visible or invisible.
- [usd_specifier](usd_specifier.html)
  Returns the primitive’s specifier.
- [usd_transformname](usd_transformname.html)
  Constructs a full name of a transform operation
- [usd_transformorder](usd_transformorder.html)
  Obtains the primitive’s transform order
- [usd_transformsuffix](usd_transformsuffix.html)
  Extracts the transform operation suffix from the full name
- [usd_transformtype](usd_transformtype.html)
  Infers the transform operation type from the full name
- [usd_typename](usd_typename.html)
  Returns the name of the primitive’s type.
- [usd_uniquetransformname](usd_uniquetransformname.html)
  Constructs a unique full name of a transform operation
- [usd_variants](usd_variants.html)
  Returns the variants belonging to the given variant set on a primitive.
- [usd_variantselection](usd_variantselection.html)
  Returns the currently selected variant in a given variant set.
- [usd_variantsets](usd_variantsets.html)
  Returns the variant sets available on a primitive.
- [usd_worldtransform](usd_worldtransform.html)
  Obtains the primitive’s world transform

Utility

## utility_group

- [assert_enabled](assert_enabled.html)
  Returns 1 if the VEX assertions are enabled (see HOUDINI_VEX_ASSERT) or 0 if assertions are disabled. Used the implement the assert macro.
- [assign](assign.html)
  An efficient way of extracting the components of a vector or matrix into float variables.
- [error](error.html)
  Reports a custom runtime VEX error.
- [forpoints](forpoints.html)
- [getcomp](getcomp.html)
  Extracts a single component of a vector type, matrix type, or array.
- [isbound](isbound.html)
  Parameters in VEX can be overridden by geometry attributes (if the attributes exist on the surface being rendered).
- [isvarying](isvarying.html)
  Check whether a VEX variable is varying or uniform.
- [opend](opend.html)
  Ends a long operation.
- [opstart](opstart.html)
  Start a long operation.
- [pack_inttosafefloat](pack_inttosafefloat.html)
  Reversibly packs an integer into a finite, non-denormal float.
- [print_once](print_once.html)
  Prints a message only once, even in a loop.
- [printf](printf.html)
  Prints values to the console which started the VEX program.
- [ramp_lookup](ramp_lookup.html)
  Evaluates a Houdini-style ramp at a specific location.
- [ramp_pack](ramp_pack.html)
  Packs a set of arrays into a string-encoded ramp.
- [ramp_unpack](ramp_unpack.html)
  Unpacks a string-encoded ramp into a set of arrays.
- [select](select.html)
  Returns one of two parameters based on a conditional.
- [set](set.html)
  Creates a new value based on its arguments, such as creating a vector from its components.
- [setcomp](setcomp.html)
  Sets a single component of a vector or matrix type, or an item in an array.
- [sleep](sleep.html)
  Yields processing for a certain number of milliseconds.
- [swizzle](swizzle.html)
  Rearranges the components of a vector.
- [unpack_intfromsafefloat](unpack_intfromsafefloat.html)
  Reverses the packing of pack_inttosafefloat to get back the original integer.
- [warning](warning.html)
  Reports a custom runtime VEX warning.

volume

## volume_group

- [volume](volume.html)
  Returns the volume of the microvoxel containing a variable such as P.
- [volumecubicsample](volumecubicsample.html)
  Samples the volume primitive’s value.
- [volumecubicsamplev](volumecubicsamplev.html)
  Samples the volume primitive’s value.
- [volumegradient](volumegradient.html)
  Calculates the volume primitive’s gradient.
- [volumeindex](volumeindex.html)
  Gets the value of a specific voxel.
- [volumeindexactive](volumeindexactive.html)
  Gets the active setting of a specific voxel.
- [volumeindexi](volumeindexi.html)
  Gets the integer value of a specific voxel.
- [volumeindexorigin](volumeindexorigin.html)
  Gets the index of the bottom left of a volume primitive.
- [volumeindexp](volumeindexp.html)
  Gets the vector4 value of a specific voxel.
- [volumeindextopos](volumeindextopos.html)
  Converts a volume voxel index into a position.
- [volumeindexu](volumeindexu.html)
  Gets the vector2 value of a specific voxel.
- [volumeindexv](volumeindexv.html)
  Gets the vector value of a specific voxel.
- [volumepostoindex](volumepostoindex.html)
  Converts a position into a volume voxel index.
- [volumeres](volumeres.html)
  Gets the resolution of a volume primitive.
- [volumesample](volumesample.html)
  Samples the volume primitive’s float value.
- [volumesamplei](volumesamplei.html)
  Samples the volume primitive’s integer value.
- [volumesamplep](volumesamplep.html)
  Samples the volume primitive’s vector4 value.
- [volumesampleu](volumesampleu.html)
  Samples the volume primitive’s vector2 value.
- [volumesamplev](volumesamplev.html)
  Samples the volume primitive’s vector value.
- [volumesmoothsample](volumesmoothsample.html)
  Samples the volume primitive’s value.
- [volumesmoothsamplev](volumesmoothsamplev.html)
  Samples the volume primitive’s value.
- [volumetypeid](volumetypeid.html)
  Gets the typeid of the data of a volume or VDB primitive.
- [volumevoxeldiameter](volumevoxeldiameter.html)
  Computes the approximate diameter of a voxel.

weightarray

## weightarray_group

- [weightarrayblend](weightarrayblend.html)
  Blends an existing name/weight array pair with another array or named item.
- [weightarrayfromname](weightarrayfromname.html)
  Initializes an index array and weight array pair with a single named entry.
- [weightarraynormalize](weightarraynormalize.html)
  Normalizes an array of floats so it sums to 1.0.
- [weightarraythreshold](weightarraythreshold.html)
  Discards any weights below a threshold from an name/weight array pair.
