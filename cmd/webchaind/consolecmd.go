// Copyright 2016 The go-ethereum Authors
// Copyright 2018 Webchain project
// This file is part of Webchain.
//
// Webchain is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Webchain is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Webchain. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/mkrufky/webchaind/console"
	"github.com/mkrufky/webchaind/node"
	"github.com/mkrufky/webchaind/rpc"
	"gopkg.in/urfave/cli.v1"
)

var (
	consoleCommand = cli.Command{
		Action: localConsole,
		Name:   "console",
		Usage:  `Webchaind Console: interactive JavaScript environment`,
		Description: `
	The webchaind console is an interactive shell for the JavaScript runtime environment
	which exposes a node admin interface as well as the Ðapp JavaScript API.
	See https://github.com/ethereum/go-ethereum/wiki/Javascipt-Console
		`,
		Flags: []cli.Flag{
			ExecFlag,
		},
	}
	attachCommand = cli.Command{
		Action: remoteConsole,
		Name:   "attach",
		Usage:  `Webchaind Console: interactive JavaScript environment (connect to node)`,
		Description: `
	The webchaind console is an interactive shell for the JavaScript runtime environment
	which exposes a node admin interface as well as the Ðapp JavaScript API.
	See https://github.com/ethereum/go-ethereum/wiki/Javascipt-Console.
	This command allows to open a console on a running webchaind node.

	<DATADIR> and <CHAINDIR> flags will be parsed as usual.
	For example:

		webchaind --chain=morden attach

	or,

		webchaind --data-dir=/path/to/webchainddata --chain privatenet attach
		`,
		Flags: []cli.Flag{
			ExecFlag,
		},
	}
	javascriptCommand = cli.Command{
		Action: ephemeralConsole,
		Name:   "js",
		Usage:  `Executes the given JavaScript files in the webchaind JavaScript VM`,
		Description: `
	The JavaScript VM exposes a node admin interface as well as the Ðapp
	JavaScript API. See https://github.com/ethereum/go-ethereum/wiki/Javascipt-Console
		`,
	}
)

// localConsole starts a new webchaind node, attaching a JavaScript console to it at the
// same time.
func localConsole(ctx *cli.Context) error {
	// Create and start the node based on the CLI flags
	node := MakeSystemNode(Version, ctx)
	startNode(ctx, node)
	defer node.Stop()

	// Attach to the newly started node and start the JavaScript console
	client, err := node.Attach()
	if err != nil {
		log.Fatal("Failed to attach to the inproc webchaind: ", err)
	}
	config := console.Config{
		DataDir: node.DataDir(),
		DocRoot: ctx.GlobalString(JSpathFlag.Name),
		Client:  client,
		Preload: MakeConsolePreloads(ctx),
	}
	console, err := console.New(config)
	if err != nil {
		log.Fatal("Failed to start the JavaScript console: ", err)
	}
	defer console.Stop(false)

	// If only a short execution was requested, evaluate and return
	//
	// --exec as command sub-flag
	if script := ctx.String(ExecFlag.Name); script != "" {
		console.Evaluate(script)
		return nil
	}

	// --exec as global flag
	if script := ctx.GlobalString(ExecFlag.Name); script != "" {
		console.Evaluate(script)
		return nil
	}

	// Otherwise print the welcome screen and enter interactive mode
	console.Welcome()
	console.Interactive()

	return nil
}

// remoteConsole will connect to a remote webchaind instance, attaching a JavaScript
// console to it.
func remoteConsole(ctx *cli.Context) error {
	// Attach to a remotely running webchaind instance and start the JavaScript console
	chainDir := MustMakeChainDataDir(ctx)
	var uri = "ipc:" + node.DefaultIPCEndpoint(chainDir)
	if ctx.Args().Present() {
		uri = ctx.Args().First()
	}
	client, err := rpc.NewClient(uri)
	if err != nil {
		log.Fatal("attach to remote webchaind: ", err)
	}

	config := console.Config{
		DataDir: chainDir,
		DocRoot: ctx.GlobalString(JSpathFlag.Name),
		Client:  client,
		Preload: MakeConsolePreloads(ctx),
	}

	console, err := console.New(config)
	if err != nil {
		log.Fatal("Failed to start the JavaScript console: ", err)
	}
	defer console.Stop(false)

	// If only a short execution was requested, evaluate and return
	//
	// --exec as command sub-flag
	if script := ctx.String(ExecFlag.Name); script != "" {
		console.Evaluate(script)
		return nil
	}

	// --exec as global flag
	if script := ctx.GlobalString(ExecFlag.Name); script != "" {
		console.Evaluate(script)
		return nil
	}
	// Otherwise print the welcome screen and enter interactive mode
	console.Welcome()
	console.Interactive()

	return nil
}

// ephemeralConsole starts a new webchaind node, attaches an ephemeral JavaScript
// console to it, and each of the files specified as arguments and tears the
// everything down.
func ephemeralConsole(ctx *cli.Context) error {
	// Create and start the node based on the CLI flags
	node := MakeSystemNode(Version, ctx)
	startNode(ctx, node)
	defer node.Stop()

	// Attach to the newly started node and start the JavaScript console
	client, err := node.Attach()
	if err != nil {
		log.Fatal("Failed to attach to the inproc webchaind: ", err)
	}
	config := console.Config{
		DataDir: node.DataDir(),
		DocRoot: ctx.GlobalString(JSpathFlag.Name),
		Client:  client,
		Preload: MakeConsolePreloads(ctx),
	}
	console, err := console.New(config)
	if err != nil {
		log.Fatal("Failed to start the JavaScript console: ", err)
	}
	defer console.Stop(false)

	// Evaluate each of the specified JavaScript files
	for _, file := range ctx.Args() {
		if err = console.Execute(file); err != nil {
			log.Fatalf("Failed to execute %s: %v", file, err)
		}
	}
	// Wait for pending callbacks, but stop for Ctrl-C.
	abort := make(chan os.Signal, 1)
	signal.Notify(abort, os.Interrupt)

	go func() {
		<-abort
		os.Exit(0)
	}()
	console.Stop(true)

	return nil
}
