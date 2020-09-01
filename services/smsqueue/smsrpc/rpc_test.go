package smsrpc_test

import (
	"RizhaoLanshanLabourUnion/services/smsqueue/smsrpc"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T){
	smsrpc.SendMessage("lsxxxx","5xxxxho","1xxxx","您的案件已经成功接受，时间为"+time.Now().Format("2006-01-02 15:04:05"))
}

func BenchmarkSendMessage(b *testing.B) {
	j := 0
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		j++
	}
	b.StopTimer()
}
