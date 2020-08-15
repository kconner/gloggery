# gloggery

gloggery is a basic static site generator for blogs in [Gemini](https://gemini.circumlunar.space), usually called gemlogs or glogs.

For posts, gloggery uses a simple post filename convention and plain text content.

When you run gloggery, it builds one post page per post, one index page to list all posts, and one Atom feed file to list recent posts. You can customize the templates for these page types.

## Setup

1. Have a [Go language](https://golang.org) environment
2. `make install`

make will build the executable, `gloggery`, and install it in `~/bin` by default. You can change that by setting the `prefix` variable; by default it is `$HOME`.

Installing will also create an empty `~/.gloggery/posts` folder and copy default page templates to `~/.gloggery/templates`.

## Writing posts

First, have a look at the [example post files](https://github.com/kconner/gloggery/tree/main/posts).

To create a post, add a file to `~/.gloggery/posts`, named like:

`2020-08-15-1427-hello-world`

This filename consists of the UTC year, month, day, and 24-hour time of day, plus a readable but URL-friendly slug title. You can get the current date and time with `date -u +%Y-%m-%d-%H%M`.

Inside the file, the first line (text up to the first `\n\n`) will be treated as the post title. The rest of the file is the post body, in which you can write Gemtext or just plain text.

> To make the most of your glog, you will want to understand Gemtext syntax, described by section 5 of the [Gemini specification](https://gemini.circumlunar.space/docs/specification.html). Link lines, which begin with `=>`, are of particular interest.

## Publishing

Given posts, templates, and a few other bits of information, gloggery will generate Gemtext page files and an Atom feed, all of which you can serve up with the Gemini protocol.

First, run `gloggery --help` and understand its arguments.

By default, gloggery consumes posts and templates as `--input` from `~/.gloggery` and emits pages as `--output` into `~/public_gemini/glog`, which is compatible with the author's home pubnix, [tilde.team](https://tilde.team).

The default site `--title` is your username prefixed with `~`.

The `--url` argument should be a `gemini://` URL corresponding to the `--output` folder.

> Make sure and get this URL right. gloggery doesn't have a reliable way to identify your host's domain name for the default value, which may cause Atom feed links to be incorrect.

Run gloggery with the options appropriate to your glog.

gloggery will emit to the output folder:

- one post page file per post that has changed
- a index file listing all posts, if any post changed
- an Atom feed file listing recent posts, if any post changed

> Post page filenames will not exactly match post input filenames. gloggery omits the time of day and adds the `.gmi` extension, but preserves the date and slug.

## Page templates

You can modify the page templates in `~/.gloggery/templates`. These are [Go text templates](https://golang.org/pkg/text/template/), but you don't need to understand that syntax if you leave the `{{ }}` terms alone.

By default, gloggery only rebuilds pages when post files change. If you change a template, you can rebuild all pages by running gloggery with the `--rebuild` argument.

## Developing gloggery

The smoothest way to work on gloggery's code is with `make watch-run`, which observes code file changes, builds the app, and runs it to publish the sample posts in `posts` to an `output` folder. This requires the very nice tool [entr](http://eradman.com/entrproject/).

If you just want to build when saving a code file, you can use `make watch`.

Before submitting changes, please run `make format`.
