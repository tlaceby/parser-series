## Overview

This series will cover and introduction of writing a lexer/parser in the Go programming language. The series will begin with a overview of what a parser and lexer are and then will proceed to cover a range of topics which will result in the end creation of a functional parser.

The language we will be building will be derivation of the golang. It will be assumed experience in the golang as a preq aswell as a general experience with complex programming, recursion, logic and some basic tree traversal.

## Language Syntax

```ts
import math; // Expected STD Package
import string;
import bar from "./foo/bar.lang"; // User Defined Package


const PI = math.PI;
let result = math.sqrt(4);

println("Hello world")
println($"sqrt(4) = {result}"); // Template Strings

foreach (value, index in [1..10]) {
    println(value);
}

if (math.randint(100) >= 21) {
    println("You may have a drink");
} else {
    println("You must be 21 to enjoy a drink");
}

// Function Declarations
fn add (x, y) {
    x + y; // Last statement evaluated inside a block is returned
}

typeof math          // object
typeof "Hello world" // string
typeof 45.5          // number
typeof [1, 2, 3]     // array
typeof add           // function
typeof null          // null

const nums = [1, 2, 3, 4, 5];
nums[2] = 10;

println(len("hello")) // 5
println(string.lowercase("HELLO")) // hello

// If is an expression not a statement.
let res = if (math.randint(10) % 2 == 0) {
    "even"
}

// Fancy assignment if null syntax
res ??= "odd";

// Objects & Classes
class Person {
    fn mount (self, name, age) {
        self.name = name;
        self.age = age;
    }
}

// This is simply a Class with no-name already created with the name,age properties which inherits from the Object class.
const person = {
    name: "Tyler",
    age: 23,
};


```

## Topics

- **Lexing**
  - Simple Iteration vs _Regex_
  - Error Generation
    - Recoverable vs Fatal Errors
  - Implimenting a basic lexer for arithmetic expressions.
  - Keywords/Reserved words
  - Floating Point & Hexadecimal
  - Extra / Msc - Template Strings - Character Escaping
- **Parsing**
  - Pratt Parser
  - Translating between AST Forms
  - Recursive Descent
