import React from "react";
import {Grid, Divider, makeStyles} from '@material-ui/core'
import ServerTextView from "./ServerTextView";

const useStyles = makeStyles((theme) => ({
  editor: {
    width: '100%',
    height: 'calc(100%-104px)',
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
  }
}))

const OnlyTextViewLayout = ({listener}) => {
  const classes = useStyles()
  return (
    <Grid>
      <Grid>
        <ServerTextView className={classes.editor} stateListener={listener}/>
      </Grid>
    </Grid>
  )
}

export default OnlyTextViewLayout;