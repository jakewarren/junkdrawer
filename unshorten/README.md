# unshorten

"unshorten" shortened URLs from the command line

## Installation

```
go install github.com/jakewarren/junkdrawer/unshorten@latest
```

## Usage

```
❯ unshorten -h
Usage of unshorten:
  -trace
    	display verbose trace information
```
```
unshorten <url>
```

## Example

```
❯ unshorten -trace https://t.co/rWFa3sbQEa
[301] https://t.co/rWFa3sbQEa
[200] https://www.cyberscoop.com/north-korea-lazarus-group-bangladesh-bank-donald-trump-xi-jinping/
```
