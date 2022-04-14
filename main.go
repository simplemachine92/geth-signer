package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

	/* "github.com/ethereum/go-ethereum/common/hexutil" */
	"github.com/ethereum/go-ethereum/core/types"
	/* "github.com/ethereum/go-ethereum/crypto" */
	"github.com/ethereum/go-ethereum/ethclient"
)

// use os package to get the env variable which is already set
func envVariable(key string) string {

	// set env variable using os package
	os.Setenv(key, "justin")

	// return the env variable using os package
	return os.Getenv(key)
}

// Using examples @ https://goethereumbook.org/en/signature-generate/

func main() {

	/* value := envVariable("name") */

	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/512122ae7cd847d4a5b78c1810cc4bc2")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}

	/* fmt.Printf("os package: name = %s \n", value)
	fmt.Printf("environment = %s \n", os.Getenv("APP_PK"))

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301 */
}
