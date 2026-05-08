package editor

import "errors"

var ErrInvalidRange = errors.New("editor: invalid range")

type Editor struct {
	text      string
	clipboard string
}

func NewEditor(text string) *Editor {
	return &Editor{text: text}
}

func (e *Editor) Text() string {
	return e.text
}

func (e *Editor) SetText(text string) {
	e.text = text
}

func (e *Editor) Clipboard() string {
	return e.clipboard
}

func (e *Editor) SetClipboard(text string) {
	e.clipboard = text
}

func (e *Editor) Insert(position int, value string) error {
	if position < 0 || position > len(e.text) {
		return ErrInvalidRange
	}

	e.text = e.text[:position] + value + e.text[position:]
	return nil
}

func (e *Editor) Delete(start int, end int) (string, error) {
	if !isValidRange(e.text, start, end) {
		return "", ErrInvalidRange
	}

	deleted := e.text[start:end]
	e.text = e.text[:start] + e.text[end:]

	return deleted, nil
}

func (e *Editor) Replace(start int, end int, value string) error {
	if !isValidRange(e.text, start, end) {
		return ErrInvalidRange
	}

	e.text = e.text[:start] + value + e.text[end:]
	return nil
}

func (e *Editor) Selection(start int, end int) (string, error) {
	if !isValidRange(e.text, start, end) {
		return "", ErrInvalidRange
	}

	return e.text[start:end], nil
}

func isValidRange(text string, start int, end int) bool {
	return start >= 0 && end >= start && end <= len(text)
}
