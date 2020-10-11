import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {makeStyles, Container, Card, CardContent} from "@material-ui/core";
import {getList} from "../../service/pictrueApi"
import {downloadPath} from "../../service/file"
import Pagination from '@material-ui/lab/Pagination';

const useStyles = makeStyles((theme) => ({
  root: {
    textAlign: 'center',
  },
  wrapper: {
    marginTop: 8,
  },
  content: {
    textAlign: 'center',
    paddingTop: 16,
    paddingBottom: 16,
    paddingLeft: 16,
    paddingRight: 16,
  },
  pagination: {
    padding: 16,
    justifyContent: 'center',
  }
}))

const MEDIA_TYPE_KEY = 'image';

const PictureViewer = function ({props}) {
  const viewerParams = useParams();
  const classes = useStyles()
  debugger
  const [upLevelDir, setUpLevelDir] = useState(viewerParams['upLevelDir'])
  const [parentId, setParentId] = useState(viewerParams['parentId'])
  const [index, setIndex] = useState(parseInt(viewerParams['index']))
  const [picList, setPicList] = useState([{fileId: undefined}, {fileId: undefined}, {fileId: undefined}, {fileId: undefined}, {fileId: undefined}])

  const handleChange = function (event, value) {
    setIndex(value - 1)
  }
  const clickContentHandler = function () {
    let nextPage = index + 1
    if (picList.length < nextPage) {
      return
    }
    setIndex(nextPage)
  }

  useEffect(() => {
    loadPictureList(parentId, MEDIA_TYPE_KEY, setPicList)
    debugger
  }, [parentId])
  return (
    <Container>
      <Card className={classes.wrapper} variant="outlined">
        <div onClick={clickContentHandler} className={classes.content}>
          <img
            src={picList[index].fileId ? downloadPath(picList[index].fileId) : `https://th.bing.com/th/id/OIP.tEUHSnG-H2fQP_XdUSzKTQHaHa?pid=Api&rs=1`}/>
        </div>
      </Card>
      <Pagination className={classes.pagination} shape="rounded" count={picList.length} page={index + 1}
                  onChange={handleChange}/>
    </Container>
  )
}

const loadPictureList = async function (parentId, mediaType, setPicList) {
  const res = await getList({parentId, mediaType})
  debugger
  setPicList(res.data === undefined ? [{fileId: undefined}, {fileId: undefined}, {fileId: undefined}, {fileId: undefined}, {fileId: undefined}] : res.data)
  return res || {}
}

export default PictureViewer;