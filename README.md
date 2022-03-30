<div align="center">

# letovoCLI

letovo is a fancy tool for accessing your [school](https://s.letovo.ru/) data.

**README Sections:** [Options](#options) — [Installation](#installation)

</div>


[//]: # (---)

[//]: # ()

[//]: # (**exa** is a modern replacement for the venerable file-listing command-line program `ls` that ships with Unix and Linux)

[//]: # (operating systems, giving it more features and better defaults. It uses colours to distinguish file types and metadata.)

[//]: # (It knows about symlinks, extended attributes, and Git. And it’s **small**, **fast**, and just **one single binary**.)

[//]: # ()

[//]: # (---)

<a id="options">
<h1>Command-line options</h1>
</a>

### letovo register [flags]

_Register your school credentials_

#### aliases = `r`, `reg`

### letovo help [flags]

_Pass any command to display manual about it_

#### aliases = _none_

### letovo schedule [flags]

_Get schedule from s.letovo.ru_

#### aliases = `s`, `sch`

- **-d {string}**, **--day {string}**: display a schedule for the specific day
- **-w**, **--week**: display a schedule for the week

_Default is for today_

### letovo homework [flags]

_Get homework from s.letovo.ru_

#### aliases = `h`, `hw`

- **-d {string}**, **--day {string}**: display a schedule for the specific day
- **-w**, **--week**: display a homework for the week

_Default is for today_

### letovo marks [flags]

_Get marks from s.letovo.ru_

#### aliases = `m`, `ma`

- **-d {string}**, **--day {string}**: display a schedule for the specific day
- **-a**, **--all**: display all marks
- **-f**, **--final**: display final marks
- **-s**, **--summative**: display summative marks

_Default is marks within one week_

### Some options accept parameters:

- Valid **--day** options are **`(?i)^mo`**, **`(?i)^tu`**, **`(?i)^we`**, **`(?i)^th`**, **`(?i)^fr`** and**`(?i)^sa`**

---

<a id="installation">
<h1>Installation</h1>
</a>

`letovo` is available for macOS, Linux and Windows.

### Download a binary

Compiled binary versions of `letovo` are uploaded to GitHub when a release is made. You can install `letovo` manually
by [downloading a release](https://github.com/arsikurin/letovoAnalyticsCLI/releases), extracting it, and copying the
binary to a directory in your `$PATH`, such as `/usr/local/bin`.

### Build from source

...

**© Made with ❤️ by arsikurin**