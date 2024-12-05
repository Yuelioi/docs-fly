---
title: Dictionaries
order: 2
---
| On this page | * [Overview](#overview) * [Dictionary Literals](#dictionary-literals) * [Declaring dictionary types](#declaring-dictionary-types) * [Dictionary values](#dictionary-values) * [Accessing and setting dictionary values](#accessing-and-setting-dictionary-values) * [Working with dictionaries](#working-with-dictionaries) |
| --- | --- |

Overview

## overview

VEX includes a dict datatype. Because VEX is strongly typed, it is difficult
to work with parameter sets whose value and type will vary per parameter.
The `dict` type provides a way to work with these.

The `dict` type will be considerably less efficient than expanding the
dictionary as a variable per entry; so should only be used where the
set of variables and types are varying.

Dictionary Literals

## dictionary-literals

The only dictionary literal supported is the empty dictionary: `{}`.

```vex
dict s = {};

```

Declaring dictionary types

## declaring-dictionary-types

To declare a dictionary variable, the general form is
`dict var_name`:

```vex
// My string is a normal dictionary
dict   mydict;

```

To declare a function that returns a dictionary:

```vex
// A function which returns a dictionary
dict rgb_table()
{
...
};    

```

Dictionary values

## dictionary-values

The index, or key, of a dictionary is always a string. The string
can have any value, but for ease of processing it is recommended
to avoid empty strings or strings with `.` in them.

The value of the dictionary can be most of the basic VEX types:
`int`, `float`, `vector2`, `vector`, `vector4`, `matrix2`,
`matrix3`, `matrix`, `string`, `int[]`, `float[]`, `string[]`,
`dict`, `dict[]`.

Accessing and setting dictionary values

## accessing-and-setting-dictionary-values

Use `dict[index]` to look up a value by its key in the dictionary.

The index is a string that is used to lookup the value. If they key
isn’t in the dictionary, the result is a default-initialized value.

VEX requires type to be known at compile time, but the actual stored
type may vary at run time. Thus you will often need an explicit function
cast to select what type to extract. Where possible, the native type
will be converted to the desired type.

```vex
dict dictionary;        // Create empty dictionary
dictionary['key'] = 3;        // Store 3 in the index key
float three = dictionary['key'];        // Extract key
dictionary['newkey'] = dictionary['key'];  // Error: Ambiguous type!
dictionary['newkey'] = int(dictionary['key']);  // Extract key as int and copy

```

The `getcomp` function can be used with a default argument to select the type
and provide a different result if the key is missing. If it is present,
but is the wrong type, the default initialization value is still used.

```vex
dict dictionary;
dictionary['key'] = 3;
int three = getcomp(dictionary, 'key', 52);
int fiftytwo = getcomp(dictionary, 'nonkey', 52);

```

The `typeid` function can be used to identify the dictionary value’s type. This
id can be compared against the `typeid()` of a specific VEX data type, allowing
you to take different code paths depending on the exact type of the value.

```vex
dict dictionary;
dictionary['key'] = set(1, 2, 3);

int value_type = typeid(dictionary, 'key');
if (value_type == typeid(vector()))
{
    // This code path is taken since the key has a 'vector' type.
    vector val = dictionary['key'];
    // ... do something with the vector value.
}
else if (value_type == typeid(vector4()))
{
    vector4 val = dictionary['key'];
    // ... do something with the vector4 value.
}

```

The contents of a key in one dictionary can be copied into another
without exposing the type to VEX through the `insert` command. This
can also be used to join two dictionaries together.

```vex
dict a, b;
a['key'] = 3;
b['oldkey'] = 5;
insert(b, 'newkey', a, 'key'); // Makes b['newkey'] == 3
insert(b, 'oldkey', a, 'nonkey'); // Removes b['oldkey'] as nonkey missing in a.
insert(b, a); // Adds all entries of a to b, overwriting on conflict.

```

Dictionaries can be built using the `set` command.

```vex
dict a;
a = set("key", 3.0, "nextkey", 5.0, "lastkey", "stringvalue");

```

Working with dictionaries

## working-with-dictionaries

The following functions let you query and manipulate dictionaries.

[len](functions/len.html "Returns the length of an array.")

Returns the number of keys in a dictionary.

[keys](functions/keys.html "Returns all the keys in a dictionary.")

Returns a string array of all the keys in a dictionary, sorted
alphabetically.

[isvalidindex](functions/isvalidindex.html "Checks if the index given is valid for the array or string given.")

Returns if an key is in a dictionary or not.

[insert](functions/insert.html "Inserts an item, array, or string into an array or string.")

Copies values from one dictionary to another. Also can merge
all of one dictionary into another.

[removeindex](functions/removeindex.html "Removes an item at the given index from an array.")

Deletes a key from a dictionary.

[set](functions/set.html "Creates a new value based on its arguments, such as creating a vector from its components.")

Builds a dictionary from an list of keys and values.
