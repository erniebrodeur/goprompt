# Go BDD Testing Strategy (New)

The following outlines our **brand-new** BDD approach for GoPrompt, focusing on concurrency, partial failures, and a guaranteed fallback prompt (`$dir() %`) if all else fails. This design ensures we never have a complete meltdown; partial or total failure simply yields `[ERR]` for individual segments or a minimal prompt overall.

## 1. Aggregator Concurrency & Fallback

1.1 **Single Segment (Happy Path)**
- Given an aggregator with one quick segment
- When we collect results
- Then the segment output appears in the final prompt
- And there is no `[ERR]`

1.2 **Single Segment (Timeout)**
- Given an aggregator with one slow segment
- When we collect results
- Then that segment yields `[ERR]`

1.3 **Multiple Segments (Mixed Speeds)**
- Given multiple segments (some quick, some slow)
- When we collect with a short timeout
- Then quick segments succeed, slow ones yield `[ERR]`

1.4 **Complete Fallback (All Fail)**
- Given all segments fail or time out
- When the aggregator collects results
- Then it falls back to `$dir() %`

## 2. Directory Segment

2.1 **Truncating the Path**
- Given the current directory is `/long/path/to/project`
- When `ShowComponents` is 2
- Then the output is `to/project`

2.2 **Home Directory -> `~`**
- Given the user is in `$HOME`
- When we render DirSegment
- Then it shows `~`

2.3 **Subdirectories Under Home**
- Given `$HOME/code`
- When `ShowComponents` is 1
- Then output might be `~/code`

2.4 **Dir Segment Fails**
- Given a failure mocking `os.Getwd()`
- When DirSegment tries to render
- Then it returns `[ERR]`

## 3. Git Segment

3.1 **Clean Repository**
- Given a clean Git repo on `main`
- When GitSegment renders
- Then it shows the branch name

3.2 **Dirty Repository**
- Given uncommitted changes
- When GitSegment renders
- Then it appends a marker (e.g., `*`)

3.3 **Detached HEAD**
- Given a commit checkout
- When GitSegment renders
- Then it shows `[DETACHED]`

3.4 **Timeout or Error**
- Given forced errors or slow operation
- When aggregator runs
- Then GitSegment returns `[ERR]`

## 4. User Segment

4.1 **Normal User**
- Given user is non-root
- When UserSegment renders
- Then it shows the username

4.2 **Root User**
- Given UID=0
- When UserSegment renders
- Then it displays `root` or a distinct symbol

4.3 **Failure or Timeout**
- Given a mocked error in user lookup
- When aggregator runs
- Then the segment yields `[ERR]`

## 5. Time Segment

5.1 **Default Format**
- Given no custom format
- When TimeSegment renders
- Then it shows a default time like `HH:MM`

5.2 **Custom Format**
- Given a format like `%H:%M` or `"15:04"`
- When TimeSegment renders
- Then output matches that format

5.3 **Error or Delay**
- Given forced error or delay
- Then it yields `[ERR]`

## 6. Layout & Theming

6.1 **Valid Layout Placeholders**
- Given a layout `"$dir(2) $git"`
- When aggregator merges outputs
- Then placeholders are replaced with segment results

6.2 **Unknown Placeholder**
- Given `"$bogus(99)"` in the layout
- When aggregator parses
- Then we see literal text or `[ERR]` for that placeholder

6.3 **Theme Application**
- Given a theme map with `git.dirty`, `dir.home`, etc.
- When segments reference those keys
- Then colored output appears
- Else fallback is no color or `[ERR]`

6.4 **All Layout or Theme Fails**
- Given invalid placeholders or missing theme keys for every segment
- Then aggregator reverts to `$dir() %`

## 7. Environment & CLI

7.1 **Default Environment**
- Given `$GOTHEME` and `$GOLAYOUT` are unset
- When the user runs `goprompt`
- Then it defaults to `$dir() %`

7.2 **Overriding with `render` Flags**
- Given environment variables
- When user passes `--theme alt_theme`
- Then aggregator uses that theme for one run

7.3 **Invalid Flags or Subcommands**
- Given user calls `goprompt --bogus`
- Then an error is shown, but fallback `$dir() %` remains

7.4 **`theme list` and `layout apply`**
- Lists known themes or layouts
- If unknown, yields `[ERR]` but prompt still works

7.5 **Shell Integration**
- `goprompt shell` prints instructions
- If something fails, fallback is `$dir() %`

## 8. Integration (Optional)

- Given multiple segments plus concurrency
- When aggregator runs with partial or total failure
- Then partial `[ERR]` or `$dir() %` if everything bombs

---
These new scenarios define how each domain is tested under the fallback-first architecture. Each segment or CLI path uses `[ERR]` on failure without jeopardizing the entire prompt, fulfilling the principle: **always** a working prompt.