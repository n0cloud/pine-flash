package blclient

import (
	"github.com/albenik/go-serial/v2"
	"go.uber.org/zap"
)

type BLClientOption interface {
	Apply(blc *BLClient)
}

type BLClientOptionFunc func(blc *BLClient)

func (fn BLClientOptionFunc) Apply(blc *BLClient) {
	fn(blc)
}

func WithLogger(l *zap.SugaredLogger) BLClientOption {
	return BLClientOptionFunc(func(blc *BLClient) {
		blc.logger = l
	})
}

func WithBaudRate(baudrate int) BLClientOption {
	return BLClientOptionFunc(func(blc *BLClient) {
		blc.baudrate = baudrate
		blc.port.Reconfigure(serial.WithBaudrate(blc.baudrate))
	})
}

func WithReadTimeout(n int) BLClientOption {
	return BLClientOptionFunc(func(blc *BLClient) {
		blc.port.SetReadTimeout(n)
	})
}

func WithWriteTimeout(n int) BLClientOption {
	return BLClientOptionFunc(func(blc *BLClient) {
		blc.port.SetWriteTimeout(n)
	})
}

func WithReadBufferSize(n int) BLClientOption {
	return BLClientOptionFunc(func(blc *BLClient) {
		if n > cap(blc.rbuf) {
			blc.rbuf = append(blc.rbuf, make([]byte, n-cap(blc.rbuf))...)
		}
		// no-op if n is less or equal
	})
}
