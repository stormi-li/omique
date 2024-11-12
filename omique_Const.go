package omique

import "time"

const const_retryWaitTime = 500 * time.Millisecond
const const_maxRetryCount = 10
const refresh_conn_interval = 2*time.Second
