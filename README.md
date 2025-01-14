[CD-R 700mb](https://github.com/thebaer/cdr): now with video
----

![](/videomixlaserdisc/screen.png)

# Mix Laserdisc

Before Spotify, Apple Music, Amazon Music Unlimited, Amazon Music HD, Google Play Music, Tidal, Deezer, Pandora, Idagio, LiveXLive, Primephonic, SiriusXM Essential, SiriusXM Premier, and YouTube Music, we listened to mixtapes. We compiled tracks on iTunes and burned them onto CD-Rs. We traded them with each other. We freely listened in our car or at home or with our portable CD player. And it was good.

Today we might spend an eternity making the perfect playlist for someone special, only to learn they're on the Z Music Unlimited service while you're on X Music Basic. Now, vendor lock-in means love lock-out. Our connections are disconnected. Our romantic hopes aren't high enough for the surrounding paywalls. What can we do?

## A Modern Mixtape (with video)

This is a static site generator for "burning" a modern mixtape. It's platform-independent and won't get scratched up in your car.

Run this program inside a directory of music files to produce a bare-bones HTML page that will simply play the music for someone. Upload it all to your website, and share the URL. Boom, you have a mixtape.

This only takes care of "burning the CD," if you will -- then it's up to you to make it special, like you would a normal mixtape. Add track notes, make it look cool or funky, do _something_ that takes a little more effort than sharing a freaking Spotify playlist. All this tech convenience has made us lazy and boring. So make something interesting! Show someone that they **mean something** to you!

## The Process

First, compile a collection of music files you'd like to put on a mixtape, and rename them according to the order you want, where each file name starts with a two-digit track number (i.e. `01`, `02`, ... `10`, `11`).

Of course, THIS PROGRAM IS FOR ILLUSTRATIVE PURPOSES ONLY AND SHOULD NOT BE USED TO DISTRIBUTE COPYRIGHTED MATERIAL YOU DO NOT OWN THE RIGHTS TO.

## `cdr` Usage

1. Install `cdr` with `go get github.com/thebaer/cdr/cmd/cdr`
1. Create a directory for your playlist
1. Add audio files into the directory
1. Make sure they're named in the order you want, and start with a two-digit track number, e.g. `01 - Track 1.mp3`
1. Run `cdr clean` in this directory to standardize the file names based on their metadata (supports ID3, MP4, OGG, FLAC)
1. Run `cdr burn` in this directory to generate your mixtape page
1. Open your new `index.html` in your browser!

This gives you the basic HTML. Now have fun with it.

### Over-achievers

You can also tweak the original template _before_ it generates the final page.

1. Copy `mixtape.tmpl` into your mixtape directory
1. Edit this file to your liking, being sure to retain the `{{template ...}}` lines in the file
1. Run `cdr burn` -- it'll generate your page from this template instead of the default! 

### Template parts

These are the full template codes you can use to include certain elements.

#### Player + Playlist

This shows the audio player with play / pause actions, plus the playlist. You don't need to use any other template codes beyond this.

```
{{template "full-player" .Tracks}}
```

#### Granular Elements

Optionally, you can use these more granular template codes for better control over your mixtape.

**IMPORTANT!** When using anything besides the `full-player` code, you will ALWAYS want to include the following code in your template. Otherwise your mixtape won't work well!

```
{{template "playlist-js"}}
```

#### Player

Shows only the audio player.

```
{{template "player" .Tracks}}
```

#### Playlist

Show the playlist of tracks on the mixtape.

```
{{template "playlist" .Tracks}}
```

### Developers

Requires Go 1.16 and above.

## Commands

```
NAME:
   CD-R 700MB - A static mixtape site generator

USAGE:
   cdr [global options] command [command options] [arguments...]

VERSION:
   v1.0

COMMANDS:
   burn     generate the static mixtape site
   preview  serve the mixtape site
   clean    clean and organize audio files in the current directory
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

```
~/ldr $ go build cdr
~/ldr/videomixlaserdisc $ go run ../cmd/cdr preview
```

### notes to self

[a helpful link about go executables](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-20-04)

#### helpful history commands for folx who don't use golang aka me
```
go -v
go version
go get github.com/thebaer/cdr/cmd/cdr
go install github.com/thebaer/cdr/cmd/cdr
go env
go help get
go list
/$HOME/go/bin/cdr
/$HOME/go/bin/cdr clean
/$HOME/go/bin/cdr preview
/$HOME/go/bin/cdr help
/$HOME/go/bin/cdr help preview
go build cdr
go help build
go build -pkgdir ../ cdr
ls /usr/local/go/src/cdr
go build cdr
sudo ln -s /$HOME/nun/yo/bisnez/cdr /usr/local/go/src/cdr
ls /usr/local/go/src/cdr
go build cdr
go build cdr
ls /usr/local/go/src/cdr
ls /$HOME/go/bin/cdr
go run -pkgdir ./ cdr
go run -pkgdir ./
go mod init ldr
rm -rf go.mod 
go mod init ldr
go mod tidy // makes go.sum
go build ldr
go run ldr preview
go run ldr
go build cdr
go build ldr
ls /usr/local/go/src/cdr
go build ldr
rm -rf go.sum
go run ../cmd/cdr
go run ../cmd/cdr preview
```