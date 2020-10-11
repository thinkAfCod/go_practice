import React from "react";
import {Grid, Container, Divider, makeStyles} from '@material-ui/core'
import MarkdownView from "./markdownView";
import ServerTextView from "./ServerTextView";

const useStyles = makeStyles((theme) => ({
  editor: {
    //height: 'calc(100%-104px)',
    width: '100%',
    padding: '10px',
    marginBottom: 0,
    resize: 'none',
    color: '#333',
    backgroundColor: 'transparent',
    fontSize: '18px',
    fontWeight: 400,
    lineHeight: '30px',
    border: 'none',
    outline: 'none',
    WebkitAppearance: 'none',
    overflow: 'auto',
    fontFamily: ['Microsoft YaHei', '微软雅黑'],
  },
  preview: {
    padding: 10,
    fontSize: '18px',
    fontWeight: 400,
    lineHeight: '30px',
    overflow: 'auto',
    fontFamily: ['Microsoft YaHei', '微软雅黑'],
  },
  ul: {
    lineHeight: '100%',
    listStyle: 'none',
  },
  li: {
    width: '50%',
    display: 'inline-flex',
  },
  dividerLi:{
    width: '2px',
  }
}))

const MdPreviewTextViewLayout = ({listener}) => {
  const classes = useStyles()
  const [content, setContent] = React.useState('')

  const handleContentChange = React.useCallback((newContent) => {
    setContent(newContent)
  }, [])

  return (
    <ul className={classes.ul}>
      <li className={classes.li}>
        <MarkdownView className={classes.preview} content={content}/>
      </li>

      <li key={'ServerTextView-editor-1'} className={classes.li}>
        <ServerTextView className={classes.editor} contentListener={handleContentChange} stateListener={listener}/>
      </li>
    </ul>
  )
}

export default MdPreviewTextViewLayout;