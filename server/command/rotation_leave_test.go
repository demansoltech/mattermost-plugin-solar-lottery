// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.
package command

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-plugin-solar-lottery/server/sl"
)

func TestCommandRotationLeave(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		SL, _ := getTestSL(t, ctrl)

		err := runCommands(t, SL, `
			/lotto rotation new test-rotation
			/lotto rotation join test-rotation @id1-username @id2-username @id3-username @id4-username
			`)
		require.NoError(t, err)

		deleted := sl.UserMap{}
		_, err = runJSONCommand(t, SL, `
			/lotto rotation leave test-rotation @id2-username @id3-username @id5-username`, &deleted)
		require.NoError(t, err)
		require.Equal(t, []string{"id2", "id3"}, deleted.IDs().Sorted())

		r := sl.NewRotation()
		_, err = runJSONCommand(t, SL, `
			/lotto rotation show test-rotation`, &r)
		require.NoError(t, err)
		require.Equal(t, []string{"id1", "id4"}, r.MattermostUserIDs.Sorted())
	})
}
