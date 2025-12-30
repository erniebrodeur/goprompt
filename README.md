# goprompt

[![CodeFactor](https://www.codefactor.io/repository/github/erniebrodeur/goprompt/badge)](https://www.codefactor.io/repository/github/erniebrodeur/goprompt)
[![Go Report Card](https://goreportcard.com/badge/github.com/erniebrodeur/goprompt)](https://goreportcard.com/report/github.com/erniebrodeur/goprompt)

## Usage

Install

    go install github.com/erniebrodeur/goprompt/cmd/goprompt

then

    goprompt

## Shell Integration

### Zsh

```
#!/bin/zsh
#------------------------------
# Prompt
#------------------------------
function precmd () {
  PROMPT=`~/go/bin/goprompt`
}
```
