package domain

import (
	apperrors "github.com/miguelramirez93/quasar_fire_operation/shared/app_errors"
)

func GetMessage(messages ...[]string) (msg string) {
	msg, err := decodeMessage(messages...)
	if err != nil {
		panic(err.Error())
	}

	return
}

func getAndCheckMessageLen(messages ...[]string) (int, error) {

	messagesToCompare := messages[1:]

	messageLen := len(messages[0])

	for _, message := range messagesToCompare {
		currMessageLen := len(message)
		if currMessageLen != messageLen {
			return messageLen, apperrors.ErrMessageLenMismatch
		}
	}

	return messageLen, nil
}

func decodeMessage(messages ...[]string) (string, error) {
	messageLen, err := getAndCheckMessageLen(messages...)
	if err != nil {
		return "", err
	}
	decodedMessage := make([]string, messageLen)
	messageFractionIndex := make(map[int]bool)

	for i := range decodedMessage {
		messageFractionIndex[i] = true
	}

	for _, message := range messages {
		for i := range messageFractionIndex {
			if message[i] != "" {
				decodedMessage[i] = message[i]
				delete(messageFractionIndex, i)
				if len(messageFractionIndex) == 0 {
					return getMessageFromArray(decodedMessage), nil
				}
			}
		}
	}

	return "", apperrors.ErrDecodingMessage

}

func getMessageFromArray(message []string) (msg string) {
	fixedMessageLen := len(message) - 1
	for i, word := range message {
		msg = msg + word
		if i != fixedMessageLen {
			msg = msg + " "
		}
	}

	return
}
