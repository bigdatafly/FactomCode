// Copyright 2015 FactomProject Authors. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package server

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	"github.com/FactomProject/FactomCode/common"
	"github.com/FactomProject/FactomCode/util"
	"github.com/FactomProject/factoid/block"
	// "github.com/davecgh/go-spew/spew"
)


func GetEntryCreditBalance(pubKey *[32]byte) (int32, error) {
	return eCreditMap[string(pubKey[:])], nil
}

func exportDChain(chain *common.DChain) {
	// if len(chain.Blocks) == 0 || procLog.Level() < factomlog.Info {
		// return
	// }

	dBlocks, _ := db.FetchAllDBlocks()
	sort.Sort(util.ByDBlockIDAccending(dBlocks))

	for _, block := range dBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				// procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportEChain(chain *common.EChain) {
	// if len(chain.Blocks) == 0 || procLog.Level() < factomlog.Info {
		// return
	// }

	eBlocks, _ := db.FetchAllEBlocksByChain(chain.ChainID)
	sort.Sort(util.ByEBlockIDAccending(*eBlocks))

	for _, block := range *eBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				// procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}

		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.%09d.block", block.Header.EBSequence, block.Header.EBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportECChain(chain *common.ECChain) {
	// if len(chain.Blocks) == 0 || procLog.Level() < factomlog.Info {
		// return
	// }

	ecBlocks, _ := db.FetchAllECBlocks()
	sort.Sort(util.ByECBlockIDAccending(ecBlocks))

	for _, block := range ecBlocks {
		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				// procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.EBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportAChain(chain *common.AdminChain) {
	// if len(chain.Blocks) == 0 || procLog.Level() < factomlog.Info {
		// return
	// }

	aBlocks, _ := db.FetchAllABlocks()
	sort.Sort(util.ByABlockIDAccending(aBlocks))

	for _, block := range aBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				// procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func exportFctChain(chain *common.FctChain) {
	// if len(chain.Blocks) == 0 || procLog.Level() < factomlog.Info {
		// return
	// }

	FBlocks, _ := db.FetchAllFBlocks()
	sort.Sort(util.ByFBlockIDAccending(FBlocks))

	for _, block := range FBlocks {

		data, err := block.MarshalBinary()
		if err != nil {
			panic(err)
		}

		strChainID := chain.ChainID.String()
		if fileNotExists(dataStorePath + strChainID) {
			err := os.MkdirAll(dataStorePath+strChainID, 0777)
			if err == nil {
				// procLog.Info("Created directory " + dataStorePath + strChainID)
			} else {
				procLog.Error(err)
			}
		}
		err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.GetDBHeight()), data, 0777)
		if err != nil {
			panic(err)
		}
	}
}

// to export individual block once at a time - for debugging ------------------------
func exportDBlock(block *common.DirectoryBlock) {
	// if block == nil || procLog.Level() < factomlog.Info {
		// return
	// }
	fmt.Println("export Dir Block: " + strconv.FormatUint(uint64(block.Header.DBHeight), 10))

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := dchain.ChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			// procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportEBlock(block *common.EBlock) {
	// if block == nil || procLog.Level() < factomlog.Info {
		// return
	// }
		fmt.Println("export Entry Block: " + strconv.FormatUint(uint64(block.Header.EBSequence), 10))

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.Header.ChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			// procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}

	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.%09d.block", block.Header.EBSequence, block.Header.EBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportECBlock(block *common.ECBlock) {
	// if block == nil || procLog.Level() < factomlog.Info {
		// return
	// }
	fmt.Println("export EC Block: " + strconv.FormatUint(uint64(block.Header.EBHeight), 10))

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.Header.ECChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			// procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.EBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportABlock(block *common.AdminBlock) {
	// if block == nil || procLog.Level() < factomlog.Info {
		// return
	// }
	fmt.Println("export Admin Block: " + strconv.FormatUint(uint64(block.Header.DBHeight), 10))

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.Header.AdminChainID.String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			// procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.Header.DBHeight), data, 0777)
	if err != nil {
		panic(err)
	}

}

func exportFctBlock(block block.IFBlock) {
	// if block == nil || procLog.Level() < factomlog.Info {
		// return
	// }
	fmt.Println("export Factoid Block: " + strconv.FormatUint(uint64(block.GetDBHeight()), 10))

	data, err := block.MarshalBinary()
	if err != nil {
		panic(err)
	}

	strChainID := block.GetChainID().String()
	if fileNotExists(dataStorePath + strChainID) {
		err := os.MkdirAll(dataStorePath+strChainID, 0777)
		if err == nil {
			// procLog.Info("Created directory " + dataStorePath + strChainID)
		} else {
			procLog.Error(err)
		}
	}
	err = ioutil.WriteFile(fmt.Sprintf(dataStorePath+strChainID+"/store.%09d.block", block.GetDBHeight()), data, 0777)
	if err != nil {
		panic(err)
	}

}

//--------------------------------------

func getPrePaidChainKey(entryHash *common.Hash, chainIDHash *common.Hash) string {
	return chainIDHash.String() + entryHash.String()
}

func copyCreditMap(
	originalMap map[string]int32,
	newMap map[string]int32) {
	newMap = make(map[string]int32)

	// copy every element from the original map
	for k, v := range originalMap {
		newMap[k] = v
	}

}

func printCreditMap() {
	procLog.Debug("eCreditMap:")
	for key := range eCreditMap {
		procLog.Debugf("Entry credit Key: %x Value %d", key, eCreditMap[key])
	}
}

func fileNotExists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return true
	}
	return err != nil
}
