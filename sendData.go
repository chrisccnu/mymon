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
	for index := range data{
		data[index].Endpoint = strings.ToLower(data[index].Endpoint)
		log.Infof("endpoint is lower %s",data[index].Endpoint)
	}

	js, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	log.Debugf("Send to %s, size: %d", cfg.FalconClient, len(data))
	for _, m := range data {
		log.Debugf("%s", m)
	}

	res, err := http.Post(cfg.FalconClient, "Content-Type: application/json", bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
