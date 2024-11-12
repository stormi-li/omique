package omique

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	manager "github.com/stormi-li/omi-v1/omi-manager"
)

type Producer struct {
	configSearcher *manager.Searcher
	channel        string
	address        string
	conn           net.Conn
}

func (producer *Producer) connect() error {
	producer.address, _ = producer.configSearcher.SearchByLoadBalancing(producer.channel)
	if producer.address == "" {
		producer.conn = nil
		return fmt.Errorf("no message queue service was found")
	}
	conn, err := net.Dial("tcp", producer.address)
	if err == nil {
		producer.conn = conn
		return nil
	}
	return err
}

func (producer *Producer) Publish(message []byte) error {
	var err error
	retryCount := 0

	//长度前缀协议
	byteMessage := []byte(string(message))
	messageLength := uint32(len(byteMessage))

	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, messageLength)

	for {
		if producer.conn != nil {
			if _, err = producer.conn.Write(append(lengthBuf, byteMessage...)); err == nil {
				break
			}
		}
		err = producer.connect()
		time.Sleep(const_retryWaitTime)
		if retryCount == const_maxRetryCount {
			break
		}
		retryCount++
	}
	return err
}
