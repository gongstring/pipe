// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, b3log.org & hacpai.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package cron includes all cron tasks.
package cron

import (
	"os"

	"github.com/b3log/pipe/log"
)

// Logger
var logger = log.NewLogger(os.Stdout)

// Start starts all cron tasks.
func Start() {
	refreshRecommendArticlesPeriodically()
	pushArticlesPeriodically()
	pushCommentsPeriodically()
}
