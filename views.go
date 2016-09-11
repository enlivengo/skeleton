package main

import "github.com/enlivengo/enliven"

func rootView(ctx *enliven.Context) {
	ctx.Template("home")
}
