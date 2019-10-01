/*************************************************************************
 * Copyright 2018 Gravwell, Inc. All rights reserved.
 * Contact: <legal@gravwell.io>
 *
 * This software may be modified and distributed under the terms of the
 * BSD 2-clause license. See the LICENSE file for details.
 **************************************************************************/

package version

import (
	"fmt"
	"io"
	"time"
)

const (
	MajorVersion = 3
	MinorVersion = 2
	PointVersion = 4
)

var (
	BuildDate time.Time = time.Date(2019, 9, 30, 0, 0, 0, 0, time.UTC)
)

func PrintVersion(wtr io.Writer) {
	fmt.Fprintf(wtr, "Version:\t%d.%d.%d\n", MajorVersion, MinorVersion, PointVersion)
	fmt.Fprintf(wtr, "BuildDate:\t%s\n", BuildDate.Format(`2006-01-02 15:04:05`))
}
