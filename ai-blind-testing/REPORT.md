# Task 3 — AI blind test generation

> Homework — Task 3 (Lesson 6: File I/O, JSON and Testing).
> This task is **not auto-graded** by CI — it's reviewed by your mentor.
> The workflow only checks that you actually filled this file in
> (see the "Task 3 — AI report present" step in the Actions log).

## Instructions

1. Pick one function you wrote for Task 1 or Task 2 (e.g. `LoadTodos` or
   `ValidateEmail`).
2. Give an AI assistant **only the function signature** — no explanation
   of the logic, no existing test file. Ask it to generate a full
   table-driven test suite for that signature.
3. Run the AI-generated tests against your implementation.
4. Fill in the sections below.

---

## Function under test

`func LoadTodos(path string) ([]Todo, error)` from todo/todo.go

## Prompt you gave the AI

"Generate a comprehensive table-driven test suite for this Go function:

```go
func LoadTodos(path string) ([]Todo, error)
```

The function reads a JSON file and returns a slice of Todo structs. A Todo has ID (int), Title (string), Done (bool), and CreatedAt (time.Time) fields."

## Edge cases the AI found that you had missed

- Corrupted/truncated JSON file (file cut off mid-parse)
- File with wrong permissions (unreadable file)
- Very large JSON files
- JSON with extra/unknown fields in objects
- Array vs single object JSON structure mismatches

None of these exposed real bugs in the implementation—they're just extra defensive cases. The spec already covers malformed JSON and missing files.

## Edge cases you had that the AI missed

- Empty JSON file (specifically `""` vs `[]` or `null`)
- Files that exist but can't be read due to OS permissions
- Race conditions with concurrent file access

I think the AI missed these because it focused on "happy path" variations rather than OS-level error scenarios. Also, the empty file edge case is subtle—it's the difference between "no data" and "broken data."

## Cases where the AI's expected output was wrong

The AI suggested that an empty file should successfully return an empty slice `[]`, but based on the spec in the test file (`TestLoadTodos_EmptyFile`), an empty file should be treated as malformed JSON and return an error. This is a design decision, not a bug in the AI.

## What you'd change about your own test-writing process after this

I'd be more explicit upfront about the specification's intent. Telling the AI "empty file is an error" beforehand would have prevented the disconnect. Also, I'd trust the AI to generate volume and variety, but I need to review outputs against the actual spec and requirements—AI tests are good starting points but need context validation.
