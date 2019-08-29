package addrtool

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"strings"
)
var wordList = strings.Split(worldlist, "\n")

var wordIndexes = make(map[string]uint16, len(wordList))

func init() {
	for i, word := range wordList {
		wordIndexes[strings.ToLower(word)] = uint16(i)
	}
}


func EncodeMnemonic(seed []byte) string {
	var buf bytes.Buffer
	for i, b := range seed {
		if i != 0 {
			buf.WriteRune(' ')
		}
		buf.WriteString(ByteToMnemonic(b, i))
	}
	checksum := checksumByte(seed)
	buf.WriteRune(' ')
	buf.WriteString(ByteToMnemonic(checksum, len(seed)))
	return buf.String()
}


func checksumByte(data []byte) byte {
	intermediateHash := sha256.Sum256(data)
	return sha256.Sum256(intermediateHash[:])[0]
}
// ByteToMnemonic returns the PGP word list encoding of b when found at index.
func ByteToMnemonic(b byte, index int) string {
	bb := uint16(b) * 2
	if index%2 != 0 {
		bb++
	}
	return wordList[bb]
}

// DecodeMnemonics returns the decoded value that is encoded by words.  Any
// words that are whitespace are empty are skipped.
func DecodeMnemonics(words []string) ([]byte, error) {
	decoded := make([]byte, len(words))
	idx := 0
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w == "" {
			continue
		}
		b, ok := wordIndexes[strings.ToLower(w)]
		if !ok {
			return nil, errors.New("word %v is not in the PGP word list")
		}
		if int(b%2) != idx%2 {
			return nil, errors.New("word %v is not valid at position %v, ")
		}
		decoded[idx] = byte(b / 2)
		idx++
	}
	return decoded[:idx], nil
}



//026fc0084f3af4509b1b744be9a0b912b2889601e7e858d2305dfeae2dae585cbb