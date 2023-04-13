import { DateTime } from "luxon"

export const formatToDate = (timestamp: number) => {
  return DateTime.fromSeconds(timestamp)
}

export const formatToDateString = (timestamp: number) => {
  return formatToDate(timestamp).toFormat("yyyy-MM-dd HH:mm:ss")
}
