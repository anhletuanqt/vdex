package user

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cxptek/vdex/config"
	"github.com/cxptek/vdex/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func encodeLoginEIP712(typedData *apitypes.TypedData) ([]byte, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))

	return crypto.Keccak256Hash(rawData).Bytes(), nil
}

func verifyLoginEIP712(address string, signature string) error {
	typeData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{
					Name: "name",
					Type: "string",
				},
				{
					Name: "version",
					Type: "string",
				},
				{
					Name: "chainId",
					Type: "uint256",
				},
				{
					Name: "verifyingContract",
					Type: "address",
				},
			},
			"VDEX": []apitypes.Type{
				{
					Name: "Action",
					Type: "string",
				},
				{
					Name: "OnlySignOn",
					Type: "string",
				},
			},
		},
		PrimaryType: "VDEX",
		Domain: apitypes.TypedDataDomain{
			Name:              "VDEX",
			Version:           "1.0",
			ChainId:           math.NewHexOrDecimal256(config.GetConfig().ChainID),
			VerifyingContract: "0x0000000000000000000000000000000000000000",
		},
		Message: apitypes.TypedDataMessage{
			"Action":     "Wallet Verification",
			"OnlySignOn": "https://vdex.finance/trade",
		},
	}
	sigBytes, err := hex.DecodeString(signature[2:])
	if err != nil {
		return err
	}
	hash, err := encodeLoginEIP712(&typeData)
	if err != nil {
		return err
	}
	sigBytes[64] -= 27
	pubKeyBytes, err := crypto.Ecrecover(hash, sigBytes)
	if err != nil {
		return err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if !strings.EqualFold(address, recoveredAddr.Hex()) {
		return errors.New("invalid address")
	}

	return nil
}

func signJWT(user *models.User) (string, error) {
	claim := jwt.MapClaims{
		"publicId":  user.PublicID,
		"address":   user.Address,
		"expiredAt": time.Now().Add(7 * 24 * time.Hour),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	return token.SignedString([]byte(config.GetConfig().JwtSecret))
}
