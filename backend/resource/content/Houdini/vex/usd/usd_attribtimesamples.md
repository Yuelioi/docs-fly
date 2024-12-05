---
title: usd_attribtimesamples
order: 21
---
| Since | 18.0 |
| --- | --- |

`float [] usd_attribtimesamples(<stage>stage, string primpath, string name)`

This function returns an array of time codes at which the attribute values are
authored.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Attribute name.

Returns

The array of time codes at which the attribute values are authored, or an empty array if the attribute does not exist or has no time samples.

Examples

## examples

```vex
// Get the time codes of a foo attribute.
float time_codes[] = usd_attribtimesamples(0, "/geo/cube", "foo");

```

```vex
// Get attribute values at authored time samples.
float[] usd_attribtimesamplevalues(const int input; const string primpath, attribname)
{
    float result[];

    float time_samples[] = usd_attribtimesamples( input, primpath, attribname );
    foreach( float time_code ; time_samples ) 
    {
        float value = usd_attrib( input, primpath, attribname, time_code );
        push( result, value );
    }

    return result;
}

```
