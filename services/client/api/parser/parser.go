package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"pgevangelidis/PortDomainService/services/client/api/models"
)

func StreamJson(fileName string, stream chan (models.Port)) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("failed to open the file, Error: %s", err.Error())
		return err
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	_, err = decoder.Token()
	if err != nil {
		log.Printf("failed to read the token, Error: %s", err.Error())
		return err
	}

	var token json.Token
	for decoder.More() {
		token, err = decoder.Token()
		if err != nil {
			if err == io.EOF {
				return err
			}
			log.Printf("failed to read the code token, Error: %s", err.Error())
			return err
		}

		var port = &models.Port{}
		err = decoder.Decode(&port)
		if err != nil {
			log.Printf("failed to decode the object, Error: %s", err.Error())
			return err
		}

		port.ID = fmt.Sprint(token)

		stream <- *port
	}
	return nil
}
