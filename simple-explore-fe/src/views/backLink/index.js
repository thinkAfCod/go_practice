import React from 'react';
import {Box, IconButton} from '@material-ui/core'
import {useParams} from "react-router";
import makeStyles from "@material-ui/core/styles/makeStyles";
import ArrowBackIosIcon from '@material-ui/icons/ArrowBackIos';

const useStyles = makeStyles((theme) => ({
  title: {},
  backButton: {}
}))

const BackLink = (props) => {
  //const navigate = useNavigate()
  const params = useParams()
  const classes = useStyles()
  return (
    <Box>
      <IconButton className={classes.backButton} onClick={window.history.back}>
        <ArrowBackIosIcon/>
      </IconButton>
      <Box flex={'1 1 auto'}>
        <span className={classes.title}>{params.name}</span>
      </Box>
    </Box>
  )
}

export default BackLink;