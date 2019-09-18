package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wavesplatform/gowaves/pkg/util"
)

const (
	rollbackEdge = 3500
	totalBlocks  = 4000
)

type historyTestObjects struct {
	stor *testStorageObjects
	fmt  *historyFormatter
}

func createHistory() (*historyTestObjects, []string, error) {
	stor, path, err := createStorageObjects()
	if err != nil {
		return nil, path, err
	}
	fmt, err := newHistoryFormatter(stor.stateDB)
	if err != nil {
		return nil, path, err
	}
	return &historyTestObjects{stor, fmt}, path, nil
}

func TestNormalize(t *testing.T) {
	to, path, err := createHistory()
	assert.NoError(t, err, "createHistory() failed")

	defer func() {
		to.stor.close(t)

		err = util.CleanTemporaryDirs(path)
		assert.NoError(t, err, "failed to clean test data dirs")
	}()

	// Create history record and add blocks.
	ids := genRandBlockIds(t, totalBlocks)
	history, err := newHistoryRecord(alias)
	assert.NoError(t, err, "newHistoryRecord() failed")
	for _, id := range ids {
		to.stor.addBlock(t, id)
		blockNum, err := to.stor.stateDB.blockIdToNum(id)
		assert.NoError(t, err, "blockIdToNum() failed")
		entry := historyEntry{nil, blockNum}
		err = history.appendEntry(entry)
		assert.NoError(t, err, "appendEntry() failed")
	}
	to.stor.flush(t)

	// Rollback some of blocks.
	for i, id := range ids {
		if i >= rollbackEdge {
			err = to.stor.stateDB.rollbackBlock(id)
			assert.NoError(t, err, "rollbackBlock() failed")
		}
	}

	// Normalize and check the result.
	changed, err := to.fmt.normalize(history, true)
	assert.NoError(t, err, "normalize() failed")
	assert.Equal(t, true, changed)
	rollbackMinHeight, err := to.stor.stateDB.getRollbackMinHeight()
	assert.NoError(t, err, "getRollbackMinHeight() failed")
	oldRecordNumber := 0
	for _, entry := range history.entries {
		blockID, err := to.stor.stateDB.blockNumToId(entry.blockNum)
		assert.NoError(t, err, "blockNumToId() failed")
		entryHeight, err := to.stor.rw.newestHeightByBlockID(blockID)
		assert.NoError(t, err, "newestHeightByBlockID() failed")
		if entryHeight < rollbackMinHeight {
			oldRecordNumber++
		}
		if entryHeight > rollbackEdge {
			t.Errorf("History formatter did not erase invalid blocks.")
		}
	}
	if oldRecordNumber != 1 {
		t.Errorf("History formatter did not cut old blocks.")
	}
}
