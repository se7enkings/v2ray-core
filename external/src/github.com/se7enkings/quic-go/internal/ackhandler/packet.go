package ackhandler

import (
	"time"

	"github.com/se7enkings/quic-go/internal/protocol"
	"github.com/se7enkings/quic-go/internal/wire"
)

// A Packet is a packet
type Packet struct {
	PacketNumber    protocol.PacketNumber
	PacketType      protocol.PacketType
	Ack             *wire.AckFrame
	Frames          []wire.Frame
	Length          protocol.ByteCount
	EncryptionLevel protocol.EncryptionLevel
	SendTime        time.Time

	largestAcked protocol.PacketNumber // if the packet contains an ACK, the LargestAcked value of that ACK

	// There are two reasons why a packet cannot be retransmitted:
	// * it was already retransmitted
	// * this packet is a retransmission, and we already received an ACK for the original packet
	canBeRetransmitted      bool
	includedInBytesInFlight bool
	retransmittedAs         []protocol.PacketNumber
	isRetransmission        bool // we need a separate bool here because 0 is a valid packet number
	retransmissionOf        protocol.PacketNumber
}

func (p *Packet) ToPacket() *protocol.Packet {
	return &protocol.Packet{
		PacketNumber: p.PacketNumber,
		PacketType:   p.PacketType,
		Length:       p.Length,
		SendTime:     p.SendTime,
	}
}
