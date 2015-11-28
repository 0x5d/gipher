# gipher

A GIF generator made with Go.

## Get it

```sh
$ go get github.com/castillobg/gipher
```

## Usage

```sh
gipher [OPTIONS] [in...] out
```

- ``in...`` is the list of absolute paths to the folders with the .png, .jpg and .jpeg files to be used.
**The files should be the same size**.
- ``out`` is the absolute path with the name of the output file.

**Options:**

- ``-d``: The delay between frames, in 100s of a second. Default: 25.


# To-dos

- Support images of different size.
- Choose between recursive/ not recursive.
