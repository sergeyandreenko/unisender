/*
 * TenderHub unisender service
 *
 * Copyright (c) 2023 TenderHub LLC. All rights reserved
 *
 * This software is the confidential and proprietary information of
 * TenderHub LLC ("Confidential Information"). You shall not disclose
 * such Confidential Information.
 *
 */

package common

import (
	"fmt"
	"strings"
	"time"
)

const ctLayout = "2006-01-02 15:04:05"

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

func (ct *CustomTime) IsSet() bool {
	return ct.Time.IsZero()
}
