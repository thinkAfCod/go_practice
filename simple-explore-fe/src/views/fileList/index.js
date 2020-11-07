import React, {useEffect, useState} from "react";
import {useParams, useNavigate} from 'react-router-dom';
import {Container, List, makeStyles} from "@material-ui/core";
import Pagination from '@material-ui/lab/Pagination';
import FileItem from "./fileItem";
import UpLevelItem from "./upLevelItem";
import {getFileItemPage, getFileItemParentId} from '../../service/file'

const useStyles = makeStyles((theme) => ({
  root: {
    minWidth: '100%',
    minHeight: '100%'
  },
  list: {},
  backIcon: {
    frontSize: 'large',
    htmlColor: '#64b5f6'
  }
}));

const defaultPage = {
  pageSize: 10,
  pageNo: 1,
  parentId: 0,
}

function FileList(props) {
  let params = useParams()
  console.log('FileList', params)
  let handledParams = params['*'] === '' ? defaultPage : {...defaultPage, parentId: parseInt(params['*'], 10)}
  const navigate = useNavigate()
  const classes = useStyles();
  const [page, setPage] = useState({
    pageSize: handledParams.pageSize === undefined ? 10 : handledParams.pageSize,
    pageNo: handledParams.pageNo === undefined ? 1 : handledParams.pageNo,
    pageTotal: 1,
    parentId: handledParams.parentId,
  });
  //let cpPage = {page}
  const [lastPage, setLastPage] = useState([{...page}])
  const [fileDataList, setFileDataList] = useState([])
  if (page.parentId !== handledParams.parentId) {
    setPage({...page, parentId: handledParams.parentId})
  }

  const requestFileList = async (page, setPage, lastPage, setLastPage, setFileDataList) => {
    let pageParam = page
    const res = await getFileItemPage(pageParam, pageParam.parentId);
    console.log(res)
    setPage({
      ...page,
      pageTotal: res.data.pageTotal,
    })
    setLastPage([...lastPage, {...page, parentId: res.data.data[0].parentItemId}])
    setFileDataList(res.data.data)
  }

  const requestUpLevelFileItemId = async (parentId) => {
    const res = await getFileItemParentId(parentId);
    console.log(res)
    return res.data
  }

  useEffect(() => {
    requestFileList(page, setPage, lastPage, setLastPage, setFileDataList)
  }, [page.pageNo, page.parentId])

  const handleChange = (event, value) => {
    console.log('curPage:' + value)
    
    setPage({
      ...page,
      pageNo: value
    });
  };

  const clickDirItemListener = (fileId) => {
    let cpPage = page;
    
    // setLastPage([...lastPage, {...cpPage}])
    cpPage.pageNo = 1
    cpPage.parentId = fileId
    setPage({...cpPage})
  }

  const clickBackListener = async () => {
    let lastPageCp = lastPage.pop()
    let parentId = await requestUpLevelFileItemId(page.parentId)
    setPage({
      ...lastPageCp,
      parentId: parentId
    })
  }

  let fileItems = [];
  if (page.parentId !== 0 || page.parentId !== '0') {
    fileItems.unshift((<UpLevelItem key={`back-up-level`} listener={clickBackListener}/>))
  }
  fileDataList.forEach((file) => {
    fileItems.push((<FileItem key={`file-item-${file.id}`} fileData={file} listener={clickDirItemListener}/>))
  })

  return (
    <Container fixed className={classes.root}>
      <span>{lastPage.parentId}</span>
      <List className={classes.root}>
        {fileItems}
      </List>
      <Pagination count={page.pageTotal} page={page.pageNo} shape="rounded" onChange={handleChange}/>
    </Container>
  )
}

export default FileList;