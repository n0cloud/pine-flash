package blclient

import (
	"bytes"
	"encoding/hex"
	"sync"
	"time"

	"github.com/albenik/go-serial/v2"
	"go.uber.org/zap"
)

type BLClient struct {
	baudrate int
	rwRetry  int
	cmdRetry int

	port *serial.Port
	mtx  sync.Mutex
	rbuf []byte

	logger *zap.SugaredLogger
}

func Dial(portName string, opts ...BLClientOption) (*BLClient, error) {
	blc := &BLClient{
		baudrate: 115200,
		rwRetry:  3,
		cmdRetry: 5,
		rbuf:     make([]byte, 256),
	}

	port, err := serial.Open(portName,
		serial.WithDataBits(8),
		serial.WithParity(serial.NoParity),
		serial.WithStopBits(serial.OneStopBit),
		serial.WithBaudrate(blc.baudrate),
	)
	if err != nil {
		return nil, err
	}
	port.SetReadTimeout(200)
	blc.port = port

	for _, opt := range opts {
		opt.Apply(blc)
	}

	err = blc.ResetDevice()
	if err != nil {
		return nil, err
	}

	err = blc.Handshake()
	if err != nil {
		blc.port.Close()
		return nil, err
	}

	return blc, nil
}

func (blc *BLClient) logInfo(tmpl string, args ...interface{}) {
	if blc.logger != nil {
		blc.logger.Infof(tmpl, args...)
	}
}

func (blc *BLClient) logDebug(tmpl string, args ...interface{}) {
	if blc.logger != nil {
		blc.logger.Debugf(tmpl, args...)
	}
}

func (blc *BLClient) Close() error {
	return blc.port.Close()
}

func (blc *BLClient) Write(data []byte) (err error) {
	blc.mtx.Lock()
	defer blc.mtx.Unlock()

	for ntry := 0; ntry < blc.rwRetry; ntry++ {
		_, err := blc.port.Write(data)
		if err == nil {
			break
		}
	}
	return
}
func (blc *BLClient) Read(length int) ([]byte, error) {
	blc.mtx.Lock()
	defer blc.mtx.Unlock()

	offset := 0
	try := blc.rwRetry
	for {
		n, err := blc.port.Read(blc.rbuf[offset:])
		try--
		if err != nil {
			if try > 0 {
				continue
			}
			return nil, err
		}

		offset += n
		if offset >= length || offset >= cap(blc.rbuf) {
			break
		}

		if try == 0 {
			break
		}
	}

	// copy to be memory safe
	res := make([]byte, offset)
	copy(res, blc.rbuf[:offset])

	return res, nil
}

func (blc *BLClient) Command(data []byte, length int) ([]byte, error) {
	err := blc.Write(data)
	if err != nil {
		return nil, err
	}
	if length >= 0 {
		length += 2
	} else {
		length = 0
	}
	res, err := blc.Read(length)
	if err != nil {
		return nil, err
	}
	if !isOk(res) {
		// Try to get the error infos
		errRes, err := blc.Read(0)
		if err == nil {
			res = append(res, errRes...)
		}
		return nil, ParseError(res)
	}
	return res[2:], nil
}

func (blc *BLClient) TryCommand(data []byte, length int) (b []byte, err error) {
	for ntry := 0; ntry < blc.cmdRetry; ntry++ {
		b, err = blc.Command(data, length)
		if err == nil {
			break
		}
	}
	return
}

func (blc *BLClient) ResetDevice() error {
	blc.port.SetRTS(true)
	time.Sleep(50 * time.Millisecond)
	blc.port.SetDTR(true)
	time.Sleep(50 * time.Millisecond)
	blc.port.SetDTR(false)
	time.Sleep(50 * time.Millisecond)
	blc.port.SetRTS(false)
	time.Sleep(50 * time.Millisecond)

	blc.port.Reconfigure(serial.WithBaudrate(blc.baudrate))
	blc.port.ResetInputBuffer()
	blc.port.ResetOutputBuffer()

	return nil
}

func (blc *BLClient) Handshake() error {
	shakeLength := (float64(blc.baudrate) / 10 / 1000 * 5) // send handshake for 5ms
	shakeCmd := bytes.Repeat([]byte{cmdHandShake}, int(shakeLength))

	blc.logDebug("5ms send count %f", shakeLength)

	_, err := blc.TryCommand(shakeCmd, 0)
	if err == nil {
		time.Sleep(20 * time.Millisecond)
	}
	return err
}

func (blc *BLClient) GetBootInfo() ([]byte, error) {
	data, err := blc.TryCommand([]byte{cmdGetBootInfo, 0x00, 0x00, 0x00}, 26)
	if err != nil {
		return nil, err
	}

	blc.logInfo("BootInfo: %s", hex.EncodeToString(data))
	blc.logInfo("ChipID: %s", hex.EncodeToString(data[len(data)-8:]))

	return data, nil
}

func (blc *BLClient) ReadJedecId() ([]byte, error) {
	data, err := blc.TryCommand([]byte{cmdReadJedecid, 0x00, 0x00, 0x00}, -1)
	if err != nil {
		return nil, err
	}
	blc.logInfo("ReadJedecId: %s", hex.EncodeToString(data))
	return data, nil
}
