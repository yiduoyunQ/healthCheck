package parser

import "bytes"

const (
	defaultSection = "default" // default section means if some ini items not in a section, make them in default section,
	bNumComment    = '#'       // number signal
	bSemComment    = ';'       // semicolon signal
	bEqual         = '='       // equal signal
	bDQuote        = '"'       // quote signal
	sectionStart   = '['       // section start signal
	sectionEnd     = ']'       // section end signal
	lineBreak      = '\n'
)

func GetValue(data []byte, key []byte) string {
	index := bytes.IndexByte(data, bEqual)
	if index < 1 {
		return ""
	}

	_key := bytes.TrimSpace(data[:index])
	if len(_key) == 0 || _key[0] == bNumComment {
		return ""
	}

	if !bytes.Equal(bytes.ToLower(_key), key) {
		return ""
	}

	right := bytes.IndexByte(data[index:], bNumComment)
	if right == -1 {
		data = data[index+1:]
	} else if right > 1 {
		data = data[index+1 : index+right]
	} else {
		return ""
	}

	return string(bytes.TrimSpace(data))
}

func GetSectionBody(data, section []byte) []byte {
	if len(section) == 0 {
		return data
	}

	section = bytes.ToLower(bytes.TrimSpace(section))
	if len(section) == 0 {
		return data
	}

	for index := 0; index < len(data); {
		// index '['
		ok, start := getSectionStart(data[index:])
		if !ok {
			return nil
		}
		index += start

		// index ']'
		end := bytes.IndexByte(data[index:], sectionEnd)
		if end == -1 {
			return nil
		}

		start = index
		index += end
		_section := bytes.ToLower(bytes.TrimSpace(data[start+1 : index]))

		// if not equal section,continue
		if !bytes.Equal(_section, section) {
			continue
		}

		// index next '['
		_, next := getSectionStart(data[index:])

		return data[start : index+next]
	}

	return nil
}

func getSectionStart(data []byte) (bool, int) {
	for index := 0; index < len(data); {

		// index '['
		start := bytes.IndexByte(data[index:], sectionStart)
		if start == -1 {
			return false, len(data)
		}

		index += start

		if valid(data[:index]) {
			return true, index
		}
	}

	return false, len(data)
}

// valid returns false if line containers '#'
func valid(line []byte) bool {
	// line start index
	start := bytes.LastIndexByte(line, lineBreak)
	if start == -1 {
		start = 0
	}

	// if containers '#',return false
	index := bytes.IndexByte(line[start:], bNumComment)
	if index == -1 {
		return true
	}

	return false
}
