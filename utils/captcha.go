package utils

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"log"
	"math"
	"time"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 6
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 1 * time.Minute
)

const (
	// Standard width and height of a captcha image.
	StdWidth  = 240
	StdHeight = 80
)


var (
	ErrNotFound = errors.New("captcha: id not found")
)


type PendingCaptcha struct{

	Challenge string `json:"challenge"`
	Time *Time `json:"timestamp"`
	ImageData string `json:"image"`
	digits []byte `json:"-"`
	PlainText string `json:"-"`

}


func parseDigitsToString(digits []byte) string{
	buf := bytes.NewBuffer([]byte(""))
	var i byte
	for _, i = range digits{
		buf.WriteByte(byte(i+48))
	}
	return buf.String()
}

func formatStringToDigits(str string) []byte{
	reader := bytes.NewReader([]byte(str))
	buffer := bytes.NewBuffer([]byte(""))
	for {
		b, err := reader.ReadByte()
		if err == nil {
			buffer.WriteByte(b-48)
		}else{
			break
		}
	}
	return buffer.Bytes()
}

func CreateCaptcha(id string) *PendingCaptcha{
	digits := captcha.RandomDigits(6)
	now := Time(time.Now())
	pendingCaptcha := &PendingCaptcha{
		Challenge: "",
		Time: &now,
		ImageData: "",
		digits: digits,
		PlainText: parseDigitsToString(digits),
	}

	timeStr, _ := pendingCaptcha.Time.MarshalJSON()

	image := captcha.NewImage(string(timeStr),
		                      pendingCaptcha.digits,
		                      StdWidth,StdHeight)

	buf := bytes.NewBuffer([]byte(""))
	buffWriter := bufio.NewWriter(buf)

	image.WriteTo(buffWriter)
	buffWriter.Flush()

	imageHexData := fmt.Sprintf("data:image/png;base64,%s",base64.StdEncoding.EncodeToString(buf.Bytes()))

	pendingCaptcha.ImageData = imageHexData
	pendingCaptcha.Challenge = fmt.Sprintf("%x",sha256.Sum256([]byte(pendingCaptcha.PlainText+id+string(timeStr))))
	//log.Println(pendingCaptcha.PlainText+id+string(timeStr))

	return pendingCaptcha
}


func CheckCaptcha(id, plainText, timeStr, challenge string) bool{
	tx, err := time.ParseInLocation(timeFormat,timeStr,time.Local)
	if err != nil{
		return false
	}
	if math.Abs(time.Now().Sub(tx).Minutes()) > 1 {
		log.Println("some one use expired captcha", CurrentTimeString(), Time(tx).FormattedString(),time.Now().Sub(tx).String())
		return false
	}
	return challenge == fmt.Sprintf("%x",sha256.Sum256([]byte(plainText+id+"\""+timeStr+"\"")))
}
