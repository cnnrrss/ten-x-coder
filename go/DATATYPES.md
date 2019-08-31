1. Integers
Integers are used to store whole numbers. Go has several built-in integer types of varying size for storing signed and unsigned integers -

**Signed Integers**
|Type	|Size	|Range|
|-------|-------|-----|
|int8	|8 bits	|-128 to 127|
|int16	|16 bits	|-215 to 215 -1|
|int32	|32 bits	|-231 to 231 -1|
|int64	|64 bits	|-263 to 263 -1|
|int	|Platform dependent	|Platform dependent|

The size of the generic int type is platform dependent. It is 32 bits wide on a 32-bit system and 64-bits wide on a 64-bit system.

**Unsigned Integers**
|Type	|Size	|Range|
|-------|-------|-----|
|uint8	|8 bits	|0 to 255|
|uint16	|16 bits	|0 to 216 -1|
|uint32	|32 bits	|0 to 232 -1|
|uint64	|64 bits	|0 to 264 -1|
|uint	|Platform dependent	|Platform dependent|
The size of uint type is platform dependent. It is 32 bits wide on a 32-bit system and 64-bits wide on a 64-bit system.

**Note**
_When you are working with integer values, you should always use the int data type unless you have a good reason to use the sized or unsigned integer types._

In Golang, you can declare octal numbers using prefix 0 and hexadecimal numbers using the prefix 0x or 0X. Following is a complete example of integer types -