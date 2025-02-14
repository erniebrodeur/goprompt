# Segment BDD Test Design

This design focuses on **two primary scenarios** for each segment in GoPrompt:

1. **Normal Output**  
   The segment yields its intended text, e.g. `/home/user` for a directory segment.

2. **Error or Timeout**  
   If the segment fails or times out, it returns `[ERR]`.

## Ginkgo Structure

- **Describe** – `Describe("DirSegment")` or `Describe("GitSegment")`  
- **Context** – “when the directory is valid,” “when the directory read fails,” etc.  
- **It** – Each scenario has a separate `It` block reflecting the BDD doc:
  - Normal path → Expect segment text
  - Error path → Expect `[ERR]`

```plaintext
Describe("DirSegment", func() {
    Context("valid directory", func() {
        It("returns /home/user", func() {
            // BDD check: segment.Render() == "/home/user"
        })
    })
    Context("read fails", func() {
        It("returns [ERR]", func() {
            // BDD check: segment.Render() == "[ERR]"
        })
    })
})