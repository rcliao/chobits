# Design docs

In this article, I'll write down a couple core design components toward the tachikoma
system. It involves *core*, *ghost* and the *shell*.

## Core

Core is the _core system_ that runs the core life cycle and meant to interact with
the operating system. Mostly will be done in the GoLang library as the GoLang can
do the cross compile between different systems (e.g. from MacOS to compile Windows
application).

An concrete example of the core system might involve media playing, graphic,
network access ... etc. These lower implementation may be different from one
operating system to another but it should be the same interface from tachikoma
point of view.

One possible idea about the core system is being able to establish a network of
knowledge for possible ghosts. In example, multiple tachikoma may be able to
download ghost from the central server as part of _learning_.

## Ghost

Ghost is the _applications_ that will be running in tachikoma. In example, if
tachikoma needs to know how to record diary. It's up to tachikoma to select *diary*
ghost for its functionality behind the scene.

## Shell

Shell is the _user interface_ for the tachikoma. In example, if tachikoma is
running mainly in command line. Shell will involve the Command Line Interface to
use ghost accordingly.
