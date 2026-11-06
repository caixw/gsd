// SPDX-FileCopyrightText: 2018-2026 caixw
//
// SPDX-License-Identifier: MIT

//go:generate gomsgs extract -l=und -m -f=yaml -tag=comment -func=github.com/issue9/localeutil.Phrase,github.com/issue9/localeutil.Error,github.com/issue9/localeutil.StringPhrase ./
//go:generate gomsgs update -src=./locales/und.yaml -dest=./locales/zh.yaml

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/issue9/localeutil"

	"github.com/caixw/gsd/locales"
)

const (
	objectUsage = localeutil.StringPhrase("export object name")
	dirUsage    = localeutil.StringPhrase("set source dir")
	outputUsage = localeutil.StringPhrase("set markdown doc path")
	langUsage   = localeutil.StringPhrase("set markdown page language")
	titleUsage  = localeutil.StringPhrase("set markdown page title")
	descUsage   = localeutil.StringPhrase("set markdown page description")
)

func main() {
	p, err := locales.NewPrinter(localeutil.DetectUserLanguage())
	if err != nil {
		panic(err)
	}

	fs := flag.NewFlagSet("gsd", flag.ExitOnError)

	v := fs.Bool("v", false, localeutil.StringPhrase("show version").LocaleString(p))
	obj := fs.String("object", "", objectUsage.LocaleString(p))
	dir := fs.String("dir", "", dirUsage.LocaleString(p))
	output := fs.String("o", "", outputUsage.LocaleString(p))
	lang := fs.String("lang", "cmn-Hans", langUsage.LocaleString(p))
	title := fs.String("title", "", titleUsage.LocaleString(p))
	desc := fs.String("desc", "", descUsage.LocaleString(p))

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if *v {
		if _, err := fmt.Fprintf(fs.Output(), "gsd: %s\nbuild with: %s\n", getAppVersion(""), runtime.Version()); err != nil {
			panic(err)
		}
		return
	}

	if err := export(*dir, *obj, *output, *lang, *title, *desc); err != nil {
		panic(err)
	}
}

func getAppVersion(v string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return v
}
