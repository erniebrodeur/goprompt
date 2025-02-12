# GoPrompt

GoPrompt is a CLI tool for building a **customizable shell prompt** using a set of small, parallel “segments.” Whenever you open or refresh your shell, GoPrompt runs each segment (like Git branch detection or directory info) in parallel with a short timeout. This design ensures your prompt remains responsive even if a particular segment (for example, a Git check on a large repository) takes too long. The tool is written in Go, so it compiles to a single, lightweight binary without extra dependencies.

## Why Use GoPrompt?

Many shell prompts either rely on shell scripts that can block your prompt if something takes a while (like querying Git status in a large repo) or require complex setups. GoPrompt avoids these pitfalls by leveraging Go’s concurrency. Each segment has about 100ms to produce an output; if it fails or times out, GoPrompt simply omits that segment instead of slowing you down.

## Installation

To install GoPrompt, first ensure you have Go installed and set up on your system. Then clone this repository and build:

```bash
git clone https://github.com/your-username/go-prompt.git
cd go-prompt
go build -o goprompt
```

You can move the resulting `goprompt` binary into a folder on your `PATH`, or run `go install` directly if you prefer:

```bash
go install
```

Verify everything worked by running:

```bash
goprompt --help
```

## Basic Usage

At its core, GoPrompt generates a **single line** of text that includes information from various segments in a user-defined order or layout. If you simply type:

```bash
goprompt
```

it immediately returns your prompt string, using default settings for theme and layout. You can capture that string in your shell’s `PS1` variable or reference it however you like.

## Making It Permanent (Shell Integration)

To automatically apply GoPrompt to each new shell session, you can place an `eval` statement in your `~/.bashrc`, `~/.zshrc`, or similar shell config file. GoPrompt offers a `shell` subcommand that prints the recommended lines for you. For instance:

```bash
eval "$(goprompt shell)"
```

Putting that line in your RC file ensures every time you open a shell, `goprompt` is called and your prompt is set accordingly.

Alternatively, you can manually set your shell prompt to the output of `goprompt`, but calling the `shell` subcommand is a convenient shortcut for most setups.

## Themes and Layout

GoPrompt uses two environment variables to determine how the prompt looks:

- `GOTHEME` — Controls color schemes (e.g., `monokai_dark`, `solarized_light`, etc.).
- `GOLAYOUT` — Defines the arrangement of segments, for example:
  ```text
  "┌ $dir(2) on $git\n└─ $ "
  ```

Whenever you run `goprompt`, it reads these variables to decide how to format each segment’s output. If you want a permanent setup, you can export them in your shell config, for example:

```bash
export GOTHEME=monokai_dark
export GOLAYOUT="┌ $dir on $git\n└─ $ "
```

After editing your RC file, open a new shell or `source` the file to see the updated prompt.

## Example Prompt

With just those two variables set, you might get something like:

```
┌ ~/go-prompt on main
└─ $
```

If Git takes too long or you’re not in a Git repo, that segment simply doesn’t appear, and the rest of your prompt remains responsive.

## Concurrency Under the Hood

One of GoPrompt’s core strengths is its parallel architecture. Each segment runs in a separate goroutine under a ~100ms deadline. If it finishes in time, the segment’s data is included; if not, the prompt is generated without it—no stalling or partial hangs. This keeps your shell feeling snappy, even when you’re working in big repositories or running other background tasks.

## License

GoPrompt is released under the **GNU Lesser General Public License v2.0 (LGPL-2.0)**. You are free to use, modify, and distribute this project according to those terms. See the included [LICENSE](LICENSE) file for more information.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have ideas for new segments, bug fixes, or documentation improvements. We aim to keep GoPrompt simple, reliable, and fast—so we appreciate discussions about how any proposed changes fit into that goal.