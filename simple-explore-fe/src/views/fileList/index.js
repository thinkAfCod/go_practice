import React, {useEffect, useState} from "react";
import {useParams, useNavigate} from 'react-router-dom';
import {Container, List, makeStyles} from "@material-ui/core";
import Pagination from '@material-ui/lab/Pagination';
import FileItem from "./fileItem";
import UpLevelItem from "./upLevelItem";
import {getFileItemPage} from '../../service/file'

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

function FileList(props) {
  let params = useParams()
  console.log('FileList', params)
  let handledParams
  try {
    handledParams = parseInt(params['*'], 10)
  } catch (e) {
    handledParams = 0
  }
  debugger
  const navigate = useNavigate()
  const classes = useStyles();
  const [page, setPage] = useState({
    pageSize: handledParams.pageSize === undefined ? 10 : handledParams.pageSize,
    pageNo: handledParams.pageNo === undefined ? 1 : handledParams.pageNo,
    pageTotal: 1,
    parentId: handledParams,
  });
  console.log('filelist', page)
  //let cpPage = {page}
  const [lastPage, setLastPage] = useState([{...page}])
  const [fileDataList, setFileDataList] = useState([])
  if (page.parentId !== handledParams) {
    setPage({...page, parentId: handledParams})
  }

  useEffect(() => {
    requestData(page, setPage, setFileDataList)
  }, [page.pageNo, page.parentId])

  const handleChange = (event, value) => {
    console.log('curPage:' + value)
    debugger
    setPage({
      ...page,
      pageNo: value
    });
  };

  const clickDirItemListener = (fileId) => {
    let cpPage = page;
    debugger
    setLastPage([...lastPage, {...cpPage}])
    cpPage.pageNo = 1
    cpPage.parentId = fileId
    setPage({...cpPage})
  }

  const clickBackListener = () => {
    debugger
    let lastPageCp = lastPage.pop()
    setPage({
      ...lastPageCp,
      parentId: lastPageCp.parentId
    })
    navigate(`/app/file/${lastPageCp.parentId}`)
  }

  let fileItems = [];
  if (page.parentId === 0 || page.parentId === '0') {
    fileItems.unshift((<UpLevelItem listener={clickBackListener}/>))
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

const requestData = async (page, setPage, setFileDataList) => {
  let pageParam = page
  const res = await getFileItemPage(pageParam, pageParam.parentId);
  setPage({
    ...page,
    pageNo: res.data.pageNo,
    pageTotal: res.data.pageTotal,
  })
  setFileDataList(res.data.data)
}

export default FileList;