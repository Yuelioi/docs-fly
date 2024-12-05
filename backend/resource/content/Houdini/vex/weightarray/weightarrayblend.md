---
title: weightarrayblend
order: 1
---
The new named item is assumed to weight 1.0.

`void  weightarrayblend(string &names[], float &weights[], string newname, float blend)`

Blends the existing arrays with the new item using the specified blend value.

`void  weightarrayblend(string &names[], float &weights[], string newnames[], float newweights[], float blend)`

Blends the existing arrays with another using the specified blend value. The result has the union of items in each input array.
