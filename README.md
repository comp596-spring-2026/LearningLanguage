# LearningLanguage
A programming language designed to serve as a tool to bridge the gap between a students understanding of block-based languages like Scratch and complex modern lanugages like Java or C++.

### Requirements
[Go Installation Instructions](https://go.dev/doc/install)<br>
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