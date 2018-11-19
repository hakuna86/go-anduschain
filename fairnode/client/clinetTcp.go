package fairnodeclient

import (
	"fmt"
	"github.com/anduschain/go-anduschain/common"
	"github.com/anduschain/go-anduschain/crypto/sha3"
	"github.com/anduschain/go-anduschain/fairnode/otprn"
	"github.com/anduschain/go-anduschain/p2p/discv5"
	"github.com/anduschain/go-anduschain/rlp"
	"net"
	"sync"
)

type TransferCheck struct {
	Otprn    otprn.Otprn
	Coinbase common.Address
	Enode    discv5.Node
}

func (tsf *TransferCheck) Hash() common.Hash {
	return rlpHash(tsf)
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func (fc *FairnodeClient) TCPtoFairNode() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("andus >> TCPtoFairNode Start")
	conn, err := net.DialTCP("tcp", nil, fc.SAddrTCP)
	if err != nil {
		fmt.Println("andus >> GETH DialTCP 에러", err)
	}

	defer conn.Close()
	defer fmt.Println("andus >> TCPtoFairNode 죽음")

	<-fc.TcpConnStartCh
	fmt.Println("andus >> TCPtoFairNode 채녈 들어옴")

	go func() {
		defer wg.Done()
		data := make([]byte, 4096)
		for {
			n, err := conn.Read(data)
			if err != nil {
				fmt.Println("andus >> Read 에러!!", err)
			}

			if n > 0 {
				fmt.Println(string(data[:n]))
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			tsf := TransferCheck{*fc.Otprn, *fc.Coinbase, *fc.Enode}

			sendFairnodeData, err := rlp.EncodeToBytes(tsf)
			if err != nil {
				fmt.Println("andus >> sendFairnodeData 에러", err)
			}
			conn.Write(sendFairnodeData)
		}
	}()

	for {
		select {
		case <-fc.tcptoFairNodeStopCh:
			return
		}
	}

	wg.Wait()

	////TODO : andus >> TCP 통신 to FairNode
	////TODO : andus >> 1. fair Node에 TCP 연결
	////TODO : andus >> 2. OTPRN, enode값 전달
	//
	//// TODO : andus >> 1. 채굴 리스 리스트와 총 채굴리그 해시 수신
	//
	//// TODO : andus >> 1.1 추후 서명값 검증 해야함...
	//
	//// TODO : andus >> 4. JoinTx 생성 ( fairnode를 수신자로 하는 tx, 참가비 보냄...)
	//
	//var fairNodeAddr common.Address // TODO : andus >> 보내는 fairNode의 Address(주소)
	//
	//// TODO : andus >> joinNonce 현재 상태 조회
	//
	//currentJoinNonce := fc.GetCurrentJoinNonce()
	//
	//signer := types.NewEIP155Signer(big.NewInt(18))
	//
	//// TODO : andus >> joinNonce Fairnode에게 보내는 Tx
	//tx, err := types.SignTx(types.NewTransaction(currentJoinNonce, fairNodeAddr, new(big.Int), 0, new(big.Int), nil), signer, fc.CoinBasePrivateKey)
	//if err != nil {
	//	log.Println("andus >> JoinTx 서명 에러")
	//}
	//
	//log.Println("andus >> JoinTx 생성 Success", tx)
	//
	//// TODO : andus >> txpool에 추가.. 알아서 이더리움 프로세스 타고 날라감....
	//fc.txPool.AddLocal(tx)
	//
	//// TODO : andus >> 2. 각 enode값을 이용해서 피어 접속
	//
	////enodes := []string{"enode://12121@111.111.111:3303"}
	////for _ := range enodes {
	////	old, _ := disco.ParseNode(boot)
	////	srv.AddPeer(old)
	////}
	//
	//select {
	//// type : types.TransferBlock
	//case signedBlock := <-fc.WinningBlockCh:
	//	// TODO : andus >> 위닝블록 TCP전송
	//	fmt.Println("위닝 블록을 받아서 페어노드에게 보내요", signedBlock)
	//}

	// TODO : andus >> 페어노드가 전송한 최종 블록을 받는다.
	// TODO : andus >> 받은 블록을 검증한다
	// TODO : andus >> worker에 블록이 등록 되게 한다

	//fc.FinalBlockCh <- &types.TransferBlock{}
}
