---
title: chr
order: 2
---
`string  chr(int value)`

Returns a string encoding the giving unicode codepoint as a UTF8 value. For
values less than 128, this is a one-byte string of that value. Higher values
will produce multiple byte strings.

If the given code point isn’t a valid codepoint, an empty string is returned.
