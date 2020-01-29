package util

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net"
	"time"
)

var TraceKey = "hade-trace-id"

// NewTraceId 生成新的traceId
func NewTraceId() string {
	// traceId 生成方法： 时间戳 + 机器ip + 随机数

	b := bytes.Buffer{}

	now := time.Now()
	timestamp := uint32(now.Unix())
	timeNano := now.UnixNano()
	b.WriteString(fmt.Sprintf("%08x", timestamp&0xffffffff))
	b.WriteString(fmt.Sprintf("%04x", timeNano&0xffff))

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	localIp := net.ParseIP("127.0.0.1")
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				localIp = ipNet.IP
				break
			}
		}
	}

	netIP := net.ParseIP(localIp.String())
	if netIP == nil {
		b.WriteString("00000000")
	} else {
		b.WriteString(hex.EncodeToString(netIP.To4()))
	}

	b.WriteString(fmt.Sprintf("%06x", rand.Int31n(1<<24)))

	return b.String()
}

//GetTraceId
func GetTraceId(c *gin.Context) string {
	val, exists := c.Get(TraceKey)
	if exists {
		return val.(string)
	}

	if traceId := getTraceIdFromHttpHeader(c); traceId != "" {
		c.Set(TraceKey, traceId)
		return traceId
	}

	if traceId := NewTraceId(); traceId != "" {
		c.Set(TraceKey, traceId)
		return traceId
	}

	return ""
}

func getTraceIdFromHttpHeader(c *gin.Context) string {
	if traceId := c.Request.Header.Get(TraceKey); traceId != "" {
		return traceId
	}
	return ""
}

//SetTraceId
func SetTraceId(c *gin.Context, traceId string) {
	c.Set(TraceKey, traceId)
}