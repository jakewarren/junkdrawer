# mailit

quickly send emails from the command line

## Prequisities

This program assumes you have a local MTA (such as Postfix) installed and listening on localhost:25 with no authenticaion.

## Installation

```
go install github.com/jakewarren/junkdrawer/mailit@latest
``` 

## Usage

```
Usage of mailit:
  -a, --attach strings   files to attach to the email
      --bcc strings      addresses to send bcc <BCC>
      --cc strings       addresses to send cc <CC>
  -n, --dry-run          dry run, don't actually send email
      --file string      read input from a file
  -f, --from string      email address to send from (defaults to system default)
      --inline           inline style information via premailer
      --pre              wrap body with a <pre> element (default true)
  -s, --subject string   email subject
  -t, --to strings       addresses to send to <To>
``` 

## Example

email the output of a command (read from STDIN):
``` 
date | mailit -f jdoe@acme.com -t jdoe@acme.com -s "current time"
```

email the contents of a file:
``` 
mailit --file /tmp/program.log -f jdoe@acme.com -t jdoe@acme.com
```

send blank email with a file attachment:
``` 
echo "" | mailit -a /tmp/program.log -f jdoe@acme.com -t jdoe@acme.com -s "log file"
```
