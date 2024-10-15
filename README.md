<div align='center'>
  <h1>JSOP</h1>
  <h3>JSON based programming language</h3>
</div>

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
<details><summary>Example</summary>

```json
123
```
</details>

### String
String value is a sequence of letters, symbols, and spaces enclosed in double quotation marks.
<details><summary>Example</summary>

```json
"this is a string"
```
</details>

### Boolean
Boolean value is either `true` or `false`.
<details><summary>Example</summary>

```json
true
```
</details>

### Array
Arrays are composed of expressions.
<details><summary>Example</summary>

```json
[1, "string", true]
```
</details>

### Identifiers
Strings beginning with the `$` symbol are considered as identifiers.
<details><summary>Example</summary>
  
```json
"$x"
```
</details>

### Assignment
To assign a value or function to an identifier, use the `set` key. Specify the `var` key for the name of the identifier and the `val` key for the value to be assigned.
<details><summary>Example</summary>

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
</details>

### Function
#### Function Definition
Functions can be defined by using `set` key and `lambda` expression`.
<details><summary>Example</summary>

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
</details>

#### Function Call
Functions can be called by using `command` key.
<details><summary>Example</summary>

```json
{
    "command": {
        "symbol": "+",
        "args": [1, 2]
    }
}
```
</details>

### Builtin Functions
Builtin functions are as follows,
1. `+`: add
2. `-`: subtract
3. `*`: multiply
4. `/`: divide
5. `%`: modulo
6. `&&`: and
7. `||`: or
8. `==`: equal
9. `!=`: not equal
10. `<`: less than
11. `<=`: less than euqal
12. `>`: greater than
13. `>=`: greater than euqal
14. `!`: negate
15. `print`: print the arguments to the terminal
16. `len`: length of the array
17. `at`: pick up the element of the array

### If
Conditional branches can be implemented by using the `if` key. The `alt` key is optional.
<details><summary>Example</summary>

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
</details>

### Loop
Iterations are handled by using the `loop` key.
<details><summary>Example</summary>

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
</details>

<details><summary>Example</summary>

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
</details>

### Macro
Macro can be defined by using `defmacro` key.
<details><summary>Example</summary>

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
</details>

### Comment
Comments can be inesrted by using `//` key.
<details><summary>Example</summary>

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
</details>

## Issues

## 🪧 License
JSOP is released under MIT License. See [MIT](https://raw.githubusercontent.com/JunNishimura/JSOP/main/LICENSE)
