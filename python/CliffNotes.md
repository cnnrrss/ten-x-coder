# Summary Intro to Python

### Data Structures
https://dbader.org/blog/fundamental-data-structures-in-python

```
>>> help()

To get a list of available modules, keywords, symbols, or topics, type
"modules", "keywords", "symbols", or "topics"

help> keywords

help> modules numpy
```
i
## Data Types

**Mutable objects**:
list, dict, set, byte array

**Immutable objects**:
int, float, complex, string, tuple, _frozen set_ (immutable version of set), bytes

### Strings
`x = 'a'` \
`x = "a"` \
`x = r'abc\n'` \



### Numbers
`x = 1 # integer` \
`x = 1.234 # float` 

**Operators**: +, -, *, /, **, %

### Lists
`x = []`

## Methods / Classes

This is my_module.py:
```python
def external_func():
    return 23

def _internal_func():
    return 42

# then if you try to do....
>>> from my_module import *
>>> external_func()
23
>>> _internal_func()
NameError: "name '_internal_func' is not defined"

```

#### Method Overwriting / Overriding / Overloading

**Overwriting**
```python
def f(x):
    return x + 42
print(f(3))

# f will be overwritten (or redefined) in the following:
def f(x):
    return x + 43
print(f(3))
```

**Overriding**
```python
class A:
    def methodA(self):
        return self + '!'

class B(A):
    def methodA(self):
        return self + '?'
```

**Overloading**

```
f(2)
f(2, 3)
```

Not supported like Java or C++ but...

```python
def f(n, m=None):
    if m:
        return n + m +42
    else:
        return n + 42
print(f(3), f(1, 3))
```

Or ...

```python
def f(*x):
    if len(x) == 1:
        return x[0] + 42
    elif len(x) == 2:
        return x[0] - x[1] + 5
    else:
        return 2 * x[0] + x[1] + 42
print(f(3), f(1, 2), f(3, 2, 1))
```
