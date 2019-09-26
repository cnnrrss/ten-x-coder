# Python

[Frameworks](FRAMEWORKS.md)

### Examples

[WebSockets](examples/sockets.py) \
[PageRank Algorithm](examples/page_rank.py)

### Invoking the Interpreter

**Unix**

`python3.7`

**Note**: I've aliased 3.7 to `python`

**Windows** if you have the `py.exe` launcher installed. You can use the `py` command.

To exit the interpreter Unix: `Ctrl+D` Windows: `Ctrl+Z` or `quit()`

**Argument Passing**:

`-m <name-of-module>`

`-c <name-of-command>`

#### Python as a Calculator

```bash
>>> 50 - 5*6 
20

>>> (50 - 5*6) / 4 # use some parenthesis
5.0

>>> 17 / 3  # classic division returns a float
5.666666666666667

>>> 17 // 3  # floor division discards the fractional part
5

>>> 17 % 3  # the % (aka modulus) operator returns the remainder of the division
2

>>> 5 * 3 + 2  # result * divisor + remainder
17

>>> 5 ** 2  # 5 squared # Calculate Powers
25

>>> 2 ** 7  # 2 to the power of 7
128

# (=) sign to assign variables
>>> width, height = 20, 5 * 9 # you can assign multiple vars per line
>>> width * height
900

```

If a variable is not “defined” (assigned a value), trying to use it will give you an error:

```bash
n  # try to access an undefined variable
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
NameError: name 'n' is not defined
```

In interactive mode, the last printed expression is assigned to the variable _. This means that when you are using Python as a desk calculator, it is somewhat easier to continue calculations, for example:

```bash
>>> tax, price = 12.5 / 100, 100.50
12.5625

>>> price + _
113.0625

>>> round(_, 2)
113.06

```

#### Strings 

You can use single or double quotes in Python with the same result.

Remember strings are immutable in Python

```bash
>>> 'doesn\'t'  # use \' to escape the single quote...
"doesn't"

>>> "doesn't"  # ...or use double quotes instead
"doesn't"

>>> s = 'First line.\nSecond line.'  # \n means newline

>>> s  # without print(), \n is included in the output
'First line.\nSecond line.'

>>> print(s)  # with print(), \n produces a new line
First line.
Second line.

```

Be careful, string _variables_ can be reassigned, but string data types
are immutable.

```bash
a = "Foo" # a var now points to "Foo"
b = a  
# b var points to the same "Foo" that a points to

a = a + a
# a var points to the new string "FooFoo", 
# but b still points to the old "Foo"

>>> print a
FooFoo

>>> print b # Outputs:
Foo
``` 

Raw Strings, preface string with r: `print(r'my string')`
```
>>> print('C:\some\name')  # here \n means newline!
C:\some
ame

>>> print(r'C:\some\name')  # note the r before the quote
C:\some\name
```

String _literals_
```bash
>>> 'Py' 'thon'
'Python'

>>> prefix = 'Py'
>>> prefix 'thon'  # can't concatenate a variable and a string literal
  File "<stdin>", line 1
    prefix 'thon'
    
>> prefix + 'thon' # use + operator to concat var and string literal
```

Strings can be indexed

```bash
>>> word = 'Python'
>>> word[0]  # character in position 0
'P'

>>> word[-1]  # last character
'n'

>>> word[-2]  # second-last character
'o'

>>> word[0:2]  # characters from position 0 (included) to 2 (excluded)
'Py'

>>> word[::-1] # reversed
'nohtyP'
```

Built-in functions: `len()` returns length of a string

```python
>>> s = 'supercalifragilisticexpialidocious'
>>> len(s)
34
```

Built-in String [Methods](https://docs.python.org/3/library/stdtypes.html#string-methods)

```python
str.capitalize()
str.endswith(suffix[, start[, end]])

'{name} was born in {country}'.format_map(Default(name='Connor'))
'Connnor was born in country'

'1,2,3'.split(',')
['1', '2', '3']
```

