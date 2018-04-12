package stat

import (
	"fmt"

	"github.com/KyberNetwork/reserve-data/common"
	ethereum "github.com/ethereum/go-ethereum/common"
)

const TESTHASH string = "0x5f29695f9978ca672a330e8e722cc3df70c3ee15adb5fc6b380818d40ad9cf44"

type LogStorageTest struct {
	storage LogStorage
}

func NewLogStorageTest(storage LogStorage) *LogStorageTest {
	return &LogStorageTest{storage}
}

func (self *LogStorageTest) TestCatLog() error {
	var err error
	var catLog = common.SetCatLog{
		Timestamp:       111,
		BlockNumber:     222,
		TransactionHash: ethereum.StringToHash("TESTHASH"),
		Index:           1,
		Address:         ethereum.StringToAddress(TESTUSERADDR),
		Category:        "test",
	}
	err = self.storage.StoreCatLog(catLog)
	if err != nil {
		return err
	}
	catLog = common.SetCatLog{
		Timestamp:       333,
		BlockNumber:     444,
		TransactionHash: ethereum.StringToHash("TESTHASH"),
		Index:           2,
		Address:         ethereum.StringToAddress(TESTUSERADDR),
		Category:        "test",
	}
	err = self.storage.StoreCatLog(catLog)
	if err != nil {
		return err
	}
	result, err := self.storage.GetCatLogs(0, 8640000)
	if err != nil {
		return err
	}
	if len(result) != 2 {
		return fmt.Errorf("GetCatLogs return wrong number of records, expected 2, got %d", len(result))
	}
	record, err := self.storage.GetFirstCatLog()
	if err != nil {
		return err
	}
	if record.BlockNumber != 222 {
		return fmt.Errorf("GetFirstCatLog return wrong record, expect BlockNumber 222, got %d", record.BlockNumber)
	}
	record, err = self.storage.GetLastCatLog()
	if err != nil {
		return err
	}
	if record.BlockNumber != 444 {
		return fmt.Errorf("GetFirstCatLog return wrong record, expect BlockNumber 444, got %d", record.BlockNumber)
	}
	return err
}

func (self *LogStorageTest) TestTradeLog() error {
	var err error
	var tradeLog = common.TradeLog{
		Timestamp:       111,
		BlockNumber:     222,
		TransactionHash: ethereum.StringToHash("TESTHASH"),
		Index:           1,
	}
	err = self.storage.StoreTradeLog(tradeLog, 111)
	if err != nil {
		return err
	}
	tradeLog = common.TradeLog{
		Timestamp:       333,
		BlockNumber:     444,
		TransactionHash: ethereum.StringToHash("TESTHASH"),
		Index:           2,
	}
	err = self.storage.StoreTradeLog(tradeLog, 333)
	if err != nil {
		return err
	}
	result, err := self.storage.GetTradeLogs(0, 8640000)
	if err != nil {
		return err
	}
	if len(result) != 2 {
		return fmt.Errorf("GetCatLogs return wrong number of records, expected 2, got %d", len(result))
	}
	record, err := self.storage.GetFirstTradeLog()
	if err != nil {
		return err
	}
	if record.BlockNumber != 222 {
		return fmt.Errorf("GetFirstCatLog return wrong record, expect BlockNumber 222, got %d", record.BlockNumber)
	}
	record, err = self.storage.GetLastTradeLog()
	if err != nil {
		return err
	}
	if record.BlockNumber != 444 {
		return fmt.Errorf("GetFirstCatLog return wrong record, expect BlockNumber 444, got %d", record.BlockNumber)
	}
	return err
}
