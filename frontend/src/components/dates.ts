import { DateTime } from "luxon";

export const formatToDate = (timestamp: number) => {
  return DateTime.fromSeconds(timestamp).toFormat("yyyy-MM-dd HH:mm:ss")
}
