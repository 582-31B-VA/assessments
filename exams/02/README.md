# Exam 2

> Weigth: 10%\
> Due: October 6

For this exam, your task is to design and implement three classes for an
online ordering system: `Order`, `Item` and `Status`. Each order has a
`status` (initially "processing", then "shipped" or "delivered"), and a
list of `items` provided when the order is created. Items have a `name`
and a `price`. The price is stored in cents and must be positive. In
addition to its `status` and `items`, we can get an order's `total` in
cents. We can also get a `status_message` that corresponds to its
current status. Here are the messages:

- We are processing your order.
- Your order has been shipped.
- Your order has been delivered.

Finally, record an 5–10 minutes video in which you explain how the
evaluation of the expression below unfolds according to the substitution
model covered in class.

```python
Item("Shirt", -100)
```

Your explanation should walk through the process step by step. Make sure
both your code and your face remain visible for the entire video.

## Starter files and submission

To access the starter files, first [accept the assignment][Classroom].
Then, clone your repository using `git clone <url>`, where `<url>` is
the link provided to you after accepting the assignment.

To submit your assignment, push your changes to the "main" branch on the
"origin" remote. Code pushed on other branches will not be assessed.

[Classroom]: https://classroom.github.com/a/ecfk6Bc1

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
