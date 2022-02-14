# Building the project

- Go to **src** subdirectory and run one of the following commands

### Auto determine the operating system

```shell
$ make build
```

### Linux

- ARM

```shell
$ make build-linux-arm64
```

- AMD64

```shell
$ make build-linux-amd64
```

### macOS

- ARM

```shell
$ make build-mac-arm64
```

- AMD64

```shell
$ make build-mac-amd64
```

### Windows

- AMD64

```shell
$ make build-windows-amd64
```

# Execution

- Go to **bin** subdirectory and run one of the following commands

### macOS and Linux

```shell
$ ./letovo
```

### Windows

```shell
$ ./letovo.exe
```

# References

- Cobra: https://github.com/spf13/cobra
- Viper: https://github.com/spf13/viper

---
**© Made with ❤️ by arsikurin**