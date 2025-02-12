# GoPrompt CLI Doc

This document describes GoPrompt’s CLI layout and subcommands. By default, running `goprompt` with no subcommand **immediately** renders your prompt using the environment variables `GOTHEME` and `GOLAYOUT`. If those aren’t set, GoPrompt applies its built-in defaults.

---

## Default Behavior

- **Command**: `goprompt`
- **Effect**: Renders the prompt using the current theme and layout. If you haven’t set `GOTHEME` or `GOLAYOUT`, GoPrompt uses simple defaults (a minimal layout and no special color scheme).
- **Flags**:
  - `--help` or `-h` to display usage.

### Example
```bash
goprompt
# Immediately produces your prompt string
```

---

## Major Subcommands

### 1. `render`
- **Purpose**: Provide overrides for theme/layout on a single run.  
- **Usage**: `goprompt render [flags]`
- **Typical Flags**:
  - `--theme <name>` (e.g., `monokai_dark`)
  - `--layout <style>` (e.g., `single_line`, or a custom placeholder string)
- **Example**:
  ```bash
  goprompt render --theme monokai_dark --layout single_line
  ```

### 2. `layout`
- **Purpose**: Lists or applies built-in layouts.  
- **Usage**: `goprompt layout [command]`
- **Commands**:
  - `list` (shows known layouts, if any are built-in)
  - `apply <layoutName>` (switches to a named layout for this run, overriding `GOLAYOUT`)
- **Example**:
  ```bash
  goprompt layout list
  goprompt layout apply single_line
  ```

### 3. `theme`
- **Purpose**: Lists or applies built-in color themes like `monokai_dark` or `solarized_light`.  
- **Usage**: `goprompt theme [command]`
- **Commands**:
  - `list` (shows known themes)
  - `apply <themeName>` (overrides `GOTHEME` for this run)
- **Example**:
  ```bash
  goprompt theme list
  goprompt theme apply monokai_dark
  ```

### 4. `shell`
- **Purpose**: Prints lines for shell integration. By adding `eval "$(goprompt shell)"` to your `~/.bashrc` or `~/.zshrc`, you can set up GoPrompt automatically every time you start a new shell.
- **Usage**: `goprompt shell`
- **Example**:
  ```bash
  eval "$(goprompt shell)"
  ```

### 5. `help`
- **Purpose**: Shows help about any command.  
- **Usage**: `goprompt help [command]`
- **Example**:
  ```bash
  goprompt help render
  goprompt help shell
  ```

---

## Notes on Configuration

GoPrompt no longer has a dedicated `config` subcommand. Instead, it reads environment variables for long-term settings:

- **`GOTHEME`** for color themes.
- **`GOLAYOUT`** for prompt layout strings (e.g., `$dir`, `$git` placeholders).

You can temporarily override them by using `render`, `theme apply`, or `layout apply`, but for permanent changes, set `GOTHEME` and `GOLAYOUT` in your shell’s RC file.

For more details on environment variables, see [config_doc.md](./config_doc.md).

---

## Conclusion

- **No File-Based Config**: Everything is environment-based or provided on the command line.
- **Fast & Parallel**: Subcommands let you quickly test or override your theme/layout without editing your shell config each time.
- **Shell Integration**: Use the `shell` subcommand or place an `eval` in your `.bashrc` / `.zshrc`.

That’s the essence of the GoPrompt CLI.

---

