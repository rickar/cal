# cal: Go (golang) calendar library for dealing with holidays and work days

This library augments the Go time package to provide easy handling of holidays
and work days (business days).

Holiday instances can be exact days, floating days such as the 3rd Monday of
the month, or yearly offsets such as the 100th day of the year.

The Calendar type provides functions for calculating workdays and dealing 
with holidays that are observed on alternate days when they fall on weekends.

