package blclient

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidResponse      = errors.New("invalid response")
	ErrInvalidBootHeader    = errors.New("invalid boot header")
	ErrInvalidSegmentHeader = errors.New("invalid segment header")
)

var _ BlfbResponseError[ErrCodeFlash] = &blfbResponseError[ErrCodeFlash]{}

type BlfbResponseError[T ~uint8] interface {
	error
	fmt.Stringer
	Code() T
}

type blfbResponseError[T blfberrcode] struct {
	code   T
	errMsg string
}

func (err *blfbResponseError[T]) String() string {
	return err.errMsg
}

func (err *blfbResponseError[T]) Error() string {
	return fmt.Sprintf("ErrorCode: %d, ErrorMessage: %s", err.code, err.errMsg)
}

func (err *blfbResponseError[T]) Code() T {
	return err.code
}

type ErrType uint8

const (
	ERR_FLASH ErrType = 0x00
	ERR_CMD   ErrType = 0x01
	ERR_IMAGE ErrType = 0x02
	ERR_IF    ErrType = 0x03
	ERR_MISC  ErrType = 0xFF
)

type (
	ErrCodeFlash uint8
	ErrCodeCmd   uint8
	ErrCodeImage uint8
	ErrCodeIF    uint8
	ErrCodeMisc  uint8
)

type blfberrcode interface {
	ErrCodeFlash | ErrCodeCmd | ErrCodeImage | ErrCodeIF | ErrCodeMisc
}

/*flash*/
const (
	ERR_FLASH_INIT_ERROR ErrCodeFlash = iota + 1
	ERR_FLASH_ERASE_PARA_ERROR
	ERR_FLASH_ERASE_ERROR
	ERR_FLASH_WRITE_PARA_ERROR
	ERR_FLASH_WRITE_ADDR_ERROR
	ERR_FLASH_WRITE_ERROR
	ERR_FLASH_BOOT_PARA_ERROR
	ERR_FLASH_SET_PARA_ERROR
	ERR_FLASH_READ_STATUS_REG_ERROR
	ERR_FLASH_WRITE_STATUS_REG_ERROR
)

func errFlash(code ErrCodeFlash) BlfbResponseError[ErrCodeFlash] {
	err := &blfbResponseError[ErrCodeFlash]{code: code}
	switch code {
	case ERR_FLASH_INIT_ERROR:
		err.errMsg = "ERR_FLASH_INIT_ERROR"
	case ERR_FLASH_ERASE_PARA_ERROR:
		err.errMsg = "ERR_FLASH_ERASE_PARA_ERROR"
	case ERR_FLASH_ERASE_ERROR:
		err.errMsg = "ERR_FLASH_ERASE_ERROR"
	case ERR_FLASH_WRITE_PARA_ERROR:
		err.errMsg = "ERR_FLASH_WRITE_PARA_ERROR"
	case ERR_FLASH_WRITE_ADDR_ERROR:
		err.errMsg = "ERR_FLASH_WRITE_ADDR_ERROR"
	case ERR_FLASH_WRITE_ERROR:
		err.errMsg = "ERR_FLASH_WRITE_ERROR"
	case ERR_FLASH_BOOT_PARA_ERROR:
		err.errMsg = "ERR_FLASH_BOOT_PARA_ERROR"
	case ERR_FLASH_SET_PARA_ERROR:
		err.errMsg = "ERR_FLASH_SET_PARA_ERROR"
	case ERR_FLASH_READ_STATUS_REG_ERROR:
		err.errMsg = "ERR_FLASH_READ_STATUS_REG_ERROR"
	case ERR_FLASH_WRITE_STATUS_REG_ERROR:
		err.errMsg = "ERR_FLASH_WRITE_STATUS_REG_ERROR"
	default:
		err.errMsg = "UNKNOWN"
	}
	return err
}

/*cmd*/
const (
	ERR_CMD_ID_ERROR ErrCodeCmd = iota + 1
	ERR_CMD_LEN_ERROR
	ERR_CMD_CRC_ERROR
	ERR_CMD_SEQ_ERROR
)

func errCmd(code ErrCodeCmd) BlfbResponseError[ErrCodeCmd] {
	err := &blfbResponseError[ErrCodeCmd]{code: code}
	switch code {
	case ERR_CMD_ID_ERROR:
		err.errMsg = "ERR_CMD_ID_ERROR"
	case ERR_CMD_LEN_ERROR:
		err.errMsg = "ERR_CMD_LEN_ERROR"
	case ERR_CMD_CRC_ERROR:
		err.errMsg = "ERR_CMD_CRC_ERROR"
	case ERR_CMD_SEQ_ERROR:
		err.errMsg = "ERR_CMD_SEQ_ERROR"
	default:
		err.errMsg = "UNKNOWN"
	}
	return err
}

/*image*/
const (
	ERR_IMG_BOOTHEADER_LEN_ERROR ErrCodeImage = iota + 1
	ERR_IMG_BOOTHEADER_NOT_LOAD_ERROR
	ERR_IMG_BOOTHEADER_MAGIC_ERROR
	ERR_IMG_BOOTHEADER_CRC_ERROR
	ERR_IMG_BOOTHEADER_ENCRYPT_NOTFIT
	ERR_IMG_BOOTHEADER_SIGN_NOTFIT
	ERR_IMG_SEGMENT_CNT_ERROR
	ERR_IMG_AES_IV_LEN_ERROR
	ERR_IMG_AES_IV_CRC_ERROR
	ERR_IMG_PK_LEN_ERROR
	ERR_IMG_PK_CRC_ERROR
	ERR_IMG_PK_HASH_ERROR
	ERR_IMG_SIGNATURE_LEN_ERROR
	ERR_IMG_SIGNATURE_CRC_ERROR
	ERR_IMG_SECTIONHEADER_LEN_ERROR
	ERR_IMG_SECTIONHEADER_CRC_ERROR
	ERR_IMG_SECTIONHEADER_DST_ERROR
	ERR_IMG_SECTIONDATA_LEN_ERROR
	ERR_IMG_SECTIONDATA_DEC_ERROR
	ERR_IMG_SECTIONDATA_TLEN_ERROR
	ERR_IMG_SECTIONDATA_CRC_ERROR
	ERR_IMG_HALFBAKED_ERROR
	ERR_IMG_HASH_ERROR
	ERR_IMG_SIGN_PARSE_ERROR
	ERR_IMG_SIGN_ERROR
	ERR_IMG_DEC_ERROR
	ERR_IMG_ALL_INVALID_ERROR
)

func errImage(code ErrCodeImage) BlfbResponseError[ErrCodeImage] {
	err := &blfbResponseError[ErrCodeImage]{code: code}
	switch code {
	case ERR_IMG_BOOTHEADER_LEN_ERROR:
		err.errMsg = "ERR_IMG_BOOTHEADER_LEN_ERROR"
	case ERR_IMG_BOOTHEADER_NOT_LOAD_ERROR:
		err.errMsg = "ERR_IMG_BOOTHEADER_NOT_LOAD_ERROR"
	case ERR_IMG_BOOTHEADER_MAGIC_ERROR:
		err.errMsg = "ERR_IMG_BOOTHEADER_MAGIC_ERROR"
	case ERR_IMG_BOOTHEADER_CRC_ERROR:
		err.errMsg = "ERR_IMG_BOOTHEADER_CRC_ERROR"
	case ERR_IMG_BOOTHEADER_ENCRYPT_NOTFIT:
		err.errMsg = "ERR_IMG_BOOTHEADER_ENCRYPT_NOTFIT"
	case ERR_IMG_BOOTHEADER_SIGN_NOTFIT:
		err.errMsg = "ERR_IMG_BOOTHEADER_SIGN_NOTFIT"
	case ERR_IMG_SEGMENT_CNT_ERROR:
		err.errMsg = "ERR_IMG_SEGMENT_CNT_ERROR"
	case ERR_IMG_AES_IV_LEN_ERROR:
		err.errMsg = "ERR_IMG_AES_IV_LEN_ERROR"
	case ERR_IMG_AES_IV_CRC_ERROR:
		err.errMsg = "ERR_IMG_AES_IV_CRC_ERROR"
	case ERR_IMG_PK_LEN_ERROR:
		err.errMsg = "ERR_IMG_PK_LEN_ERROR"
	case ERR_IMG_PK_CRC_ERROR:
		err.errMsg = "ERR_IMG_PK_CRC_ERROR"
	case ERR_IMG_PK_HASH_ERROR:
		err.errMsg = "ERR_IMG_PK_HASH_ERROR"
	case ERR_IMG_SIGNATURE_LEN_ERROR:
		err.errMsg = "ERR_IMG_SIGNATURE_LEN_ERROR"
	case ERR_IMG_SIGNATURE_CRC_ERROR:
		err.errMsg = "ERR_IMG_SIGNATURE_CRC_ERROR"
	case ERR_IMG_SECTIONHEADER_LEN_ERROR:
		err.errMsg = "ERR_IMG_SECTIONHEADER_LEN_ERROR"
	case ERR_IMG_SECTIONHEADER_CRC_ERROR:
		err.errMsg = "ERR_IMG_SECTIONHEADER_CRC_ERROR"
	case ERR_IMG_SECTIONHEADER_DST_ERROR:
		err.errMsg = "ERR_IMG_SECTIONHEADER_DST_ERROR"
	case ERR_IMG_SECTIONDATA_LEN_ERROR:
		err.errMsg = "ERR_IMG_SECTIONDATA_LEN_ERROR"
	case ERR_IMG_SECTIONDATA_DEC_ERROR:
		err.errMsg = "ERR_IMG_SECTIONDATA_DEC_ERROR"
	case ERR_IMG_SECTIONDATA_TLEN_ERROR:
		err.errMsg = "ERR_IMG_SECTIONDATA_TLEN_ERROR"
	case ERR_IMG_SECTIONDATA_CRC_ERROR:
		err.errMsg = "ERR_IMG_SECTIONDATA_CRC_ERROR"
	case ERR_IMG_HALFBAKED_ERROR:
		err.errMsg = "ERR_IMG_HALFBAKED_ERROR"
	case ERR_IMG_HASH_ERROR:
		err.errMsg = "ERR_IMG_HASH_ERROR"
	case ERR_IMG_SIGN_PARSE_ERROR:
		err.errMsg = "ERR_IMG_SIGN_PARSE_ERROR"
	case ERR_IMG_SIGN_ERROR:
		err.errMsg = "ERR_IMG_SIGN_ERROR"
	case ERR_IMG_DEC_ERROR:
		err.errMsg = "ERR_IMG_DEC_ERROR"
	case ERR_IMG_ALL_INVALID_ERROR:
		err.errMsg = "ERR_IMG_ALL_INVALID_ERROR"
	default:
		err.errMsg = "UNKNOWN"
	}
	return err
}

/*IF*/
const (
	ERR_IF_RATE_LEN_ERROR ErrCodeIF = iota + 1
	ERR_IF_RATE_PARA_ERROR
	ERR_IF_PASSWORDERROR
	ERR_IF_PASSWORDCLOSE
)

func errIF(code ErrCodeIF) BlfbResponseError[ErrCodeIF] {
	err := &blfbResponseError[ErrCodeIF]{code: code}
	switch code {
	case ERR_IF_RATE_LEN_ERROR:
		err.errMsg = "ERR_IF_RATE_LEN_ERROR"
	case ERR_IF_RATE_PARA_ERROR:
		err.errMsg = "ERR_IF_RATE_PARA_ERROR"
	case ERR_IF_PASSWORDERROR:
		err.errMsg = "ERR_IF_PASSWORDERROR"
	case ERR_IF_PASSWORDCLOSE:
		err.errMsg = "ERR_IF_PASSWORDCLOSE"
	default:
		err.errMsg = "UNKNOWN"
	}
	return err
}

/*MISC*/
const (
	ERR_PLL_ERROR ErrCodeMisc = iota + 252
	ERR_INVASION_ERROR
	ERR_POLLING
	ERR_FAIL
)

func errMisc(code ErrCodeMisc) BlfbResponseError[ErrCodeMisc] {
	err := &blfbResponseError[ErrCodeMisc]{code: code}
	switch code {
	case ERR_PLL_ERROR:
		err.errMsg = "ERR_PLL_ERROR"
	case ERR_INVASION_ERROR:
		err.errMsg = "ERR_INVASION_ERROR"
	case ERR_POLLING:
		err.errMsg = "ERR_POLLING"
	case ERR_FAIL:
		err.errMsg = "ERR_FAIL"
	default:
		err.errMsg = "UNKNOWN"
	}
	return err
}

func bytesToError(errMsb byte, errLsb byte) error {
	switch ErrType(errMsb) {
	case ERR_FLASH:
		return errFlash(ErrCodeFlash(errLsb))
	case ERR_CMD:
		return errCmd(ErrCodeCmd(errLsb))
	case ERR_IMAGE:
		return errImage(ErrCodeImage(errLsb))
	case ERR_IF:
		return errIF(ErrCodeIF(errLsb))
	case ERR_MISC:
		return errMisc(ErrCodeMisc(errLsb))
	}
	return ErrInvalidResponse
}

func ParseError(data []byte) error {
	if isOk(data) {
		return nil
	}
	if len(data) >= 2 && string(data[:2]) == "FL" {
		return bytesToError(data[3], data[2])
	}
	return ErrInvalidResponse
}
