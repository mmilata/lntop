package models

import "github.com/edouardparis/lntop/logging"

type ChannelBalance struct {
	Balance            int64
	PendingOpenBalance int64
}

func (m ChannelBalance) MarshalLogObject(enc logging.ObjectEncoder) error {
	enc.AddInt64("balance", m.Balance)
	enc.AddInt64("pending_open_balance", m.PendingOpenBalance)

	return nil
}

type Channel struct {
	ID                  uint64
	Active              bool
	RemotePubKey        string
	ChannelPoint        string
	Capacity            int64
	LocalBalance        int64
	RemoteBalance       int64
	CommitFee           int64
	CommitWeight        int64
	FeePerKiloWeight    int64
	UnsettledBalance    int64
	TotalAmountSent     int64
	TotalAmountReceived int64
	UpdatesCount        uint64
	CSVDelay            uint32
	Private             bool
	PendingHTLC         []*HTLC
}

func (m Channel) MarshalLogObject(enc logging.ObjectEncoder) error {
	enc.AddUint64("id", m.ID)
	enc.AddBool("active", m.Active)
	enc.AddString("remote_pubkey", m.RemotePubKey)
	enc.AddString("channel_point", m.ChannelPoint)
	enc.AddInt64("capacity", m.Capacity)
	enc.AddInt64("local_balance", m.LocalBalance)
	enc.AddInt64("remote_balance", m.RemoteBalance)
	enc.AddInt64("commit_fee", m.CommitFee)
	enc.AddInt64("commit_weight", m.CommitWeight)
	enc.AddInt64("unsettled_balance", m.UnsettledBalance)
	enc.AddInt64("total_amount_sent", m.TotalAmountSent)
	enc.AddInt64("total_amount_received", m.TotalAmountReceived)
	enc.AddUint64("updates_count", m.UpdatesCount)
	enc.AddBool("private", m.Private)

	return nil
}