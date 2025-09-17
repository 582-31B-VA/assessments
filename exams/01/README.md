# Exam 1

> Weigth: 10%\
> Due: September 22 (before class)

In the `main.py` file included with your repository, create a `Zombie`
class that models a zombie in a 2D game. The class should define the
attributes listed below. It is up to you to determine which ones are
instance variables and which ones are class variables.

- `x`: the zombie's horizontal coordinate.
- `y`: the zombie's vertical coordinate.
- `hp`: the zombie's health points (all zombies start with 3 health
  points).
- `count`: the total number of zombies created.

In addition to the constructor, the class should also define the
following methods:

- `hit`: substracts 1 from the zombie's health points and returns the
  zombie.
- `is_alive`: returns `True` if the zombie has more than 0 health
  points.
- `from_list`: returns a list of `Zombie` objects created from a list of
  coordinate tuples (e.g., `[(1, 2), (3, 3)]` produces one `Zombie` with
  `x` set to `1` and `y` set to `2`, and another with `x` set to `3` and
  `y` set to `3`).

It is up to you to determine whether to use an instance method, a class
method or a static method.

Finally, record an 5–10 minutes video in which you explain how the
evaluation of the expression below unfolds according to the substitution
model covered in class. Your explanation should walk through the process
step by step. Make sure both your code and your face remain visible for
the entire video.

```python
Zombie(1, 2).is_alive()
```

## Starter files and submission

To access the starter files, first [accept the assignment][Classroom].
Then, clone your repository using `git clone <url>`, where `<url>` is
the link provided to you after accepting the assignment.

To submit your assignment, push your changes to the `main` branch on the
`origin` remote. Code pushed on other branches will not be assessed.

[Classroom]: https://classroom.github.com/a/9dAcagJm

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
