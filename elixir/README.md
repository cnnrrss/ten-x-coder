# Elixir

https://elixir-lang.org/getting-started/basic-types.html

## Installation

`brew install elixir`

Or Docker 

`docker run -it --rm elixir`


### PATH env var

`export PATH="$PATH:/path/to/elixir/bin"`

Example: (/usr/local/bin/elixir)

Verify

`elixir --version`

## Interactive Mode

You have 3 executables `iex`, `elixir` and `elixirc`

`iex`

```bash
Erlang/OTP 20.0 [64-bit] [smp:2:2] [...]

Interactive Elixir (1.9.1) - press Ctrl+C to exit
iex(1)> 40 + 2
42
iex(2)> "hello" <> " world"
"hello world"
```

Press Ctrl+C to exit

`iex -S SCRIPTNAME`


## Basic Types

```bash
iex> 1          # integer
iex> 0x1F       # integer
iex> 1.0        # float
iex> true       # boolean
iex> :atom      # atom / symbol
iex> "elixir"   # string
iex> [1, 2, 3]  # list
iex> {1, 2, 3}  # tuple
```