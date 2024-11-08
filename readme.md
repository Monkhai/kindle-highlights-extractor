# Kindle Highlight Exporter

This is a simple script that will try to extract all of your kindle highlights out into markdown files.

## Requirments

in order for this to work you need golang installed. In addition you need to add a `tokens.go` file with your tokens. The following is the format you must follow for this to work

```golang
var XMain string = "<your string token>"
var UbidMain string = "<your string token>"
var AtMain string = "<your string token>"
var SessionId string = "<your string token>"
```

To get these tokens you need go to the [kindle reader](reader.amazon.com) and get there from the cookies tab in the application section of the devtools

## Other

I made this script for personal use but then realized others might also not want to pay for something this simple. If this does not work for you please open an issue and i will be happy to try and fix it!
