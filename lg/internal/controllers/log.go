package controllers

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"time"

	"github.com/tmobaird/dv/core"
)

type LogArgs struct {
	Before string
	After  string
}

func (args LogArgs) BeforeDate() (time.Time, error) {
	return time.Parse(time.DateOnly, args.Before)
}

func (args LogArgs) AfterDate() (time.Time, error) {
	return time.Parse(time.DateOnly, args.After)
}

func (args LogArgs) HasBefore() bool {
	return args.Before != ""
}

func (args LogArgs) HasAfter() bool {
	return args.After != ""
}

func (controller Controller) RunLog(args LogArgs) (OutputTarget, error) {
	writeTo := OutputTarget{}
	before, after, err := parseBeforeAndAfter(args)
	if err != nil {
		return writeTo, err
	}

	cmd, stdin, err := startPager()
	if err != nil {
		writeTo = NewSimpleOutputTarget()
	} else {
		writeTo = NewCmdOutputTarget(stdin, cmd)
	}

	entries, err := os.ReadDir(core.LogDirectoryPath())
	if err != nil {
		writeTo.Close()
		return writeTo, err
	}

	filtered := filterLogs(entries, before, after)
	sortLogs(filtered)

	if len(filtered) == 0 {
		writeTo.Close()
		return writeTo, errors.New("no dev logs exist")
	}

	for i, entry := range filtered {
		day, err := core.LogFileNameToTime(entry.Name())
		output := logEntryOutput(entry.Name(), i == 0)
		if err == nil {
			diff := time.Since(day) / time.Hour / 24
			arg := strconv.FormatInt(int64(diff), 10) + "d"
			if arg == "0d" {
				arg = "today"
			} else if arg == "1d" {
				arg = "yesterday"
			}
			output += fmt.Sprintf("To view: dv lg show %s\n\n", arg)
		}
		writeTo.Write([]byte(output))
	}

	writeTo.Close()

	return writeTo, err
}

func parseBeforeAndAfter(args LogArgs) (time.Time, time.Time, error) {
	var err error
	before := time.Now().Add(24 * time.Hour)
	after := time.Now().Add(-1 * 20000 * 24 * time.Hour)
	if args.HasBefore() {
		before, err = args.BeforeDate()
		if err != nil {
			return before, after, err
		}
	}
	if args.HasAfter() {
		after, err = args.AfterDate()
		if err != nil {
			return before, after, err
		}
	}

	return before, after, nil
}

func filterLogs(files []os.DirEntry, before time.Time, after time.Time) []os.DirEntry {
	filtered := []os.DirEntry{}
	for _, entry := range files {
		t, err := core.LogFileNameToTime(entry.Name())
		if err == nil && entry.Type().IsRegular() && t.Before(before) && t.After(after) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

func sortLogs(files []os.DirEntry) {
	sort.Slice(files, func(i, j int) bool {
		a, _ := core.LogFileNameToTime(files[i].Name())
		b, _ := core.LogFileNameToTime(files[j].Name())

		return a.After(b)
	})
}

func startPager() (*exec.Cmd, io.WriteCloser, error) {
	pager := getPager()

	cmd := exec.Command(pager)
	cmd.Stdout = os.Stdout
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return cmd, stdin, err
	}

	if err := cmd.Start(); err != nil {
		return cmd, stdin, err
	}

	return cmd, stdin, nil
}

func getPager() string {
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less"
	}
	return pager
}
