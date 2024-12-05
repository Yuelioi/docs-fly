---
title: abspath
order: 1
---
| Since | 18.0 |
| --- | --- |

`string  abspath(string relpath)`

Returns the supplied path converted to an absolute path. Relative paths
are treated as relative to Houdini’s current working directory. If the
supplied path is already absolute, the path is returned unchanged. The file
does not need to exist.
