# How to use

[Build from source](#building-the-project-from-source) or
[download the latest release](https://github.com/arsikurin/letovoAnalyticsCLI/releases)

## Building the project from source

- Go to **src** subdirectory and run one of the following commands

##### Auto determine the operating system

```shell
$ make build
```

##### Linux

- ARM

```shell
$ make build-linux-arm64
```

- AMD64

```shell
$ make build-linux-amd64
```

#### macOS

- ARM

```shell
$ make build-mac-arm64
```

- AMD64

```shell
$ make build-mac-amd64
```

##### Windows

- AMD64

```shell
$ make build-windows-amd64
```

## Execution

- Go to **bin** subdirectory and run one of the following commands

##### macOS and Linux

```shell
$ ./letovo
```

##### Windows

```shell
$ ./letovo.exe
```

---
**© Made with ❤️ by arsikurin**