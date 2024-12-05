---
title: metastart
order: 4
---
`int  metastart(string filename, vector p)`

Open a geometry file and return a “handle” for the metaballs of
interest, at the position p. You can then use
[metanext](metanext.html "Iterate to the next metaball in the list of metaballs returned by the metastart() function.") to move the handle to the next metaball for
evaluation, and [metaimport](metaimport.html "Once you get a handle to a metaball using metastart and metanext, you
can query attributes of the metaball with metaimport.") to query attributes of the
metaball.
