package transformer

import (
	"encoding/json"
	"testing"

	"github.com/qtumproject/janus/pkg/eth"
	"github.com/qtumproject/janus/pkg/internal"
	"github.com/qtumproject/janus/pkg/qtum"
)

func TestChainIdMainnet(t *testing.T) {
	testChainIdsImpl(t, "main", "0x1")
}

func TestChainIdTestnet(t *testing.T) {
	testChainIdsImpl(t, "test", "0x2")
}

func TestChainIdRegtest(t *testing.T) {
	testChainIdsImpl(t, "regtest", "0x22ba")
}

func TestChainIdUnknown(t *testing.T) {
	testChainIdsImpl(t, "???", "0x1")
}

func testChainIdsImpl(t *testing.T, chain string, expected string) {
	//preparing request
	requestParams := []json.RawMessage{}
	request, err := internal.PrepareEthRPCRequest(1, requestParams)
	if err != nil {
		t.Fatal(err)
	}

	mockedClientDoer := internal.NewDoerMappedMock()

	//preparing client response
	getBlockCountResponse := qtum.GetBlockChainInfoResponse{Chain: chain}
	err = mockedClientDoer.AddResponseWithRequestID(2, qtum.MethodGetBlockChainInfo, getBlockCountResponse)
	if err != nil {
		t.Fatal(err)
	}

	qtumClient, err := internal.CreateMockedClientForNetwork(mockedClientDoer, qtum.ChainAuto)
	if err != nil {
		t.Fatal(err)
	}

	//preparing proxy & executing request
	proxyEth := ProxyETHChainId{qtumClient}
	got, jsonErr := proxyEth.Request(request, internal.NewEchoContext())
	if jsonErr != nil {
		t.Fatal(jsonErr)
	}

	want := eth.ChainIdResponse(expected)

	internal.CheckTestResultEthRequestRPC(*request, want, got, t, false)
}
