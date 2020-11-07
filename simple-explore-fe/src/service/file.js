import {get, put} from "../utils/request";

export async function getFile(fileId) {
  const rs = await get(downloadPath(fileId))
  return rs || ''
}

export async function putFile(fileId, params) {
  const url = `/api/file?id=${fileId}`
  const rs = await put(url, params)
  return rs || {}
}

export async function getFileItemPage(page, parentId) {
  const url = `/api/item?pageSize=${page.pageSize}&pageNo=${page.pageNo}&parentId=${parentId}`
  const rs = await get(url);
  return rs || {};
}

export async function getFileItemParentId(parentId) {
  const url = `/api/item/parent/id?id=${parentId}`
  const rs = await get(url);
  return rs || {};
}

export function downloadPath(fileId) {
  return `/api/file?id=${fileId}`
}