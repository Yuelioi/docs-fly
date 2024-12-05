---
title: sprintf
order: 36
---
`string  sprintf(string format, ...)`

Formats a string like [printf](printf.html "Prints values to the console which started the VEX program.") but returns the result as a
string instead of printing it.

Tip
You can use this function pad a number (such as the frame number in a filename) with zeros (similarly to the [padzero](../../expressions/padzero.html "Returns a string padding a number to a given length with zeros.") expression function) using a format such as `sprintf("texture%04d.exr", @Frame)`.
