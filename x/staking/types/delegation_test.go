package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestDelegationEqual(t *testing.T) {
	d1 := NewDelegation(sdk.AccAddress(addr1), addr2, sdk.NewDec(100))
	d2 := d1

	ok := d1.Equal(d2)
	require.True(t, ok)

	d2.ValidatorAddress = addr3
	d2.Shares = sdk.NewDec(200)

	ok = d1.Equal(d2)
	require.False(t, ok)
}

func TestDelegationString(t *testing.T) {
	d := NewDelegation(sdk.AccAddress(addr1), addr2, sdk.NewDec(100))
	require.NotEmpty(t, d.String())
}

func TestUnbondingDelegationEqual(t *testing.T) {
	ubd1 := NewUnbondingDelegation(sdk.AccAddress(addr1), addr2, 0,
		time.Unix(0, 0), sdk.NewInt(0))
	ubd2 := ubd1

	ok := ubd1.Equal(ubd2)
	require.True(t, ok)

	ubd2.ValidatorAddress = addr3

	ubd2.Entries[0].CompletionTime = time.Unix(20*20*2, 0)
	ok = ubd1.Equal(ubd2)
	require.False(t, ok)
}

func TestUnbondingDelegationString(t *testing.T) {
	ubd := NewUnbondingDelegation(sdk.AccAddress(addr1), addr2, 0,
		time.Unix(0, 0), sdk.NewInt(0))

	require.NotEmpty(t, ubd.String())
}

func TestRedelegationEqual(t *testing.T) {
	r1 := NewRedelegation(sdk.AccAddress(addr1), addr2, addr3, 0,
		time.Unix(0, 0), sdk.NewInt(0),
		sdk.NewDec(0))
	r2 := NewRedelegation(sdk.AccAddress(addr1), addr2, addr3, 0,
		time.Unix(0, 0), sdk.NewInt(0),
		sdk.NewDec(0))

	ok := r1.Equal(r2)
	require.True(t, ok)

	r2.Entries[0].SharesDst = sdk.NewDec(10)
	r2.Entries[0].CompletionTime = time.Unix(20*20*2, 0)

	ok = r1.Equal(r2)
	require.False(t, ok)
}

func TestRedelegationString(t *testing.T) {
	r := NewRedelegation(sdk.AccAddress(addr1), addr2, addr3, 0,
		time.Unix(0, 0), sdk.NewInt(0),
		sdk.NewDec(10))

	require.NotEmpty(t, r.String())
}

func TestNewDelegationResp(t *testing.T) {
	cdc := codec.New()
	dr1 := NewDelegationResp(sdk.AccAddress(addr1), addr2, sdk.NewDec(5), sdk.NewInt(5))
	dr2 := NewDelegationResp(sdk.AccAddress(addr1), addr3, sdk.NewDec(5), sdk.NewInt(5))
	drs := DelegationResponses{dr1, dr2}

	bz1, err := json.Marshal(dr1)
	require.NoError(t, err)

	bz2, err := cdc.MarshalJSON(dr1)
	require.NoError(t, err)

	require.Equal(t, bz1, bz2)

	bz1, err = json.Marshal(drs)
	require.NoError(t, err)

	bz2, err = cdc.MarshalJSON(drs)
	require.NoError(t, err)

	require.Equal(t, bz1, bz2)

	var drs2 DelegationResponses
	require.NoError(t, cdc.UnmarshalJSON(bz2, &drs2))
	require.Equal(t, drs, drs2)
}

func TestNewRedelegationEntryResp(t *testing.T) {
	cdc := codec.New()
	entries := []RedelegationEntryResp{
		NewRedelegationEntryResp(0, time.Unix(0, 0), sdk.NewDec(5), sdk.NewInt(5), sdk.NewInt(5)),
		NewRedelegationEntryResp(0, time.Unix(0, 0), sdk.NewDec(5), sdk.NewInt(5), sdk.NewInt(5)),
	}
	rdr1 := NewRedelegationResp(sdk.AccAddress(addr1), addr2, addr3, entries)
	rdr2 := NewRedelegationResp(sdk.AccAddress(addr2), addr1, addr3, entries)
	rdrs := RedelegationResponses{rdr1, rdr2}

	bz1, err := json.Marshal(rdr1)
	require.NoError(t, err)

	bz2, err := cdc.MarshalJSON(rdr1)
	require.NoError(t, err)

	require.Equal(t, bz1, bz2)

	bz1, err = json.Marshal(rdrs)
	require.NoError(t, err)

	bz2, err = cdc.MarshalJSON(rdrs)
	require.NoError(t, err)

	require.Equal(t, bz1, bz2)

	var rdrs2 RedelegationResponses
	require.NoError(t, cdc.UnmarshalJSON(bz2, &rdrs2))
	require.Equal(t, rdrs, rdrs2)
}
