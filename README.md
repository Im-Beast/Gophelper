# Gophelper
[![Go report](https://goreportcard.com/badge/github.com/Im-Beast/Gophelper)](https://goreportcard.com/report/github.com/Im-Beast/Gophelper)

Simple bot written in Golang using [DiscordGo](https://github.com/bwmarrin/discordgo).

## Introduction
This is my first project using Golang, originally written just for fun in midnight, rewritten to be flexible and easily scalable next 2 days

## Features
* Gophelper has simple rate limiter that can has different config per command
* Gophelper provides easy way to add commands with ability to use arguments
* Commands can contain spaces e.g. `play pingpong` may be just one command
* Nothing stops you also from creating multiple routers which may be used for a different purpose e.g. providing different language based on router's prefix
* Routers support middleware to control what passes through

## How do I install this
1. Clone repo
2. Edit config files to your preferences and add `BOT_TOKEN` env variable which will store your bot token
3. `cd` to cloned repo to directory `src/Main`
4. Run `go build .` command
5. Launch built executable

## Config files
Gophelper provides easy to edit JSON config files. <br>
It also provides way to translate bot into multiple languages without much effort, in this repo you can see two languages: polish and english, nothing stops you by expanding it by more languages

## Pulls, forks and issues
Be free to fork this repo, if you have any comment regarding to quality of my Golang code - go ahead and open issue/add pull request. My first meet with go was 3 days ago, so I know many things can be done better.

## Commands
Gopher comes with some commands out of a box, all of them are pretty simple <br>
Here's the list of some of them: <br>

| Alias                       | Description                                                                                                                                                                                                                                                                                                                                                                                             |
|-----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| help                        | Provides either list of categories, commands or information about them.<br>You may spot characters like `_` in Usage, there's how to interpret them:<br> * `[]` simply means attribute<br> * `_` before argument means it's non obligatory, you may skip it or set some random things in its place<br> * `{}` e.g. `[_user{mention/id}]` means you can put either user id or user mention in this place |
| game pingpong               | Play ping pong match versus bot                                                                                                                                                                                                                                                                                                                                                                         |
| 8ball                       | *Magically* answers your questions                                                                                                                                                                                                                                                                                                                                                                      |
| stats                       | Shows information about you or given user                                                                                                                                                                                                                                                                                                                                                               |
| hug/kitty/doggie/waifu etc. | Gives you specified thing                                                                                                                                                                                                                                                                                                                                                                               |
| lang                        | Sets language config of router to given file located in "Languages/" e.g. `go lang english.json`                                                                                                                                                                                                                                                                                                        |

## Credits
* DiscordGo developers
* Creators of images used in kittie/doggie/kiss/hug/waifu and other commands :D

## Media
![Gophelper](https://github.com/Im-Beast/Gophelper/blob/main/docs/showcase.gif)