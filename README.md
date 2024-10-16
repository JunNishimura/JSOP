<div align='center'>
  <h1>JSOP</h1>
  <h3>write program in JSON</h3>
</div>

<p align='center'>
  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/JunNishimura/JSOP">
  <img alt="GitHub" src="https://img.shields.io/github/license/JunNishimura/JSOP">
  <a href="https://goreportcard.com/report/github.com/JunNishimura/jsop"><img src="https://goreportcard.com/badge/github.com/JunNishimura/jsop" alt="Go Report Card"></a>
</p>

```json
{
    "command": {
      "symbol": "print",
      "args": "Hello, World!"
    }
}
```

## 💾 Installation
### Homebrew Tap
```
brew install JunNishimura/tap/JSOP
```

### go intall
```
go install github.com/JunNishimura/jsop@latest
```

### GitHub Releases
Download exec files from [GitHub Releases](https://github.com/JunNishimura/JSOP/releases).

## 💾 How to use
Since REPL is not provided, you can write your program in any file and pass the file path as a command line argument to execute the program.

```bash
jsop ./path/to/file.jsop.json
```

## 📖 Language Specification
1. Everything is an expression.
2. Only `.jsop` and `.jsop.json` are accepted as file extensions.

### Integer
Integer value is a sequence of numbers.
<details open><summary>Example</summary>

```json
123
```
</details>

### String
String value is a sequence of letters, symbols, and spaces enclosed in double quotation marks.
<details open><summary>Example</summary>

```json
"this is a string"
```
</details>

### Boolean
Boolean value is either `true` or `false`.
<details open><summary>Example</summary>

```json
true
```
</details>

### Array
Arrays are composed of expressions.
<details open><summary>Example</summary>

```json
[1, "string", true]
```
</details>

### Identifiers
Strings beginning with the `$` symbol are considered as identifiers.
<details open><summary>Example</summary>
  
```json
"$x"
```
</details>

Embedding the Identifier in a string is accomplished by using curly brackets.
<details open><summary>Example</summary>
  
```json
"{$hello}, world!"
```
</details>

### Assignment
To assign a value or function to an identifier, use the `set` key. 
| parent key | children key | explanation |
| ---- | ---- | ---- |
| set |  | declaration of assignment |
|  | var | identifier name |
|  | val | value to assign |
<details open><summary>Example</summary>

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
| parent key | children key | explanation |
| ---- | ---- | ---- |
| lambda |  | declaration |
|  | params | parameters(optional) |
|  | body | body of function |
<details open><summary>Example</summary>

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
| parent key | children key | explanation |
| ---- | ---- | ---- |
| command |  | declaration of function calling |
|  | symbol | function to call |
|  | args | arguments(optional) |
<details open><summary>Example</summary>

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
| 関数 | explanation |
| ---- | ---- |
| + | addition |
| - | subtraction |
| * | multiplication |
| / | division |
| % | modulo |
| ! | negation |
| && | and operation |
| \|\| | or operation |
| == | equation |
| != | non equation |
| > | greater than |
| >= | greater than equal |
| < | smaller than |
| >= | smaller than equal |
| print | print to standard output |
| len | length of array |
| at | access to the element of array |

### If
Conditional branches can be implemented by using the `if` key.
| parent key | children key | explanation |
| ---- | ---- | ---- |
| if |  | declaratoin of if |
|  | cond | condition |
|  | conseq | consequence(the program to execute when cond is true) |
|  | alt | alternative(the program to execute when cond is false) |
<details open><summary>Example</summary>

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
| parent key | children key | explanation |
| ---- | ---- | ---- |
| loop |  | declaration of loop |
|  | for | the identifier for loop counter |
|  | from | the initial value of loop counter |
|  | until | loop termination condition (break when loop counter equals this value) |
|  | do | Iterative processing body |
<details open><summary>Example</summary>

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

You can also perform a loop operation on the elements of an Array. Unlike the example above, the `in` key specifies an Array.
<details open><summary>Example</summary>

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

Also, insert `break` and `continue` as keys as follows.
```json
[
    {
        "set": {
            "var": "$sum",
            "val": 0
        }
    },
    {
        "loop": {
            "for": "$i",
            "from": 1,
            "until": 15,
            "do": {
                "if": {
                    "cond": {
                        "command": {
                            "symbol": ">",
                            "args": ["$i", 10]
                        }
                    },
                    "conseq": {
                        "break": {}
                    },
                    "alt": {
                        "if": {
                            "cond": {
                                "command": {
                                    "symbol": "==",
                                    "args": [
                                        {
                                            "command": {
                                                "symbol": "%",
                                                "args": ["$i", 2]
                                            }
                                        },
                                        0
                                    ]
                                }
                            },
                            "conseq": {
                                "set": {
                                    "var": "$sum",
                                    "val": {
                                        "command": {
                                            "symbol": "+",
                                            "args": ["$sum", "$i"]
                                        }
                                    }
                                }
                            },
                            "alt": {
                                "continue": {}
                            }
                        }
                    }
                }
            }
        }
    },
    "$sum"
]
```

### Retrun
Use `return` key when you exit the program with return.
```json:return.jsop.json
[
    {
        "set": {
            "var": "$f",
            "val": {
                "lambda": {
                    "body": [
                        {
                            "set": {
                                "var": "$sum",
                                "val": 0
                            }
                        },
                        {
                            "loop": {
                                "for": "$i",
                                "from": 1,
                                "until": 11,
                                "do": {
                                    "if": {
                                        "cond": {
                                            "command": {
                                                "symbol": ">",
                                                "args": ["$i", 5]
                                            }
                                        },
                                        "conseq": {
                                            "return": "$sum"
                                        },
                                        "alt": {
                                            "set": {
                                                "var": "$sum",
                                                "val": {
                                                    "command": {
                                                        "symbol": "+",
                                                        "args": ["$sum", "$i"]
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    ]
                }
            }
        }
    },
    {
        "command": {
            "symbol": "$f"
        }
    }
]
```

### Macro
Macro can be defined by using `defmacro` key.
| parent key | children key | explanation |
| ---- | ---- | ---- |
| defmacro |  | declaration of macro definition |
|  | name | name of macro |
|  | keys | keys |
|  | body | the body of macro |

You can also call the `quote` symbol for quoting, and unquote by adding backquotes to the beginning of the string.
<details open><summary>Example</summary>

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

Defining function can be much simpler if you use Macro.
<details open><summary>Example</summary>

```json
[
    {
        "defmacro": {
            "name": "defun",
            "keys": ["name", "params", "body"],
            "body": {
                "command": {
                    "symbol": "quote",
                    "args": {
                        "set": {
                            "var": ",name",
                            "val": {
                                "lambda": {
                                    "params": ",params",
                                    "body": ",body"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    {
        "defun": {
            "name": "$add",
            "params": ["$x", "$y"],
            "body": {
                "command": {
                    "symbol": "+",
                    "args": ["$x", "$y"]
                }
            }
        }
    },
    {
        "command": {
            "symbol": "$add",
            "args": [1, 2]
        }
    }
]
```
</details>

### Comment
Comments can be inesrted by using `//` key.
<details open><summary>Example</summary>

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

## 🤔 FizzBuzz Problem
Finally, I have included an example of solving a FizzBuzz problem in JSOP.
<details open><summary>Example</summary>

```json
[
    {
        "set": {
            "var": "$num",
            "val": 31
        }
    },
    {
        "loop": {
            "for": "$i",
            "from": 1,
            "until": "$num",
            "do": {
                "if": {
                    "cond": {
                        "command": {
                            "symbol": "&&",
                            "args": [
                                {
                                    "command": {
                                        "symbol": "==",
                                        "args": [
                                            {
                                                "command": {
                                                    "symbol": "%",
                                                    "args": ["$i", 3]
                                                }
                                            },
                                            0
                                        ]
                                    }
                                },
                                {
                                    "command": {
                                        "symbol": "==",
                                        "args": [
                                            {
                                                "command": {
                                                    "symbol": "%",
                                                    "args": ["$i", 5]
                                                }
                                            },
                                            0
                                        ]
                                    }
                                }
                            ]
                        }
                    },
                    "conseq": {
                        "command": {
                            "symbol": "print",
                            "args": "{$i}: FizzBuzz"
                        }
                    },
                    "alt": {
                        "if": {
                            "cond": {
                                "command": {
                                    "symbol": "==",
                                    "args": [
                                        {
                                            "command": {
                                                "symbol": "%",
                                                "args": ["$i", 3]
                                            }
                                        },
                                        0
                                    ]
                                }
                            },
                            "conseq": {
                                "command": {
                                    "symbol": "print",
                                    "args": "{$i}: Fizz"
                                }
                            },
                            "alt": {
                                "if": {
                                    "cond": {
                                        "command": {
                                            "symbol": "==",
                                            "args": [
                                                {
                                                    "command": {
                                                        "symbol": "%",
                                                        "args": ["$i", 5]
                                                    }
                                                },
                                                0
                                            ]
                                        }
                                    },
                                    "conseq": {
                                        "command": {
                                            "symbol": "print",
                                            "args": "{$i}: Buzz"
                                        }
                                    },
                                    "alt": {
                                        "command": {
                                            "symbol": "print",
                                            "args": "$i"
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
]
```
</details>

## 🪧 License
JSOP is released under MIT License. See [MIT](https://raw.githubusercontent.com/JunNishimura/JSOP/main/LICENSE)
