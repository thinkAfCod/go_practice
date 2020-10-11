import React, {useState} from 'react';
// import {useParams} from 'react-router-dom';
// import {TextareaAutosize, Box, Grid, Divider, makeStyles} from '@material-ui/core';
// import {getFile, putFile} from '../../service/file'
import MdPreviewTextViewLayout from "./MdPreviewTextViewLayout";
import OnlyTextViewLayout from "./OnlyTextViewLayout";

// const useStyles = makeStyles((theme) => ({
//   container: {
//     overflow: 'hidden',
//     padding: 4,
//     backgroundColor: '#aaa',
//     width: '100%',
//     height: '100%',
//   },
// }))

function TextDisplayer(props) {
  //const classes = useStyles()
  const [readState, setReadState] = useState(true)

  const listenState = React.useCallback((state) => {
    setReadState(state)
  }, [])

  return (
    readState ? <MdPreviewTextViewLayout listener={listenState}/> : <OnlyTextViewLayout listener={listenState}/>
  )
}

export default TextDisplayer;