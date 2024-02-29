# Build Your Own wc Tool

This challenge is a part of the <b>Coding Challenges</b> by <i>John Crickett</i>, which helps software engineers level up their skills through deliberate practice. 

Link to the Coding Challenegs: <a href="https://codingchallenges.fyi/challenges/intro">Click Here</a>


## Table Of Contents
- [Description](#description)
- [Usage](#usage)
- [Examples](#examples)
- [Learnings](#learnings)


## Description

This challenge is to build out own version of the Unix command line tool wc! We need to take the input from a file passed in arguments or stdin, and print the number of bytes, words, lines and charcters according to the flag present.

| Flag | Info                                         |
| ---- | -------------------------------------------- |
| -c   | Outputs the number of bytes in the file      |
| -l   | Outputs the number of lines in the file      |
| -w   | Outputs the number of words in the file      |
| -m   | Outputs the number of characters in the file |



## Usage

Execute `go build -o ccwc` to create the executable. Now run the executable `ccwc` with the required flag and file name. Following are the possible arguments.



Basic Command will be of type

```shell
./ccwc <flag> <fileName>
```

or

```sh
cat <fileName> | ./ccwc <flag>
```



## Examples

#### From File

| Output Requirment      | Command              |
| ---------------------- | -------------------- |
| Bytes                  | `./ccwc -c test.txt` |
| Lines                  | `./ccwc -l test.txt` |
| Words                  | `./ccwc -w test.txt` |
| Characters             | `./ccwc -m test.txt` |
| Lines, Words and Bytes | `./ccwc test.txt`    |

#### From Stdin

| Output Requirment      | Command                   |
| ---------------------- | ------------------------- |
| Bytes                  | `cat test.txt |./ccwc -c` |
| Lines                  | `cat test.txt |./ccwc -l` |
| Words                  | `cat test.txt |./ccwc -w` |
| Characters             | `cat test.txt |./ccwc -m` |
| Lines, Words and Bytes | `cat test.txt |./ccwc`    |



## Learnings

- It was a nice opportunity for me to learn about the `flag` package of Go, which helped me read the flag and file name from the command line.
- Implemented Modules for `Arguments` and `Calculations`, which added a practice for me to work with custom modules, exporting Types and Functions in GoLang.
- Utilized `os` and `io` packages to read the file and stdin respectively.
- Calculating the Bytes, Lines, Words and Charcters in Go was a bit challenging part, which needed a bit of lookup on how the `bufio` package works. (Although I didn't use it eventually for the solution)
- Got to know more about `Runes` and `Fields` in GoLang.

## Contribution

Contributions are welcome! If you find any bugs or want to enhance the app, feel free to open issues or submit pull requests. Please make sure to follow the coding standards and guidelines.

Happy coding! If you have any questions or need assistance, don't hesitate to reach out.

## Designed and Implemented By
<b>Harsh Bardolia</b>
-   [Github](https://github.com/HarshBardolia01)
-   [LinkedIn](https://www.linkedin.com/in/harsh-bardolia-0a0407217/)

## License

[MIT](/LICENSE) © Harsh Bardolia