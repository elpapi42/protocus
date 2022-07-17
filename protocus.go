package protocus

import (
	"encoding/binary"
	"encoding/json"
)

type CommandCodes int8

const (
	broadcast CommandCodes = 1
)

func intToBytes(i int, lenght int) []byte {
	output := make([]byte, lenght)
	binary.BigEndian.PutUint64(output, uint64(i))
	return output
}

func getCommandCodeBytes(commandCode CommandCodes) []byte {
	output := intToBytes(int(commandCode), 8)
	return output
}

func Broadcast(channel string, file []byte) ([]byte, error) {
	command := getCommandCodeBytes(broadcast)

	headers, err := json.Marshal(map[string]string{"channel": channel})
	if err != nil {
		return nil, err
	}

	headersLen := intToBytes(len(headers), 16)

	fileLen := intToBytes(len(file), 32)

	output := append(append(append(append(command, headersLen...), fileLen...), headers...), file...)

	return output, nil
}
