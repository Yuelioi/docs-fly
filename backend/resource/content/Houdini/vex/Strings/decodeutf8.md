---
title: decodeutf8
order: 7
---
| Since | 19.0 |
| --- | --- |

`int [] decodeutf8(string str)`

Converts a UTF8 string into a series of code points. VEX treats its strings as UTF8 by default, but this means indexing into strings can have unusual behavior. Generating code points ensures there
is one index for each logical character.
