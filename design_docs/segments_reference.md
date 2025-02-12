# GoPrompt Segments Reference

This document describes the **built-in segments** for GoPrompt, detailing what they do, which parameters (if any) they accept, and how they handle timeouts or errors. Since the project is **feature-frozen**, only these segments are officially supported at this time.

---

## 1. Shared Behavior

- **Concurrency & Timeout**: Each segment runs in its own goroutine, with about a 100ms deadline. If it doesn’t produce a result in time, GoPrompt omits it rather than blocking.
- **Error Handling**: Non-critical errors (like “not in a Git repo”) yield an empty string. Serious internal failures return an error, which typically causes the segment to be omitted.
- **Theming**: If `GOTHEME` is set, segments may wrap their outputs in ANSI color codes. If no theme is set, they typically return plain text.

---

## 2. `dir` Segment

**Purpose**  
Displays the current working directory (or a truncated version if you specify a parameter).

**Placeholder**  
```
$dir
```
Optionally, you can pass a numeric parameter in parentheses to indicate how many levels of the path to show:
```
$dir(2)
```
For instance, if your directory is `~/projects/go-prompt`, `$dir(2)` might return `projects/go-prompt`, while `$dir` alone might return the full path.

**Edge Cases**  
- If the user’s directory is `/`, `$dir(2)` returns `/`.  
- If the user is in a very deep path, `$dir(n)` simply shows the last `n` components.

---

## 3. `git` Segment

**Purpose**  
Displays Git repository information (branch name, dirty indicator, etc.) using [`go-git`](https://github.com/go-git/go-git).

**Placeholder**  
```
$git
```

**Behavior**  
- Returns the current branch name if inside a Git repo (e.g., `main`, `feature-xyz`).  
- May optionally show a `*` or similar symbol if the repo is “dirty.”  
- If not in a Git repo, it returns an empty string.

**Timeout Handling**  
If Git scanning takes too long (large repo, slow I/O), the segment times out and is omitted.

---

## 4. `time` Segment

**Purpose**  
Displays the current local time in a user-specified format.

**Placeholder**  
```
$time("<format>")
```
For instance:
```
$time("%H:%M")
```
might output `14:05` for 2:05 PM in 24-hour format. We use [`time.Format`](https://pkg.go.dev/time#Time.Format) patterns or strftime-like placeholders, depending on how we implement it.

**Edge Cases**  
- If no format is provided, we might use a default like `15:04` (hh:mm 24-hour).  
- If time can’t be retrieved for some reason (rare), it returns empty.

---

## 5. `fill` Segment

**Purpose**  
Produces a repeated character or pattern for alignment, horizontal rules, etc.

**Placeholder**  
```
$fill(<length>[,char='<c>'])
```
**Examples**:
1. `$fill(5)` → `     ` (5 spaces).  
2. `$fill(10,char='-')` → `----------` (10 dashes).

**Behavior**  
- Reads the numeric argument to know how many times to repeat.  
- If `char` is omitted, defaults to space.  
- You can use this for fancy ASCII separators or line-filling in multi-line layouts.

---

## 6. `host` Segment

**Purpose**  
Displays hostname info or indicates if the session is remote (e.g., over SSH).

**Placeholder**  
```
$host
```
**Behavior**  
- Typically calls `os.Hostname()` or checks SSH environment variables.  
- If you’re on a local machine, it may show just the short hostname. If you’re on a remote session, it might say `remote@hostname`, depending on your config.

**Edge Cases**  
- If hostname detection fails, returns empty.

---

## 7. `user` Segment

**Purpose**  
Shows the current username, often highlighting if it’s `root`.

**Placeholder**  
```
$user
```
**Behavior**  
- Returns your local username (`alice`, `bob`, etc.).  
- Could highlight `root` in a different color if a theme is active.  
- If username is unavailable, returns empty.

---

## 8. (Optional) `ruby` / `go` Segments

If you’ve seen references to a `ruby` or `go` segment in older docs:

- **`ruby`**: Might detect Ruby version or Gemfile if you’re in a Ruby project.  
- **`go`**: Similar for Go modules, returning something like `go1.19`.  

Since we’re **feature-frozen** for now, these are not guaranteed to exist or be fully supported. If you want them, you may need to check or implement them locally.

---

## 9. Usage in Layout

When writing `GOLAYOUT`, you can combine these placeholders with literal text, line breaks, or ASCII separators. For example:

```
export GOLAYOUT="┌ $dir(2) on $git\n└─ $ "
```
Here, `$dir(2)` and `$git` are replaced by their respective segment outputs, with a line break (`\n`) between lines.

**Another Example**:
```
export GOLAYOUT="[$time(\"%H:%M\")] $user@$host : $fill(5,char='*') : $git"
```
This might produce:
```
[14:05] alice@myhost : ***** : main*
```