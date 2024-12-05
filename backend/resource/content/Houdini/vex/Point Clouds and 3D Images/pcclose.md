---
title: pcclose
order: 4
---
`void  pcclose(int &handle)`

This function closes the handle associated with a pcopen function. VEX
will close handles automatically, however, it’s good practice to call
pcclose. When there are pcopen calls made from within a loop, VEX may
consume additional memory if pcclose isn’t called when the handle is no
longer required.
