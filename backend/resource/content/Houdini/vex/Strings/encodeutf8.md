---
title: encodeutf8
order: 11
---
| Since | 19.0 |
| --- | --- |

`string  encodeutf8(int codepoints[])`

Converts a series of code points into a UTF8 string. VEX treats its strings as UTF8 by default, but this means indexing into strings can have unusual behavior. Generating code points ensures there
is one index for each logical character.
