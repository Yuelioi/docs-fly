---
title: usd_primvartimesamples
order: 114
---
| Since | 18.0 |
| --- | --- |

`float [] usd_primvartimesamples(<stage>stage, string primpath, string name)`

This function returns an array of time codes at which the values are authored for a primvar found directly on the given primitive.

Show/hide arguments

`<stage>`

When running in the context of a node (such as a wrangle LOP), this argument can be an integer representing the input number (starting at 0) to read the stage from. The integer is equivalent to the string form referencing a particular input, e.g., “opinput:0”.

You can also use this argument to refer to a USD file (e.g., “/path/to/file.usd”), or to another LOP node’s cooked stage using the `op:` as the path prefix (e.g., “op:/stage/lop_node”).

`primpath`

The path to the primitive.

`name`

Primvar name.

Returns

The array of time codes at which the primvar values are authored, or an empty array if the primvar does not exist or has no time samples.

Examples

## examples

```vex
// Get the time codes of a foo primvar.
float time_codes[] = usd_primvartimesamples(0, "/geo/cube", "foo");

```

```vex
// Get primvar values at authored time samples on the given primitive.
float[] usd_primvartimesamplevalues(const int input; const string primpath, primvarname)
{
    float result[];

    float time_samples[] = usd_primvartimesamples( input, primpath, primvarname );
    foreach( float time_code ; time_samples ) 
    {
        float value = usd_primvar( input, primpath, primvarname, time_code );
        push( result, value );
    }

    return result;
}

```
