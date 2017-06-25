package task

import (
	"regexp"
	"strings"
	"time"
)

type Task struct {
	bytes []byte
}

func NewTask(bytes []byte) *Task {
	return &Task{bytes: bytes}
}

var (
	// time could be "2017-06-21T13:09:18.311643" or "2017-06-21 13:15:31"
	genericLine = regexp.MustCompile("^([DIEW]), \\[([\\d+T\\- \\.:]+) #\\d+\\] \\[(.*)\\] \\s?[A-Z]+ -- [A-Za-z]+: (.*)$")

	dirVersion = regexp.MustCompile("^Director Version: (.+)$")
	taskNum    = regexp.MustCompile("^Starting task: (.+)$")
)

func (t *Task) Details() (Details, error) {
	var details Details
	var errs []error

	for lineNum, lineStr := range strings.Split(string(t.bytes), "\n") {
		if m := genericLine.FindStringSubmatch(lineStr); len(m) > 0 {
			t1, err := t.parseTime(m[2])
			if err != nil {
				errs = append(errs, err)
				continue // todo
			}

			line := &Line{
				Number:  lineNum,
				Level:   m[1],
				Time:    t1,
				Group:   m[3], // todo empty group?
				Content: m[4],
			}

			details.Lines = append(details.Lines, line)

			if m := dirVersion.FindStringSubmatch(line.Content); len(m) > 0 {
				details.DirectorVersion = m[1]
				continue
			}

			if m := taskNum.FindStringSubmatch(line.Content); len(m) > 0 {
				details.TaskID = m[1]
				continue
			}

			continue
		}

		lenLines := len(details.Lines)

		if lenLines > 0 {
			prevLine := details.Lines[lenLines-1]
			prevLine.Content += "\n" + lineStr
			details.Lines[lenLines-1] = prevLine
		} else {
			line := UnknownLine{Number: lineNum, Content: lineStr}
			details.UnknownLines = append(details.UnknownLines, line)
		}
	}

	if len(errs) > 0 {
		return details, errs[0]
	}

	return details, nil
}

func (t *Task) parseTime(str string) (time.Time, error) {
	t1, err := time.Parse("2006-01-02T15:04:05", str) // eg "2017-06-21T13:09:18.311595"
	if err != nil {
		t1, err = time.Parse("2006-01-02 15:04:05", str) // eg "2017-06-21 13:09:18"
		if err != nil {
			return time.Time{}, err
		}
	}

	return t1, nil
}
