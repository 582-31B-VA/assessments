# Project

> Weight: 20%\
> Due date: October 22

"Twenty-one" is the name given to a family of popular card games where
the aim is to score exactly twenty-one points or, failing that, to come
as close to twenty-one as possible. Number cards count as their number,
face cards (the jack, queen and king) count as 10, and aces count as 1
or 11.

The game is divided into four phases:

1. **Dealing phase**: First, two cards are dealt to the player, who is
   known as the "punter".

2. **Punter's turn**: The punter, having examined their cards, announces
   whether they will stay with the cards they have ("stand") or draw
   another card ("hit"). This phase is repeated until the punter stands,
   scores exactly twenty-one, or exceeds twenty-one. In the latter two
   cases, the game ends immediately.

3. **Banker's turn**: Once the punter has announced they will stay with
   their cards, the banker takes their turn. The banker is dealt cards
   until their hand achieves a total of 17 or higher. If the banker
   scores exactly twenty-one or exceeds twenty-one, the game ends.

4. **Resolving phase**: Whoever's hand is closer to twenty-one wins. If
   the score is even, the banker wins.

The game is played with a standard 52-card deck comprising 13 ranks in
each of the four suits: clubs, diamonds, hearts, and spades. Each suit
includes ten numeral cards, from one (ace) to ten, and three face cards:
king, queen and jack. The deck is shuffled _before_ the dealing phase.

## Requirements

In teams of two, create a command-line program to play twenty-one. When
executed, the program should print the punter's hand, and wait for them
to input "hit" to receive another card, or "stand" to stay with the
cards they have. The program should then print the banker's hand, and
declare a winner.

Here is an example of a full game:

```
--- Punter's Turn --------------------------------------------------------------

Your hand: [7H], [4C] (Total: 11)
Hit or Stand?
> Hit

Your hand: [7H], [4C], [6D] (Total: 17)
Hit or Stand?
> Stand

--- Banker's Turn --------------------------------------------------------------

Banker shows: [10S], [10C] (Total: 20)

--- Resolving Outcome ----------------------------------------------------------

Punter's hand: 17
Banker's hand: 20

Result: Banker wins!
```

Your repository includes a `main.py` file to get you started. It defines
a `Game` class that represents the game in progress, including the
`deck` of cards and the hands of both the `punter` and the `banker`.

The class also defines one method per phase: `dealing_phase`,
`punter_phase`, `banker_phase`, and `resolving_phase`. Their signature
is provided, but you must implement them. These methods must be pure,
meaning they cannot call `prompt` or `print`. Instead, they must return
a `PlayerStatus` enum that represents the game's status (ongoing, won or
lost) from the player's point of view.

The `punter_phase` method should take one argument: a callback function
named `hit` whose return value determines if the punter wants to hit or
not. When called, the callback function should receive the punter's
current score as its argument. This callback can be used to prompt the
player while keeping `punter_phase` pure.

Note that cards are always dealt from the _top_ of the deck. If you are
using a list to represent the deck of cards, dealing a card means
removing it from the _front_ of the list.

The `start` method is used to start a new game. It should shuffle the
deck, then call each phase method one after another. This is where most
of the printing should happen.

In addition to these instance methods, the `Game` class provides a class
method named `from_cards` which should construct a `Game` instance with
a deck composed of the specified cards _in order_. It should take a
single argument: a list of tuples, where each tuple is a card of the
form `(<rank>, <suit>)`, with `<rank>` being a number from 1 to 13, and
`<suit>` being either a number from 1 to 4 (clubs, diamonds, hearts,
spades). For instance, `(1, 1)` is the ace of clubs. Again, the method's
signature is provided, but you are responsible for its implementation.
You are not expected to call this method; it's only used for testing.

Your program should determine the value of aces based on the total value
of the hand. An ace is worth 11 unless it would cause the total to
exceed 21, in which case it is worth 1 instead.

## Starter files and submission

Only one student per team should [accept the assignment][Classroom].
That student can then give their teammate access to the repository
(Settings → Collaborators and teams → Manage access). Please add both
your names to the README file.

To make collaboration smoother and encourage code review, you will share
your code with your teammate using GitHub's pull request workflow (see
below). Note that only code on the remote `main` branch will be
assessed. Please do not close the "feedback" pull request.

[Classroom]: https://classroom.github.com/a/I_Vv-YL6

## Assessment criteria

- Program design [5]
  - requirements are met
  - program flow is decomposed into manageable, logical pieces
  - data structures are appropriate
  - common code is unified, not duplicated
  - appropriate algorithms are used, and coded cleanly
  - code is lint-free (see `uv run ruff check`)
  - no global variables

- Readability [3]
  - constants are used instead of hard-coded values
  - complex or meaningful expressions are named
  - naming is consistent and descriptive
  - inline comments are used to explain reasoning
  - documentation is used to explain interfaces
  - code is formatted (see `uv run ruff format`)
  - type hints are correct (see `uv run pyright`)

- Version control [1]
  - commits contain related changes
  - messages are consistent and informative
  - changes are merged using pull requests

- Teamwork [1]
  - workload is divided evenly
  - comments on pull requests are useful
  - peer assessment

## Pull request workflow

1. Make sure your local main branch is up to date with the remote main
   branch by running `git pull`.

2. From your local main branch, create a local feature branch with
   `git checkout -b <branch-name>`, where `<branch-name>` corresponds to
   the name of the feature you will implement.

3. Commit your changes with `git add <path>` and `git commit` (many
   times if necessary).

4. Once your code is ready, push your commits to GitHub with
   `git push -u origin <branch-name>`.

5. Using GitHub's web interface, create a pull request (Pull requests →
   New pull request → feature branch's name → Create pull request).

6. Let your teammate review your pull request. If necessary, make
   changes to your pull request by pushing new commits (repeat step #2
   and #3).

7. Let your teammate rebase and merge your pull request (Pull requests →
   pull request's name → Rebase and merge). If your pull request is
   out-of-date with main, GitHub will not let you merge it. In that
   case:

   1. Back in your terminal, switch to your local main branch with
      `git checkout main`.

   2. Sync your local main branch with the remote main branch using
      `git pull`.

   3. Switch back to your local feature branch with
      `git checkout <branch-name>`.

   4. Incorporate the latest changes from the main branch into your
      local feature branch with `git rebase main`.

   5. Fix merge conflicts, if any.

   6. Update your remote feature branch with `git push --force`.

8. Delete the remote feature branch (Code → main → view all branches →
   trash icon).

9. Back in your terminal, switch to your local main branch with
   `git checkout main`.

10. Sync your local main branch with the remote main branch using
    `git pull`.

11. Delete your local feature branch with `git branch -D <branch-name>`.

12. Repeat for every feature.
