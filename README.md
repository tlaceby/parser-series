# Writing a Lexer/Parser in Go: A YouTube Series

## Overview

Welcome to my comprehensive YouTube series on building a lexer/parser using the Go programming language. We will start with the basics of what lexers and parsers do, gradually moving towards creating a functional parser for a language derived from Go/Typescript/C#.

By the end of this series, you will learn to design your own language syntax, from conceptualization to creating an Abstract Syntax Tree (AST). This AST will represent a complex language that you can further extend for code generation or to build an interpreter/VM.

### Prerequisites

- Experience with Go programming
- Understanding of complex programming concepts
- Familiarity with recursion and logic
- Basic knowledge of tree traversal

## Topics Covered

### Lexing

- Simple Iteration vs. Regex
- Error Generation (Recoverable vs. Fatal Errors)
- Implementing a Basic Lexer for Arithmetic Expressions
- Keywords/Reserved Words
- Floating Points & Hexadecimals
- Miscellaneous: Template Strings, Character Escaping

### Parsing

- Pratt Parser
- Translating Between AST Forms
- Recursive Descent

## Language Syntax Example

Here's a sneak peek into the syntax of the language you'll be working with:

### Imports

Imports are used to include external modules or libraries in your program, allowing you to use their functionality within your own code. Here `fs` and `tasks` could be standard library modules while `myLib` is an external module in a nearby folder.

```ts
import fs;
import tasks;
import myLib from "../lib/myLib.lang";

```

### Variable Declarations

In this language, you can declare variables using const for constants that cannot be reassigned, and let for variables whose values can change. Types can be explicitly defined, such as []number for an array of numbers, or inferred by the type checker if not specified.

You will also see a custom syntax `MIN..MAX` when constructing the numbers array. This series will also demonstrate how easy it can be to add custom and unique syntax similar to the range-expression in Rust.

```ts
const MIN = 1;
const MAX = 100;

let numbers: []number;
numbers = MIN..MAX; // returns the numbers [1, 2, 3, ..., 99, 100] as an array.

```

### Conditional Statements & Loops

The language supports conditional statements like if-else for branching logic based on conditions. Loops, such as foreach and while, are used for iterating over collections or executing a block of code multiple times until a condition is met.

```ts
if random.selectOne(choices) == 50 {
  println("Your number was selected!");
} else {
  println("Your number was not selected");
}

foreach value, index in choices {
    println(value, index);
}

foreach value in choices {
    println(value);
}

let x = 0;
while x < 10 {
    x += 1;
    println(x);
}

```

### Classes & OOP

The language adopts an object-oriented programming approach, allowing the definition of classes with properties and methods. Classes support encapsulation of data and behavior, with syntax for constructors (fn mount), methods (fn birthday, fn greet), and property declarations. Object instances are created using the new keyword.

```ts
class Person {
  let name: string;
  let age: number;
  let languages: []string = ["Go", "Javascript"];

  fn mount (name: string, age: number) {
    this.age = age;
  }

  fn birthday (): number {
    this.age += 1;
  }

  fn greet () {
    println("Hello my name is ", this.name);
  }
}

const p1 = new Person ("John Doe", 43);
p1.greet();

```

### Functions & Function Expressions

Functions are defined using the fn keyword, followed by the function name, parameters, and return type. The language also supports anonymous functions (or function expressions), which can be assigned to variables, passed as arguments, or returned from other functions. Anonymous functions use the fn keyword without a name.

```rust
fn abs (n: number): number {
  if n >= 0 {
    n;
  }

  -n;
}

const add = fn(x: number, y: number): number {
  x + y;
};

tasks.interval(fn(task: TaskInfo){
  if task.time > time.second * 10 {
    tasks.kill(task.id);
  }
}, 1000);

```

### Code Example

Puting it all together here is a sample program:

```ts
import fs;
import path;
import time;

class DirectoryReader {
  let directoryPath: string;

  fn mount(directoryPath: string) {
    this.directoryPath = directoryPath;
  }

  fn readRecentFiles() {
    let allFiles: []string = fs.readDir(this.directoryPath);
    let recentFiles: []string = [];

    foreach file in allFiles {
      let fullPath: string = path.join(this.directoryPath, file);
      let fileInfo: FileInfo = fs.stat(fullPath);
      if this.isFileRecent(fileInfo.creationTime) {
        recentFiles.push(fullPath);
      }
    }

    foreach file in recentFiles {
      println(file, fs.stat(file).creationTime);
    }
  }

  fn isFileRecent(creationTime: Time): boolean {
    let twentyFourHoursAgo: Time = time.now() - time.hours(24);
    creationTime > twentyFourHoursAgo;
  }
}

fn main() {
  const directory: string = "/path/to/directory";
  const reader = new DirectoryReader();
  reader.mount(directory);
  reader.readRecentFiles();
}

main();

```

This sample program demonstrates the versatility and power of the parser we're building together. It showcases several language features, including:

- Imports: It begins by importing necessary modules fs, path, and time, which are essential for file and time operations.

- Class Definition: The DirectoryReader class is defined with methods for mounting a directory and reading files that were created within the last 24 hours. This demonstrates the object-oriented capabilities of our language.

- Variable Declarations: Variables are declared using let for mutable variables and const for constants, showing the language's capability to handle different types of data storage.

- Conditional Logic and Loops: The program uses foreach loops to iterate over files and an if statement to filter recent files, illustrating how our language handles control flow.

- Function Definition and Invocation: The main function is defined and invoked, which uses an instance of DirectoryReader to perform its operations. This shows how functions are first-class citizens in our language.

By parsing this program, our parser will generate an Abstract Syntax Tree (AST) that represents the program's structure. This AST can then be used for further code generation, interpretation, or compilation, depending on the end goal of your language project.

This example serves as a practical application of all the concepts we've covered in this series. It ties together lexing, parsing, AST generation, and gives a glimpse into how these components work together to interpret or compile a high-level program written in our custom language.

By following this series and implementing your own lexer/parser, you're not just learning the mechanics of these processes but also how to apply them to create a language that can serve real-world needs, similar to how Go, Typescript, and C# are used today.
Next Steps

-------

In upcoming series, we may dive deeper into error handling, optimization techniques for our parser, typechecking, code-generation, and interpretation.

Stay tuned, and happy coding!
