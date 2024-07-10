# bgate
```
A terminal interface to Bible Gateway

Usage:
  bgate [flags] <query>
  bgate [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  download    Download a translation of the Bible for local usage rather than reaching out to BibleGateway
  help        Help about any command
  list        List all books of the Bible and how many chapters they have

Flags:
  -c, --config string        Config file to use. (default "~/.config/bgate/config.json")
      --force-local          Force the program to crash if there isn't a local copy of the translation you're trying to read.
      --force-remote         Force the program to use the remote searcher even if there is a local copy of the translation.
  -h, --help                 help for bgate
  -p, --padding int          Horizontal padding in character count.
  -t, --translation string   The translation of the Bible to search for. (default "ESV")
  -w, --wrap                 Wrap verses, this will cause it to not start each verse on a new line.

Use "bgate [command] --help" for more information about a command.
```

## Install
To install, you must have golang installed on your machine. You can just run:
```
go install github.com/Shhwip/bgate-scraper@latest
```

## Examples
An example would be:
```
bgate -t LSB -i 1cor1
```
which would pull up 1 Corinthians 1 in an interactive session.


## Config
Config values use the same name as the flag. Below is my personal config.
``` json
{
	"translation": "NRSVUE",
	"delay": 10000
}
```

## Note
Currently, the local querying is not as feature rich as remote querying.

Downloads are saved in ~/.bgate as sqlite files
