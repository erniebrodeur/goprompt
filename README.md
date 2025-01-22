# goprompt

A customizable, segment-based command prompt in Go, featuring truecolor support via [termenv](https://github.com/muesli/termenv), modular segments (directory, Git, user, time, etc.), and an extensible theming system.

## Features

- **Segment-Based Design**  
  Each segment is an independent piece of logic (e.g., directory info, Git status).
- **Truecolor & Theming**  
  Uses `termenv` to detect terminal capabilities and apply 24-bit color.
- **Modular & Configurable**  
  Easily reorder or disable segments, define custom themes, and extend with new segments.
- **Performance-Oriented**  
  Uses [`go-git`](https://github.com/go-git/go-git) to check repo status without spawning extra processes.

## Installation

```bash
go install github.com/erniebrodeur/goprompt/cmd/goprompt@latest
```

> **Note:** If you prefer a specific version, replace `@latest` with a [tag](#versioning--releases), like `@v0.1.0`.

## Usage

**Integrate with your shell** by calling the `goprompt` binary in the prompt initialization. For example, in **Zsh**:

```bash
function precmd() {
  PROMPT="$(goprompt)"
}
```

In **Bash**:

```bash
PS1='$(goprompt)'
```

## Segments

- **Directory**: Shortens the current path to `.../<last>/<dir>` and supports `~/` home replacement.  
- **Git**: Displays the current branch and adds `*` if there are uncommitted changes.  
- **User**: Shows current username, indicates root user, and appends `@host` if via SSH.  
- **Time/Date**: Splittable into separate segments (e.g., `time` vs. `date`) so you can format or color them independently.

## Theming

`goprompt` supports a basic theme interface for coloring each segment. By default, it assigns one color per segment (e.g., directory = blue, Git = yellow). You can:

1. Define your own theme by implementing the `theme.Theme` interface.
2. Use environment variables or config files (in progress) to override colors.
3. Color the filler line (`─────`) separately from segments if desired.

## Configuration

(Planned) The system can load a YAML/JSON config file specifying:

- **Segment Order**  
- **Theme Colors**  
- **Enable/Disable** segments

Future updates will detail how to provide a `config.yaml` to reorder and recolor segments dynamically.

## Versioning & Releases

`goprompt` follows [Semantic Versioning](https://semver.org/). You can install specific tagged versions:

```bash
go install github.com/erniebrodeur/goprompt/cmd/goprompt@v0.1.0
```

For unreleased work, you can pin a commit hash or use `@latest`.

## Contributing

1. **Fork** the repo and clone locally.  
2. **Create** a feature branch.  
3. **Commit** and push your changes, then **open a PR**.  
4. If you add a segment or theme, please **include** a short example in the docs.

## License

This project is licensed under the [MIT License](LICENSE).
