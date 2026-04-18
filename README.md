# LearningLanguage
A programming language designed to serve as a tool to bridge the gap between a students understanding of block-based languages like Scratch and complex modern lanugages like Java or C++.

### Requirements
[Go Installation Instructions](https://go.dev/doc/install)
Clone the repo to its own directory

### Running the REPL
To run an interactive REPL on the command line, in the directory with main.go run

```bash
go run ./main.go
```

If you want to use file input, use
```bash
go run ./main.go -i <filename>.txt
```

If you want to use file output, use
```bash
go run ./main.go -o <filename>.txt
```
Both of these parameters can be used together.

### Language Documentation
#### Variable Creation/Assignment
To create a variable, use the following syntax:

```
create int/bool/string/float <identity>;
```

In order to assign a value to a variable, it first needs to be created, then use:
```
set <identity> = <expression>;
```

#### Valid Expressions
The following are the valid types of expressions within the language:
* Literals (123, true/false, 3.14, "Text")
* Prefix Expressions (-4, !false)
* Arithmetic Infix Expressions (1+1, 4-2, 10/5, 4*4)
* Comparison Infix Expressions (>/>=, ==/!=, </<=)

#### If Statements
The language operates on if/else pairs only, no else if. If you wish to achieve if else logic, nested if statements are needed.
Any statement that requires curly braces or newlines in other languages will use ```begin;``` to indicate the start of the body and ```end;``` to indicate the ending of the body.
```
if (1 > 2) begin;
print("1 is less than 2");
end;
else begin;
print("1 is not less than 2");
end;
```

#### Structures
To create a structure, use the following syntax:
```
struct myStruct(
    int x,
    bool y,
    string z
) [
    x: 123,
    y: true,
    z: "test"
];
```
Assignment of attributes of the structure are optional, allowing for the following:
```
struct myStruct(
    int x,
    bool y,
    string z
);
set myStruct.x = 123;
set myStruct.y = true;
set myStruct.z = "test";
```
To access attributes of a structure:
```
myStruct.x;
```

### Printing
In order to have your code output anything, you need a print statement:
```
print(1+2);
print("Hello World");
```
The parameters for printing are any expression.