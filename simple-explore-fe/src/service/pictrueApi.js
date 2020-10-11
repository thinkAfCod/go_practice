import {get} from "../utils/request";

export async function getList(params) {
  const url = `/api/file?parentId=${params.parentId}&mediaType=${params.mediaType}`;
  const rs = await get(url);
  debugger
  return rs || {};
}