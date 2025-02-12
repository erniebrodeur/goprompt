# GoPrompt TDD & Testing Approach

This document explains **how** we’re structuring tests in GoPrompt, following a test-driven development (TDD) mindset. We write test stubs and scenarios first, then implement the aggregator, segments, and CLI commands to satisfy those tests.

---

## 1. Overview of TDD Philosophy

We adopt TDD for GoPrompt to ensure each feature is well-defined and validated. This means:
1. **Write Tests First**: Create stub tests for new features or segments.
2. **Implement Just Enough**: Write the minimal code to pass these tests.
3. **Refactor & Extend**: Improve code structure, add new segments, or refine concurrency once tests confirm functionality.

---

## 2. Test Files & Directories

We use Go’s standard testing layout:

- **Aggregator Tests**: `internal/aggregator/aggregator_test.go`
- **Segment Tests**: Each segment in `internal/segments/` has a corresponding `_test.go` file (e.g., `git_segment_test.go`, `dir_segment_test.go`).
- **Integration Tests**: Optional higher-level tests (e.g., end-to-end from CLI to aggregator) can live in `cmd/goprompt/` or a dedicated test directory. We might name these `cli_integration_test.go` or similar.

Running `go test ./...` from the project root will discover and execute all these tests.

---

## 3. Aggregator Test Plan

### 3.1 Timeout Behavior

- **TestAggregatorTimeout**:  
  Simulate a segment that deliberately sleeps longer than 100ms. The aggregator should exclude it from the final prompt string without blocking overall output.

### 3.2 Ordering & Merging

- **TestAggregatorOrdering**:  
  Have multiple segments that finish in random order. Confirm we preserve the layout’s logical sequence in the final string.

### 3.3 Error Handling

- **TestAggregatorSegmentError**:  
  A segment returns an error. By default, we omit that segment from the final prompt (or produce an empty string). The aggregator itself shouldn’t crash or stall.

### 3.4 Basic Integration

- **TestAggregatorSimple**:  
  All segments complete quickly and successfully, producing a combined prompt matching the layout placeholders.

---

## 4. Segment Test Plan

### 4.1 GitSegment

- **TestGitSegmentInRepo**:  
  Create a temporary Git repo, check that `git_segment` returns the current branch name (e.g., `main`).
- **TestGitSegmentNoRepo**:  
  With no `.git` directory, we expect an empty string.
- **TestGitSegmentTimeout**:  
  Optionally simulate a slow Git operation to confirm aggregator timeouts, though this might be tested more thoroughly at the aggregator level.

### 4.2 DirSegment

- **TestDirSegmentBasic**:  
  If the user is in `~/projects/go-prompt`, ensure we get `go-prompt` or a truncated version depending on layout params (e.g., `$dir(2)`).
- **TestDirSegmentRoot**:  
  If the path is `/`, confirm we handle edge cases gracefully.

### 4.3 Additional Segments

Other segments (time, fill, host, etc.) follow a similar pattern:  
1. **Normal Scenario** (returns a meaningful string).  
2. **Edge / Missing** scenario (returns empty).  
3. **(Optional)** Timeout or error scenario.

---

## 5. CLI & Integration Tests

### 5.1 Subcommand Tests

We may write `_test.go` files to verify Cobra commands: `render`, `theme`, `layout`, `shell`. Each can be tested with typical Go subcommand test patterns—spawning the command in a test environment and checking outputs.

### 5.2 End-to-End

We could have an integration test that:
1. Sets environment variables (`GOTHEME`, `GOLAYOUT`).
2. Runs `goprompt` or `goprompt render`.
3. Verifies the returned string includes or excludes expected segments.

This helps ensure the entire pipeline (env reading → aggregator → layout parsing) behaves correctly in a real scenario.

---

## 6. Mocks & Utilities

For segments that rely on external I/O (like Git detection), we use:
- **Temporary Directories**: `t.TempDir()` to initialize a minimal `.git` structure.
- **Fakes/Stubs**: If needed, to replace slow or network-based calls with local stubs.

---

## 7. Coverage Goals

We aim for high coverage on aggregator logic, segment behaviors, and subcommand handling. At a minimum:
- **Aggregator**: Full coverage of concurrency/timeout paths.
- **Segments**: Tests for both common and edge cases.
- **CLI**: Basic coverage to ensure subcommands parse flags and call aggregator or environment lookups properly.

```bash
go test -cover ./...
```
will give us an overview.

---

## 8. Workflow Summary

1. **Create Test Cases**: Write `_test.go` with clear, descriptive function names.  
2. **Implement Minimal Code**: Just enough aggregator/segment code to pass tests.  
3. **Refactor**: Improve concurrency logic, layout parsing, or code clarity once tests are green.  
4. **Repeat**: For each new segment or feature, add tests first, then implement.

---

## Conclusion

This TDD/Test Doc outlines how we’ll systematically validate each piece of GoPrompt, from concurrency timeouts in the aggregator to segment-specific logic. By maintaining consistent, well-organized tests, we ensure that new features and refactors don’t break existing functionality. The final result is a robust, easily maintainable prompt system that meets our performance and user experience goals.