# Xopet

A tool to unpack the content of file using zstd encoder.

## Install

```shell
$ go mod download
```

## Build

```shell
$ go build -ldflags "-s -w"
```

## Usage

Peek the list of the content

```shell
xopet -l <zstd_file>
```

Dump all the contents from the zstd file. By default, it will dump to the current directory.
To specify the output dir, add flag `-o <output_dir>`.

```shell
xopet <zstd_file>
```

## License

[MIT](LICENSE)
