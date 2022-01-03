package pay

import (
	"context"
	"crypto/ecdsa"
	"errors"

	"fmt"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yimsoijoi/cryptx/datamodel"
	"golang.org/x/crypto/sha3"
)

// PayERC20 pays token to toAddress
func PayERC20(
	ctx context.Context,
	token *datamodel.Token,
	toAddress common.Address,
	amountStr string,
) error {
	client, err := ethclient.Dial("https://speedy-nodes-nyc.moralis.io/ac87e6329b8601865ea39581/bsc/mainnet")
	if err != nil {
		log.Println("init client failed", err.Error())
		return errors.New("init client failed")
	}

	privateKey, err := crypto.HexToECDSA("MySecretPrivateKey")
	if err != nil {
		log.Println("can't change privatekey to ECDSA")
		return errors.New("parse ECDSA failed")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("can't assert type: publickey is not of type *ecdsa.Publickey")
		return errors.New("assert type failed")
	}

	tokenAddress := common.HexToAddress(string(token.Address))
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Println("can't pending nonce", err.Error())
		return errors.New("pending nonce failed")
	}
	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Println("can't estimate gas price", err.Error())
		return errors.New("estimate gas price failed")
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	amount, ok = amount.SetString(amountStr, token.Decimal)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Println("can't get gaslimit:", err.Error())
		return errors.New("get gaslimit failed")
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Println("can't get chainID:", err.Error())
		return errors.New("get chainId failed")
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println("can't sign transaction:", err.Error())
		return errors.New("sign transaction failed")
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Println("can't send transaction:", err.Error())
		return errors.New("send transaction failed")
	}
	fmt.Printf("tx sent: %v", signedTx.Hash().Hex())
	return nil
}
