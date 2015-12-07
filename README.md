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


## Example

Let's say you have this folder structure:
```sh
home/
└── Pictures/
    └── ducks/
        ├── duck_1.jpg
        ├── ducks.png
        ├── ducks_2.jpeg
        └── img1231.jpg
```

To create a GIF slideshow of all the duck images, with a 1/2 second delay between each one, you'd run:
```sh
$ gipher -d 50 ~/Pictures/ducks/ ~/Pictures/ducks.gif
```

``ducks.gif`` will be created in ``~/Pictures/``!

![Ducks](https://i.imgur.com/gtPR7fh.gif "How awesome are duck gifs?")

# To-dos

- Add tests.
- Support images of different size.
- Add quiet mode.
