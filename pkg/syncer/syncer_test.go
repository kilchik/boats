package syncer

import (
	"database/sql"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindNextFreeSlot(t *testing.T) {
	assert := assert2.New(t)

	date := func(str string) time.Time {
		d, _ := time.Parse("02.01.2006", str)
		return d
	}

	for _, tcase := range []struct{
		reserved *reservSlots
		expectedFrom sql.NullTime
		expectedTo sql.NullTime
	} {
		{
			reserved:&reservSlots{
				from:[]time.Time{
					date("29.06.2019"),
				},
				to:[]time.Time{
					date("29.07.2019"),
				},
			},
			expectedFrom:sql.NullTime{},
			expectedTo:sql.NullTime{},
		},
		{
			reserved:&reservSlots{
				from:[]time.Time{
					date("29.10.2019"),
				},
				to:[]time.Time{
					date("27.11.2019"),
				},
			},
			expectedFrom:sql.NullTime{date("28.11.2019"), true},
			expectedTo:sql.NullTime{},
		},
		{
			reserved:&reservSlots{
				from:[]time.Time{
					date("15.11.2019"),
				},
				to:[]time.Time{
					date("15.12.2019"),
				},
			},
			expectedFrom:sql.NullTime{date("30.10.2019"),true},
			expectedTo:sql.NullTime{date("14.11.2019"), true},
		},
		{
			reserved:&reservSlots{
				from:[]time.Time{
					date("15.10.2019"),
					date("16.12.2019"),
					date("02.03.2020"),
				},
				to:[]time.Time{
					date("15.12.2019"),
					date("16.01.2020"),
					date("02.04.2020"),
				},
			},
			expectedFrom:sql.NullTime{date("17.01.2020"), true},
			expectedTo:sql.NullTime{date("01.03.2020"), true},
		},
	} {
		from, to := findNextFreeSlot(date("30.10.2019"), tcase.reserved)
		assert.Equal(tcase.expectedFrom, from)
		assert.Equal(tcase.expectedTo, to)
	}
}
