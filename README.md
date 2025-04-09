# Driven

A to-do list app for grinders

---

Driven is a to-do list app I built for personal use with simplicity, speed and security
in mind. Driven tried to stay out of your way and let you focus on your tasks.

## 🌟 Features overview
- Access your weekly and daily task views
- Prioritize tasks with simple numeric levels
- Stay logged in for 2 weeks with secure access/refresh token auth
- Add notes to tasks for extra context
- Fast and responsive UX powered by HTMX
Try it out [here](https://mfirx7bkff.us-east-1.awsapprunner.com/)

## 🛠️ Some of the technologies I used
1. [Go](https://go.dev/) + [Echo](https://echo.labstack.com/): light weight and fast HTTP server.
2. [HTMX](https://htmx.org/): enables good interactivity with very small code.
3. [Templ](https://templ.guide/): fast Go-HTML compile-time template.
4. [TailwindCSS](https://tailwindcss.com/): CSS framework for consistent + clean UI

## 🧭 Usage Guide
### Getting Started
- On landing page, click "Get Started"
- Create an account with a username, email, and password
- Log in using your email and password
  (Passwords are securely hashed, of course!)

*Good NEWS, The login session is 2 weeks thanks do access/refresh token mechanism*

### Home page
This page is the main page you operate inside.

#### ➕ New tasks box
At the left side of the screen, you will the "New tasks" box. This is for creating
a new task. You can create task by simply putting the task name (don't include too
much detail. Driven has feature for that), the priority (default to 0) and task date
(default to today).

*Note that Driven don't try to name your priority but simply give it number. You should
be the one deciding what metrics to gauge your task priority*

At the top bar center, you will see that there are 2 modes, "week" and "to day".

#### 📅 Week view
Driven represent your tasks normally in the week view, so you can see what you have
to do in an entire week. You can try navigating the week using the ">" and "<" at the
top of calendar.

#### ☀️ Day view
Now if you click at "to day" mode or click at any task in the week mode, you will be
navigated to another format of the task which only focus on the one day. You can
- navigate each day using either
  1. the ">" and "<" at the top
  2. the quick navigation tool at the bottom left of the screen

Driven also enables user to note any key information yourself need to know when you do that task.
You can simply write on the textarea and click "save".

(For the sake of simplicity, in the day view, you cannot create a task outside of that day
because Driven assumes that you are focusing on that day.)

#### Profile
At the top right in the top bar. You can sign out.

## Contribution
You are more than welcome to contribute to my project. Please clone + branch and
open a pull request for me to review. You can debug/enhance anything in the [issues](https://github.com/WaronLimsakul/Driven/issues)
section or come up with any features that you think will help boosting our productivity.

### Local Development Setup
#### Prerequisites
  - Go 1.24+
  - PostgreSQL (or update the config to match your setup)
  - [Tailwind CLI](https://tailwindcss.com/docs/installation/tailwind-cli)

#### Env setup
```ini
ENV=development
PORT=1234

DB_URL=postgres://...

GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR=sql/schema
GOOSE_DBSTRING=postgres://...

JWT_SECRET=randome_some_string
```
