# LIA

> Weight: 50%\
> Due: November 12\
> Modality: in teams of 2

A social news website is a website that features user-posted stories
where users can comment on news posts. [Lobsters] and [Hacker News] are
popular examples of social news websites geared toward programming.

For this course's Learning Integration Assessment (LIA), you will use
Flask to build your own social news website called _Shrimp_. Here are
the requirements:

- The website's navigation includes a link to a page where users can
  create a story. A story is a news link to another website. It also
  includes a title and an optional description. The date at which the
  story was posted is recorded, as well as the account of the user who
  posted it. Once their story has been created, users are redirected to
  the story's detail page (see below).

- No two stories can share the same news link. Titles may not be blank
  and may contain up to 80 characters. Descriptions may contain up to
  10,000 characters.

- On the home page, users can see a list of all the stories sorted from
  most to least recent. It shows each story's title, the name of the
  account who shared it, the news link, and when it was posted.

- From the home page, users can click on a story to go to another page
  where its description is shown. This detail page also lists the
  comments for the story sorted from oldest to newest.

- On the detail page, there is a form to leave a new comment. A comment
  has a body, which may contain up to 10,000 characters. We also store
  its date of creation, and the account of the user who authored it.

- Users are required to create an account to post a story or leave a
  comment. The website's navigation includes a link to a registration
  page where users can create an account. Accounts have a unique
  username (which may not be blank, may not include whitespace, and may
  contain up to 50 characters), a valid and unique email, and a password
  (minimum 8 characters, maximum 20, including at least one digit and
  one uppercase letter).

- Users that are logged in don't see the link to the registration page.
  Instead, the website's navigation includes a link to their account
  page, and a link to log out.

- On their account page, users can change their email and password.

- When users click the name of the author of a story or comment, they go
  to that account's profile page. The profile page shows the name of the
  account, and lists all its stories and comments.

- A flash message is displayed when users create an account, when they
  change their email or password, when they log in or log out, and when
  they leave a comment on a story.

[Lobsters]: https://lobste.rs/
[Hacker News]: https://news.ycombinator.com/news

## Starter files and submission

Only one student per team should [accept the assignment][Classroom].
That student can then give their teammate access to the repository
(Settings → Collaborators and teams → Manage access). Please add both
your names to the README file.

To make collaboration smoother and encourage code review, you will share
your code with your teammate using GitHub's pull request workflow (see
below). Note that only code on the remote `main` branch will be
assessed. Please do not close the "feedback" pull request.

[Classroom]: https://classroom.github.com/a/4YNO0cJn

## Assessment criteria

- Program design [5]
  - requirements are met
  - program flow is decomposed into manageable, logical pieces
  - data structures are appropriate
  - common code is unified, not duplicated
  - appropriate algorithms are used, and coded cleanly
  - markup is semantic
  - code is lint-free (see `uv run ruff check`)
  - project is well-structured, as seen in class
  - separation of concerns (handler, model, template) is respected
  - no global variables

- Readability [3]
  - constants are used instead of hard-coded values
  - complex or meaningful expressions are named
  - naming is consistent and descriptive
  - inline comments are used to explain reasoning
  - documentation is used to explain public interfaces
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
