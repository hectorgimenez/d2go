package main

import (
	"fmt"
	"time"

	"github.com/hectorgimenez/d2go/pkg/packet"
)

func main() {
	sniffer := packet.NewPacketSniffer()
	interceptor := packet.NewPacketInterceptor(sniffer)

	sniffer.Start()

	// Simulate some packets (in a real scenario, these would come from the game)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			interceptor.OnPacketReceive([]byte{1, 2, 3, 4}, 0, false)
			interceptor.OnPacketSend([]byte{5, 6, 7, 8}, 0, "")
		}
	}()

	time.Sleep(time.Second * 15)

	sniffer.Stop()

	for _, packet := range sniffer.Queue {
		fmt.Printf("Packet: %s, OpCode: %d, Length: %d, IsIncoming: %v\n",
			packet.Name, packet.OpCode, packet.Length, packet.IsIncoming)
	}
}
