## Overview

This series will cover and introduction of writing a lexer/parser in the Go programming language. The series will begin with a overview of what a parser and lexer are and then will proceed to cover a range of topics which will result in the end creation of a functional parser.

The language we will be building will be derivation of the golang. It will be assumed experience in the golang as a preq aswell as a general experience with complex programming, recursion, logic and some basic tree traversal.

The language will also be able to handle static types and explicit/ types. Through this series you will learn how to come up with your own syntax from idea to ast. At the end of the series you will have a AST made which represents a complex language which you can take and extend further and either use for code-generation or a interpreter/vm.


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

## Language Syntax

```ts
import fs;
import tasks;
import myLib from "../lib/myLib.lang";

const MIN = 1;
const MAX = 100;
let numbers: []number;
numbers = MIN..MAX;

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
