# GoPrompt Overview

> **Internal Document**  
> This is the **entry point** for GoPrompt’s **core design**—covering purpose, constraints, and high-level usage.  
> Deeper technical details live in references like [data_domain.md](./data_domain.md), [module_diagram.md](./module_diagram.md), and [theming_doc.md](../../design/references/theming_doc.md).

## 1. Purpose

GoPrompt is a CLI for customizable, **fast** prompts on Unix-like shells. It runs multiple **segments** in parallel (aggregator concurrency), each returning a piece of info (Git branch, current directory, user, etc.). After a short timeout (~100ms), it merges all segments into a final prompt string.

## 2. Key Concepts

- **Segments**: Small Go structs that detect environment info (Git, directory, user). They can return `[ERR]` if something goes wrong.  
- **Aggregator**: Spawns each segment in a goroutine, enforces a timeout, and assembles final output.  
- **Theming**: Uses [theming_doc.md](../../design/references/theming_doc.md) approach with hex color codes (`#RRGGBB`) to produce 24-bit ANSI coloring for each segment/state.  
- **Layout**: A functional string that references placeholders like `$git`, `$dir(2)`, etc., letting users shape a single or multi-line prompt.

## 3. Configuration

- **No “config” Subcommand**: We rely on environment variables for permanent defaults (`GOTHEME`, `GOLAYOUT`).  
- **One-Off Overrides**: Run `goprompt render --theme monokai_dark --layout "..."` to override defaults for a single invocation.

## 4. Cobra CLI

- **Root**: Running `goprompt` alone renders the prompt using current defaults.  
- **Subcommands**:  
  - `render`: Temporary theme/layout overrides.  
  - `layout`: List or apply known layouts.  
  - `theme`: List or apply built-in themes.  
  - `shell`: Print shell integration lines.  

See [cli_doc.md](./cli_doc.md) and [cli_help_output.md](./cli_help_output.md) for the latest command details.

## 5. Internal Docs

- **[data_domain.md](./data_domain.md)**: How aggregator, segments, and theming data flow.  
- **[module_diagram.md](./module_diagram.md)**: Code organization and how packages fit together.  
- **[theming_doc.md](../../design/references/theming_doc.md)**: Hex-based color approach for multiple states per segment.  
- **[tdd_doc.md](./tdd_doc.md)**: Testing strategy, concurrency checks, etc.

## 6. Release & Usage

While Moya remains internal, GoPrompt itself can be open-sourced. We plan to meet best OSS practices (license, readme, etc.). For final user-facing docs or repos, see your local repository’s top-level README and any open source guidelines.

---

**This updated overview ensures no leftover references to a “config” command** and clarifies environment-based configuration plus the new references. Once we’re ready, we can store it as `overview_doc.md` in `library/projects/goprompt/`.
