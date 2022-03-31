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

### letovo help [flags]

Pass any command to display manual about it

#### aliases = _none_

<br/>

### letovo register [flags]

Register your school credentials

#### aliases = `r`, `reg`

<br/>

### letovo schedule [flags]

Get schedule from s.letovo.ru _(Default is for today)_

#### aliases = `s`, `sch`

- **-d {string}**, **--day {string}**: display a schedule for the specific day
- **-w**, **--week**: display a schedule for the week

<br/>

### letovo homework [flags]

Get homework from s.letovo.ru _(Default is for today)_

#### aliases = `h`, `hw`

- **-d {string}**, **--day {string}**: display a schedule for the specific day
- **-w**, **--week**: display a homework for the week

<br/>

### letovo marks [flags]

Get marks from s.letovo.ru _(Default is marks within one week)_

#### aliases = `m`, `ma`

- **-d {string}**, **--day {string}**: display a schedule for the specific day
- **-a**, **--all**: display all marks
- **-f**, **--final**: display final marks
- **-s**, **--summative**: display summative marks

<br/>

### Some options accept parameters:

- Valid **--day** options are monday **`(?i)^mo`**, tuesday **`(?i)^tu`**, wednesday **`(?i)^we`**,
  thursday **`(?i)^th`**, friday **`(?i)^fr`** and saturday **`(?i)^sa`**

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