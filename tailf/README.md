# tailf

This tool replicates `tail -f` functionality with the addition of two features:  
* clear the screen while continuing to tail the file  
* drop a horizontal rule marker to mark your place

## Installation

```
go get github.com/jakewarren/junkdrawer/tailf
```

## Usage

Keyboard Shortcuts:
| Shortcut | Action               |
|----------|----------------------|
| Esc      | Quit                 |
| Ctrl+c   | Quit                 |
| q        | Quit                 |
| h        | Draw horizontal rule |
| c        | Clear Screen         |
| Ctrl+l   | Clear Screen         |


## Examples

Read from stdin:
```
cat /var/log/syslog | tailf
```
Read from single file:
```
tailf /var/log/syslog
```
Read from multiple files:
```
tailf /var/log/syslog /tmp/app.log
```

## Known Issues

For an unknown reason, occasionally the program becomes unresponsive to the keyboard shortcuts. I believe I have fixed the issue but be advised it could happen.