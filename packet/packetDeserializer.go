package packet

import "fmt"

// deserializePacket Will contain parameters to index on elasticsearch
type deserializedPacket struct {
	FlowHash      uint64
	NetFlow       string
	Payload       string
	IPv4Length    uint16
	TCPpacketType string
	SrcPort       uint32
	DstPort       uint32
	DstMAC        string
	SrcMAC        string
	ttl           uint8
	srcIP         string
	dstIP         string
	dnsType       string
	dnsRespCode   string
	dnsContent    string
}

// takes DecodedPacket and returns a deserializedPacket with the necessary fields to index on elasticsearch
func deserializePacket(decodedpacket *DecodedPacket) *deserializedPacket {

	deserializedpacket := deserializedPacket{

		Payload:     decodedpacket.Payload.String(),
		FlowHash:    decodedpacket.FlowHash,
		IPv4Length:  decodedpacket.ipv4.Length,
		SrcPort:     uint32(decodedpacket.TCP.SrcPort),
		DstPort:     uint32(decodedpacket.TCP.DstPort),
		SrcMAC:      decodedpacket.eth.SrcMAC.String(),
		DstMAC:      decodedpacket.eth.DstMAC.String(),
		ttl:         decodedpacket.ipv4.TTL,
		srcIP:       decodedpacket.ipv4.SrcIP.String(),
		dstIP:       decodedpacket.ipv4.DstIP.String(),
		dnsType:     decodedpacket.DNS.LayerType().String(),
		dnsRespCode: decodedpacket.DNS.ResponseCode.String(),
		dnsContent:  string(decodedpacket.DNS.Contents),
	}
	fmt.Printf("%+v", deserializedpacket)
	return &deserializedpacket
}
