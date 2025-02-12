# GoPrompt Comprehensive CLI Help

This file shows sample output a user might see when running `goprompt --help` or `goprompt help <subcommand>`. References to any `config` subcommand have been removed; we now rely on environment variables or the subcommands below.

---

## Top-Level Help

```
$ goprompt --help
GoPrompt is a CLI tool that builds a customizable prompt by running multiple segments in parallel.

USAGE:
    goprompt [flags]
    goprompt [command]

DESCRIPTION:
    - Running 'goprompt' alone immediately renders your prompt using
      the GOTHEME and GOLAYOUT environment variables (or defaults).
    - Subcommands let you temporarily override theme/layout or list
      built-in layouts/themes.

EXAMPLES:
    goprompt
    goprompt render --theme monokai_dark --layout single_line
    goprompt layout list
    goprompt shell     # Prints shell integration instructions

AVAILABLE COMMANDS:
    render   Renders the prompt with optional flags for theme/layout
    layout   Lists or applies known layouts (single_line, multi_line, etc.)
    theme    Lists or applies built-in color themes
    shell    Prints shell integration lines
    help     Help about any command

FLAGS:
    -h, --help   Display help information

Use "goprompt [command] --help" for more information about a command.
```

---

## Subcommand: `render`

**Help Output**:
```
$ goprompt render --help
Renders the GoPrompt output, allowing quick flags for theme or layout.

USAGE:
    goprompt render [flags]

FLAGS:
        --theme <themeName>     Temporarily override GOTHEME
        --layout <layoutStyle>  Temporarily override GOLAYOUT
    -h, --help                 Display help information for 'render'

EXAMPLES:
    goprompt render --theme monokai_dark
    goprompt render --layout single_line
    goprompt render --theme solarized_light --layout multi_line
```

---

## Subcommand: `layout`

**Help Output**:
```
$ goprompt layout --help
Manages prompt layout (e.g., single_line, multi_line).

USAGE:
    goprompt layout [command]

AVAILABLE COMMANDS:
    list   Show all known layouts
    apply  Apply a layout name for this session

FLAGS:
    -h, --help   Display help information

EXAMPLES:
    goprompt layout list
    goprompt layout apply single_line
```

---

## Subcommand: `theme`

**Help Output**:
```
$ goprompt theme --help
Manages built-in color themes.

USAGE:
    goprompt theme [command]

AVAILABLE COMMANDS:
    list   Show all known themes
    apply  Apply a theme for this session

FLAGS:
    -h, --help   Display help for 'theme'

EXAMPLES:
    goprompt theme list
    goprompt theme apply monokai_dark
```

---

## Subcommand: `shell`

**Help Output**:
```
$ goprompt shell --help
Prints shell integration lines so you can source them in your .bashrc or .zshrc.

USAGE:
    goprompt shell

DESCRIPTION:
    Typically used as:
        eval "$(goprompt shell)"

    This command prints a snippet that runs 'goprompt' for each new
    shell session, using your environment variables (GOTHEME, GOLAYOUT).

FLAGS:
    -h, --help   Display help
```

---

## Subcommand: `help`

**Help Output**:
```
$ goprompt help --help
Help about any command.

USAGE:
    goprompt help [command]

FLAGS:
    -h, --help   Display help for 'help'

EXAMPLES:
    goprompt help render
    goprompt help shell
```

---

## Summary

This updated reference no longer includes a `config` subcommand. Environment variables (`GOTHEME` / `GOLAYOUT`) or command-line flags handle configuration. The `shell` subcommand assists with integrating GoPrompt into your shell startup files.