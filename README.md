
# ics2txt

Display ics/ical files as text

## Installations

```
$ go get  github.com/yzzyx/ics2.txt
```

## Usage

```
$ ics2txt meeting.ics
```

Or in a mailcap-file, for usage with e.g. alot or mutt

```asciidoc
text/calendar: ics2txt '%s'; copiousoutput
```

## Credits

The actual parsing of the ics-files are done by the go-ical library by
emersion: https://github.com/emersion/go-ical