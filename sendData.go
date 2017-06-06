package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"strings"
)

func sendData(data []*MetaData) ([]byte, error) {

	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	log.Debugf("Send to %s, size: %d", cfg.FalconClient, len(data))
	for _, m := range data {
		log.Debugf("%s", m)
	}

	log.Infof("send upper %s",string(js))
	js = bytes.NewBufferString(strings.ToLower(string(js))).Bytes()

	res, err := http.Post(cfg.FalconClient, "Content-Type: application/json", bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}
	log.Info("success11111")
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
