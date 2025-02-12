# GoPrompt Implementation Doc

This document explains how GoPrompt assembles a final prompt string using **parallel segment execution**, how it handles **layouts**, and how each segment integrates. It’s intended for developers who want to understand or modify the internal code structure.

---

## 1. Overview

GoPrompt’s core functionality revolves around a **central aggregator** that collects data from various “segments” (like Git branch info, current directory, etc.). Each segment runs under a short timeout (~100ms) using Go’s concurrency. If a segment responds in time, its result is included in the prompt; if not, GoPrompt omits it without blocking.

We also rely on two environment variables, `GOTHEME` and `GOLAYOUT`, to determine how results are styled (theme) and arranged (layout). When a user runs `goprompt`, we read these variables, parse the layout string, spawn each segment in parallel, then merge their outputs to produce the final prompt line.

---

## 2. Aggregator and Concurrency

### 2.1 Context & Timeout
- We create a `context.WithTimeout` (usually 100ms) at the start of prompt generation.  
- Each segment is invoked with this context so it can check `ctx.Err()` if time runs out or if the aggregator cancels.

### 2.2 Segment Execution
- The aggregator stores a list of segments, each with a `Render(ctx context.Context) (string, error)` method.
- We launch all segments concurrently (one goroutine per segment).  
- Once the timeout expires or all segments complete, we collect the results.

### 2.3 Ordering and Merging
- Even though segments run in parallel, we maintain **logical order** based on how they appear in the layout.  
- If a segment times out or returns an error, we treat it as an empty string.  
- We then join the outputs into a single string (potentially with line breaks or other separators) as specified by the layout.

---

## 3. Segments

Each “segment” is a small Go struct or function that implements:
```go
type Segment interface {
	Render(ctx context.Context) (string, error)
}
```

**Examples**:  
- **GitSegment**: Checks the current repo’s branch/dirtiness using `go-git`, returning a short string like `main*`.  
- **DirSegment**: Returns the last `n` directories of the path (e.g., `~/projects/moya`), possibly truncated.  
- **TimeSegment**: Returns the current time with a specified format.  

Segments must finish quickly—if they exceed the context deadline, their result is omitted.

---

## 4. Layout Parsing

### 4.1 Reading `GOLAYOUT`
- We read `GOLAYOUT` (if set) or fall back to a default (e.g., `"$dir $git"`).  
- The aggregator scans the layout string for **placeholders** (e.g., `$dir(2)`, `$git`, `$time("%H:%M")`), plus any literal text (ASCII separators, line breaks, etc.).

### 4.2 Mapping Placeholders to Segments
- For each placeholder, we look up or instantiate the corresponding segment.  
- For instance, `$dir(2)` means “call DirSegment with the parameter 2,” `$git` means “call GitSegment,” etc.  
- We store them in an ordered slice, matching the sequence they appear in the layout string.

### 4.3 Substitution
- After each segment completes, we substitute its result back into the layout string, preserving the original placeholder positions.  
- If a segment is missing or timed out, we insert nothing (or `""`).

---

## 5. Themes (GOTHEME)

### 5.1 Theming Approach
- If `GOTHEME` is defined, we pass a “theme object” or name to each segment so it can wrap text in ANSI color codes.  
- If `GOTHEME` is empty, segments produce uncolored output.

### 5.2 Minimal Integration
- We keep theming optional; each segment checks if color is requested.  
- No advanced templating is required—just some ANSI codes if desired.

---

## 6. Putting It All Together

1. **Read Environment**: Check `GOLAYOUT` and `GOTHEME`.  
2. **Parse Layout**: Identify placeholders (segments) plus any literal text.  
3. **Aggregator**: Create a `context.WithTimeout(100 * time.Millisecond)`, spin up each segment’s `Render` in a goroutine.  
4. **Collect Results**: Wait for all goroutines or for the timeout.  
5. **Assemble Prompt**: Merge each segment’s result into the layout’s placeholders in order.  
6. **Return String**: Print to stdout if run directly, or pass to a shell integration if using `goprompt shell`.

---

## 7. Code Structure (Suggested)

- **`cmd/goprompt/`**: Main entry point with Cobra or a similar CLI framework.  
- **`internal/aggregator/`**: Houses the aggregator logic (context, concurrency, result merging).  
- **`internal/segments/`**: Each segment in its own file (e.g. `git_segment.go`, `dir_segment.go`).  
- **`internal/theme/`**: Optional directory for theme-handling code, if we want more than just a few color mappings.  
- **`main.go`** or `cmd/goprompt/main.go` uses Cobra commands, references aggregator.

This layout keeps the concurrency logic distinct from the CLI parsing, making it easier to test each piece.

---

## 8. Testing

We rely on Go’s standard testing with `_test.go` files:

1. **Aggregator Tests**  
   - Confirm timeouts work correctly (e.g., simulate a slow segment).  
   - Verify the final string respects segment order.

2. **Segment Tests**  
   - Each segment has a small `_test.go` verifying it returns the expected string for normal cases, and an empty string or error for a problem or timeouts.  
   - Possibly mock external calls (like Git) using temp directories.

3. **Integration Tests**  
   - Optionally test the full chain: parse layout, run aggregator, confirm final output.

---

## 9. Next Steps

- **Implement** the aggregator, define a `RenderSegments(context.Context, []Segment) string` or similar function.  
- **Create** a set of common segments (DirSegment, GitSegment, etc.).  
- **Connect** it to the CLI subcommands: `render`, `shell`, etc.  
- **Iterate** on theming if needed.

---

## Conclusion

GoPrompt’s implementation relies on a clean separation between **CLI commands** and the **aggregator** that executes segments concurrently. By using environment variables and a straightforward segment interface, we keep it simple yet performant. Each developer should find it easy to add or modify segments, test concurrency, and expand theming as the project evolves.