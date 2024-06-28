package minecraft

import (
	"bytes"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// PacketHandler handles incoming packets
type PacketHandler interface {
	HandlePacket(pk packet.Packet, conn *Conn, handle *bool)
}

// PacketDataHandler handles packet data before before gophertunnel decodes the packet
type PacketDataHandler interface {
	HandleData(h *packet.Header, full []byte, payload *bytes.Buffer, conn *Conn, handle *bool)
}

type nopPacketHandler struct{}

func (n nopPacketHandler) HandlePacket(pk packet.Packet, conn *Conn, handle *bool) {}

type nopPacketDataHandler struct{}

func (n nopPacketDataHandler) HandleData(h *packet.Header, full []byte, payload *bytes.Buffer, conn *Conn, handle *bool) {
}
