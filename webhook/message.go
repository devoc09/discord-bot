package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ReadConfig(fname string) (username string, avatar_url string, webhook_url string, err error) {
	bytes, err := os.ReadFile(fname)
	if err != nil {
		fmt.Printf("Read Config Error!! %s\n", err)
		return "", "", "", err
	}

	var config []Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return "", "", "", err
	}
	conf := config[0]
	return conf.UserName, conf.Abatar_url, conf.WebhookUrl, nil
}

func getresp(resp *http.Response) uint64 {
	if resp.StatusCode == 204 {
		return 0
	} else {
		return 1
	}
}

func SendMinMessage(url string, minmessage *MinMessage) (<-chan uint64, <-chan error, error) {
	msg, err := json.Marshal(minmessage)
	if err != nil {
		return nil, nil, fmt.Errorf("JSON marshal error: %s", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, nil, fmt.Errorf("Create Request error: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	doneCh, errCh := make(chan uint64), make(chan error)

	go func() {
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errCh <- err
			return
		}

		switch getresp(resp) {
		case 0:
			doneCh <- 0
		case 1:
			errCh <- err
		}
	}()

	return doneCh, errCh, nil
}

func SendEmbedMessage(url string, embedmessage *EmbedMessage) (<-chan uint64, <-chan error, error) {
	msg, err := json.Marshal(embedmessage)
	if err != nil {
		return nil, nil, fmt.Errorf("JSON marshal error: %s", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
	if err != nil {
		return nil, nil, fmt.Errorf("Create Request error: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	doneCh, errCh := make(chan uint64), make(chan error)

	go func() {
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errCh <- err
			return
		}

		switch getresp(resp) {
		case 0:
			doneCh <- 0
		case 1:
			errCh <- err
		}
	}()

	return doneCh, errCh, nil
}
