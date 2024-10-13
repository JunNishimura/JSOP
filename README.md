<h1 align='center'>
  JSOP <br/>JSON based programming language
</h1>

<p align='center'>
  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/JunNishimura/JSOP">
  <img alt="GitHub" src="https://img.shields.io/github/license/JunNishimura/JSOP">
  <a href="https://goreportcard.com/report/github.com/JunNishimura/JSOP"><img src="https://goreportcard.com/badge/github.com/JunNishimura/JSOP" alt="Go Report Card"></a>
</p>

```json
{
    "command": {
      "symbol": "print",
      "args": "Hello, World!"
    }
}
```

## 💻 Installation
### Homebrew Tap
```
brew install JunNishimura/tap/JSOP
```

### go intall
```
go install github.com/JunNishimura/jsop@latest
```

## 📖 Language Specification
1. Everything is an expression.
2. Only `.jsop` and `.jsop.json` are accepted as file extensions.

### Integer
Integer value is a sequence of numbers.
```json
123
```

### String
String value is a sequence of letters, symbols, and spaces enclosed in double quotation marks.
```json
"this is string"
```

### Boolean
Boolean value is either `true` or `false`.
```json
true
```

### Array
Arrays are composed of expressions.
```json
[1, "string", true]
```

### Identifiers
Strings beginning with the `$` symbol are considered as identifiers.
```json
"$x"
```

### Assignment
To assign a value or function to an identifier, use the `set` key. Specify the `var` key for the name of the identifier and the `val` key for the value to be assigned.

```json
[
    {
        "set": {
            "var": "$x",
            "val": 10
        }
    },
    "$x"
]
```

### Function
#### Function Definition
Functions can be defined by using `set` key and `lambda` expression`.
```json
{
    "set": {
        "var": "$add",
        "val": {
            "lambda": {
                "params": ["$x", "$y"],
                "body": {
                    "command": {
                        "symbol": "+",
                        "args": ["$x", "$y"]
                    }
                }
            }
        }
    }
}
```

#### Function Call
Functions can be called by using `command` key.
```json
{
    "command": {
        "symbol": "+",
        "args": [1, 2]
    }
}
```

### Builtin Functions
Builtin functions are as follows,
1. `+`: add
2. `-`: subtract
3. `*`: multiply
4. `/`: divide
5. `==`: equal
6. `!=`: not equal
7. `<`: less than
8. `<=`: less than euqal
9. `>`: greater than
10. `>=`: greater than euqal
11. `!`: negate
12. `print`: print the arguments to the terminal
13. `len`: length of the array
14. `at`: pick up the element of the array

### If
Conditional branches can be implemented by using the `if` key.
```json
{
    "if": {
        "cond": {
            "command": {
                "symbol": "==",
                "args": [1, 2]
            }
        },
        "conseq": {
            "command": {
                "symbol": "+",
                "args": [3, 4]
            }
        },
        "alt": {
            "command": {
                "symbol": "*",
                "args": [5, 6]
            }
        }
    }
}
```

### Loop
Iterations are handled by using the `loop` key.
```json
{
    "loop": {
        "for": "$i",
        "from": 0,
        "until": 10,
        "do": {
            "command": {
                "symbol": "print",
                "args": "$i"
            }
        }
    }
}
```

```json
[
    {
        "set": {
            "var": "$arr",
            "val": [10, 20, 30]	
        }
    },
    {
        "loop": {
            "for": "$element",
            "in": "$arr",
            "do": {
                "command": {
                    "symbol": "print",
                    "args": "$element"
                }
            }
        }
    }
]
```

### Macro
Macro can be defined by using `defmacro` key.
```json
{
    "defmacro": {
        "name": "unless",
        "keys": ["cond", "conseq", "alt"],
        "body": {
            "command": {
                "symbol": "quote",
                "args": {
                    "if": {
                        "cond": {
                            "command": {
                                "symbol": "!", 
                                "args": ",cond"
                            }
                        },
                        "conseq": ",conseq",
                        "alt": ",alt"
                    }
                }
            } 
        }
    }
}
```


### Comment
Comments can be inesrted by using `//` key.
```json
{
    "//": "this is a function to add two values",
    "set": {
        "var": "$add",
        "val": {
            "lambda": {
                "params": ["$x", "$y"],
                "body": {
                    "command": {
                        "symbol": "+",
                        "args": ["$x", "$y"]
                    }
                }
            }
        }
    }
}
```

## Issues

## 🪧 License
JSOP is released under MIT License. See [MIT](https://raw.githubusercontent.com/JunNishimura/JSOP/main/LICENSE)