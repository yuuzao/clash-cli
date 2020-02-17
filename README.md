## clash-cli
Use command line to manage clash proxy settings. 

## Installation
Windows platform is not supported yet.

#### Manual
Download for release page, then extract it to you PATH.

#### Use GO
```
go get github.com/Sisylocke/clash-cli
```
## Usage

```
clash-cli [flags]

Flags:
  -a, --add      add a new piece of rule
  -d, --delete   delete an existed piece of rule
  -f, --find     find a specific piece of rule
  -h, --help     help for clash-cli
  -m, --mode     change the proxy mode: GLOBAL, Rule or Direct
  -n, --node     switch to another proxy node
  -s, --status   show the current clash status
```
