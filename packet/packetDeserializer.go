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
	dnsType       uint64
	dnsRespCode   uint8
	dnsContent    string
}

// takes DecodedPacket and returns a deserializedPacket with the necessary fields to index on elasticsearch
func deserializePacket(decodedpacket *DecodedPacket) *deserializedPacket {

	deserializedpacket := deserializedPacket{

		Payload:     string(decodedpacket.Payload),
		FlowHash:    decodedpacket.FlowHash,
		IPv4Length:  decodedpacket.ipv4.Length,
		SrcPort:     uint32(decodedpacket.TCP.SrcPort),
		DstPort:     uint32(decodedpacket.TCP.DstPort),
		SrcMAC:      string(decodedpacket.eth.SrcMAC),
		DstMAC:      string(decodedpacket.eth.DstMAC),
		ttl:         decodedpacket.ipv4.TTL,
		srcIP:       string(decodedpacket.ipv4.SrcIP),
		dstIP:       string(decodedpacket.ipv4.DstIP),
		dnsType:     uint64(decodedpacket.DNS.LayerType()),
		dnsRespCode: uint8(decodedpacket.DNS.ResponseCode),
		dnsContent:  string(decodedpacket.DNS.Contents),
	}
	fmt.Printf("%#v", deserializedpacket)
	return &deserializedpacket
}
