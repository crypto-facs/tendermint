package p2p_test

import (
	"context"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/internal/p2p"
)

// Common setup for P2P tests.

var (
	ctx    = context.Background()
	chID   = p2p.ChannelID(1)
	chDesc = p2p.ChannelDescriptor{
		ID:                  byte(chID),
		Priority:            5,
		SendQueueCapacity:   10,
		RecvMessageCapacity: 10,
		MaxSendBytes:        1000,
	}

	selfKey  crypto.PrivKey = ed25519.GenPrivKeyFromSecret([]byte{0xf9, 0x1b, 0x08, 0xaa, 0x38, 0xee, 0x34, 0xdd})
	selfID                  = p2p.NodeIDFromPubKey(selfKey.PubKey())
	selfInfo                = p2p.NodeInfo{
		NodeID:     selfID,
		ListenAddr: "0.0.0.0:0",
		Network:    "test",
		Moniker:    string(selfID),
	}

	peerKey  crypto.PrivKey = ed25519.GenPrivKeyFromSecret([]byte{0x84, 0xd7, 0x01, 0xbf, 0x83, 0x20, 0x1c, 0xfe})
	peerID                  = p2p.NodeIDFromPubKey(peerKey.PubKey())
	peerInfo                = p2p.NodeInfo{
		NodeID:     peerID,
		ListenAddr: "0.0.0.0:0",
		Network:    "test",
		Moniker:    string(peerID),
	}
)
