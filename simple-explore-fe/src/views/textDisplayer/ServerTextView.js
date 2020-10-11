import React, {useState, useEffect} from "react";
import {useParams} from "react-router-dom";
import {TextareaAutosize, Grid, makeStyles} from "@material-ui/core";
import CreateIcon from '@material-ui/icons/Create';
import SaveIcon from '@material-ui/icons/Save';
import SubjectIcon from "@material-ui/icons/Subject";
import {getFile, putFile} from '../../service/file';

const useStyles = makeStyles((theme) => ({
  toolbarLine: {
    width: '100%',
    height: '40px',
    backgroundColor: '#d9d9d9',
    borderBottom: '1px solid #ccc',
    listStyle: 'none',
    margin: 0,
  },
  liItem: {
    display: 'inline-block',
    padding: '8px 10px 8px 10px',
    height: '40px',
    width: '44px',
  },
}))

const ServerTextView = (props) => {
  const {className, contentListener, stateListener} = props
  const classes = useStyles()
  const pathParams = useParams()
  const screenResize = () => {
    console.log('window.innerHeight - 104',window.innerHeight - 124)
    return (window.innerHeight - 124) / 30
  }
  const [content, setContent] = useState('')
  const [rowsMax, setRowsMax] = useState(() => screenResize())
  const [readState, setReadState] = useState(true)

  const requestNovel = React.useCallback(async () => {
    const novel = await getFile(pathParams['fileId'])
    setContent(novel ? novel.toString() : '1')
    if (contentListener) {
      contentListener(novel)
    }
  }, [])

  const modifyReadState = React.useCallback(() => {
    //stateListener(!readState)
    setReadState(!readState)
  }, [])

  const postNovel = React.useCallback(async () => {
    setReadState(true)
    stateListener(true)
    console.log('save', content.length)
    //await putFile(pathParams['fileId'], {content: content})
  }, [])

  const custResize = React.useCallback(() => {
    setRowsMax(screenResize())
  }, [setRowsMax])

  const handleChange = React.useCallback((event) => {
    setContent(event.target.value)
    if (contentListener) {
      contentListener(event.target.value)
    }
  }, [])

  useEffect(() => {
    requestNovel()
    window.addEventListener('resize', custResize)
    return () => {
      window.removeEventListener('resize', custResize);
    }

  }, [requestNovel, custResize])

  console.log('ServerTextView', (window.innerHeight - 104) / 30)
  return (
    <Grid container>
      <ul className={classes.toolbarLine}>
        <li className={classes.liItem} onClick={modifyReadState}>
          {readState
            ? <CreateIcon style={{width: '24px', height: '24px'}}/>
            : <SubjectIcon style={{width: '24px', height: '24px'}}/>
          }
        </li>
        <li className={classes.liItem} style={{float: 'right'}}
            onClick={() => postNovel()}><SaveIcon
          style={{width: '24px', height: '24px'}}/>
        </li>
      </ul>
      <TextareaAutosize rowsMax={rowsMax} defaultValue={content} className={className} readOnly={readState}
                        onChange={handleChange}/>
    </Grid>
  )
}

export default ServerTextView;