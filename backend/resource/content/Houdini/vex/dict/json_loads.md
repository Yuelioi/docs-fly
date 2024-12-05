---
title: json_loads
order: 2
---
| Since | 18.5 |
| --- | --- |

`dict  json_loads(string json, int flags)`

Tries to create a dictionary out of a JSON-style string. If a parsing
error occurs, an empty dictionary is created. There are two formats
for JSON recognized. The fully typed format consists of maps for
each element that stores the type and value. This allows proper
round-tripping of VEX dictionaries, but is verbose. The compact
format requires the dictionary to guess the type of JSON elements,
so things like vectors will be created as float arrays.

The VEX dictionaries need to have arrays of uniform type. If an
array has varying type, the elements will be created as strings
of the raw JSON values.

The possible flag values are:

| `0` | Auto-detect if the string is a typed or typeless JSON |
| --- | --- |
| `1` | Require a fully typed JSON |
| `2` | Require a typeless JSON. If a fully-typed JSON is read, an extra layer of dictionaries will be created. |
