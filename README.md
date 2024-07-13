# bgate-scraper
```
A tool for downloading the public domain bibles from BibleGateway

Usage:
  bgate-scraper [flags] <query>
  bgate-scraper [command]

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

Use "bgate-scraper [command] --help" for more information about a command.
```

## Install
To install, you must have golang installed on your machine. Do that [here](https://go.dev/doc/install). 

Then run:
```
go install github.com/Shhwip/bgate-scraper@latest
```

## Examples
To download the KJV run:
```
bgate-scraper -t KJV download
```
to pull up 1 Corinthians 1 using the KJV translation
```
bgate-scraper -t KJV 1cor1
```

## Config
Config values use the same name as the flag. Below is my personal config.
``` json
{
	"translation": "WEB",
  "padding": 100
}
```

## Note
Currently, the local querying is not as feature rich as remote querying.

Downloads are saved in ~/.bgate as sqlite files
