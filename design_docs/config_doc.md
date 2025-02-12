# GoPrompt Config (Environment Variables)

This document explains how GoPrompt handles **themes** and **layouts** using simple environment variables—no file-based config or commands required.

## Environment Variables

### `GOTHEME`
- Specifies the color scheme for segments.
- Example: `export GOTHEME=monokai_dark`

### `GOLAYOUT`
- Controls the arrangement or placeholders for each segment.
- Example:  
  ```bash
  export GOLAYOUT="┌ $dir on $git\n└─ $ "
  ```
  Here, `$dir` and `$git` are placeholders for their respective segments.

## Defaults

If `GOTHEME` or `GOLAYOUT` is **unset**, GoPrompt uses built-in defaults:
- A simple color scheme (or no colors).
- A minimal layout with just directory info.

## Updating the Shell

To make changes permanent, add these exports to your shell’s RC file (`~/.bashrc`, `~/.zshrc`, etc.):
```bash
export GOTHEME=solarized_light
export GOLAYOUT="─ $dir(1) : $git : $ "
```
After saving, open a new shell or run `source ~/.bashrc` (or equivalent) to apply.

## Shell Subcommand

GoPrompt also provides a `shell` subcommand that prints lines to integrate the prompt automatically into your shell. Adding:
```bash
eval "$(goprompt shell)"
```
to your RC file calls GoPrompt whenever a new shell starts, generating a prompt using your chosen environment variables.

## Summary

- **No config file** or dedicated `config` command.
- **Use environment variables** to define theme/layout.
- **Optional** `shell` subcommand for quick integration.

That’s all there is to configuring GoPrompt. By staying environment-driven, setup remains simple and easily portable between different machines or shell configurations.