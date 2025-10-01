# Exam 3

> Weigth: 10%\
> Due: October 27

In computing, a file system is a hierarchical tree structure made of
**nodes**, each with its unique **name**. **Files** _are_ nodes that can
store **content**, whereas **directories** _are_ nodes that can have
**children**, each of which is either a file or a directory node.

Your task for this exam is to model this structure using three classes:
`Node`, `File` and `Directory`. Nodes are considered equal only if their
name is identical (see the `__eq__` special method). At no point can two
child of a directory can be equal.

In addition, record an 5–10 minutes video in which you explain how the
evaluation of the expression below unfolds according to the substitution
model covered in class.

```python
File("greet.txt", "hello world").name
```

Your explanation should walk through the process step by step. Make sure
both your code and your face remain visible for the entire video.

## Starter files and submission

To access the starter files, first [accept the assignment][Classroom],
then clone the GitHub Classroom repository on your computer.

To submit your assignment, push your changes to the `main` branch on the
`origin` remote. Code pushed on other branches will not be assessed.

[Classroom]: https://classroom.github.com/a/ZPd8SXi_

## Assessment criteria

- Program design [3]
  - requirements are met
  - tests pass (see `uv run pytest`)
  - code is lint-free (see `uv run ruff check`)
  - program flow is decomposed into manageable, logical pieces
  - data structures are appropriate
  - common code is unified, not duplicated
  - appropriate algorithms are used, and coded cleanly
  - no global variables

- Readability [3]
  - constants are used instead of hard-coded values
  - complex or meaningful expressions are named
  - naming is consistent and descriptive
  - code is formatted (see `uv run ruff format`)
  - documentation is correct (see `uv run pyright`)
  - inline comments are used to explain reasoning when necessary

- Version control [1]
  - commits contain related changes
  - messages are consistent and informative

- Video [3]
  - reasoning is clear, logical, and accurate
  - terminology is used appropriately
  - explanation proceeds step by step, without skipping key parts
