# time

    format := "1/2/2006 15:04:05"
    parsed := time.Parse( format, some_date_string)
    hour := parsed.Hour()
    dow := parsed.Weekday()