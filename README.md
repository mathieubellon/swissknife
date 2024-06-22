
# 🇨🇭 Swissknife

A little utility tool to manage various tasks

```bash
NAME:
   Swissknife - A multi-purposes utility command-line tool for managing detectors

USAGE:
   Swissknife [global options] command [command options] 

VERSION:
   0.1

COMMANDS:
   markdown  Generate markdown changelog links from the specified Detection Engine version
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```



## Usage/Examples

Generate markdown for a given detection engine version

```bash
swissknife markdown --version 2.115.0
```

With absolute URL

```bash
swissknife markdown --version 2.115.0 --absolute-url
```


## License

[MIT](https://choosealicense.com/licenses/mit/)

