package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/v3/assert"

	appparams "github.com/umee-network/umee/v4/app/params"
)

func TestKeyExchangeRate(t *testing.T) {
	testCases := []struct {
		denom string
		// KeyPrefixExchangeRate | []byte(denom) | 0
		expectedKey []byte
	}{
		{
			denom:       appparams.BondDenom,
			expectedKey: []byte{0x1, 0x75, 0x75, 0x6d, 0x65, 0x65, 0x0},
		},
		{
			denom:       IbcDenomLuna,
			expectedKey: []byte{0x1, 0x69, 0x62, 0x63, 0x2f, 0x30, 0x45, 0x46, 0x31, 0x35, 0x44, 0x46, 0x32, 0x46, 0x30, 0x32, 0x34, 0x38, 0x30, 0x41, 0x44, 0x45, 0x30, 0x42, 0x42, 0x36, 0x45, 0x38, 0x35, 0x44, 0x39, 0x45, 0x42, 0x42, 0x35, 0x44, 0x41, 0x45, 0x41, 0x32, 0x38, 0x33, 0x36, 0x44, 0x33, 0x38, 0x36, 0x30, 0x45, 0x39, 0x46, 0x39, 0x37, 0x46, 0x39, 0x41, 0x41, 0x44, 0x45, 0x34, 0x46, 0x35, 0x37, 0x41, 0x33, 0x31, 0x41, 0x41, 0x30, 0x0},
		},
		{
			denom:       IbcDenomAtom,
			expectedKey: []byte{0x1, 0x69, 0x62, 0x63, 0x2f, 0x32, 0x37, 0x33, 0x39, 0x34, 0x46, 0x42, 0x30, 0x39, 0x32, 0x44, 0x32, 0x45, 0x43, 0x43, 0x44, 0x35, 0x36, 0x31, 0x32, 0x33, 0x43, 0x37, 0x34, 0x46, 0x33, 0x36, 0x45, 0x34, 0x43, 0x31, 0x46, 0x39, 0x32, 0x36, 0x30, 0x30, 0x31, 0x43, 0x45, 0x41, 0x44, 0x41, 0x39, 0x43, 0x41, 0x39, 0x37, 0x45, 0x41, 0x36, 0x32, 0x32, 0x42, 0x32, 0x35, 0x46, 0x34, 0x31, 0x45, 0x35, 0x45, 0x42, 0x32, 0x0},
		},
	}

	for i, testCase := range testCases {
		actualKey := KeyExchangeRate(testCase.denom)
		t.Run(fmt.Sprintf("test %d - expected key: %s should be the same as actual key: %s", i, testCase.expectedKey, actualKey), func(t *testing.T) {
			assert.DeepEqual(t, testCase.expectedKey, actualKey)
		})
	}
}

func TestKeyFeederDelegation(t *testing.T) {
	testCases := []struct {
		val sdk.ValAddress
		// KeyPrefixFeederDelegation | lengthPrefixed(addr)
		expectedKey []byte
	}{
		{
			val:         []byte("addr________________"),
			expectedKey: []byte{0x2, 0x14, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f},
		},
	}

	for i, testCase := range testCases {
		actualKey := KeyFeederDelegation(testCase.val)
		t.Run(fmt.Sprintf("test %d - expected key: %s should be the same as actual key: %s", i, testCase.expectedKey, actualKey), func(t *testing.T) {
			assert.DeepEqual(t, testCase.expectedKey, actualKey)
		})
	}
}

func TestKeyMissCounter(t *testing.T) {
	testCases := []struct {
		val sdk.ValAddress
		// KeyPrefixMissCounter | lengthPrefixed(addr)
		expectedKey []byte
	}{
		{
			val:         []byte("addr________________"),
			expectedKey: []byte{0x3, 0x14, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f},
		},
	}

	for i, testCase := range testCases {
		actualKey := KeyMissCounter(testCase.val)
		t.Run(fmt.Sprintf("test %d - expected key: %s should be the same as actual key: %s", i, testCase.expectedKey, actualKey), func(t *testing.T) {
			assert.DeepEqual(t, testCase.expectedKey, actualKey)
		})
	}
}

func TestKeyAggregateExchangeRatePrevote(t *testing.T) {
	testCases := []struct {
		val sdk.ValAddress
		// KeyPrefixAggregateExchangeRatePrevote | lengthPrefixed(addr)
		expectedKey []byte
	}{
		{
			val:         []byte("addr________________"),
			expectedKey: []byte{0x4, 0x14, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f},
		},
	}

	for i, testCase := range testCases {
		actualKey := KeyAggregateExchangeRatePrevote(testCase.val)
		t.Run(fmt.Sprintf("test %d - expected key: %s should be the same as actual key: %s", i, testCase.expectedKey, actualKey), func(t *testing.T) {
			assert.DeepEqual(t, testCase.expectedKey, actualKey)
		})
	}
}

func TestKeyAggregateExchangeRateVote(t *testing.T) {
	testCases := []struct {
		val sdk.ValAddress
		// KeyPrefixAggregateExchangeRateVote | lengthPrefixed(addr)
		expectedKey []byte
	}{
		{
			val:         []byte("addr________________"),
			expectedKey: []byte{0x5, 0x14, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f},
		},
	}

	for i, testCase := range testCases {
		actualKey := KeyAggregateExchangeRateVote(testCase.val)
		t.Run(fmt.Sprintf("test %d - expected key: %s should be the same as actual key: %s", i, testCase.expectedKey, actualKey), func(t *testing.T) {
			assert.DeepEqual(t, testCase.expectedKey, actualKey)
		})
	}
}

func TestParseDenomAndBlockFromHistoricPriceKey(t *testing.T) {
	denom := "umee"
	blockNum := uint64(4567)
	key := KeyHistoricPrice(denom, blockNum)

	parsedDenom, parsedBlockNum := ParseDenomAndBlockFromKey(key, KeyPrefixHistoricPrice)
	assert.Equal(t, denom, parsedDenom)
	assert.Equal(t, blockNum, parsedBlockNum)
}

func TestParseDenomAndBlockFromMedianKey(t *testing.T) {
	denom := "umee"
	blockNum := uint64(4567)
	key := KeyMedian(denom, blockNum)

	parsedDenom, parsedBlockNum := ParseDenomAndBlockFromKey(key, KeyPrefixMedian)
	assert.Equal(t, denom, parsedDenom)
	assert.Equal(t, blockNum, parsedBlockNum)
}

func TestParseDenomAndBlockFromMedianDeviationKey(t *testing.T) {
	denom := "umee"
	blockNum := uint64(4567)
	key := KeyMedianDeviation(denom, blockNum)

	parsedDenom, parsedBlockNum := ParseDenomAndBlockFromKey(key, KeyPrefixMedianDeviation)
	assert.Equal(t, denom, parsedDenom)
	assert.Equal(t, blockNum, parsedBlockNum)
}
